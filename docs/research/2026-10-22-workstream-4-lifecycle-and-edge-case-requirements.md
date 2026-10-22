# Pinchd Workstream 4 — Product Lifecycle and Edge-Case Requirements
**Date:** 22 October 2026  
**Scope:** Account lifecycle, onboarding, profile visibility, discovery eligibility, likes, matches, messaging, pause, block, unmatch, deletion, appeals and failure states.  
**Status:** Requirements synthesis from approved Pinchd decisions and existing project research. This is not a final API contract or UI specification. Mobile UI design remains paused.

## 1. Purpose and requirements language
This workstream turns approved product policy into testable lifecycle behavior. **MUST** means required for the approved product direction or a safety/privacy invariant. **SHOULD** means recommended unless a documented reason says otherwise. **OPEN GATE** means legal, privacy, vendor, technical or empirical validation is still required. Product approval does not itself establish that an implementation is safe, lawful or technically feasible.

Principles: hard eligibility gates precede ranking; never silently relax a hard preference; server state is authoritative for safety and consent; retries must not duplicate side effects; user-visible status must match backend truth; and failure should default to preserving privacy and preventing unwanted contact.

## 2. Canonical account lifecycle
Suggested account states:
- **UNREGISTERED:** no account record.
- **REGISTRATION_IN_PROGRESS:** account creation or required age-assurance flow has begun but is incomplete.
- **AGE_CHECK_PENDING:** age assurance is not yet complete or a vendor result is pending.
- **ACTIVE:** account meets required eligibility and can use permitted features.
- **PAUSED:** owner has temporarily hidden their profile from new discovery; existing matches and conversations follow the published pause policy.
- **RESTRICTED:** one or more capabilities are limited for safety, integrity, compliance or investigation.
- **SUSPENDED:** account cannot use normal service while a time-limited or reviewable enforcement action applies.
- **DELETION_PENDING:** deletion requested; sign-in and discovery behavior are restricted according to a documented grace/security policy.
- **DELETED / DE-IDENTIFIED:** normal account use is disabled and personal data has been deleted or de-identified subject to lawful retention exceptions.

Do not overload a single status field with age-check, enforcement, subscription and visibility states. Model independent dimensions where needed (e.g., account status, age eligibility, profile visibility, enforcement restrictions, subscription status) and derive permissions centrally.

### Transition invariants
- An account cannot become discovery-eligible until minimum age and other mandatory eligibility checks pass.
- A pending, failed, expired or indeterminate age check must not be treated as a pass. Retry and appeal paths must be documented.
- Pausing is not a ban, unmatch or deletion. It must not silently destroy conversations.
- A restriction must be enforceable server-side across all entry points, including old sessions, queued jobs and APIs.
- Deletion requests must trigger an orchestrated process across primary data, media, search/indexes, caches, analytics, notifications, vendors and backups according to the documented retention schedule and lawful exceptions.
- A deleted account must not be silently reactivated by a stale client, webhook or retry.
- Subscription cancellation and account deletion are separate workflows; explain any required app-store subscription action without using it to obstruct deletion.

## 3. Registration and onboarding
### Required behavior
1. Explain the product purpose, minimum age, core community rules and essential privacy information before the user is asked to invest substantial effort.
2. Collect only fields needed for eligibility, profile creation, matching or an explicitly explained feature.
3. Record consent and notices with version/time evidence where appropriate; do not use one blanket consent for unrelated processing.
4. Complete the required age-assurance flow before allowing discovery or other restricted functionality. Do not confuse optional profile verification with legal age assurance.
5. Provide clear pending, success, failure and retry states. Vendor timeouts and callback delays must not create a temporary eligibility bypass.
6. If the check fails, show a neutral explanation and available appeal/support route without revealing fraud-detection signals that enable evasion.
7. Prevent duplicate account creation and repeated vendor billing or verification requests on retry.
8. Do not make optional AI, voice, profile enrichment or marketing permission a prerequisite for basic use.

