# Pinchd — Full-Fold Lifecycle, Privacy and Architecture Reconciliation
**Date:** 22 October 2026  
**Status:** Consolidated recommendation and decision register; not legal sign-off, a final retention schedule, production architecture approval or UI specification.  
**Purpose:** Reconcile lifecycle edge cases, data classification/retention, API/schema implications and the approved technical direction into one implementation-ready decision baseline, while keeping unresolved legal/vendor/empirical decisions explicit.  
**Product positioning:** “Serious about compatibility, flexible about relationship goals.”  
**Design status:** Mobile UI remains paused.

## 1. Executive decision summary
This document resolves contradictions and proposes conservative defaults for engineering planning. It does not convert a recommendation into an approved product policy where the existing decision register explicitly marks an open gate.

### Reconciled decisions
1. **Architecture direction:** ADR-0001 is authoritative: Go backend; React Native + Expo + TypeScript; AWS; PostgreSQL; private S3 media; managed async work when justified; role-restricted admin; one primary AWS region; Multi-AZ managed PostgreSQL in production where feasible. No microservices or multi-region database by default.
2. **State model:** Account lifecycle, age eligibility, profile visibility, enforcement restrictions and subscription entitlement are separate dimensions. Effective permissions are derived server-side.
3. **Eligibility:** Age and other mandatory gates, visibility/enforcement, blocks/pair suppression and mutual hard preferences are evaluated before soft ranking. No ranking, Premium entitlement, exploration or low supply can override a hard gate.
4. **Pause:** Recommended default is to hide from new discovery and prevent new likes while paused; do not unmatch or close existing conversations solely because of pause. Freeze existing outgoing likes from presentation and re-evaluate on resume. Product confirmation remains needed before public wording/implementation.
5. **Pass:** Recommended default is non-reversible at launch. A pass is not a block and does not create a permanent safety exclusion.
6. **Outgoing likes:** Recommended default is no silent expiry. Re-evaluate on display and action; eligibility or safety changes can invalidate a like regardless of its age.
7. **Matches and messaging:** Match creation is idempotent and requires valid reciprocal interest plus current eligibility. Block/unmatch immediately closes prohibited contact server-side and suppresses queued side effects.
8. **Unmatch history:** Approved baseline: messaging closes immediately; read-only past-match history remains visible for 30 days, with report/block available and no messaging. This is a user-visible display window, not a backend evidence-retention period.
9. **Pair suppression:** Approved baseline is permanent exclusion of the pair after unmatch. Minimum post-deletion suppression data, lawful basis, retention and user disclosure are still a legal/privacy gate; do not implement indefinite retention without review.
10. **Age assurance:** Unknown/pending/failed/expired states never grant age-gated access. Provide accessible retry/support/appeal; vendor callback must be authenticated, replay-safe and auditable.
11. **Deletion:** Account deletion is an orchestrated, idempotent workflow across primary data, media, indexes/caches, analytics, notifications, vendors and backup restore. Retention exceptions must be narrow, documented and reviewed.
12. **Premium:** Server-authoritative entitlement based on verified store state; retries and notifications idempotent. Safety, consent, core matching/messaging, blocking/reporting, age assurance required for eligibility and deletion remain free.
13. **Data governance:** Every field and event needs purpose, sensitivity class, access owner, retention/deletion behavior and vendor/data-flow traceability. No global retention period is invented here.
14. **Operations:** One outbox-based event pattern for durable asynchronous effects; workers re-check current authoritative state immediately before sending or exposing data.
15. **Launch gate:** Technical correctness, legal/privacy approval, age assurance, staffing, accessibility, vendor due diligence and tested recovery remain separate required gates.

## 2. Authority and conflict resolution
Use this hierarchy when project documents conflict:
1. Explicit user-approved product policy in the reconciled Workstream 1 baseline.
2. ADR-0001, the approved technical direction.
3. This reconciliation, for proposed lifecycle defaults and cross-document consistency, subject to the status labels below.
4. Detailed requirements and supporting research documents.
5. ADR-0002's older provisional recommendations, except where its 22 October update explicitly says ADR-0001 supersedes them.

