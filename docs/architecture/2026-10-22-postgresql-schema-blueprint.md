# Pinchd — PostgreSQL Schema Blueprint and API Contract (Draft)
**Date:** 22 October 2026  
**Status:** Architecture proposal for review; not a migration, final contract, or implementation authorization.  
**Scope:** Relational model, invariants, transaction boundaries, API surface, idempotency and event delivery.  
**Depends on:** Workstream 4 lifecycle requirements; Workstream 5 open architecture decisions; data classification and retention map; `../research/2026-10-22-full-fold-lifecycle-privacy-architecture-reconciliation.md`; ADR-0001.

## 1. Design constraints

- PostgreSQL is the source of truth for account, eligibility, preferences, likes, matches, blocks, conversations, messages, entitlements and deletion workflow state.
- Go API owns business rules and authorization. Clients never determine eligibility, match creation, block effectiveness or Premium entitlement.
- Model independent state dimensions independently: account lifecycle, age assurance, visibility, enforcement and subscription entitlement.
- Apply hard discovery gates before ranking. Ranking must not bypass hard preferences or enforcement.
- Treat all API input as untrusted. Use authenticated identity, object-level authorization, allowlisted response DTOs and server-side validation.
- Use database constraints for invariants where practical, plus transaction-level checks for rules that span rows.
- Use an outbox for reliable event publication; consumers must be idempotent. Queues, caches and search indexes are projections, not authority.
- Every table needs an owner, purpose, classification, lawful-purpose assessment, retention/deletion rule and access policy before production.
- Names and types below are illustrative. Confirm UUID strategy, extensions, enum strategy, partitioning, RLS, provider limits and migration tooling during a technical proof-of-concept.

## 2. Proposed schema groups

### 2.1 Identity, lifecycle and consent

| Table | Core columns (illustrative) | Constraints / notes |
|---|---|---|
| `accounts` | `id uuid PK`, `auth_subject text UNIQUE NOT NULL`, `lifecycle_state`, `created_at`, `updated_at`, `deleted_at NULL`, `row_version bigint` | One account per auth subject; lifecycle state is not age or enforcement state. Do not store passwords if an identity provider owns authentication. |
| `account_age_assurance` | `account_id PK/FK`, `state`, `provider`, `workflow_version`, `checked_at`, `expires_at`, `appeal_state`, `evidence_ref NULL` | Store minimal result and reference only; avoid copying identity documents/biometrics into product DB. Exact fields depend on Didit evidence and legal review. |
| `account_visibility` | `account_id PK/FK`, `visibility_state`, `changed_at`, `row_version` | Visibility such as visible/paused is separate from account lifecycle. |
| `account_restrictions` | `id uuid PK`, `account_id FK`, `restriction_type`, `starts_at`, `ends_at NULL`, `status`, `case_id NULL`, `created_by` | Multiple restrictions may coexist. Restricted evidence belongs in the separate safety domain. |
| `consent_receipts` | `id uuid PK`, `account_id FK`, `purpose_code`, `notice_version`, `decision`, `recorded_at`, `withdrawn_at NULL` | Purpose-specific, versioned evidence; not a substitute for determining lawful basis. |
| `account_deletion_jobs` | `id uuid PK`, `account_id`, `requested_at`, `state`, `policy_version`, `completed_at NULL`, `failure_code NULL` | Durable orchestration state; idempotent per deletion request and each downstream target. Keep suppression evidence only if counsel approves necessity and duration. |

### 2.2 Profile, preferences and media