### Edge cases
- Age provider unavailable: hold the account at the restricted/pending boundary, communicate status and retry policy; do not bypass.
- Provider callback duplicated or arrives out of order: idempotently apply the valid latest state and audit changes.
- User abandons onboarding: expire incomplete sessions and vendor tokens; apply retention/deletion rules.
- User changes legal name, date of birth or other eligibility field: invoke a risk-based re-check rather than silently overwriting authoritative assurance evidence.
- User disputes a result: provide a controlled appeal path; ensure an appeal does not automatically expose restricted service access.
- OPEN GATE: final minimum age, permitted assurance methods, vendor configuration, lawful basis, biometric handling, accuracy/error thresholds and deletion schedule require legal/privacy/vendor evidence.

## 4. Profile creation, editing and visibility
### Required behavior
- Required profile fields must be minimal and clearly explained. Optional fields must remain optional.
- Draft changes should not become public until the user saves/publishes them, except where the UI explicitly explains autosave.
- Validate media type, size and safety policy server-side; strip EXIF/location metadata from public media; store originals only if needed under a documented restricted process.
- Media processing/moderation failures must not publish unprocessed content by accident.
- Editing relationship goals, gender identity, discovery preferences, children/family plans or other matching attributes must recompute eligibility/ranking without revealing private preference settings to others.
- Profile completeness is not a desirability score and must not secretly determine romantic exposure.
- A moderation restriction on one photo/prompt should have a defined effect on the profile and appeal path; do not silently replace user content with misleading content.
- Public profile output must use a field allowlist, not serialize the entire private account object.
- No exact home/work location, live movement or unnecessary metadata is shown. Location is approximate and must not enable triangulation.

### Edge cases
- Upload succeeds but metadata/database save fails: use a recoverable upload state and cleanup orphaned media.
- User removes a photo currently referenced by a like/comment: preserve only the minimum evidence needed for an existing report/appeal; do not keep it publicly visible.
- User edits a profile after being liked: display the current profile by default, but make clear which content a comment refers to if that content is removed; never fabricate a quote or visual reference.
- Profile becomes ineligible after editing age/location/preferences: immediately exclude from new discovery and prevent new likes; existing matches follow the applicable restriction policy.
- User reports their own compromised account: provide a recovery and session-revocation path.
- OPEN GATE: moderation vendor, media processing, image retention, appeal SLAs and profile field visibility matrix.

## 5. Discovery and eligibility
### Required evaluation order
1. Account and age eligibility.
2. Active/visible profile state and enforcement restrictions.
3. Safety exclusions: either account blocked the other, pair permanently excluded after unmatch, or other authoritative suppression applies.
4. Mutual hard preference compatibility, including the user-defined age range, selected genders, hard distance radius and any field explicitly marked hard.
5. Candidate availability and anti-abuse limits.
6. Reciprocal soft ranking and controlled exploration among eligible candidates.
7. Explanation generation using only truthful, permitted reasons.

Hard constraints MUST never be overridden by ranking, premium status, scarcity, experiments or exploration. A small candidate pool is not permission to show an excluded person. A user may choose to change their own preference, but the system must show the change and likely effect before applying it.

### Required behavior
- A candidate is eligible only if the relevant criteria are mutually satisfied where the policy requires reciprocity.
- Missing or undisclosed data is not equivalent to “open-minded,” consent, a match, or a positive preference.
- Relationship intentions remain distinct categories; do not infer intention from messages or behavior.
- Soft preference changes affect ranking, not hard eligibility.
- The system should explain profile recommendations with accurate, non-sensitive reasons. Never expose another user's private dealbreakers or internal safety flags.
- Do not use a person's attractiveness, profile completeness, revenue value or willingness to pay as a hidden proxy for compatibility.
- Travel mode, if offered, must be user-controlled, clearly time-bounded/visible as appropriate, and protected against stalking. It must not disclose a person's real-time location.
- If eligible supply is low, explain that the pool is limited and offer voluntary preference adjustments; do not fabricate profiles or likes.
- Cache and precomputed recommendations must re-check current eligibility before display and before an action is accepted.

