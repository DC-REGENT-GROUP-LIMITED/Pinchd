# Pinchd — Data Classification and Retention Map
**Date:** 22 October 2026  
**Scope:** Data categories, access boundaries, purposes, lifecycle events, deletion propagation and unresolved retention decisions.  
**Status:** Working privacy-by-design artefact for review; not legal advice or an approved retention schedule. Exact periods require documented purpose, lawful basis, legal review and vendor evidence. Mobile UI design remains paused.

## 1. Purpose and handling rules
This map supports the Workstream 4 lifecycle requirements and Workstream 5 architecture decisions. It is a data inventory starter, not proof that all data flows or processors have been discovered.

- Collect only fields needed for a documented purpose; keep optional fields optional.
- Separate public profile projections from private account data, safety evidence, age-assurance records, billing records and security/audit logs.
- Assign a purpose, owner, lawful basis, sensitivity class, access role, retention trigger and deletion behavior to every field/category before production.
- Avoid storing raw identity documents, biometric data, exact GPS histories or intimate imagery unless a documented necessity and approved workflow require it.
- A vendor reference is not automatically harmless: it may be linkable to an individual and must be classified accordingly.
- Do not put message bodies, tokens, identity evidence, exact location or sensitive profile attributes into routine logs, traces or async-event payloads.
- Define lawful holds narrowly, restrict access, record the authority and scope, review periodically and delete when the hold ends.
- Retention periods must be based on purpose and legal requirements, not chosen merely because storage is cheap or deletion is difficult.

## 2. Classification levels
| Class | Description | Examples | Baseline controls |
|---|---|---|---|
| **C0 — Public profile projection** | Information deliberately shown to eligible users | Published photos after metadata stripping; display name; approved prompts | Explicit allowlist; visibility and eligibility checks; no permanent public media URL |
| **C1 — Personal account data** | Identifiable operational/profile data not generally public | Email/phone, date of birth if held, private preferences, approximate location, account settings | Encryption, access control, purpose limitation, user-access/deletion workflow |
| **C2 — Sensitive/high-impact data** | Data with elevated harm if exposed or misused | Relationship/sexual-life inferences where applicable, sensitive identity attributes, private messages, precise location if ever collected | Strict purpose review, minimisation, narrow access, DPIA/legal review where required, no routine logs |
| **C3 — Restricted safety/identity evidence** | Evidence whose exposure could cause serious harm | Reports, intimate imagery submitted as evidence, suspected exploitation evidence, identity/age-assurance images or biometric material if processed | Dedicated restricted storage/access, MFA, audit every access/export, specialist playbooks, strict retention and legal review |
| **C4 — Security and accountability records** | Records used to secure systems and demonstrate actions | Authentication/security events, admin audit logs, moderation decision history, deletion task evidence | Tamper-resistant controls where feasible, redaction, least privilege, retention justified by security/legal need |

Classification is a handling aid, not a substitute for assessing the legal category of each field or processing purpose.