**Important:** “Recommended default” means the project can use it to plan schema/tests, but it is not a final user-facing promise or a substitute for privacy/legal approval. “Approved baseline” refers only to policies already approved. “Open gate” blocks the relevant production behavior until the named owner resolves it.

## 3. Decision status vocabulary
- **APPROVED:** explicit product/architecture direction already agreed.
- **RECOMMENDED DEFAULT:** conservative proposed behavior; confirm before public launch.
- **OPEN — LEGAL/PRIVACY:** requires documented advice/assessment and, where relevant, user-facing disclosure.
- **OPEN — VENDOR:** requires contract/configuration/evidence from provider.
- **OPEN — TECHNICAL PoC:** must be demonstrated in an isolated test.
- **OPEN — PRODUCT RESEARCH:** needs usability or market evidence.
- **DEFERRED:** intentionally not needed for initial release unless evidence changes.
- **RELEASE BLOCKER:** must be complete before the affected feature or public launch.

## 4. Canonical state model
Do not put every state into one users.status enum. Model independent state dimensions and derive capabilities from them.

### 4.1 Account lifecycle
- registration_in_progress
- age_check_pending
- active
- deletion_pending
- deleted_or_deidentified

Paused is preferably a profile-visibility state, not an account lifecycle state. Restricted and suspended are enforcement dimensions rather than replacements for account identity/lifecycle. Avoid a single enum that makes combinations impossible to represent.

### 4.2 Independent dimensions
- **Account lifecycle:** whether normal account data exists and whether deletion is underway/completed.
- **Age eligibility:** unknown, pending, eligible, failed, expired, appeal_pending; exact vendor result mapping requires review.
- **Profile visibility:** draft, visible, paused, moderation_hidden, deletion_hidden.
- **Enforcement:** zero or more capability restrictions, each with scope, reason code, effective time, review/expiry, actor and appeal state. Do not expose internal risk signals.
- **Subscription entitlement:** derived from verified platform state, including active, billing retry/grace where applicable, cancelled-but-current, expired, refunded/revoked and unknown reconciliation state.
- **Relationship state:** outgoing like, incoming-like projection, match, active conversation, unmatch, block and pair suppression are separate records/states.
- **Media moderation:** pending, approved, rejected, removed, quarantined; only approved/allowed media can be publicly served.
- **Deletion workflow:** pending, processing, waiting on vendor, exception review, complete; state must remain truthful.

### 4.3 Permission derivation
A single server-side policy function should decide each capability: sign in, edit profile, appear in discovery, receive new likes, create likes, create match, send message, receive message, upload media, use Premium convenience features, report, block, appeal, request deletion and access a past-match record.

No client flag, cached feed, stale JWT claim, store callback or queue message is sufficient authority for safety, eligibility, blocking, deletion or entitlement.
## 5. Lifecycle decisions — reconciled matrix