### Edge cases
- Candidate changes age, gender, relationship intention, location or a hard preference after being queued: re-evaluate at display and like time.
- User changes their own hard preference: recompute eligible set; explain that existing candidates may disappear.
- GPS permission denied or stale: provide an explicit fallback/manual region where policy allows; do not guess exact location.
- User enters travel mode then pauses or is restricted: remove or limit travel exposure according to authoritative state.
- A profile is deleted while in a recommendation cache: suppress it immediately on read/action and invalidate caches asynchronously.
- OPEN GATE: ranking weights, controlled-exploration policy, location precision, anti-abuse thresholds and fairness monitoring require PoC/pilot validation.

## 6. Likes, comments and incoming-like queue
### Required behavior
- A like signals interest only. It does not create a match, authorize a message, or override the recipient's preferences.
- The sender may withdraw an outgoing like before a match; withdrawal is idempotent.
- An optional comment must attach to a specific photo or prompt when applicable and be authored by the user. The UI must not imply the comment was AI-written if it was not.
- A comment must not bypass blocks, reporting restrictions, rate limits or content policy.
- A thoughtful, relevant commented like may receive the approved modest transparent ordering advantage. No advantage for comment length, payment or AI usage. Monitor abuse and fairness.
- Free users see five eligible incoming likes unblurred at a time in the rolling queue; like-back/pass advances the queue. Remaining eligible likes stay blurred while queued. Premium users can see the full eligible queue. Five is an approved launch hypothesis, not proven optimal.
- Queue order must not be manipulated with fake likes, fabricated activity or deceptive urgency.
- An incoming like is eligible to show only if both accounts remain eligible and no block, unmatch exclusion, deletion or enforcement state suppresses it.

### Edge cases
- Duplicate like requests: one logical like, no duplicate notifications.
- Like arrives after sender withdraws, blocks, pauses, deletes or is restricted: reject or suppress it based on current authoritative state.
- Recipient passes: define whether the pass is reversible; if not, record the exclusion consistently and do not later re-present the same like as new.
- Comment's target photo/prompt is removed: do not show broken/misleading content; preserve only minimum necessary reference for a valid report or appeal.
- Premium expires while queue is open: apply entitlement at next server-authoritative refresh; do not leak the remaining blurred queue or delete incoming likes.
- Five-slot queue is empty due to eligibility filtering: show a truthful empty state, not a fake count.
- OPEN GATE: define pass reversibility, like expiration policy, notifications, rate limits and the exact premium entitlement enforcement.

## 7. Match creation and lifecycle
### Required behavior
- A match is created only when two valid reciprocal interests exist and both accounts remain eligible at the transaction boundary.
- Match creation MUST be idempotent and transactionally safe. Concurrent reciprocal likes must create one match and one logical event.
- Match notifications are secondary side effects and may be retried without duplicating the match or notification beyond the documented deduplication policy.
- No match may be created if either user blocked the other, pair is permanently excluded after unmatch, account is deleted/suspended, or mandatory age/eligibility fails.
- Matching must not be a paid privilege; a premium entitlement cannot create consent or override hard eligibility.
- If an account becomes restricted after a match, the system must apply a defined conversation/access policy and not let a stale match cache bypass the restriction.

### Edge cases
- Reciprocal likes race concurrently: unique pair invariant plus transaction/retry handling.
- Duplicate webhook or job: idempotency key or natural unique constraint prevents duplicate match.
- Block and like-back occur concurrently: safety state wins; no new match should be created after a block is authoritative.
- One like is withdrawn while the other is submitted: transaction ordering must yield a single valid result and explain the final state consistently.
- Push delivery fails: match remains valid; client can refresh state.
- OPEN GATE: pair key canonicalisation, transactional design and notification deduplication are implementation decisions to prove in tests.