## 3. Data inventory and lifecycle matrix
| Data group | Purpose | Default access | Key lifecycle events | Deletion/retention rule to define |
|---|---|---|---|---|
| Account identifiers and contact details | Sign-in, recovery, essential service communications | Account services; narrowly authorised support | Account creation, change, closure | Delete/de-identify after account closure unless a documented legal/security exception applies |
| Age-assurance workflow/results | Enforce minimum-age eligibility and handle disputes | Eligibility service; restricted compliance roles | Attempt, pass/fail, appeal, vendor deletion | Keep only minimum result/reference and evidence required; exact period and vendor copies require legal/vendor review |
| Raw identity/biometric evidence, if processed | Only if required by chosen assurance workflow | Highly restricted; avoid copying locally | Vendor workflow and appeal | Prefer vendor-held evidence where lawful and appropriate; establish deletion proof, transfer and appeal rules before use |
| Public profile projection | Show profile to eligible users | Eligible discovery/match surfaces | Publish, edit, pause, restriction, deletion | Remove from public service immediately when no longer eligible; downstream copies invalidated |
| Private profile fields/preferences | Personalisation and matching | Account owner and tightly scoped services | Edit, preference change, account closure | Delete with account or field removal, subject to narrowly documented exceptions |
| Profile media and derivatives | Display profile and optional user-shared content | Private object store; signed access only after policy check | Upload, moderation, edit/remove, report, deletion | Remove public access immediately; delete originals/derivatives per schedule; retain restricted evidence only when justified |
| Discovery/cache/index data | Efficient retrieval and ranking | Discovery service only | Preference change, block, pause, restriction, deletion | Invalidate promptly; never authoritative for access; rebuild must honour suppression/deletion tombstones |
| Likes, passes and incoming queue | Record interest and queue state | Relevant account surfaces and matching services | Like, withdraw, pass, eligibility change, account closure | Define expiry and pass semantics; remove ineligible queue items and avoid retaining unnecessary pair history |
| Matches and pair suppression | Enforce reciprocal interest and prevent re-matching after unmatch | Matching service; minimal UI projections | Match, unmatch, block, account deletion | Permanent pair exclusion is approved after unmatch; retention after account deletion requires explicit privacy/legal decision |
| Messages and attachments | Enable conversation and safety reporting | Participants while authorised; restricted case access for valid reports | Send, edit/delete if supported, block, unmatch, restriction, account deletion | Separate user-visible history from evidence/legal retention; define message/attachment policy and processor/backups |
| Reports, cases and enforcement | Investigate abuse, protect users and manage appeals | Assigned T&S staff; stricter access for sensitive evidence | Report, triage, action, appeal, closure | Retain only for defined safety/legal purpose; restrict access; set review/expiry and lawful-hold process |
| Blocks | Prevent unwanted contact and discovery | Blocker controls; enforcement/policy services | Block, possible unblock, account deletion | Maintain authoritative suppression while policy requires; decide minimal retention and unblock semantics |
| Billing and entitlements | Verify purchases, restore purchases and grant Premium features | Billing service; limited finance/support access | Purchase, renewal, refund, expiry, account closure | Retain only required transaction/audit data; separate platform records from app profile data |
| Notifications and push tokens | Deliver service and safety notifications | Notification service | Register, revoke, delivery, account deletion | Revoke tokens on logout/deletion; avoid sensitive message contents in push payloads |
| Analytics and experiments | Understand performance and improve product | Approved analytics roles | Event creation, aggregation, account deletion request | Minimise/pseudonymise; define purpose and retention; ensure deletion/opt-out semantics where applicable |
| Logs, traces and crash reports | Diagnose reliability/security | Engineering/SRE with redaction | Event, incident, expiry | No raw secrets/content; retention tied to diagnostic/security need; tightly control exports |
| Admin access and audit | Accountability for privileged operations | Security/privacy/T&S leads as appropriate | View/export/change/override | Record actor, purpose, target, time and outcome; protect against tampering; define retention and review |
| Async jobs and deletion tasks | Reliable background processing | Worker/service roles | Enqueue, retry, dead-letter, complete | Minimise payload; expire completed jobs when audit needs permit; ensure retries cannot resurrect deleted data |
| Backups and disaster recovery | Restore service after failure | Restricted infrastructure/security roles | Backup, restore, expiry | Document expiry, encryption/key recovery, access, deletion tombstone replay and restore tests |
| Vendor/subprocessor copies | Deliver identity, media, messaging, billing or analytics service | Vendor roles under contract | Transfer, processing, deletion request, incident | Maintain processor register, contract terms, transfer review and deletion/return evidence |

## 4. Required retention schedule fields
Do not fill exact durations by guesswork. For every data category, record:
- Category and specific fields.
- Purpose(s) and lawful basis; Article 9 condition if special-category data is processed.
- Whether the data is optional, public, sensitive or restricted evidence.
- Data controller/processor and storage region.
- Who can read, change, export and delete it.
- Start trigger for the retention clock (creation, last activity, account deletion, case closure, transaction, legal hold release).
- Standard retention period and rationale.
- Deletion or de-identification action, including replicas and derivatives.
- Backup expiry and restore behavior.
- Legal hold/exception criteria, approval, scope, review date and end condition.
- Processor/vendor deletion mechanism and evidence.
- User-facing explanation and rights-handling route.
- Owner, reviewer and last validation date.

