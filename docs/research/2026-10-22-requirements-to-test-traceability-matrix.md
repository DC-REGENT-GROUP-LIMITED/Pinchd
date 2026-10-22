# Pinchd — Requirements-to-Test Traceability Matrix
**Date:** 22 October 2026  
**Status:** Working verification plan; tests are proposed until executed. No test is marked passed without recorded evidence.  
**Related:** Full-fold lifecycle/privacy/architecture reconciliation; Workstream 4 lifecycle requirements; Workstream 5 open architecture decisions; PostgreSQL schema blueprint; technical proof-of-concept plan.

## 1. Purpose and evidence rules
This matrix links product requirements and POC-01 through POC-12 to observable proof, accountable owner roles, pass/fail conditions and release gates. It is not proof that the application exists or that tests have run.

Evidence must be reproducible, dated and tied to a commit/build, dependency/vendor workflow version and test environment. Prefer synthetic accounts, synthetic profile/media/message content and simulated provider callbacks. Never place identity documents, biometric payloads, private user messages or safety case evidence in general test logs. Record conditional passes with limitations and remediation owners; no undocumented waiver for a safety/privacy invariant.

### Status vocabulary
- NOT RUN: no evidence supplied.
- PASS: all assertions passed and evidence is attached/referenced.
- CONDITIONAL: bounded limitation, owner and deadline recorded; must not waive a critical safety/legal gate.
- FAIL: one or more assertions failed.
- BLOCKED: external decision, vendor evidence, legal review or environment prevents valid testing.
- NOT APPLICABLE: written rationale and approver recorded.

## 2. Release gates
- **G0 — Design/schema freeze:** lifecycle defaults confirmed or explicitly left out of scope; data classification and retention decisions have owners; API/data model reviewed. No migration is authorised solely by this matrix.
- **G1 — Controlled technical beta:** core identity/age gates, eligibility, block/unmatch, match idempotency, private media, deletion safety, accessibility of critical paths and minimum T&S operations pass for the beta scope.
- **G2 — Public launch:** applicable legal scope/risk assessments, DPIA/lawful basis/retention, vendor due diligence, subscription review, accessibility, incident operations, support coverage, cost controls and recovery evidence signed off.
- **G3 — Paid launch:** store billing and consumer subscription obligations, pricing disclosures, restore/refund/expiry reconciliation and support process pass in store sandbox and counsel review.
- **G4 — Scale/expansion:** capacity, reliability, staffing, data residency, regional recovery and unit economics justify the change.

A test passing a technical assertion does not prove legal compliance or operational readiness.

## 3. Lifecycle requirements traceability