## 8. Messaging and conversation
### Required behavior
- Only matched, currently eligible users may send messages, subject to restrictions, blocks and pair state.
- Message authorization must be server-side on every send; delivery, push, retries, background workers and media links must re-check current authorization.
- No unsolicited messages before a match. Do not sell unsolicited messaging privileges.
- Safety reporting must be possible from relevant conversation context without requiring the user to reopen harmful content.
- Message state should distinguish pending, sent, delivered and failed only where the product can substantiate those states; never claim a read receipt if none exists.
- Use idempotency for send retries to prevent duplicate messages.
- Apply documented rate limits and abuse detection without requiring users to reveal more private data than necessary.
- User-written AI-assisted drafts, if offered, must be opt-in, editable and never auto-send or auto-publish.
- Retain message evidence only under a documented privacy, safety, legal and deletion policy.

### Edge cases
- Block during message send: once block is authoritative, reject new sends and prevent later queued delivery; race behavior must be explicitly tested.
- Message accepted by API but push worker runs after block/unmatch: worker re-checks authorization and suppresses delivery.
- Sender edits/deletes a message: define what the recipient sees and what evidence remains for a report; do not promise total erasure if lawful safety retention applies.
- Attachment upload completes after conversation is closed: reject attachment association and clean up orphaned media.
- User loses network: retry with same idempotency key; no duplicate send.
- One account is restricted/suspended/deleted: stop new sends immediately; define whether historic messages remain visible.
- OPEN GATE: encryption architecture, retention period, message deletion behavior, read receipts, media moderation and legal preservation workflow.

## 9. Pause, hide, block and unmatch are different actions
### Pause / hide my profile
- Stops the account from appearing in new discovery and from receiving new discovery likes, subject to a clear and consistent policy.
- Does not itself block, unmatch or delete existing conversations.
- Existing matches may remain visible and messaging may continue unless the user selects a separate control or safety restriction applies.
- On resume, re-check age/account status, profile visibility, blocks, pair exclusions and eligibility before discovery.
- Explain whether already-sent likes remain active, expire or are withdrawn; this is an OPEN GATE to decide before implementation.

### Block
- Immediate and free. Removes the blocked pair from discovery, likes, matching, messages, notifications, recommendations, search, caches and queues.
- Block is authoritative on the server; clients may optimistically hide the user but must reconcile with server state.
- Existing match/conversation is closed according to the product policy; no future matching or contact.
- Do not tell the blocked user who blocked them. Do not reveal whether a person viewed or reported them.
- A block must not depend on the blocked user being currently discoverable or in the blocker's matches; provide an abuse-report route for a profile reached through another valid surface.
- Test reverse-direction actions and account re-creation/evasion safeguards.

### Unmatch
- Immediate; removes active match lists and closes messaging; no initiator reveal.
- Read-only past-match history remains for 30 days, with report/block available while visible and no messaging. Underlying retention is separate and must follow the retention policy.
- Permanently exclude the pair from matching again, separately from blocking.
- Must not reveal who initiated unmatch through notifications, API responses or state differences.
- Any pending message, notification, attachment or worker must re-check pair state and be suppressed as needed.

### Edge cases
- Pause while likes or messages are queued: re-check visibility and conversation permissions before processing.
- Block and unmatch race: block remains the stronger safety state; both users are excluded.
- Unmatch followed by new likes: permanent pair exclusion rejects future match creation.
- Account deletion after block: preserve only the minimum pair suppression token needed to honour safety/policy if lawful; document this in retention disclosures and legal review.
- Blocked person changes profile photo/name: block must not rely only on mutable profile attributes.
- OPEN GATE: exact semantics for incoming/outgoing likes on pause, existing matches while paused, and the minimal data needed for permanent pair exclusion.