| Table | Core columns (illustrative) | Constraints / notes |
|---|---|---|
| `profiles` | `account_id PK/FK`, `display_name`, `birth_date` or minimized age representation, `bio`, `profile_state`, `updated_at`, `row_version` | Birth date is high-impact personal data; decide whether to retain exact date after age check. Never return internal fields in public DTOs. |
| `profile_gender_identities` | `account_id FK`, `identity_code` | Separate from discovery preferences; controlled vocabulary and privacy review. |
| `discovery_preferences` | `account_id PK/FK`, `min_age`, `max_age`, `radius_km`, `location_mode`, `updated_at`, `row_version` | Age and radius are hard boundaries by default. Separate future-family and existing-children preferences. |
| `discovery_gender_preferences` | `account_id FK`, `gender_code` | Selected genders are hard eligibility constraints; define empty-set behavior explicitly (no results, not wildcard). |
| `profile_attribute_preferences` | `account_id FK`, `attribute_code`, `desired_value`, `strictness` | Only use supported fields; strictness is per-field. Sensitive attributes require documented purpose and review. |
| `profile_prompts` | `id uuid PK`, `account_id FK`, `prompt_code`, `answer_text`, `position`, `updated_at` | Prompt catalog should be versioned separately; cap length and sanitize rendering. |
| `profile_media` | `id uuid PK`, `account_id FK`, `object_key`, `media_state`, `position`, `moderation_state`, `created_at`, `deleted_at NULL` | Private object storage; never store public bucket URLs. Strip metadata, scan/moderate and issue short-lived signed access only after authorization. |
| `location_private` | `account_id PK/FK`, `coarse_geography`, `updated_at` | Minimize precision and retention; no client-visible exact coordinates. Geospatial extension/index only if PoC justifies it. |
| `profile_projection` (optional) | `account_id PK`, safe discovery fields, `projection_version` | Rebuildable projection; must not become authority. Omit private/sensitive fields by construction. |

### 2.3 Likes, matches, conversations and safety boundaries

| Table | Core columns (illustrative) | Constraints / notes |
|---|---|---|
| `likes` | `id uuid PK`, `actor_id FK`, `recipient_id FK`, `state`, `created_at`, `withdrawn_at NULL`, `comment_text NULL`, `comment_target_type NULL`, `comment_target_id NULL` | CHECK actor != recipient; one current logical like per ordered pair, enforced by unique key or a separate current-like table. Preserve minimal history only under an approved retention policy. |
| `matches` | `id uuid PK`, `pair_low_id FK`, `pair_high_id FK`, `state`, `created_at`, `ended_at NULL`, `end_reason NULL`, `row_version` | Canonical pair order: pair_low_id < pair_high_id; unique pair key prevents duplicate active match. Decide whether one row is reused or match history is append-only; recommended append-only match records plus a pair guard. |
| `pair_exclusions` | `pair_low_id FK`, `pair_high_id FK`, `reason`, `created_at`, `source_match_id NULL` | Canonical pair key UNIQUE; unmatch creates permanent exclusion per approved policy. Account-deletion suppression duration and lawful basis remain open. |
| `blocks` | `blocker_id FK`, `blocked_id FK`, `created_at`, `reason_code NULL` | CHECK blocker != blocked; unique directed pair. Blocking must immediately gate all reads, sends, discovery and async deliveries in either direction. |
| `conversations` | `id uuid PK`, `match_id FK UNIQUE`, `state`, `created_at`, `closed_at NULL` | Conversation requires a valid match; never treat conversation ID as authorization. |
| `messages` | `id uuid PK`, `conversation_id FK`, `sender_id FK`, `client_message_id`, `body`, `created_at`, `moderation_state`, `deleted_at NULL` | UNIQUE(sender_id, client_message_id); validate sender is a current participant and eligible to message within transaction. Retention, edits, deletion and evidence handling are open policy gates. |
| `reports` | `id uuid PK`, `reporter_id FK`, `subject_id FK NULL`, `object_type`, `object_id`, `category`, `created_at`, `case_id` | Minimal user-facing report record; sensitive case evidence should live in a separately permissioned safety system/schema. |
| `safety_cases` | `id uuid PK`, restricted case fields, assigned role, severity, state | Separate access controls, access audit, retention schedule and incident workflow; never expose through ordinary account APIs. |

### 2.4 Entitlements, audit and reliable work

| Table | Core columns (illustrative) | Constraints / notes |
|---|---|---|
| `store_transactions` | `id uuid PK`, `provider`, `provider_transaction_id`, `account_id FK`, `product_code`, `state`, `verified_at` | UNIQUE(provider, provider_transaction_id); verify signed/provider server notifications and deduplicate renewals/refunds. |
| `entitlements` | `account_id FK`, `entitlement_code`, `state`, `valid_from`, `valid_until NULL`, `source_transaction_id` | Server-authoritative; safety and consent features never depend on paid entitlement. |
| `outbox_events` | `id uuid PK`, `aggregate_type`, `aggregate_id`, `event_type`, `schema_version`, `payload`, `created_at`, `published_at NULL`, `attempts` | Insert in same transaction as domain change; payload contains minimum necessary data, not secrets or full sensitive objects. |
| `consumer_inbox` | `consumer_name`, `event_id`, `processed_at` | UNIQUE(consumer_name, event_id); supports idempotent processing. |
| `idempotency_records` | `account_id`, `route_key`, `idempotency_key`, `request_hash`, `response_status`, `response_ref`, `created_at`, `expires_at` | UNIQUE(account_id, route_key, idempotency_key); same key + same payload replays result; same key + different payload returns conflict. Retention/TTL to be set. |
| `security_audit_events` | `id uuid PK`, actor, action, target type/ID, timestamp, correlation ID, safe metadata | Append-oriented, restricted access, tamper-resistance strategy and retention approved before production; avoid message bodies, credentials and unnecessary sensitive data. |

