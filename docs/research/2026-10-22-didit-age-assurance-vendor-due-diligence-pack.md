# Pinchd — Didit Age-Assurance Vendor Due-Diligence and Evidence Pack
**Date:** 22 October 2026  
**Status:** Evidence request and validation plan. Not a statement of Didit's current capabilities, contract terms, security posture or compliance status. Complete only from current vendor documents, configuration inspection, written answers and Pinchd's own tests.  
**Scope:** Required age assurance for a UK-first adult dating service, separate from optional profile-photo verification.

## 1. Purpose and non-negotiable boundary
Pinchd has identified Didit as the intended age-assurance vendor/workflow, but the project evidence available in the research documents does not prove the final workflow configuration, end-to-end integration, data handling, contract terms, accuracy, appeals or deletion behavior. The integration is **not considered complete** because a dashboard or liveness setting is configured.

Age assurance must be a server-authoritative eligibility gate. Unknown, pending, failed, expired, contradictory or unavailable provider state must never grant age-gated access. Optional profile verification is a separate feature and must not be represented as proof of age unless counsel and the technical evidence explicitly establish that it meets the required age-assurance purpose.

Do not submit real identity documents, biometric material or real-user data to a development/test flow until the data flow, vendor terms, storage/access, region, retention and approvals are documented. Use synthetic callbacks and vendor sandbox/test methods where available.

## 2. Evidence status at creation
- Intended provider: Didit, based on current project direction.
- Exact product/module, workflow ID/version and configuration: **not evidenced in this pack**.
- Contract/DPA/subprocessors/regions: **not evidenced in this pack**.
- End-to-end mobile/backend integration: **not proven**.
- Callback signature/replay protection and state mapping: **not proven**.
- Retention/deletion evidence: **not proven**.
- Accuracy, false rejection, accessibility and appeal performance: **not proven**.
- Approval status: **BLOCKED pending evidence and review**.

## 3. Vendor evidence request checklist

### A. Product, workflow and contractual scope
Request and retain:
- Exact product name, service tier, modules enabled and workflow identifier/version.
- Workflow diagram from initiation through result, retry, expiry, appeal/manual review and deletion.
- Current configuration export or screenshots plus configuration-change history and owner.
- Written definition of what each result means: verified age threshold, age band, pass/fail, inconclusive, liveness result, document authenticity, face match and fraud/risk indicators.
- Clarification whether the configured workflow estimates/assures an age threshold or verifies identity for a different purpose; do not infer equivalence.
- Current contract/order form, service description, SLA/support terms, DPA and relevant security schedules.
- Subprocessor list and change-notification terms.
- Vendor use of data for model training, analytics, product improvement or secondary purposes; opt-outs and contractual restrictions.
- Service termination/export/deletion assistance and post-termination retention.
- Support escalation route, incident notification commitments, maintenance windows and outage communications.

### B. Data inventory and data flow
Obtain a field-by-field inventory covering:
- Inputs collected from the user, device, app and identity document.
- Images/video/selfie frames, liveness signals, face geometry/templates or other biometric-like material, extracted document fields, date of birth/age result, document number, name, address and metadata, where applicable.
- Derived scores, fraud signals, reason codes, age bands, confidence measures and manual-review notes.
- Data sent from mobile client to vendor, vendor to Pinchd backend, and vendor/subprocessors to any other destination.
- Whether Pinchd receives raw images, only a token/result, or a mixture; exact API response and webhook payload examples with sensitive fields redacted.
- Purpose, controller/processor role, lawful basis assumptions, access roles and retention/deletion for every field.
- Processing/storage regions, support-access locations, subprocessors and any international transfer mechanism/assessment required.
- Encryption in transit/at rest, tenant separation, key management, access logging, breach response and deletion guarantees.
- Which data fields are mandatory versus optional and whether the workflow can be configured to minimise collection.
- Whether user data is used to train or improve models and whether a no-training/no-secondary-use commitment can be contracted.

**Required output:** a versioned data-flow diagram plus data dictionary and retention table. Hosting in a UK/EU region alone is not proof that every access or transfer is local.

### C. Accuracy, bias and accessibility
Request vendor documentation and independently test:
- Measured performance for the exact workflow and age threshold being used, with methodology, sample composition, confidence intervals and limitations.
- False acceptance and false rejection rates, including uncertainty/indeterminate results.
- Evidence by relevant demographic groups and document types, where lawfully and ethically measured; limitations of the evaluation and any known disparities.
- Performance under low light, poor camera, older device, disability/accessibility constraints, network interruption and supported languages.
- Supported devices, OS versions, SDK versions, webview/native requirements, background/deep-link behavior and fallback routes.
- User-facing explanation, accessible instructions, alternative method, manual review and appeal route for users unable to complete the default flow.
- Vendor incident history or known failure modes relevant to this exact product, with remediation information.
- Change management for model/workflow updates and notification before a change that could affect results.

