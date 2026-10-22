# Research index

Keep dated research notes with:
- Research date and scope.
- Source URL, publisher and publication/update date.
- Key finding and why it matters for Pinchd.
- Confidence/limitations.
- Product or engineering implication.
- Follow-up owner and validation date.

See `../../research/2026-folder-architecture-research.md`.

- `2026-10-21-messaging-and-conversation-design.md` — evidence review, proposed messaging requirements, safety/privacy considerations, post-launch metrics and experiment hypotheses.
- `2026-10-21-matching-algorithms-fairness-explainability.md` — reciprocal matching, preference integrity, exposure fairness, explainability, cold-start handling and post-launch metrics.
- `2026-10-21-privacy-uk-gdpr-age-assurance-data-governance.md` — UK GDPR, sensitive data, privacy-by-design, DPIA, age assurance, biometrics, vendor governance, retention and launch gates.
- `2026-10-21-trust-safety-operations-incident-response.md` — reporting, triage, severity targets, incident response, appeals, moderation capacity, ban evasion and launch gates.
- `2026-10-21-launch-strategy-local-density-cold-start-economics.md` — focused market selection, waitlist validation, eligible-supply metrics, staged beta, launch gates and unit economics.
- `2026-10-21-monetisation-subscription-economics-ethical-pricing.md` — free/premium boundaries, subscription lifecycle, ethical pricing, UK legal/platform checks and unit economics.

- `2026-10-21-hinge-redesign-feature-review.md` — assessment of the Hinge “Making Confident Moves” designer case study: private saved profiles, lightweight music prompts, and why public reviews are not recommended.
- `2026-10-21-appthetics-hinge-ui-patterns-review.md` — preliminary review of Hinge-style UI patterns; the supplied Appthetics URL returned 404, so exact article-specific claims remain unverified.

- `2026-10-21-research-gap-audit-and-product-decision-register.md` — audit of existing research coverage, decision-status vocabulary, reconciled open decisions, prioritised remaining workstreams and research completion gates. Start here before beginning Workstream 1.

- `2026-10-21-workstream-1-reconciled-decision-baseline.md` — reconciled confirmed principles, proposed/hypothesis items, working MVP scope, conflicts and status vocabulary for Workstream 1.

- 2026-10-22-workstream-2-uk-legal-regulatory-readiness.md — UK-first legal/regulatory readiness desk review covering Online Safety Act, UK GDPR/DPA, age assurance and Children's Code, subscriptions/consumer law, equality/accessibility, vendor governance and launch gates.

- `2026-10-22-full-fold-research-status-and-launch-readiness.md` — current research workstream status, locked direction, critical compliance evidence pack, stop conditions and design-readiness gate.
- `2026-10-22-local-fixture-harness-test-record.md` — dated result for six local Go fixture-validation tests, vet and whitespace checks, with explicit limits and next environment gate.
- `2026-10-22-poc01-postgres-baseline-record.md` — dedicated loopback-only PostgreSQL cluster, least-privilege role, transaction/constraint/rollback/restart checks and explicit limitations; partial POC-01 evidence, not full completion.
- `2026-10-22-poc01-go-api-baseline-record.md` — minimal Go API liveness/readiness endpoints, configuration/error handling, unit/vet checks, isolated database integration and loopback HTTP smoke results; conditional evidence only, not production readiness.