## 3. Pair identity, uniqueness and indexes

- Canonical pair helper: sort the two UUIDs into `pair_low_id` and `pair_high_id`; reject self-pairs. Apply the same canonicalization to matches and pair exclusions.
- Suggested constraints: UNIQUE(auth_subject); CHECK actor_id <> recipient_id; CHECK pair_low_id < pair_high_id; UNIQUE(pair_low_id, pair_high_id) on pair exclusions; UNIQUE(blocker_id, blocked_id); UNIQUE(match_id) on conversations; UNIQUE(sender_id, client_message_id); UNIQUE(provider, provider_transaction_id); UNIQUE(consumer_name, event_id).
- Likes: use a partial unique index for active/pending logical likes if state transitions permit it, e.g. unique actor/recipient for rows whose state is pending/active. Confirm exact lifecycle and history model before writing SQL.
- Index FK columns used in joins, especially likes by recipient/state/created_at, likes by actor/state, matches by each participant/state, blocks by blocker and blocked, messages by conversation/created_at, reports by case/state/created_at, deletion jobs by state/created_at, outbox by published_at/created_at.
- Avoid indexes containing unnecessary sensitive attributes. Use composite indexes only after query-plan testing with realistic data.
- Use UTC timestamptz for instants. Prefer CHECK constraints and reference tables for controlled codes; choose native enums only if migration evolution is operationally acceptable.
- No DDL should be generated until the team chooses retention, state transitions, migration tooling, and provider-specific capabilities.

## 4. State dimensions and derived eligibility

Do not create a single `is_active` flag as a proxy for all permissions. Proposed independent dimensions:

1. **Account lifecycle:** registration_in_progress, active, deletion_pending, deleted/de-identified.
2. **Age assurance:** unknown, pending, passed, failed, expired, appeal_pending. Only current passed result satisfies the age gate.
3. **Visibility:** visible, paused, hidden_by_policy.
4. **Enforcement:** no active restriction, restricted capabilities, suspended. Store scoped restrictions with start/end and case reference.
5. **Profile/media moderation:** pending, approved, rejected, removed.
6. **Subscription entitlement:** inactive, active, grace/pending confirmation, expired, refunded/revoked (provider mapping must be validated).
7. **Match/conversation:** active, ended/closed; unmatch and block are not equivalent.

A centralized eligibility function/service should evaluate account state, current age assurance, moderation, both parties' preferences, visibility, restrictions, blocks, pair exclusions, deletion state and any required operational policy. It must run for discovery and re-run at like, match and message mutations. Cached discovery results are candidates only and must be revalidated before display/action.

## 5. Transaction boundaries and race conditions

### 5.1 Like and reciprocal match
1. Authenticate actor and validate request schema, target visibility and current eligibility.
2. Begin transaction; lock or serialize on canonical pair key (e.g. pair guard row or transaction-scoped advisory lock chosen after PoC).
3. Re-check both accounts, preferences, age assurance, restrictions, block in either direction, pair exclusion and existing match.
4. Insert/update the actor's like with an idempotency key; reject a self-like.
5. Check valid reciprocal like under the same serialization boundary.
6. If reciprocal and still eligible, insert exactly one match and its conversation, with a unique pair constraint as the final guard.
7. Insert outbox events in the same transaction; commit. Publish asynchronously and idempotently.
8. Return authoritative state. Concurrent reciprocal likes must result in one match/conversation, never duplicate notifications.