## 5. Deletion orchestration requirements
1. Authenticate and validate the deletion request with proportionate account-recovery controls.
2. Commit deletion-pending state and immediately stop normal discovery, likes, matches and new messages.
3. Revoke sessions and invalidate access tokens, push tokens, signed media access and cached public projections as applicable.
4. Create an idempotent deletion task for each relevant store/system/processor.
5. Delete or de-identify primary account/profile data and media variants according to the approved schedule.
6. Remove discovery indexes, caches, recommendation queues and analytics identifiers where required.
7. Send vendor deletion/return requests and track their status and evidence.
8. Retain only approved legal/security/safety exceptions with restricted access, explicit purpose and expiry/review.
9. Apply deletion tombstones before serving restored backups; do not silently recreate deleted users.
10. Report completion only when all required tasks are complete or a documented exception has been reviewed and accurately described.

## 6. Special handling: safety reports and evidence
- Reports can contain third-party personal data; avoid promising unrestricted deletion of all evidence where a lawful safety/legal need may require retention, but do not use this as a blanket reason to keep everything indefinitely.
- Separate the report narrative from the underlying content where feasible, and keep references to the minimum object/version needed for investigation.
- Restrict sensitive image evidence; never copy suspected child sexual abuse material into ordinary ticketing, email or chat tools.
- Log access/export with a case ID and reason; require approval for bulk export.
- Avoid using case evidence for unrelated product analytics or model training without a separate documented review and lawful basis.
- Ensure account deletion does not disable the ability to report a serious incident that occurred before deletion where a lawful route exists.

## 7. Special handling: location and preference data
- Prefer coarse location; do not store location trails or expose exact location without a clear, reviewed necessity.
- Keep discovery preferences private; do not return hard/soft settings to other users or include them in routine analytics.
- Treat sensitive attributes and inferences as requiring explicit purpose and legal/privacy review before collection or matching use.
- Missing or undisclosed values must not be transformed into positive consent or compatibility signals.
- When preferences change, invalidate stale candidate caches before they can expose a now-ineligible profile.

## 8. Vendor and cross-border evidence pack
For each processor/vendor, document:
- Data categories, purpose and controller/processor roles.
- Subprocessors, storage/processing regions and international transfers.
- Security measures, access controls, incident notification and support access.
- Retention defaults, deletion APIs/processes, backup expiry and deletion evidence.
- Contract, data-processing terms and any transfer mechanism/assessment required.
- Service outage, fallback, export/return and termination plan.
- Age-assurance-specific workflow/version, collected fields, image/biometric handling, accuracy/error/appeal process and user-impact testing.
- Whether data is used for vendor training, analytics or secondary purposes and what controls apply.

## 9. Open decisions and required owners
1. Exact retention period for account/contact data — privacy/legal owner.
2. Exact age-assurance evidence/result retention and vendor deletion — privacy/legal/vendor owner.
3. Message and attachment retention after unmatch, suspension and deletion — product plus privacy/legal.
4. Minimum pair-suppression data after account deletion — privacy/legal; implementation blocked pending decision.
5. Safety case/evidence retention, lawful holds and external disclosure — T&S plus legal.
6. Analytics event retention and deletion/opt-out semantics — product analytics plus privacy.
7. Backup retention, RPO/RTO and tombstone replay — engineering/security.
8. Processor deletion SLAs and evidence — vendor owner.
9. Log/audit retention and privileged export controls — security/privacy.
10. Lawful basis and any Article 9 condition for each sensitive field or matching use — privacy/legal.
11. Whether any precise location, biometrics or novel AI processing is actually necessary — product/privacy/security.
12. User-facing privacy notice and deletion wording — legal/product.

## 10. Acceptance criteria
- Every production data field has a documented purpose, classification, owner and retention rule.
- Every processor and data transfer is in the register and contractually reviewed as required.
- Public profile responses use an explicit allowlist.
- Logs and async events contain no unnecessary sensitive payloads.
- Blocks, deletion and hard eligibility invalidate or bypass caches safely.
- Deletion tasks are idempotent, monitored and recoverable; restore tests prove deleted accounts are not resurrected.
- Restricted evidence access is role-controlled and audited.
- Retention exceptions are narrow, approved, reviewable and time-bounded where possible.
- Privacy/legal review signs off on the schedule and notices before public launch.

## Completion status
This map identifies categories and required controls but intentionally does not invent retention periods or legal bases. It must be reconciled against the actual data model, vendor contracts, privacy impact assessment, legal advice and implemented data flows before it can become Pinchd's approved retention schedule.