| ID | Requirement / invariant | Verification scenario | Pass criteria | Owner role | Gate |
|---|---|---|---|---|---|
| LC-01 | Required age eligibility precedes gated access | Attempt discovery during unknown, pending, failed, expired and appeal-pending states | Every gated route denies access consistently; no client-side bypass | Identity + security | G1/G2 |
| LC-02 | Vendor callback authenticity and binding | Invalid signature, wrong account/workflow, replayed and duplicated callback | Invalid/unbound callbacks rejected; valid duplicate is idempotent and audited | Identity + vendor owner | G1 |
| LC-03 | Out-of-order age results | Deliver older pending/failure after newer valid state and vice versa | State transition follows documented provider/version rules; no stale callback grants access | Identity + vendor owner | G1 |
| LC-04 | Accessible age-check failure/appeal | Provider outage, false reject, no-device/assistive-tech path | Neutral error, retry/support/appeal reachable outside restricted area; no fraud signals exposed | Product + accessibility + T&S | G1/G2 |
| LC-05 | Independent state dimensions | Combine paused profile, eligible age state, Premium expiry and scoped restriction | Derived capabilities follow policy; one dimension does not overwrite another | Backend + QA | G0/G1 |
| LC-06 | Publish only valid profile/media | Publish with missing fields, unapproved media or processing failure | Publish denied or restricted according to explicit policy; no unprocessed content exposed | Profile/media + QA | G1 |
| LC-07 | Public field allowlist | Probe profile response with private preferences, internal flags and account fields | Only approved public fields returned; private fields absent from API and caches | Backend + security | G1 |
| LC-08 | Approximate location/no triangulation | Compare displayed location and repeated candidate outputs across test coordinates | No exact coordinates or location trail exposed; re-identification/triangulation risk reviewed | Product privacy + security | G1/G2 |
| LC-09 | Hard eligibility before ranking | Generate candidates violating age, selected gender, hard radius or hard preference | Zero excluded candidates appear, regardless of score, Premium, exploration or low supply | Matching + QA | G1 |
| LC-10 | Reciprocal hard eligibility | Candidate A qualifies B but B does not qualify A | No recommendation/action that violates the approved reciprocity rule | Matching + product | G1 |
| LC-11 | Unknown is not positive intent | Null/undisclosed relationship goal, family plan or lifestyle field | Unknown never treated as open-minded, consent or positive compatibility | Matching + QA | G1 |
| LC-12 | Eligibility rechecked on action | Candidate changes preference/location/age state after feed page created | Stale item removed or action rejected before like/match commit | Matching + backend | G1 |
| LC-13 | Profile pause behavior | Pause while in discovery, incoming queue, outgoing likes and active match | Behavior matches the finally approved pause decision; no accidental unmatch or new exposure | Product + backend + QA | G0/G1 |
| LC-14 | Existing outgoing likes on pause/resume | Pause, change hard preference, resume, attempt reciprocal match | No silent withdrawal; re-evaluate before action according to final policy | Product + matching | G0/G1 |
| LC-15 | Pass semantics | Pass same candidate repeatedly; attempt undo/re-discovery; compare with block | Behavior matches approved reversibility decision; pass never masquerades as safety block | Product + QA | G0/G1 |
| LC-16 | Outgoing-like expiry | Keep valid like beyond normal test duration; change eligibility then act | No unapproved silent expiry; changed eligibility still prevents invalid match | Matching + QA | G0/G1 |
| LC-17 | Five-like free queue | Seed >5 eligible incoming likes plus blocked/ineligible/stale items | Five eligible items are unblurred; queue advances correctly; suppressed items never leak | Backend + billing + QA | G1/G3 |
| LC-18 | Premium full eligible queue | Active, expired, refunded and unverified Premium states | Only verified current entitlement expands visible queue; safety remains unchanged | Billing + backend | G3 |
| LC-19 | Like-with-comment reference | Comment targets a prompt/photo that is edited or removed | Reference remains truthful or neutral unavailable state; no fabricated content | Product + privacy + QA | G1 |
| LC-20 | Match idempotency | Parallel reciprocal likes, retry after timeout, duplicate idempotency key | Exactly one active match and conversation per pair; conflicting key reuse rejected safely | Backend + QA | G1 |
| LC-21 | Like withdrawal before match | Withdraw own outgoing like before reciprocal interest; try after match | Own like can be withdrawn only under approved state rules; no unauthorised target mutation | Backend + QA | G1 |
| LC-22 | Messaging requires active match | Like without match, unmatch, block, suspension or deletion then send | No message accepted when policy disallows; error does not leak hidden block status | Messaging + security | G1 |
| LC-23 | Message-delivery race | Queue message/push, then block/unmatch/delete before worker executes | Worker rechecks state and suppresses prohibited delivery/notification | Messaging + T&S | G1 |
| LC-24 | Block propagation | Block while discovery page, queue item, message and push are cached/queued | Block becomes authoritative before success; all relevant surfaces suppress pair | Backend + T&S + QA | G1 |
| LC-25 | Unmatch history | Unmatch, inspect active list, past-match history, messaging and report/block options | Immediate messaging closure; no initiator reveal; read-only history matches approved 30-day display policy | Product + messaging + privacy | G1 |
| LC-26 | Pair exclusion after unmatch | Attempt rediscovery/matching after unmatch; separately unblock | Pair remains excluded regardless of block state; post-deletion behavior not assumed until legal gate closes | Matching + privacy/legal | G0/G1 |
| LC-27 | Restriction scope | Apply messaging-only restriction and full suspension | Only permitted capabilities remain; queued work respects updated restrictions | T&S + backend | G1 |
| LC-28 | Deletion request durability | Retry request, lose client response, concurrent sessions and stale webhook | One deletion workflow; new activity stopped according to policy; stale events cannot reactivate account | Privacy + backend | G1/G2 |
| LC-29 | Deletion propagation | Fail media/vendor/index/cache task, retry and restore backup | Tasks durable/idempotent; tombstones reapplied; completion status truthful; exception audited | Privacy + platform | G1/G2 |
| LC-30 | Retention separation | Inspect visible past-match data, messages, safety evidence, audit logs and backups | Each class follows separately approved retention schedule; 30-day display is not misused as global retention | Privacy + security | G0/G2 |
| LC-31 | Optional feature minimisation | Decline AI/voice/marketing/profile enrichment | Basic product remains usable; optional processing not silently enabled | Product + privacy + QA | G1/G2 |
| LC-32 | Accessibility of critical safety actions | Screen reader, text scaling, keyboard/switch and error states for report/block/appeal/delete | Critical actions are discoverable and operable; severe defects block release | Accessibility + QA | G1/G2 |
| LC-33 | Staff evidence access | Staff roles attempt case access/export outside assignment/permission | Least privilege enforced; access reason/audit recorded; bulk export restricted | T&S + security | G1/G2 |
| LC-34 | Analytics minimisation | Inspect emitted events and logs under representative workflows | No raw message, ID/biometric payload, exact location, secrets, private preferences or reporter identity in general analytics | Privacy + security | G1/G2 |
| LC-35 | Incident and escalation readiness | Tabletop urgent threat, NCII, extortion, account takeover and provider outage | Named decision-maker, escalation route, evidence handling and communication steps are executable | T&S lead + operations | G1/G2 |
| LC-36 | Recovery and restore | Restore database/media metadata from backup, reapply deletion tombstones and verify integrity | Restore succeeds inside formally approved RPO/RTO; deleted accounts not resurrected | Platform + security | G2 |