### 5.2 Message send
1. Authenticate and validate conversation membership and message bounds.
2. In a transaction, re-check active match/conversation, sender restrictions, recipient state and block in either direction.
3. Enforce sender/client-message idempotency; store message and outbox event atomically.
4. Commit, then asynchronously fan out notifications after re-checking block, account state and notification eligibility.
5. A stale push or cached thread must not enable a message send or reveal content after block/unmatch.

### 5.3 Block and unmatch
- **Block:** write block and any required match/conversation state atomically; publish invalidation event via outbox. The API response must not claim success before the authoritative transaction commits. Every read/send/worker path checks current block state. Cache invalidation is defence-in-depth, not the enforcement boundary.
- **Unmatch:** end match and close conversation, create permanent canonical pair exclusion, and commit one audit/outbox record. No initiator reveal. Remove from active lists immediately; preserve only the approved 30-day read-only past-match history. Report/block while visible; no further messaging.
- Clarify and test whether a pre-existing match remains ended when only a block is later removed; recommended default is it does not automatically restore.

### 5.4 Account deletion
- Record deletion request and freeze normal access according to approved policy.
- Create durable deletion tasks for each target: primary tables, object storage, derived projections/search, caches, notifications, analytics, vendor systems and downstream processors.
- Each task has a stable operation ID, attempts, last error category, completion time and evidence. Retry safely; never let stale webhooks resurrect deleted data.
- Separate data eligible for deletion from records held under a documented lawful exception (e.g., specific safety evidence or financial records). Record the exception, owner, access control and expiry/review date.
- Verify deletion across replicas/projections and vendors. Backups need a documented expiry/restore-and-reapply-deletions strategy.
- Permanent pair exclusion after unmatch is approved. Suppression following account deletion remains an OPEN GATE; do not retain identifiers indefinitely without an approved necessity and retention basis.

## 6. Proposed API surface (versioned under /v1)

All endpoints require TLS. Authenticated endpoints derive account identity from the verified token, not a client-supplied account ID. Use request schemas, response DTO allowlists, per-object authorization, rate limits, correlation IDs and consistent error envelopes. This matrix is a proposed starting point, not final OpenAPI.

| Method and route | Purpose | Auth / key checks | Idempotency / response |
|---|---|---|---|
| POST /v1/registration | Create/continue account | Anonymous or partially authenticated; abuse controls | Registration request key; returns next required step, never discovery access prematurely |
| GET /v1/me | Current account projection | Account owner | No; minimal private DTO |
| PATCH /v1/me/profile | Update profile/prompts | Owner; editable-field allowlist; media ownership | Idempotency optional; row version / If-Match for conflict control |
| PUT /v1/me/preferences | Replace preference set | Owner; validate hard boundaries and strictness | Versioned update; return canonical preferences |
| POST /v1/age-assurance/sessions | Start approved assurance workflow | Eligible registration state; rate limit | Stable operation ID; vendor cost deduplication |
| POST /v1/age-assurance/callbacks/{provider} | Process vendor callback | Vendor signature, replay defense, workflow binding | Deduplicate provider event; never trust browser callback alone |
| POST /v1/media/upload-intents | Request private upload | Owner; quotas; type/size limits | Stable upload intent; object key not client-chosen |
| POST /v1/media/{id}/complete | Finalize upload processing | Owner and upload state | Idempotent; media remains unavailable until safety pipeline clears it |
| GET /v1/discovery | Get eligible candidate page | Active, age-passed, not paused/restricted; server computes eligibility | Cursor; results are ephemeral candidates and revalidated on action |
| POST /v1/likes | Like/comment on profile or prompt | Actor eligible; target eligible; no block/exclusion | Idempotency-Key; returns pending like or match result |
| DELETE /v1/likes/{id} | Withdraw outgoing unmatched like | Owner; only withdrawable states | Idempotent; conflict if already matched or otherwise terminal |
| GET /v1/likes/incoming | Read eligible incoming queue | Owner; eligibility re-check; free/premium projection policy | Cursor; never expose blurred-queue contents via metadata |
| POST /v1/passes | Pass candidate | Actor eligible; target validation | Idempotency-Key; reversibility policy OPEN |
| GET /v1/matches | List current matches | Participant only; no deleted/blocked content leaks | Cursor; current authoritative state |
| POST /v1/conversations/{id}/messages | Send message | Participant, active match, no block, restrictions permit | Client message ID/idempotency key; atomic write + outbox |
| GET /v1/conversations/{id}/messages | Read message page | Participant; visibility/safety rules | Cursor; no cache bypass of block/unmatch |
| POST /v1/users/{id}/block | Block account | Authenticated actor; target validation | Idempotent; block state commits before success |
| POST /v1/matches/{id}/unmatch | End match | Participant only | Idempotent; creates permanent pair exclusion |
| POST /v1/reports | Submit report | Authenticated or supported limited reporting path | Idempotent; return safe receipt, not investigation detail |
| POST /v1/me/pause | Pause discovery | Owner; policy for likes/outgoing likes OPEN | Idempotent; existing conversations follow published policy |
| POST /v1/me/resume | Resume discovery | Owner; revalidate all eligibility | Idempotent; stale likes are not blindly reactivated |
| POST /v1/me/deletion-requests | Request account deletion | Owner and step-up auth if risk requires | Idempotent; returns clear state and next steps |
| GET /v1/me/deletion-status | View deletion progress | Owner until access cutoff; safe support route after | No; disclose status without exposing internal security details |
| GET /v1/me/entitlements | Read entitlement projection | Owner | No; server-derived |
| POST /v1/billing/store-notifications/{provider} | Process store event | Signature/provider verification, replay defense | Provider event/transaction deduplication |
| POST /v1/appeals | Appeal eligible restriction/age result | Owner; validate allowed appeal type | Idempotency key; restricted case record, safe receipt |

