# Pinchd — Technical Proof-of-Concept Plan
**Date:** 22 October 2026  
**Status:** Proposed validation plan; not an implementation commitment or production architecture sign-off.  
**Goal:** Resolve the highest-risk technical, privacy, vendor and cost assumptions with the smallest reproducible experiments. Do not build mobile UI during this phase.

## 1. Decision principles
- Test risks that could invalidate the chosen architecture before adding feature breadth.
- Use synthetic accounts and synthetic content unless real data is essential and explicitly approved.
- Do not send real identity documents, biometrics or user messages to a test system without a documented data flow, vendor terms, access controls and approval.
- Record results, versions, configuration, costs, failures and limitations. “Dashboard configured” is not proof of end-to-end integration.
- Keep the proof of concept disposable and isolated from production credentials/data.
- No production migration or public beta is authorised by passing a technical PoC alone; legal, privacy, safety and operational gates remain separate.

## 2. PoC work packages

| ID | Experiment | Minimum evidence | Pass condition |
|---|---|---|---|
| POC-01 | Go API + PostgreSQL transaction baseline | Reproducible local/test environment, migration path, typed validation, auth middleware and test harness | Constraints and rollback behavior work; errors are mapped safely |
| POC-02 | Concurrent reciprocal likes | Parallel requests, pair serialization, unique constraints and outbox | Exactly one match and conversation; retries never duplicate side effects |
| POC-03 | Block/unmatch enforcement | Block and unmatch races against message send, discovery fetch, cached profile and queued push | No prohibited message or notification escapes after authoritative state commits; race behavior documented |
| POC-04 | Outbox/inbox reliability | Simulate process crash between commit and publish, duplicate delivery, retry and dead letter | No lost committed event; duplicate delivery is idempotent; backlog is observable |
| POC-05 | Age assurance vendor workflow | Inspect Didit contract, workflow version, fields, regions, callback signature, replay, expiry, appeal, failure and deletion | Age-gate state is server-authoritative; pending/failed/expired/outage cannot grant gated access |
| POC-06 | Private media pipeline | Signed upload, size/type limits, metadata stripping, scan/moderation state, private object access and deletion | No public bucket/object access; unapproved media is not shown; deletion propagates |
| POC-07 | Push notification safety | Queue message/match notification, then block/unmatch/delete before worker executes | Worker rechecks current state and suppresses prohibited content; payload reveals minimal data |
| POC-08 | Subscription entitlement | Store sandbox purchase, renewal, cancellation, refund, revoke, delayed/duplicate notification and restore | Server verifies provider state, deduplicates events and never gates safety features |
| POC-09 | Account deletion orchestration | Fail one downstream deletion target, retry, stale webhook and restore-from-backup scenario | Workflow is durable, auditable, idempotent; deletion cannot be undone by retry; exceptions are documented |
| POC-10 | Eligibility and geo query | Synthetic age/radius/gender/hard-preference combinations and query-plan fixtures | Zero hard-gate violations; query cost is measured; no exact-location leakage |
| POC-11 | Accessibility tooling baseline | Component test approach, VoiceOver/TalkBack device smoke tests and large-text pass | Critical flows have a feasible test path; third-party SDK gaps are logged |
| POC-12 | Cost and operational visibility | Load profile, logs/metrics, backup/restore, alerts and budget estimates | Costs are attributable, failure is observable and restore succeeds within provisional targets |

## 3. Suggested minimal test stack
- Go API with clear domain packages and dependency boundaries.
- PostgreSQL matching the intended managed-service version/extension assumptions.
- Migration tooling selected after team review; migrations are forward/reverse strategy documented.
- Automated unit, integration, property-based and concurrency tests.
- A local or isolated test queue/outbox worker; no premature microservice split.
- Infrastructure-as-code prototype for only the resources actually exercised.
- Synthetic vendor callbacks and billing notifications with signature/replay validation.
- Structured logs with correlation IDs and safe fields only.
- A disposable test environment with least-privilege credentials and spend caps.
## 4. Required test scenarios

### Database and API
- Two reciprocal likes arrive at the same instant.
- The same request is retried after the server commits but before the client receives a response.
- The same idempotency key is reused with a different payload.
- A user attempts to read or mutate another user's object by changing its identifier.
- A block commits while a message request is waiting on a lock.
- An outbox consumer crashes after performing its side effect but before recording completion.
- A migration fails halfway; the documented recovery path is exercised.
- Database connection loss occurs during transaction commit; client retry does not duplicate the operation.