## 4. Technical PoC traceability (POC-01 through POC-12)

| PoC | Hypothesis and minimum experiment | Pass criteria / required evidence | Owner role | Gate / dependency |
|---|---|---|---|---|
| POC-01 Go API + PostgreSQL transaction baseline | Disposable environment; migrations; typed validation; auth middleware; safe error mapping; connection/commit failure | Reproducible setup; constraints and transaction behavior verified; no sensitive error leakage; migration recovery documented | Backend/platform | G0; Go/Postgres versions selected |
| POC-02 Concurrent reciprocal likes | Parallel reciprocal requests, retries, unique constraints and pair locking | Exactly one match and conversation; idempotency works; conflicting key payload rejected; evidence includes concurrency run and DB state | Backend/QA | G1; matching invariants defined |
| POC-03 Block/unmatch enforcement | Race block/unmatch against message send, discovery, cached profile and queued push | No prohibited action accepted after authoritative state commit; no stale exposure; race outcomes documented | Backend/messaging/T&S | G1; block policy |
| POC-04 Outbox/inbox reliability | Crash before publish, duplicate delivery, consumer crash after side effect, retry/dead letter | No lost committed event; duplicate effects prevented; backlog/dead letter observable and alertable | Backend/platform | G1; event schemas |
| POC-05 Age assurance vendor workflow | Inspect actual Didit agreement/configuration/workflow, fields, regions, callback, expiry, appeal and deletion | Evidence pack complete; valid callback bound to account/workflow; pending/failed/expired/outage never grants access; false-reject path tested | Identity/vendor/privacy/legal | G1/G2; vendor evidence and legal review |
| POC-06 Private media pipeline | Signed upload, file validation, metadata stripping, moderation, object access and deletion | Private by default; unapproved media not served; no cross-account object access; derived assets deleted/revoked as specified | Media/platform/security | G1; storage choice and media policy |
| POC-07 Push notification safety | Queue notification, then block/unmatch/delete or change preferences before worker runs | Current-state check suppresses prohibited delivery; payload minimised; preferences and token revocation honored | Messaging/platform/T&S | G1; event/outbox |
| POC-08 Subscription entitlement | Store sandbox purchase/renewal/cancel/refund/revoke/restore with delay and duplicates | Verified provider state wins over client; event dedupe works; expiry updates free queue; safety unaffected | Billing/backend/legal | G3; store account and legal review |
| POC-09 Account deletion orchestration | Inject failure into each downstream target; retry; stale webhook; restore backup | Durable idempotent workflow, auditable exceptions, truthful completion and no account resurrection | Privacy/platform/security | G1/G2; approved retention/exception schedule |
| POC-10 Eligibility and geo query | Synthetic age/radius/gender/hard-preference combinations, nulls and query plans | Zero hard-gate violations; latency/cost measured; no exact-location leakage; indexes justified by measurements | Matching/data | G1; approved hard/soft policy |
| POC-11 Accessibility tooling baseline | Automated component checks plus VoiceOver/TalkBack, large text, focus/error paths and vendor SDK smoke tests | Repeatable manual/automated path for critical flows; third-party gaps recorded with owner and mitigation | Mobile/accessibility/QA | G1/G2; device matrix |
| POC-12 Cost and operational visibility | Load profile, logs/metrics, backup restore, alerts, budget estimates and failure exercise | Spend attributable; failure visible; restore evidence matches approved RPO/RTO; alert ownership and thresholds documented | Platform/finance/operations | G2; workload assumptions and regional quote |
## 5. Evidence package required for each row
Attach or link:
1. Requirement ID and PoC ID; test date; responsible tester and independent reviewer.
2. Code commit/build hash, dependency versions, database version/extensions, infrastructure configuration and relevant vendor workflow/version.
3. Environment identifier, synthetic fixture IDs, test input matrix and exact steps.
4. Expected result and actual result, including failure traces and relevant metrics.
5. Database state or sanitised logs that demonstrate the assertion without exposing sensitive payloads.
6. Cost per test/run and projected monthly cost with explicit workload assumptions.
7. Data processed, processing/storage region, processors, retention/deletion behavior and security controls.
8. Result: NOT RUN, PASS, CONDITIONAL, FAIL, BLOCKED or NOT APPLICABLE.
9. Limitations, remediation ticket/owner/date, reviewer sign-off and next revalidation trigger.