## 7. Error contract and authorization

Use stable machine-readable codes plus safe user messages. Proposed codes:
- `AUTH_REQUIRED`, `ACCOUNT_NOT_ELIGIBLE`, `AGE_CHECK_REQUIRED`, `AGE_CHECK_PENDING`
- `PROFILE_UNAVAILABLE`, `PREFERENCE_CONFLICT`, `BLOCKED_OR_UNAVAILABLE` (avoid confirming a hidden block to the other party)
- `PAIR_EXCLUDED`, `MATCH_NOT_ACTIVE`, `CONVERSATION_CLOSED`, `MESSAGE_NOT_ALLOWED`
- `IDEMPOTENCY_CONFLICT`, `VERSION_CONFLICT`, `RATE_LIMITED`, `VALIDATION_ERROR`
- `ENTITLEMENT_REQUIRED`, `DELETION_IN_PROGRESS`, `TEMPORARY_FAILURE`

Return consistent error envelope with code, safe message, correlation ID and optional field errors. Do not leak whether an account exists, exact age-check failure reason, block status, moderation signals, safety case status or vendor fraud flags. Use 401 for unauthenticated, 403/404 according to a consistent non-disclosure policy, 409 for state/idempotency conflicts, 422 for valid JSON with invalid domain state, 429 for limits, and 5xx for unexpected/transient failures. Final status mapping requires API security review.

## 8. Event catalogue (initial)

| Event | Produced when | Consumer examples | Safety/consistency rule |
|---|---|---|---|
| account.registered | Account row committed | Onboarding orchestration | Minimal payload; no PII beyond identifiers required |
| age_assurance.updated | Verified provider result accepted | Eligibility projection | Verify signature, workflow binding, event uniqueness and current account state |
| profile.published / profile.hidden | Visibility/moderation changes | Discovery projection, caches | Projection may lag; action paths re-check source of truth |
| like.created / like.withdrawn | Like state commits | Incoming queue, notification worker | Re-check block/eligibility before display or notification |
| match.created / match.ended | Match transition commits | Conversation projection, notifications | Unique match invariant; no duplicate fan-out |
| block.created | Block commits | All read/send/discovery/notification paths | Priority invalidation; each operation still checks canonical block state |
| message.created | Message commits | Push delivery, moderation pipeline | Message content only to authorized consumers; notifications avoid leaking content |
| report.created / case.updated | Report/case changes | Safety queue | Restricted channel and audit access |
| deletion.requested / deletion.target.completed | Deletion workflow progresses | Deletion workers, support status | Never restore data on retry; lawful exceptions explicit |
| entitlement.updated | Verified billing event processed | Entitlement projection | Deduplicate provider events; purchase never overrides safety restrictions |

## 9. Operational requirements