Do not assume that vendor-wide accuracy claims apply to Pinchd's configured flow, UK user population, threshold or supported devices.

### D. API, callback and security evidence
Request exact documentation for:
- Authentication method and credential scopes.
- Callback/webhook signature algorithm, signed fields, timestamp tolerance, secret rotation and key compromise response.
- Event IDs, replay behavior, retry schedule, ordering guarantees, idempotency and delivery logs.
- Workflow/account/session binding and how Pinchd correlates an event to the correct user without trusting user-supplied identifiers.
- Token lifetime, expiry, cancellation, one-time use and protection against session reuse.
- Rate limits, error taxonomy, timeout semantics and safe retry guidance.
- SDK permissions, collected device data, local storage/cache behavior and third-party SDK dependencies.
- Security assurance reports/certifications if available, penetration-test summary and vulnerability disclosure process.
- Incident notification timing, investigation cooperation and evidence access.
- Vendor status page, outage notification and support response route.

Pinchd must verify callback authenticity, workflow/account binding, event freshness and allowed state transitions server-side. A client redirect or client-side success screen is not authoritative evidence of a pass.

### E. Retention, deletion and user rights
Request written, specific answers:
- Retention period for raw images/video, extracted document fields, biometrics/templates, result tokens, event logs, support tickets, backups and fraud/appeal records.
- What triggers each retention clock and whether any periods are configurable.
- Deletion API/process, required identifiers, completion time, retry semantics and confirmation evidence.
- What remains in backups and for how long; how deleted data is prevented from reappearing after restore.
- Legal holds, fraud/security exceptions and how they are authorised, scoped, reviewed and disclosed.
- Whether Pinchd can retrieve a deletion receipt or audit evidence without retaining the original sensitive data.
- How data subject access/erasure requests are routed and how vendor identity matching is safely handled.
- Contractual treatment of data at service termination and subprocessors.
- What happens to age eligibility if the vendor deletes its underlying evidence: whether a minimal result token can remain valid, whether re-check is required and on what basis.

Do not invent or adopt a retention period until the privacy/legal owner has reviewed actual vendor terms, Pinchd's purpose and the full lifecycle. Minimise raw identity/biometric data held by Pinchd where possible.

### F. Operational resilience and support
Obtain:
- Availability/support commitments and exclusions.
- Known maintenance/outage behavior, regional failover and status communications.
- Rate limits and capacity constraints relevant to expected beta volume.
- Support hours, severity definitions, escalation contacts and incident bridge process.
- Recovery/back-up commitments for vendor-side data and configuration.
- Export, exit and alternative-vendor migration plan.
- Cost schedule including per-check, retry, manual review, failed check, storage, minimum commitment, tax, overage and support charges.
- Rules for sandbox/test use and any restrictions on automated tests.
- Change notification and version pinning options.

## 4. Pinchd internal architecture and implementation requirements
- Create a server-side assurance session bound to the authenticated account, intended workflow version and an expiring nonce/session identifier.
- Do not place provider secrets in the mobile app.
- Keep eligibility state separate from the raw vendor result; store only the minimum result/evidence pointer needed for the documented purpose.
- Verify callback signature and freshness, bind it to the expected session/account/workflow, and process it idempotently.
- Reject replayed, unbound, malformed, expired and contradictory events; audit transitions without copying unnecessary payloads.
- Define explicit state mapping for eligible, pending, failed, expired, inconclusive, provider-unavailable and appeal-pending. Unknown is denied by default.
- A delayed or unavailable vendor must leave the user in a truthful restricted/pending state; never create a temporary bypass.
- Client redirect, local storage, push notification or UI state cannot set age eligibility.
- Do not disclose vendor fraud scores or detailed anti-evasion signals to users.
- Provide an accessible retry/support/appeal path that does not require entering restricted discovery.
- Re-check age eligibility on protected API actions and on relevant state transitions.
- Define re-check triggers, workflow version changes and result expiry before launch.
- Record notice/consent version where appropriate, but do not treat consent as a substitute for lawful basis or a required Article 9 condition.
- Keep age assurance separate from optional profile-photo verification and any matching/ranking score.