| Decision | Status | Reconciled behavior | Required proof / owner |
|---|---|---|---|
| Minimum age and age assurance | Approved gate; legal/vendor open | No discovery or restricted use until the required check is eligible; pending/unknown/failure/expiry is not a pass | Legal/privacy + Didit owner: exact workflow, evidence, appeal, accuracy, data flows and deletion |
| Optional profile verification | Distinct from age assurance | Never use optional profile verification as a substitute for required age assurance or as a romantic desirability score | Product/security: define badge meaning and misuse controls |
| Pause visibility | Recommended default | Hide from new discovery and prevent new likes; preserve existing matches/conversations unless separately blocked, unmatched or restricted | Product owner: confirm public wording; test all feeds and queues |
| Existing outgoing likes during pause | Recommended default | Freeze presentation while paused; do not silently withdraw; re-evaluate eligibility on resume and before action | Product + engineering: tests for preference changes and stale likes |
| Existing matches during pause | Recommended default | Pause alone does not close matches or messaging | Product/T&S: confirm user expectation and abuse edge cases |
| Pass reversibility | Recommended default | Non-reversible in launch version; pass is not a block and does not create permanent pair exclusion | Product research; document accidental-pass recovery policy |
| Outgoing-like expiry | Recommended default | No silent time-based expiry; re-evaluate current eligibility on every display/action | Product/engineering: monitor stale queues and provide later expiry only with published policy |
| Incoming-like free queue | Approved baseline | Five eligible incoming likes visible at a time in rolling queue; like-back/pass advances; remaining eligible likes blurred while queued; Premium sees full eligible queue | Engineering: presentation entitlement, not data deletion; filter current eligibility on read/action |
| Commented likes | Approved baseline | Optional user-written comment on a specific photo/prompt; only modest transparent ordering advantage if approved policy permits; no length/AI/payment advantage | Product/T&S: abuse and fairness monitoring; preserve reference integrity if content is removed |
| Match creation | Approved baseline | Only valid reciprocal interest while both accounts remain eligible; exactly one active match/conversation per pair | Engineering: transaction/unique constraint/concurrency tests |
| Messaging permission | Approved baseline | Like alone is not permission; messaging requires valid active match and current policy eligibility | Engineering: authorize on send and delivery; race tests |
| Unmatch | Approved baseline | Immediately close messaging and remove active lists; do not reveal initiator; 30-day read-only past-match display, report/block available, no messaging | Product/privacy: ensure visible-history window and backend retention are distinct |
| Pair exclusion after unmatch | Approved baseline | The pair must not be matched again, separately from block | Privacy/legal: decide minimum suppression data and post-deletion behavior |
| Block | Approved safety behavior | Immediate server-authoritative exclusion across discovery, likes, matches/messages as specified, notifications, caches, queues and future account interaction | Engineering/T&S: end-to-end and race tests |
| Account deletion | Approved required direction; details open | Stop new activity immediately and start durable deletion orchestration; no stale client/webhook can reactivate account | Privacy/legal + engineering: grace period, exceptions, vendor and backup handling |
| Suspension/restriction | Approved safety principle | Capability-scoped restriction may disable messaging; full suspension disables normal use; preserve reporting/appeal paths as policy permits | T&S/legal: enforcement matrix, notice, appeal and review |
| Premium expiry/refund | Recommended default | Reconcile store state; apply current free entitlement on authoritative refresh; never remove safety or consent controls | Billing owner: sandbox tests and legal/store review |
| Like/comment target removed | Open policy, recommended handling | Do not fabricate a reference; show current profile or a neutral “content no longer available” state; retain minimum evidence only if lawful and needed | Product/privacy/T&S: define report/appeal evidence rules |
| Profile becomes ineligible | Approved invariant | Remove from new discovery and reject new like/match actions; existing matches follow restriction policy, not automatic silent deletion | Engineering/T&S: all state transitions and queues tested |

## 6. Relationship and messaging state rules

### 6.1 Pair identity and uniqueness
- Represent a pair using stable internal account IDs sorted into canonical order (low_account_id, high_account_id) for pair-level suppression, active match uniqueness and unmatch history.
- Reject self-pairs.
- Maintain block as directional actor-to-target records (both directions are possible); effective blocking is true if either direction is active.
- Pair suppression after unmatch is pair-level and independent of directional blocks.
- A deleted account's pair-suppression record is only retained if the privacy/legal gate approves the minimal design, retention and disclosure.

### 6.2 Reciprocal-like transaction
1. Authenticate the actor and derive account ID server-side.
2. Confirm account lifecycle, age eligibility, enforcement, visibility, target eligibility and no block/pair suppression.
3. Confirm reciprocal hard-preference compatibility and that the target can receive the action.
4. Insert or idempotently update the actor's like under a unique actor/target constraint.
5. In the same transaction, detect a valid reciprocal like and create one active match plus one conversation under unique constraints.
6. Write domain events to the transactional outbox before commit.
7. Commit; publish events asynchronously.
8. Before any push or downstream side effect, re-check current block, match, account, deletion and notification state.
9. If the transaction outcome is uncertain to the client, retry using the same idempotency key; never create duplicates.

A match event is not permission to ignore later block/unmatch/delete state. Event consumers must verify current authoritative state.

### 6.3 Message send and delivery
- Authorize sender and recipient against the current active match/conversation, account lifecycle, restrictions, block state and unmatch state.
- Store message and outbox event in one transaction.
- Delivery workers re-check that the conversation remains active and contact remains permitted before delivery/notification.
- After block/unmatch, no new message is accepted and queued content/notifications must not be delivered if the current policy prohibits it.
- Keep notification payloads minimal; avoid message text/profile details on lock screens by default.
- Define edit/delete behavior for sent messages separately; do not silently rewrite conversation evidence.
- Attachments must be private, scanned/moderated as required, metadata-stripped and access-controlled.
## 7. Privacy, retention and deletion reconciliation