- Metrics: transaction conflict rate, match duplicate prevention, idempotency replays/conflicts, outbox lag, queue retries/dead letters, block enforcement latency, deletion completion/error by target, age callback delay, database saturation and restore tests.
- Alerts: block enforcement failure, unexpected message after block/unmatch, duplicate match/conversation, age-gate bypass, deletion backlog, outbox backlog, repeated vendor callback failures, elevated 5xx and spend anomaly.
- Logs use correlation IDs and allowlisted structured fields; never log access tokens, identity images, full message bodies or unrestricted sensitive profile fields.
- Define and test RPO/RTO, backup encryption/access, restore procedure, migration rollback/forward-fix, secrets rotation and incident response before launch.
- Production access least privilege, MFA for staff, auditable privileged actions and separate safety-case permissions.
- Rate-limit account creation, likes, messages, reports, age sessions, uploads and appeals; calibrate with abuse testing, not arbitrary permanent values.

## 10. Traceability to existing requirements

| Requirement | Schema/API enforcement | Acceptance evidence |
|---|---|---|
| Hard eligibility precedes ranking | Preference tables + central eligibility service; discovery and mutation checks | LIFE-04 / discovery tests: hard age, radius, gender and strict fields cannot be overridden |
| Age gate cannot fail open | `account_age_assurance`; registration/discovery checks | Pending, expired, callback replay and provider timeout tests |
| Like is not a match | `likes`, `matches`, `conversations`; POST /likes | One-sided like cannot create conversation or permit messages |
| Reciprocal match exactly once | Canonical pair + unique constraint + transaction/outbox | Concurrent reciprocal requests create one match and one conversation |
| Block immediately enforced | `blocks`; every read/send/worker check | Block races with send, push, cache and queued notification |
| Unmatch ends contact | Match/conversation state + `pair_exclusions` | No messages; 30-day read-only history; permanent rediscovery exclusion |
| Deletion is orchestrated | `account_deletion_jobs` and target tasks | Retries, partial failure, vendor deletion, backups and stale webhook tests |
| Premium never buys safety | `entitlements` only gates approved convenience features | Entitlement loss does not disable blocking, reporting or core consent |
| Audit and events are reliable | Outbox/inbox and security audit | Crash between commit/publish; consumer replay; audit access review |
| Privacy by design | Classification map, DTO allowlists, scoped retention | API contract tests prevent sensitive field leakage |

## 11. Open decisions and launch gates

Do not silently turn these into final policy:
1. **Pause semantics:** proposed default hides profile from new discovery/new likes, does not unmatch; outgoing likes freeze/re-evaluate on resume. Confirm product copy and state behavior.
2. **Pass reversibility:** proposed initial default is non-reversible. Validate against user research.
3. **Like expiry:** proposed no silent expiry; re-check eligibility on display/action. Decide whether operational expiry is needed.
4. **Message and safety evidence retention:** set per data class with legal basis, purpose, access, expiry and deletion exceptions; no invented universal duration.
5. **Account deletion grace/access cutoff:** choose security policy and user explanation; define what happens to matches and pair suppression.
6. **Pair suppression after deletion:** legal/privacy decision required before retaining identifiers.
7. **Age-assurance data model:** inspect Didit contract/workflow, data flow, storage/regions, callbacks, appeals and deletion evidence; do not duplicate identity evidence without need.
8. **Match history storage:** choose append-only ended match vs reusable pair row while preserving the 30-day read-only history behavior and permanent pair exclusion.
9. **API authentication, authorization, RLS and staff access:** security architecture review and threat model.
10. **Rate limits, pagination limits, idempotency TTL, error mapping and RPO/RTO:** derive from abuse cases, load tests and operations capacity.
11. **PostGIS, partitioning and cache/search technology:** only adopt after measured query/volume requirements.
12. **Go/PostgreSQL proof-of-concept:** concurrency tests, migration/restore test, deletion workflow test, signed URL/media processing, outbox reliability and cost estimate.
13. **Retention map and this schema must agree** before production migrations. Run privacy review and data-flow inventory against every column and event payload.
14. **No mobile UI design yet:** this is backend contract groundwork only. Research and product requirements remain in progress.

## 12. Next implementation-free deliverable

Before writing migrations, produce a matching-engine specification with hard eligibility pseudocode, soft ranking factors, unknown/missing-value semantics, deterministic tie-breakers, bounded exploration, explainability rules and synthetic fairness/edge-case tests. Then validate API and schema choices with a narrow Go/PostgreSQL proof-of-concept. Keep UI work paused until research, legal/privacy review, vendor due diligence and core technical acceptance gates are complete.