## 5. Required Pinchd proof-of-concept tests
1. Valid signed callback for the correct account/session/workflow changes state only according to documented result mapping.
2. Invalid signature, altered payload, wrong account, wrong workflow, expired nonce and replay are rejected.
3. Duplicate callback produces no duplicate side effect; out-of-order callback cannot overwrite a newer valid state incorrectly.
4. Pending, inconclusive, failed, expired and unknown results deny gated access.
5. Vendor outage during registration does not expose discovery; recovery permits safe retry without duplicate charges/sessions.
6. Client fakes a success redirect or tampers with local state; backend still denies access without verified server result.
7. Account deletion during a pending vendor session prevents later callback from reactivating account.
8. Workflow configuration changes are versioned and tested before rollout.
9. Appeal and manual review are accessible, role-restricted and audited.
10. Vendor deletion request is tracked to evidence; backup and exception behavior matches approved retention policy.
11. Network loss, app termination, device back navigation and deep-link replay do not bypass the gate.
12. Test logs and analytics contain no unnecessary identity image, biometric payload, document number or vendor fraud signal.

## 6. Decision and acceptance checklist
- [ ] Exact product/module, workflow ID and version captured.
- [ ] Contract, DPA, subprocessors and security terms reviewed.
- [ ] Data dictionary, data-flow diagram, regions and transfer assessment completed.
- [ ] Retention/deletion schedule and backup behavior documented for every field.
- [ ] Accuracy/error/false-reject evidence reviewed for the actual workflow and threshold.
- [ ] Accessibility, alternatives, retry, manual review and appeal paths tested.
- [ ] Callback signature, account/workflow binding, replay and idempotency tests pass.
- [ ] Outage and stale-callback tests pass without eligibility bypass.
- [ ] Privacy/legal owner records lawful basis and any required special-category condition, DPIA outcome and notice changes.
- [ ] Vendor owner records support escalation, cost model, outage behavior and exit plan.
- [ ] Independent reviewer signs the evidence pack and lists residual risks.
- [ ] POC-05 result recorded with reproducible evidence and no unapproved waiver.

## 7. Required written answers from vendor
Ask the vendor to answer each point in writing and attach authoritative documentation:
1. Which exact workflow and product capability will be used to determine that a person meets Pinchd's minimum age? What does the result certify, and what does it not certify?
2. What exact fields, images, video, device signals, derived scores and biometric data are collected, transmitted, stored and returned in this configuration?
3. Which legal entity acts as processor/controller for each purpose, and which subprocessors and countries can access or process the data?
4. What are the retention periods for each data category, including backups and support logs? Can each be configured? What is the deletion evidence?
5. What are the documented false acceptance, false rejection and inconclusive rates for the actual age threshold/workflow, and what are the study limitations?
6. How are disabled users, unsupported devices, poor connectivity and false rejections accommodated? What accessible alternative and appeal route is available?
7. How are webhooks authenticated, replay-protected, ordered, retried and bound to a workflow/account? Provide sample signed test events and rotation procedures.
8. What happens to in-flight checks, callbacks and existing results during outage, configuration change, termination or account deletion?
9. Is any submitted data used for model training, product improvement, advertising, profiling or another secondary purpose? What contractual controls prevent this?
10. What are all per-check, retry, failure, manual review, support, storage and overage costs, and what minimum commitments apply?
11. What security assurance, penetration testing, vulnerability management and incident notification evidence can be provided?
12. What exact service/version change triggers a new assessment or requires Pinchd to re-run validation?

## 8. Decision ownership and release rule
- **Product owner:** confirms minimum age, intended user experience, fallback and user-facing wording.
- **Privacy/legal owner:** documents lawful basis, special-category condition where applicable, DPIA decision, privacy notice, transfer assessment, retention and deletion exceptions.
- **Vendor owner:** obtains contract, workflow/configuration evidence, subprocessors, regions, support and cost schedule.
- **Backend/security owner:** implements server-authoritative state, callback verification, access control, audit and idempotency.
- **Accessibility/T&S owner:** validates alternate paths, appeals, user safety and operational handling of false rejection/outage.
- **Independent reviewer:** confirms evidence is complete, test results are reproducible and residual risks are accepted by the correct accountable role.

**Release rule:** POC-05 is BLOCKED until the current vendor evidence is obtained and reviewed. It cannot be marked PASS based on a dashboard setting, liveness configuration, a vendor marketing statement or a successful client redirect. Production use additionally requires legal/privacy sign-off and an end-to-end test record. This pack does not assert that Didit currently meets any requirement listed here.