### 7.1 Retention model
No universal retention period is approved. Assign a rule per data class and purpose:
- **C0 public profile projection:** visible only while the profile and each field are eligible/approved for display; invalidate caches on hide/block/delete.
- **C1 account/contact data:** purpose-limited; retention linked to account lifecycle, security and legal obligations.
- **C2 sensitive/high-impact data:** identity/age-assurance results, location/preferences, relationship goals and other sensitive fields; strict purpose and access controls, minimal collection and review of legal basis.
- **C3 restricted safety/identity evidence:** reports, case evidence, verification artifacts or sensitive media; role-restricted, audited, purpose-bound and retained only under an approved schedule/legal hold.
- **C4 security/accountability records:** authentication, access, billing reconciliation and audit events; minimise payloads and separate from product analytics.

Every retained class needs: purpose, lawful basis, owner, storage location, processor/subprocessor, access roles, retention trigger, deletion method, backup behavior, legal exception, user-facing disclosure and verification evidence.

### 7.2 30-day unmatch history is not 30-day message retention
The approved 30-day period defines how long the user-facing read-only past-match history remains visible. It does not automatically mean that all message bodies, attachments, safety reports, audit logs, backups or vendor copies are erased on day 30. Those require separate retention rules and clear notice. Conversely, the 30-day display rule does not justify retaining all underlying content indefinitely.

### 7.3 Deletion workflow
Recommended orchestration:
1. Authenticate and validate the request; record request ID, account, timestamp and policy version.
2. Immediately hide profile from discovery, prevent new likes/matches and revoke active sessions/refresh tokens as specified.
3. Set a durable deletion tombstone/state so stale clients, retries and webhooks cannot reactivate the account.
4. Create idempotent deletion tasks for profile projection, account/contact data, likes, active matches/conversations, messages/attachments according to the approved schedule, media objects/thumbnails, search indexes, caches, push tokens, analytics mappings and vendor data.
5. Remove from candidate pools and invalidate precomputed feeds; enforce tombstone checks after backup restore.
6. Send vendor deletion/return requests and track evidence and failure/retry states.
7. Retain only explicitly approved legal/safety/security exceptions with restricted access, reason, reviewer, expiry/review date and lawful basis.
8. On partial failure, keep deletion pending and surface truthful internal status; retry without recreating deleted records.
9. Complete only when required tasks are confirmed or an authorised, documented exception is approved.
10. Verify the result with automated tests and periodic deletion audits.

Do not promise instant complete erasure where documented lawful exceptions or backup cycles apply. Do not use exceptions as a blanket reason to retain everything.

### 7.4 Pair suppression after account deletion — release blocker
Product policy requires permanent pair exclusion after unmatch, but the minimum data needed after deletion is unresolved. Before implementation, privacy/legal must decide:
- whether a minimal pseudonymous pair token is lawful and necessary;
- whether it can be retained after account deletion and for how long;
- whether hashing/pseudonymisation actually reduces risk in this context;
- what the privacy notice says and how users exercise rights;
- whether pair suppression can be enforced without retaining a re-identifiable account mapping;
- how restoration, future account creation and appeal work.

**Do not implement indefinite post-deletion suppression storage until this gate is signed off.** If the gate is not resolved, the product must not claim that the permanent-exclusion guarantee survives account deletion.

### 7.5 Location and preferences
- Store the minimum location representation needed for discovery; approximate public location and prevent triangulation.
- Avoid location history/trails unless separately justified and reviewed.
- Keep hard/soft preferences private; never expose another user's dealbreakers or use them in ordinary analytics.
- Sensitive attributes/inferences need documented purpose and legal/privacy review before collection or matching use.
- Missing/undisclosed data never becomes consent, positive compatibility or open-minded intent.
- Invalidate cached recommendations immediately when relevant preferences, blocks, visibility or eligibility change.
## 8. Architecture reconciliation and target boundaries