## 10. Account deletion and data lifecycle
### Required behavior
- Provide in-app account deletion initiation in line with applicable store requirements.
- Confirm consequences clearly; do not use guilt, dark patterns, repeated retention offers or unnecessary support contact as a barrier.
- Re-authentication may be required proportionately to prevent account takeover, but must have a recovery path.
- Immediately revoke sessions and stop discovery, likes, matching and new messages once deletion is authoritative.
- Orchestrate deletion across database records, S3 media, thumbnails/transcodes, search indexes, caches, analytics identifiers, push tokens, email/SMS vendors, age-assurance vendors and other processors as applicable.
- Separate data deleted, data de-identified, data scheduled for deletion and data retained under documented lawful exception. Do not promise “all data immediately erased” if backups or lawful holds prevent that.
- Backups must have a defined expiry/restore process so deleted data is not silently reintroduced after recovery.
- Keep only the minimum data needed for lawful obligations, security, abuse prevention or pair exclusion, with a documented purpose, access restriction and retention period.
- Account deletion does not automatically cancel an app-store subscription unless the relevant platform workflow does so; explain any necessary separate action without obstructing deletion.

### Edge cases
- User changes mind while deletion is pending: define a short, disclosed cancellation window if one exists; do not silently reactivate after completion.
- Vendor deletion API fails: retry, record exception, alert owner and prevent the account from becoming usable again.
- Backups restore a deleted profile: replay deletion tombstones/suppression records before serving traffic.
- Legal hold applies: restrict access, record authority and scope, periodically review, and delete when the hold ends.
- User requests data access/erasure while suspended: route through privacy process; do not make access rights contingent on account standing.
- OPEN GATE: retention schedule, lawful holds, backup expiry, processor deletion evidence and app-store subscription copy require privacy/legal review.

## 11. Appeals and enforcement lifecycle
- Explain the action, relevant rule category and available appeal route, subject to safety and legal restrictions.
- Separate temporary protective restriction from final enforcement where practical.
- Provide appeal intake that works when the account is locked; protect against appeal flooding and malicious doxxing.
- Review new evidence and whether policy was applied consistently; log reviewer, rationale, result and any reinstatement.
- Restore access and correct downstream states when a decision is reversed, without reactivating deleted accounts or restoring unsafe contact against current blocks.
- A reporter's safety should not depend on whether the reported user's appeal is successful; protect reporter identity.
- Appeals must not reveal internal detection signals that facilitate evasion.
- OPEN GATE: appeal time targets, escalation levels, external complaint route and exact disclosure obligations require legal review and pilot capacity data.

## 12. Cross-cutting failure and race-condition rules
Every workflow must be tested under duplicate requests, concurrent updates, delayed callbacks, retries, stale caches, worker crashes, network loss, partial database/media writes and account state changes mid-operation.

- **Authoritative state:** safety, consent, eligibility and entitlement decisions are made server-side from current state.
- **Idempotency:** retrying an operation does not create duplicate likes, matches, messages, deletion jobs or charges.
- **Fail closed for access:** if the system cannot confirm age eligibility, block state or message permission, do not grant access that depends on it.
- **Fail safe for user data:** partial failures must not publish private media or expose sensitive profile fields.
- **Outbox/retry pattern:** external notifications and asynchronous effects should be recoverable without duplicating core transactions.
- **Cache invalidation:** cache is an optimization, never the authority for block, deletion, suspension or hard eligibility.
- **Auditability:** high-impact state transitions have actor, timestamp, reason, source and correlation ID; avoid logging unnecessary sensitive content.
- **Observability:** alert on elevated failed reports, delayed block propagation, duplicate matches/messages, stuck deletion jobs, age-check bypasses and unbounded queues.
- **Recovery:** document rollback/repair procedures; test backups and deletion tombstones during restore.