### Vendor and age assurance
- Callback signature is invalid, replayed, delayed, duplicated or bound to a different workflow/account.
- Provider returns pending, failed, expired, unknown or contradictory results.
- Provider is unavailable during registration and later recovers.
- User appeals an outcome; staff can review within least-privilege access.
- User requests deletion; vendor-side deletion/retention evidence is recorded.
- A provider configuration change occurs; workflow version and test evidence remain traceable.

### Media and notifications
- Uploaded file has invalid type, oversized payload, malicious metadata or a failed scan.
- A signed URL expires and cannot be used to access another member's object.
- A profile is blocked, deleted or moderated after a feed is cached.
- A push is queued before unmatch/block but executes after it.
- Notification lock-screen payload does not reveal sensitive profile or message content.
- Object deletion and derived thumbnail deletion are retried after partial failure.

### Billing and deletion
- Purchase is delayed, duplicated, refunded, revoked or restored on a new device.
- Entitlement expires while the app remains open.
- Store and backend disagree; safety remains available and Premium is reconciled without client override.
- Deletion partially fails, retries and completes; status remains truthful.
- A stale billing webhook arrives after deletion and cannot reactivate the account.
- Backup restore is followed by the documented deletion-reapplication process.

## 5. Evidence record template
For every POC ID, record:
- Owner and reviewer.
- Question/hypothesis and why it matters.
- Exact code, dependency, SDK, vendor workflow and infrastructure versions.
- Test environment and synthetic fixture identifiers.
- Steps, load/concurrency parameters and expected outcome.
- Actual result, logs/metrics and reproducible evidence.
- Failure modes and known limitations.
- Data processed, region, retention and deletion behavior.
- Cost per run and estimated recurring cost at conservative volume.
- Security/privacy/legal dependencies.
- Pass / conditional pass / fail and remediation owner.
- Decision made and date for revalidation.

## 6. Provisional reliability targets
Do not treat these as final service-level objectives. Establish measured baselines first:
- Block state becomes authoritative in the database before success is returned.
- No new message is accepted after block/unmatch rules prohibit it.
- Match creation is idempotent under retries and concurrency.
- Age-assurance pending/unknown/failed/expired states never pass the eligibility gate.
- Committed domain events eventually reach consumers or a visible dead-letter state.
- Deletion jobs remain visible until complete or an approved exception is documented.
- Backup restore is tested; define formal RPO/RTO based on business impact and budget before production.
- Metrics and alerts exist for blocked-contact incidents, duplicate matches, age-gate bypass, deletion backlog, event backlog, database saturation and abnormal spend.

## 7. Cost and architecture guardrails
- Keep one Go modular monolith and one primary AWS region initially unless PoC evidence requires a change.
- Avoid microservices, Kubernetes, multi-region replication, bespoke ML and complex ranking infrastructure until measured requirements justify them.
- Compare managed service minimum costs with usage-based costs; include logs, backups, data transfer, vendor minimums, support and compliance—not only compute.
- Use a spend limit and alert before every cloud test. Shut down disposable resources after collecting evidence.
- Do not lower age assurance, moderation, privacy, security or accessibility requirements to meet a cloud budget.
- Document which assumptions require a production-like environment; a local test cannot prove cloud IAM, network, backup, vendor SLA or store behavior.

## 8. Go / hold / fail
**Pass:** expected invariants hold, evidence is reproducible, known limitations are accepted by the correct owner and recurring cost is plausible.

**Conditional pass:** core invariant holds but operational, vendor or cost evidence is incomplete; the missing gate has an owner and deadline.

**Fail:** a hard safety/privacy/consent invariant is violated, results are not reproducible, or the design relies on unsupported vendor behavior. Rework the design before continuing.

A PoC does not establish legal compliance, a vendor's contractual suitability, accessibility conformance or operational staffing readiness. Those require separate evidence.

## 9. Recommended order
1. POC-01, POC-02, POC-03 and POC-04: transaction and enforcement correctness.
2. POC-05 and POC-09: age assurance and deletion gates.
3. POC-06 and POC-07: private media and notification race safety.
4. POC-08: store entitlement lifecycle.
5. POC-10 and POC-11: eligibility performance and accessibility.
6. POC-12: cost, observability and recovery.
7. Consolidate decisions into ADRs and update the research status register before any production build decision.

**Do not start mobile UI design from this plan.** The research and requirements completion gate remains in force.