### 8.1 Authoritative stack
**ADR-0001 is the source of truth.** ADR-0002 remains useful as a list of domain boundaries, risks and PoC gates, but its older TypeScript backend recommendation is superseded.

- Client: React Native + Expo + TypeScript.
- Backend: Go modular monolith with explicit domain packages and typed interfaces.
- Cloud: AWS, one primary region initially.
- Transactional store: PostgreSQL; use PostGIS only if measured geo-query requirements justify it.
- Media: private S3 with short-lived scoped access, upload validation, metadata stripping and moderation lifecycle.
- Async processing: managed queue/workers when justified, backed by a transactional outbox and idempotent consumers.
- Admin: separate role-restricted console, MFA, least privilege and audit trails.
- Operations: CI/CD, automated tests, infrastructure as code, secrets management, dependency/secret scanning, redacted observability, budgets, backups, restore drills and incident response.
- Production database: managed PostgreSQL with Multi-AZ resilience where feasible; define RPO/RTO and prove recovery before claiming objectives.
- Deferred: microservice fleet, multi-region database, Redis/search/ML infrastructure unless measured needs justify them.

### 8.2 Domain boundaries inside the modular monolith
1. Identity and account lifecycle.
2. Age assurance and eligibility.
3. Profile and media.
4. Preferences and discovery policy.
5. Likes, pair state and matches.
6. Conversations, messages and notifications.
7. Trust, safety, moderation and appeals.
8. Billing and Premium entitlement.
9. Privacy rights and data lifecycle.
10. Analytics/experimentation.
11. Admin/support and audit.

These are logical ownership boundaries, not a requirement to deploy eleven services. Controllers handle transport, authentication and validation; domain services enforce policy; repositories own persistence; workers own durable asynchronous tasks.

### 8.3 Core tables/entities (conceptual, not migration DDL)
- accounts: stable internal ID, lifecycle state, created/updated/deletion timestamps; avoid putting every state dimension here.
- account_credentials / sessions: authentication and revocation data with separate security controls.
- age_assurance_cases: vendor, workflow/version, external reference, state, timestamps, appeal state and minimal evidence pointer; no unnecessary identity images or biometric payloads in the app database.
- profiles: owner, publish/visibility state and version.
- profile_fields or typed profile tables: field values with explicit field allowlist and classification; prefer typed columns for frequently queried policy fields.
- profile_media: owner, private object key, media type, moderation state, metadata processing state, created/removed timestamps.
- discovery_preferences: versioned owner preferences; identity separate from discovery selections; explicit strictness per field where allowed.
- account_enforcement: capability scope, state, start/end/review, reason code and case reference; restricted visibility.
- blocks: actor, target, created/revoked state; unique active actor-target pair.
- pair_suppressions: canonical pair IDs, reason (unmatched or other approved reason), timestamps and policy version; post-deletion fields blocked pending legal/privacy decision.
- likes: actor, target, optional prompt/media reference, optional user-written comment, state and timestamps; unique active actor-target rule.
- matches: canonical pair, state, created/unmatched timestamps, unique active pair.
- conversations: match reference and state; one per match.
- messages / message_attachments: conversation, sender, payload reference, timestamps, moderation/deletion state and classification; retention unresolved by data class.
- incoming_like_queue should preferably be a query/projection, not a source of truth; any materialised projection must be rebuildable and re-filtered.
- reports / safety_cases / evidence_references: restricted case access, state, owner, severity, access audit and retention trigger.
- subscription_events / entitlements: provider transaction identifiers, event dedupe keys, verified state and entitlement history.
- deletion_requests / deletion_tasks / deletion_tombstones: durable orchestration, retries, exceptions and completion evidence.
- outbox_events / inbox_deduplication: durable event publishing and consumer idempotency.
- audit_events: security/administrative action references, not general analytics.
- analytics_events: minimised event schema, separated from direct identifiers and safety evidence.

Physical names, types and indexes require a schema review against the current blueprint and implementation framework before migrations are written.
## 9. API contract rules

