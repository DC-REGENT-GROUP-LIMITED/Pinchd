# Pinchd — Full-fold research status and launch-readiness register

- **As of:** 2026-10-22
- **Purpose:** Single navigation point for research completion, evidence needed and launch gates. Not a substitute for the detailed workstream documents.
- **Decision rule:** Approved product direction is not the same as empirical validation, technical proof, vendor confirmation or legal sign-off.

## Locked product and engineering direction

- Product positioning: “Serious about compatibility, flexible about relationship goals.”
- Workstream 1 product decisions are approved; see `2026-10-21-workstream-1-reconciled-decision-baseline.md` and the approval update at the end of that file.
- Stack: Go backend; React Native + Expo + TypeScript; AWS; PostgreSQL; private S3 media; managed asynchronous workers/queues when justified; role-restricted admin; CI/CD, automated tests, infrastructure as code, secrets management, monitoring, cost alerts, backup/restore and incident response.
- Launch infrastructure: one primary region; Multi-AZ managed PostgreSQL in production where feasible. Multi-region database deferred.
- Cost ranges are planning estimates only: development/testing $25–$100/month; lean production $150–$500+/month; growth $500–$2,000+/month. Exclusions: labour, legal/compliance, age assurance, paid vendors, app-store fees and unusual usage. Validate a current regional quote before provisioning.
- Mobile UI design remains paused.

## Workstream status

| Workstream | Current evidence | Status / exit gate |
|---|---|---|
| 1. Product decisions and MVP direction | Reconciled baseline and explicit user approval | Product direction approved; validate hypotheses and preserve legal/technical gates |
| 2. UK legal and regulatory readiness | `2026-10-22-workstream-2-uk-legal-regulatory-readiness.md` | Desk research completed; counsel/applicability and operational evidence still required |
| 3. Trust & safety operations | `2026-10-21-trust-safety-operations-incident-response.md` | Research exists; choose coverage model, staffing, severity targets, escalation ownership and prove report/block flows |
| 4. Lifecycle and edge cases | `2026-10-22-workstream-4-lifecycle-and-edge-case-requirements.md`; `2026-10-22-full-fold-lifecycle-privacy-architecture-reconciliation.md` | Lifecycle state dimensions and recommended defaults reconciled; pause/pass/expiry remain proposed defaults; deletion/retention and post-deletion pair suppression remain open gates |
| 5. Data/API/event contracts | `2026-10-22-data-classification-and-retention-map.md`; `2026-10-22-full-fold-lifecycle-privacy-architecture-reconciliation.md`; `../architecture/2026-10-22-postgresql-schema-blueprint.md` | Reconciliation drafted across lifecycle, privacy, schema/API, transactions and events; retention schedule and post-deletion suppression remain legal/privacy blockers; PoCs unexecuted |
| 6. Matching specification and fairness | `2026-10-21-matching-algorithms-fairness-explainability.md`; `2026-10-22-matching-engine-specification-and-fairness-tests.md` | Policy translation drafted: eligibility pseudocode, ranking boundaries, unknown handling, 24 synthetic tests, metrics and launch gates; weights remain unapproved |
| 7. Launch market and growth | `2026-10-21-launch-strategy-local-density-cold-start-economics.md`; `2026-10-22-launch-market-selection-scorecard-and-validation-plan.md` | Market scorecard and capped validation plan drafted; shortlist, actual channel tests, eligible-supply evidence and budget remain open |
| 8. Monetisation and unit economics | `2026-10-21-monetisation-subscription-economics-ethical-pricing.md`; `2026-10-22-premium-pricing-and-unit-economics-model.md` | Ethical boundaries and unit-economics model drafted with current official fee/VAT sources; exact price, quotes, retention evidence and legal/store review remain open |
| 9. Competitor and differentiation | `2026-10-22-competitor-landscape-and-differentiation-review.md`; Hinge research notes | Public-source comparison and differentiation thesis drafted; user-demand validation and hands-on UK feature/price verification remain open |
| 10. Accessibility and inclusive design | `2026-10-22-accessibility-and-inclusive-product-acceptance-criteria.md` | Accessibility acceptance baseline and device/assistive-tech test matrix drafted; legal applicability, disabled-user testing and final standard scope remain open |
| 11. Vendor and technical proof-of-concept | ADR-0001/0002; `../architecture/2026-10-22-technical-proof-of-concept-plan.md`; `2026-10-22-didit-age-assurance-vendor-due-diligence-pack.md`; `2026-10-22-requirements-to-test-traceability-matrix.md`; `2026-10-22-local-fixture-harness-test-record.md`; synthetic fixture and isolated environment specifications | Didit evidence request pack, isolated PoC runbook and LC-01–LC-36 / POC-01–POC-12 traceability drafted; fixture JSON/reference validation and six local Go harness tests passed, but this is fixture-validation only; PostgreSQL 17.10 is installed/running but not approved for PoC use, psql is not on PATH, isolated database/user and environment controls remain open; all product PoCs remain NOT RUN, vendor evidence/measured costs and legal/privacy gates remain open |
| 12. Design readiness | All above; `2026-10-22-lifecycle-defaults-and-privacy-decision-log.md` | Product defaults still need explicit confirmation; legal/privacy/vendor gates remain open; only begin mobile UI after research, evidence, test and design-readiness gates are signed off |

## Critical legal evidence pack

Before public registration, maintain one traceable folder/register containing:

1. Written Online Safety Act scope determination and current Ofcom tool outputs.
2. Applicable illegal-content risk assessment, child-access assessment and required policy/recordkeeping materials.
3. Data inventory, flow map, lawful-basis register (Article 6 and Article 9 where applicable), DPIA decision/report and retention schedule.
4. Didit/vendor due diligence, age assurance evidence, data minimisation and failure/appeal test results.
5. Moderation/reporting/complaints, urgent-harm, evidence preservation, appeals and incident runbooks; named owner and capacity model.
6. Subscription and consumer-law commencement review; store billing/cancellation/refund/restore requirements.
7. Equality/accessibility review, including age-check false rejects and relevant fairness risks.
8. Vendor/subprocessor register, processor contracts and transfer assessment where required.
9. Security threat model, access model, penetration/security test report, backup restoration evidence and breach-response exercise.
10. Approved terms, privacy notice, community standards, support contacts, marketing claim substantiation and deletion flow.

## Launch gates and stop conditions

Do not open public registration if any of these is true:
- The service's OSA scope or required assessments are unknown and not owned by counsel/compliance.
- Age assurance can be trivially bypassed, the vendor data flow is unknown, or failure/appeal routes are absent.
- Reporting, blocking or urgent-harm escalation is not operationally staffed and tested.
- Blocks can be bypassed via caches, queued notifications, messaging, discovery or account lifecycle.
- Sensitive data is used without documented purpose and legal basis, or the DPIA/legal gate is unresolved for high-risk processing.
- Account deletion has no tested propagation/retention plan.
- Paid entitlements can be client-forged or subscription terms/cancellation have not been legally/store reviewed.
- No monitoring, incident ownership, recovery plan, budget alerts or tested backup restore exists.
- Accessibility and safety-critical flows have not been tested.

## Research quality rules

- Cite primary legislation/regulators/official provider docs; record publication/update dates and applicability limits.
- Separate sourced fact, inference, product choice, hypothesis, legal review and vendor confirmation.
- Record negative evidence and uncertainties; never call a vendor integration complete based only on dashboard configuration.
- No mobile UI design until requirements and research gates are met.