## 13. Acceptance-test matrix
| ID | Scenario | Expected result |
|---|---|---|
| LIFE-01 | Age check pending or vendor unavailable | No discovery/matching access that requires eligibility |
| LIFE-02 | User pauses profile | Profile excluded from new discovery; existing match behavior follows disclosed pause policy |
| DISC-01 | Candidate violates any hard preference | Candidate never shown, liked or matched |
| DISC-02 | Candidate becomes ineligible after caching | Rejected at display/action boundary; stale cache cannot bypass |
| LIKE-01 | Duplicate like request | One logical like and deduplicated side effects |
| LIKE-02 | User withdraws outgoing like before reciprocal interest | Like becomes inactive; no match from stale state |
| LIKE-03 | Block exists when like-back is processed | No match; pair remains suppressed |
| MATCH-01 | Two reciprocal likes arrive concurrently | Exactly one match record and logical match event |
| MSG-01 | User sends without a valid match | Server rejects send |
| MSG-02 | Block occurs before queued message delivery | Delivery/push suppressed |
| MSG-03 | Client retries send after timeout | No duplicate message |
| SAFE-01 | Block from profile/message surface | Immediate suppression across discovery, likes, matches, messaging, notifications and queues |
| SAFE-02 | Unmatch then new reciprocal likes | Pair cannot match again |
| DEL-01 | Account deletion becomes authoritative | New discovery/likes/matches/messages stop; deletion orchestration begins |
| DEL-02 | Deleted account appears in restored backup | Deletion tombstone is applied before traffic is served |
| DEL-03 | Processor deletion fails | Exception logged, retry/owner alert raised, account remains unusable |
| MOD-01 | High-impact restriction appealed | Review is logged; reversal repairs eligible downstream state without overriding blocks |
| PRIV-01 | Public profile response requested | Only allowlisted public fields returned; no private preferences or sensitive flags |
| SYS-01 | Worker runs after state changes | Worker re-checks authoritative state and suppresses invalid side effect |
| SYS-02 | Cache is stale after block/deletion | Server-side authorization wins; stale cache cannot expose or contact the account |

## 14. Evidence required before design/build sign-off
- Approved state diagrams and exact transitions for account, age eligibility, visibility, enforcement, likes/matches, conversation and deletion.
- API/data contracts with authorization invariants, idempotency rules and error semantics.
- Data-flow/retention map covering messages, reports, media, age assurance, analytics, notifications, backups and vendors.
- Threat model and test evidence for block propagation, stale caches, race conditions, deletion, account recovery and privilege boundaries.
- UK counsel/privacy review for age assurance, Online Safety Act scope, user-generated content, reporting, retention, appeals, consumer subscription and store rules.
- Pilot evidence for eligible-supply effects, queue size five, like expiry/pass semantics, moderation workload and ranking fairness.
- Vendor proof for identity/age, media storage/processing, notification delivery, observability and deletion.
- No mobile UI design until these requirements and research completion gates are accepted.

## 15. Open decisions register
1. Does pause preserve, expire or withdraw outgoing likes? What happens to incoming likes while paused?
2. Can a user reverse a pass, and under what constraints?
3. Do likes expire? If yes, what is the disclosed policy and how are expiry notifications handled?
4. How long do messages and attachments remain visible after unmatch, suspension and deletion?
5. What exact deletion grace window, if any, is justified and disclosed?
6. What minimal pair-exclusion record can be retained after deletion, on what lawful basis and for how long?
7. What is the exact age-assurance failure/appeal path and vendor outage behavior?
8. What happens to existing matches and messaging during pause or non-safety restriction?
9. What precise premium entitlement refresh rules apply when a subscription expires or is refunded?
10. Which ranking, rate-limit and anti-evasion parameters are empirically tuned rather than hard-coded?
11. What message retention, read-receipt and deletion semantics are supported by the privacy/security design?
12. What response and appeal targets can the actual launch team reliably meet?

## Completion status
This is a lifecycle and edge-case requirements synthesis, not a final UI, API, schema or legal specification. All open gates above must be resolved or explicitly accepted with an owner before implementation sign-off. Preserve the approved product policy and do not silently convert unresolved choices into defaults.