### 9.1 Contract invariants
- Version APIs and use explicit request/response DTOs; never serialize ORM/database records wholesale.
- Derive actor identity from the authenticated session/token, never a request-body user ID.
- Enforce object-level authorization on every read and write.
- Use idempotency keys for likes, match-triggering actions, message sends, deletion requests and provider callbacks where retries can duplicate side effects.
- Return stable error codes and correlation IDs; never disclose private preferences, reporter identity, internal risk signals or another user's block state.
- Re-check current policy on every sensitive action; a previously returned feed item is not proof of continued eligibility.
- Pagination cursors must be opaque, scoped to the viewer and policy/version context, and safe against enumeration.
- Rate-limit authentication, recovery, likes, messages, uploads, reports and other abuse-sensitive routes.
- Staff/admin APIs are a separate audience with MFA, least privilege, explicit case access and audited actions.

### 9.2 Initial route inventory
| Method / route | Purpose | Key authorization and invariants |
|---|---|---|
| POST /v1/registrations | Begin account registration | Rate-limited; idempotent; no discovery access |
| POST /v1/age-assurance/sessions | Start required check | Bound to authenticated account/workflow version; vendor gate |
| GET /v1/age-assurance/status | Read status | Own account only; no vendor fraud details |
| POST /v1/age-assurance/appeals | Appeal result | Own account; restricted review queue |
| GET/PATCH /v1/me | Read/update account/profile fields | Explicit allowlist; version/conflict handling |
| POST /v1/profile/publish | Publish profile | Required fields/media/moderation/age and policy gates pass |
| POST /v1/media/upload-sessions | Get scoped upload permission | Owner-bound, short-lived; size/type constraints |
| POST /v1/media/{id}/complete | Finalise media processing | Ownership, scan/moderation state and metadata processing |
| GET /v1/discovery | Retrieve candidates | Eligibility first; opaque cursor; current block/suppression check |
| POST /v1/likes | Like target/profile element | Idempotent; mutual hard-gate re-check; optional comment validation |
| DELETE /v1/likes/{id} | Withdraw outgoing like | Own outgoing like only; only before match as approved |
| POST /v1/passes | Pass a candidate | Idempotent; not a block; non-reversible default pending confirmation |
| GET /v1/incoming-likes | Queue of eligible incoming likes | Current entitlement and eligibility; five unblurred free items; never expose stale suppressed item |
| GET /v1/matches | Active and permitted past matches | Own relationship data; active vs read-only past history separated |
| POST /v1/matches/{id}/unmatch | Unmatch | Immediate state change; idempotent; pair exclusion event |
| POST /v1/blocks / DELETE /v1/blocks/{id} | Block/unblock | Actor authorization; effective block if either direction active; unblock does not remove pair suppression |
| GET /v1/conversations/{id}/messages | Read messages | Active or approved read-only visibility state; current access check |
| POST /v1/conversations/{id}/messages | Send message | Active match, no block/unmatch/restriction; idempotency key |
| POST /v1/reports | Report user/content/message | Reporter identity protected; evidence reference controlled |
| POST /v1/account/pause / POST /v1/account/resume | Toggle visibility | Policy-consistent; no destructive relationship transition |
| POST /v1/account/deletion-requests | Request deletion | Re-auth/risk control; durable idempotent workflow |
| GET /v1/account/deletion-status | Read deletion progress | Own account; truthful status and exception wording |
| POST /v1/billing/provider-events/{provider} | Receive verified store event | Signature/receipt validation, replay-safe, provider-specific contract |
| POST /v1/billing/restore | Reconcile purchases | Authenticated account; provider verification; no client-only entitlement |
| POST /v1/admin/cases/{id}/actions | Staff case action | Separate admin auth, role/case scope, reason, audit and review rules |

This is an initial route inventory, not final OpenAPI. Request/response schemas, exact status codes, pagination shape, authentication provider, and versioning/deprecation policy remain implementation choices.

## 10. Event and transaction model

### 10.1 Transactional outbox
Persist domain state and its outbox event in the same PostgreSQL transaction. A publisher retries until acknowledged or moved to an observable dead-letter/review state. Consumers must be idempotent and record event IDs in an inbox/deduplication table. Never rely on a best-effort in-process callback for a safety-critical side effect.