- `2026-10-22-workstream-3-operating-readiness.md` — trust-and-safety operating model, severity targets, case lifecycle, evidence controls, staffing/capacity, metrics, incident stop conditions and pre-launch acceptance checklist.
- `2026-10-22-workstream-4-lifecycle-and-edge-case-requirements.md` — testable lifecycle requirements and edge cases for onboarding, profiles, discovery, likes, matching, messaging, pause, block, unmatch, deletion, appeals and failure/race conditions.
- `2026-10-22-workstream-5-open-architecture-decisions.md` — proposed conservative defaults and evidence/legal gates for pause, pass, like expiry, messaging retention, pair exclusion after deletion, age-assurance outages, subscription entitlement refresh and recovery objectives.
- `2026-10-22-data-classification-and-retention-map.md` — working data inventory, sensitivity classes, access boundaries, deletion orchestration, vendor evidence and retention decisions requiring legal/privacy review.
- `2026-10-22-matching-engine-specification-and-fairness-tests.md` — draft hard-eligibility pseudocode, soft ranking inputs, unknown-value handling, tie-breaks, exploration, explanations and synthetic acceptance tests.
- `2026-10-22-launch-market-selection-scorecard-and-validation-plan.md` — weighted candidate-market scorecard, evidence confidence ladder, bottom-up activation model, capped validation experiments and GO/HOLD/STOP gates.
- `2026-10-22-competitor-landscape-and-differentiation-review.md` — official-source review of Hinge, Bumble, Tinder and category/App Store constraints, with Pinchd differentiation thesis and validation plan.
- `2026-10-22-accessibility-and-inclusive-product-acceptance-criteria.md` — WCAG-informed and platform-specific accessibility requirements, inclusive product principles, test matrix and release gates.
- `2026-10-22-premium-pricing-and-unit-economics-model.md` — ethical free/Premium boundaries, sourced platform-fee/VAT assumptions, contribution formulas, scenario inputs and pricing launch gates.
- `2026-10-22-full-fold-lifecycle-privacy-architecture-reconciliation.md` — consolidated state model, lifecycle defaults, privacy/deletion gates, schema/API/event rules, transaction boundaries, acceptance gates and decision owners.
- `2026-10-22-requirements-to-test-traceability-matrix.md` — LC-01–LC-36 lifecycle acceptance matrix, POC-01–POC-12 verification criteria, evidence rules, owner roles and release gates; all tests currently NOT RUN.
- `2026-10-22-didit-age-assurance-vendor-due-diligence-pack.md` — vendor evidence request, data-flow/retention/security questions, callback requirements, PoC tests and age-assurance release checklist; vendor facts remain unverified until evidence is received.
- `2026-10-22-lifecycle-defaults-and-privacy-decision-log.md` — proposed lifecycle defaults requiring product confirmation, approved policy baseline, privacy/legal blockers and closure workflow.
- `2026-10-22-full-fold-execution-and-evidence-register.md` — prioritised action register ACT-01–ACT-25, evidence register EV-01–EV-17, dependencies, accountable roles and execution guardrails; no named owners or passes are invented.
- `../architecture/2026-10-22-postgresql-schema-blueprint.md` — draft PostgreSQL domain schema, constraints/indexes, API route matrix, transaction boundaries, outbox events, lifecycle/eligibility model and unresolved implementation gates.
- `../architecture/2026-10-22-technical-proof-of-concept-plan.md` — staged Go/PostgreSQL, age assurance, block/unmatch, outbox, media, push, billing, deletion, accessibility and cost validation plan.
- `../../tests/poc-harness/README.md` — local standard-library Go checks for synthetic fixture safety, references and selected policy invariants; not production logic or POC completion evidence.
- `../architecture/2026-10-22-poc-execution-runbook.md` — ordered local-only PoC procedure, evidence rules, stop conditions and explicit separation between mocks and real vendor tests.
- `../../tests/fixtures/poc/core-scenarios.json` — deterministic synthetic-only account, profile, relationship, callback and scenario fixtures; JSON syntax validated, privacy/engineering review and behavioral tests pending.
- `2026-10-22-counsel-legal-briefing-pack.md` — draft counsel instruction pack with focused UK legal/privacy questions, evidence checklist and decision log; no legal advice or sign-off is represented.
- `2026-10-22-synthetic-test-data-and-fixture-specification.md` — proposed synthetic-only fixture catalogue and invariants for POC-01–POC-12; initial deterministic JSON fixtures exist, with review and behavioral tests pending.
- `2026-10-22-isolated-poc-environment-specification.md` — proposed environment separation, security, spend, execution, stop and evidence controls; a dedicated local cluster is provisioned, but engineering/security/privacy review and environment approval remain open.