## 6. Release decision rules
- Never mark a row PASS because a dashboard, vendor setting or configuration toggle exists; prove end-to-end behavior.
- Conditional pass is not available for age-gate bypass, hard-preference violation, block bypass, prohibited messaging, public private media, account resurrection or unauthorised evidence access.
- Legal/privacy decisions require the appropriate qualified reviewer and documented evidence; engineering cannot infer lawful basis or retention periods.
- A vendor assertion is supporting evidence, not a replacement for Pinchd's own due diligence and integration tests.
- If a prerequisite policy is unresolved, mark affected tests BLOCKED or test only the invariant that is already approved.
- Record failures honestly; do not change acceptance criteria retrospectively just to obtain a pass.
- POC completion permits an architecture decision review, not automatic production approval.

## 7. Execution order
1. POC-01 baseline and CI test harness.
2. POC-02 concurrency and POC-04 outbox reliability.
3. POC-03 block/unmatch plus POC-07 push race.
4. POC-10 eligibility and geo query.
5. POC-06 private media.
6. POC-05 age assurance only after vendor evidence and data-flow review permit the test.
7. POC-09 deletion after the retention/exception policy is defined enough to test without inventing retention.
8. POC-08 billing in store sandbox before paid launch.
9. POC-11 accessibility throughout, not only at the end.
10. POC-12 cost, backup/restore, operational and incident tests before controlled beta/public launch.

## 8. Current baseline
All rows are **NOT RUN** unless a separate dated evidence record proves otherwise. This matrix does not claim code exists, a test environment has passed, Didit has been integrated, or legal/operational approval has been granted.