### 10.2 Proposed event names
- account.registration_started
- age_assurance.status_changed
- profile.published
- profile.visibility_changed
- profile.eligibility_changed
- media.moderation_changed
- like.created
- like.withdrawn
- match.created
- match.unmatched
- block.created
- block.revoked
- message.created
- message.delivery_suppressed
- report.created
- enforcement.changed
- subscription.entitlement_changed
- account.deletion_requested
- account.deletion_task_completed
- account.deletion_completed

Events should carry IDs, event version, occurred time, correlation ID and minimal references. Avoid full profile, message, precise location, age-assurance evidence or safety narrative in event payloads.

### 10.3 Current-state recheck
Before any notification, delivery, candidate exposure, media URL issuance or queued action, load/check authoritative account, block, match, enforcement, visibility and deletion state. A committed event signals that something happened; it does not grant permanent authorization.

### 10.4 Transaction boundaries
- Like + reciprocal match/conversation creation + outbox event: one transaction.
- Message creation + outbox event: one transaction.
- Block/unmatch + active-state closure + outbox event: one transaction; queued consumers recheck current state.
- Deletion request + tombstone/state + first tasks + outbox event: one transaction.
- Subscription provider event + deduplication + entitlement reconciliation + outbox event: one transaction after provider verification.
- Vendor API calls, object deletion, push delivery and external service operations: asynchronous idempotent tasks, not held inside long database transactions.
## 11. Security, reliability and operational controls
- Authentication identity is server-derived; sessions are revocable; privileged staff use MFA.
- Object-level authorization is tested for every API resource.
- Private media only; scoped short-lived access; check current state before issuing URLs and serving objects.
- No secrets, identity images, biometric payloads, raw message bodies or precise coordinates in routine logs/analytics.
- Encrypt in transit and at rest as appropriate; use managed secrets, key rotation and least privilege.
- Audit privileged evidence access and exports with case ID, purpose and actor; restrict bulk export.
- Separate product analytics from safety case evidence and security audit.
- Test rate limits, abuse detection and account recovery without exposing hidden risk signals.
- Keep staging and production separate; do not copy production data into development by default.
- Versioned migrations require a tested forward-recovery/rollback strategy.
- Monitor outbox backlog, deletion backlog, block propagation failures, duplicate matches, age-gate anomalies, provider outages, message delivery suppression and abnormal spend; assign named owners.
- Define RPO/RTO per critical system before production; test backup restore and deletion-tombstone replay. Do not advertise unproven SLOs.
- Use one primary AWS region and managed PostgreSQL Multi-AZ where feasible; add Redis, search, microservices or multi-region only with measured justification.

## 12. Reconciled requirements and acceptance gates

### Must pass before controlled beta
- Hard-preference and age-gate tests show no normal-flow violations.
- Concurrent reciprocal likes produce one match and one conversation.
- Block/unmatch prevents prohibited new messages and suppresses queued notifications after authoritative state commits.
- Profile/feed/media caches cannot expose a blocked, hidden, deleted or ineligible profile.
- Pending/unknown/failed/expired age assurance never grants gated access.
- Account deletion cannot be undone by stale client, webhook, worker retry or backup restore.
- Private media cannot be fetched by another user or after access is revoked under the current policy.
- Premium is server-authoritative; safety remains free.
- Critical reporting, block, appeal, account recovery and deletion flows are accessible and tested.
- Incident ownership, moderation capacity, support routes and urgent-harm playbooks are real and exercised.
- Legal/privacy/vendor gates for the actual beta scope are signed off.

### Must pass before public launch
- Written legal scope and required Online Safety Act assessments/controls are complete as applicable.
- Privacy data-flow map, lawful-basis register, DPIA decision/report, processor/transfer review and retention schedule are approved.
- Didit workflow, data handling, errors/appeals, callbacks, retention and deletion are verified.
- Subscription terms, cancellation/refund, store rules and consumer-law applicability are reviewed for launch date.
- Accessibility manual testing with assistive technologies and disabled users is documented.
- Security testing, backup restore, incident exercise, budget alerts and operational coverage are evidenced.
- Apple/Google review requirements, in-app deletion and meaningful differentiation evidence are ready.
- The chosen launch market passes demand, eligible-supply, cost and operational gates.

### Release blockers
Any known hard-preference bypass, age-gate bypass, block bypass, prohibited messaging after unmatch/block, deleted-account resurrection, public private-media access, client-forged Premium entitlement, inaccessible critical safety path or unresolved high-risk legal/privacy gate blocks the affected release.
## 13. Decision register — owners and gates

| ID | Decision | Reconciled status | Owner / gate |
|---|---|---|---|
| REC-01 | Go + RN/Expo/TS + AWS + PostgreSQL + private S3 | APPROVED direction | ADR-0001; implementation gated by PoC |
| REC-02 | Modular monolith before service extraction | APPROVED direction | Engineering; revisit on measured need |
| REC-03 | Pause behavior and outgoing likes | RECOMMENDED DEFAULT | Product confirmation + user comprehension |
| REC-04 | Pass non-reversible | RECOMMENDED DEFAULT | Product research; before final API contract |
| REC-05 | No silent like expiry | RECOMMENDED DEFAULT | Product confirmation; monitor queue health |
| REC-06 | Five-like free queue | APPROVED hypothesis | Product analytics; fairness and comprehension review |
| REC-07 | Unmatch 30-day read-only display | APPROVED baseline | Privacy review must keep display vs storage separate |
| REC-08 | Permanent pair exclusion after unmatch | APPROVED baseline | Post-deletion retention remains legal/privacy blocker |
| REC-09 | Pair suppression after account deletion | OPEN — LEGAL/PRIVACY | Privacy/legal owner; block indefinite retention |
| REC-10 | Message/attachment and safety-evidence retention | OPEN — LEGAL/PRIVACY | Data-class-specific schedule and user notice |
| REC-11 | Age-assurance evidence, appeals and vendor deletion | OPEN — VENDOR + LEGAL/PRIVACY | Didit evidence pack and due diligence |
| REC-12 | Premium entitlement reconciliation | RECOMMENDED DEFAULT | Billing owner; store sandbox proof and legal review |
| REC-13 | RPO/RTO and backup retention | OPEN — TECHNICAL/OPERATIONS | Define and prove restore before production |
| REC-14 | API schema, errors and cursor contracts | OPEN — TECHNICAL | Freeze after lifecycle review and PoC |
| REC-15 | Exact AWS services/versions and cost | OPEN — TECHNICAL/COST | Regional quote, threat model, load/cost test |
| REC-16 | Formal retention periods and lawful bases | OPEN — LEGAL/PRIVACY | Must be approved before public launch |
| REC-17 | Pause/restriction interaction with existing matches | RECOMMENDED DEFAULT | Product/T&S review; end-to-end tests |
| REC-18 | Pass history and accidental-pass recovery | RECOMMENDED DEFAULT | User research; not a safety exclusion |
| REC-19 | Post-deletion re-registration and pair exclusion | OPEN — LEGAL/PRIVACY + PRODUCT | Align deletion transparency and abuse prevention |
| REC-20 | Admin access and evidence export | APPROVED security direction; details open | T&S/security; role tests and audit evidence |

## 14. Required follow-up artifacts
1. Update ADR-0002 status to **Superseded for stack selection; retained as supporting feasibility checklist**.
2. Use this document as the lifecycle/privacy/architecture reconciliation index, while preserving detailed source docs.
3. Update the PostgreSQL blueprint to reference state dimensions, pair suppression gate, deletion workflow and transaction rules.
4. Build a requirements-to-test traceability matrix for lifecycle requirements and POC-01 through POC-12.
5. Obtain privacy/legal decisions on retention and post-deletion pair suppression before freezing the data schema.
6. Obtain Didit workflow and data-flow evidence before implementing age assurance.
7. Run isolated technical PoCs before implementation/production approval.
8. Keep mobile UI design paused until research workstreams, legal/vendor gates and design-readiness review are complete.

## 15. Final conclusion
The reconciled baseline is a server-authoritative, consent-first, privacy-minimised Go/PostgreSQL system with explicit lifecycle dimensions, idempotent transactions and durable asynchronous processing. The most important unresolved blockers are legal retention and post-deletion pair suppression, age-assurance evidence and appeal behavior, exact lifecycle policy confirmation, verified billing/vendor behavior, operational capacity and measured technical recovery.

This document reconciles project direction and supplies conservative defaults. It does not claim unresolved decisions have legal approval or that the proposed system has been implemented or tested.
