# Pinchd research gap audit and product decision register
**Audit date:** 21 October 2026  
**Scope:** Read-only audit of the current documentation tree; no application code or mobile UI changed.  
**Status:** Audit baseline — use as the working register for Research Workstream 1.

## Executive finding
Pinchd has a strong first-pass research foundation across matching, messaging, privacy, trust and safety, local launch economics, monetisation, and Hinge-inspired feature review. The main gap is not absence of research everywhere; it is that research notes, requirements, policy decisions, architecture proposals, and external legal obligations are not yet reconciled into one approved decision set.

The consolidated requirements remain a **draft**. The stack and deployment ADRs remain **proposed / decision required**. The preference-policy log is **in progress**. Mobile UI design remains intentionally paused. No production implementation should be inferred from the repository scaffold or from research recommendations.

## 1. What the current documents cover
| Area | Existing evidence in repository | Audit assessment |
|---|---|---|
| Product direction and MVP | `product/consolidated-requirements-v1.md` | Strong draft baseline; not final approved scope. |
| Preference integrity and relationship goals | `product/preference-policy-working-decisions.md`; consolidated requirements | Many policies explicitly decided; remaining category and unknown-value details need a final register. |
| Matching and fairness | `research/2026-10-21-matching-algorithms-fairness-explainability.md` | Good conceptual foundation; algorithm, thresholds, exposure policy and measurement remain to be operationalised. |
| Messaging and conversation design | `research/2026-10-21-messaging-and-conversation-design.md` | Strong direction; exact product limits, delivery states, anti-abuse thresholds and support operations need decisions. |
| Privacy, age assurance and biometrics | `research/2026-10-21-privacy-uk-gdpr-age-assurance-data-governance.md`; `security/privacy-foundations.md` | Broad desk research; launch-specific legal assessment, DPIA decisions, retention map and supplier approval remain open. |
| Trust and safety | `research/2026-10-21-trust-safety-operations-incident-response.md`; `trust-and-safety/operating-model.md` | Broad policy and threat coverage; operational capacity, rota, severity SLAs and escalation ownership are not approved. |
| Launch strategy and local density | `research/2026-10-21-launch-strategy-local-density-cold-start-economics.md` | Decision framework exists; no launch market, budget, density gate or acquisition plan chosen. |
| Monetisation | `research/2026-10-21-monetisation-subscription-economics-ethical-pricing.md` | Ethical boundaries and candidate features researched; price, entitlement, storefront and forecast assumptions remain unvalidated. |
| Hinge/product pattern reviews | `research/2026-10-21-hinge-redesign-feature-review.md`; `research/2026-10-21-appthetics-hinge-ui-patterns-review.md` | Useful inspiration; not all concepts are validated or decided. Appthetics source limitations are already noted. |
| Architecture and feasibility | `architecture/decisions/ADR-0001-stack-and-deployment.md`; `ADR-0002-technical-feasibility.md` | Detailed provisional architecture and proof gates; stack, providers, deployment and proof-of-concept results still open. |
| Data/API/operations | `data/README.md`, `api/README.md`, `operations/README.md` | These are checklists/conventions, not completed data dictionary, API contracts, retention schedule or runbooks. |
| Verification provider integration | Requirements sections 16–18 and current Didit workflow work | Liveness workflow is configured externally; selfie-to-profile-photo comparison, backend validation, data flow, legal review and test evidence remain open. |

## 2. Decision status vocabulary
Use the following status labels in future decisions:
- **Confirmed:** explicitly agreed by the product owner; changes require a recorded revision.
- **Proposed:** recommended by research, not yet approved.
- **Hypothesis:** needs validation with public evidence, prototype testing, or controlled beta measurement.
- **Legal review required:** do not treat as a settled legal conclusion.
- **Vendor confirmation required:** dependent on documented provider capability, contract or support response.
- **Deferred:** intentionally excluded from current scope.
- **Blocked:** cannot safely progress until a named dependency is resolved.

Do not promote a research recommendation to “confirmed” just because it appears in a draft requirements document.

## 3. Confirmed product principles to preserve during reconciliation
- Positioning: “Serious about compatibility, flexible about relationship goals.”
- Adult audience; support varied relationship intentions without morally ranking them.
- Hard constraints precede ranking; never silently relax them.
- Age range is a hard discovery boundary, separate from legal age eligibility.
- Distance is hard by default, with explicit user-controlled changes and approximate displayed location.
- Reciprocal recommendations and honest explanations; unknown values are not silently assumed compatible.
- Blocks, safety restrictions, legal eligibility and permanent unmatched-pair exclusions cannot be overridden by ranking or payment.
- Safety risk is separate from romantic desirability.
- Essential matching, mutual messaging, blocking, reporting, privacy controls and legally required age assurance remain free.
- No fake profiles/activity, deceptive scarcity, hidden pay-to-win compatibility or outcome promises.
- Optional selfie verification and legal age assurance are distinct; badges must describe only the check actually performed.
- Mobile UI design remains paused until the research and requirements are reconciled and approved.

## 4. Open decisions that must be reconciled first
### Product scope
1. Freeze the MVP and explicitly list post-MVP candidates.
2. Decide whether private saved profiles are MVP, beta or later; avoid artificial expiry/urgency.
3. Decide initial discovery presentation only at the product-principle level (no screen design yet); preserve an accessible alternative if a single-profile feed is selected.
4. Decide profile minimums, prompt taxonomy, photo limits, media types and profile visibility states.
5. Define onboarding fields, skip/decline behaviour, editability and when a profile becomes discoverable.
6. Finalise likes, comment-on-prompt/photo, pass, undo (if any), incoming-like queue, match and conversation limits.
7. Confirm notification categories, quiet hours, pause/discovery pause, read receipts/typing indicators and per-conversation controls.
8. Define account lifecycle: deactivate, pause, suspend, appeal, delete, reactivation and permanent pair exclusions.

### Preference and matching policy
9. Reconcile the preference-policy log with the consolidated requirements and mark each row confirmed/proposed/legal review.
10. Define behaviour for unknown/missing/undisclosed values for every hard field.
11. Define candidate-generation and ranking baseline, tie-breaking, exploration/diversity, reciprocal compatibility estimation and cold-start handling.
12. Define fair-exposure metrics and guardrails without turning fairness into a desirability score.
13. Define the exact scope and limits of recommendation explanations.
14. Specify cache invalidation, concurrent preference changes and how eligibility is rechecked across discovery, incoming likes, notification queues and match creation.

### Trust, privacy and compliance
15. Maintain a jurisdiction-specific obligation register based on the product's actual service and launch plan; preserve Didit's confirmed findings as the current workflow-planning input while legal applicability and interpretation are documented by qualified review where needed.
16. Complete the age-assurance decision separately from optional profile verification.
17. Confirm the supported Didit flow for profile-photo matching; document API inputs, score semantics, thresholds, webhooks, retries, retention and failure handling before integration.
18. Decide which verification results and evidence are stored by Pinchd, for how long, and which roles can access them.
19. Complete the data inventory, purpose/lawful-basis map, special-category analysis where applicable, DPIA decision, retention schedule and deletion propagation plan.
20. Define moderation coverage, support capacity, severity response targets, emergency escalation, appeals and vendor outage fallback.
21. Complete a threat model covering stalking/location inference, scraping, account takeover, scams, ban evasion, abusive reports, insider access and sensitive-data exposure.
22. Approve community standards, terms, privacy notice, age/verification explanations, report/appeal messaging and support/refund policy.

### Commercial, launch and architecture
23. Select a launch market using eligible active supply, acquisition reach, operating capacity and cost—not population alone.
24. Define waitlist validation and launch experiments, success metrics, guardrails, stop criteria and expansion gates.
25. Decide initial subscription price and feature entitlements only after willingness-to-pay and cost modelling; no free trial remains the stated decision.
26. Model conservative/base/upside unit economics including acquisition, app-store fees, taxes, verification, moderation, support, infrastructure, refunds and churn.
27. Approve stack ADR after small proof-of-concept tests and team capability review.
28. Choose hosting, authentication, media storage, notification, analytics, billing and observability providers after privacy/security/exit review.
29. Create actual data dictionary/ERD, API contracts, event taxonomy, test plan, release plan and operational runbooks.
30. Define accessibility acceptance criteria and test device/platform coverage before UI design begins.

## 5. Prioritised research completion plan
| Priority | Work package | Output / exit criterion |
|---|---|---|
| P0 | Reconcile research and product decisions | One decision register; conflicts and duplicates resolved; MVP and explicit exclusions agreed. |
| P0 | Legal/privacy launch gates | Obligation register, age-assurance decision path, DPIA determination, data-flow/retention requirements and unresolved counsel questions. |
| P0 | Safety operations readiness | Launch staffing/capacity model, severity targets, escalation ownership, report/appeal procedures and incident drills. |
| P1 | Product lifecycle and edge cases | Complete state-transition/acceptance criteria for onboarding, profiles, discovery, likes, matches, messaging, pause, blocks, unmatch and deletion. |
| P1 | Data and API design | Data dictionary, ERD, field sensitivity, API/event contracts, retention/deletion map. |
| P1 | Matching specification | Eligibility gates, reciprocal ranking baseline, explanations, unknown-value behaviour, fairness metrics and synthetic test cases. |
| P1 | Launch market and growth model | Evidence-based market scorecard, opt-in demand validation, density targets and acquisition plan. |
| P1 | Unit economics and subscription | Transparent price scenarios and cost model; free/Premium matrix; storefront/legal review. |
| P2 | Competitor/differentiation refresh | Source-backed comparison of Hinge, Bumble, Tinder and relevant alternatives; distinguish observed features from marketing claims. |
| P2 | Accessibility and inclusive design research | Accessibility requirements for onboarding, profiles, messaging, location, verification and support. |
| P2 | Support and customer operations | Support channel, refunds, complaints, data requests, response expectations and escalation runbooks. |
| P2 | Architecture proof of concept | Evidence for mobile framework, provider SDKs, media, push, blocking, deletion, billing and matching performance. |
| Gate | Design-readiness review | All critical decisions approved or explicitly risk-accepted; no hidden critical safety/privacy blocker. Then mobile UI design may begin. |

## 6. Research quality and source hygiene
- Every material external claim should identify source, publication/update date, scope, limitations and the specific Pinchd implication.
- Label company/vendor claims separately from independent evidence and user anecdotes.
- Reddit and app-store reviews are qualitative, self-selected signals; use them to generate hypotheses, not estimate population prevalence.
- Check the live legal text, regulator guidance and effective dates at the time of each launch decision. A compliance tool's matched requirements are important inputs, but keep the underlying evidence and the organisation's applicability determination traceable.
- Record conflicting evidence rather than silently selecting the convenient source.
- Prefer primary regulator, legislation, peer-reviewed research and official product documentation; use secondary reporting for context.
- Do not create research for a feature already explicitly deferred unless a new decision depends on it.

## 7. Audit limitations
This audit is based on the repository documents and their recorded status as read on 21 October 2026. It does not prove that external legal obligations have been satisfied, that a provider configuration is production-ready, that user demand has been validated, or that technical proof-of-concept tests have passed. Those require their own evidence and sign-off.

## 8. Next action
Workstream 1 is to use this register to reconcile each existing product decision into **Confirmed / Proposed / Hypothesis / Legal review required / Vendor confirmation required / Deferred / Blocked**, remove contradictions between documents, and produce a clean MVP decision baseline. Do not design screens or start application implementation during this workstream.


## Status update — 2026-10-22

This audit's original sections 1 and 8 describe the state on 2026-10-21 and are historical. Subsequent decisions supersede its original “Workstream 1 approval pending” and “next action: Workstream 1” language.

- Workstream 1: the user explicitly approved all previously proposed/unresolved product decisions on 2026-10-21/22. The approval was recorded in `2026-10-21-workstream-1-reconciled-decision-baseline.md`. Hypotheses still require validation; legal/vendor/technical gates remain gates rather than being magically resolved by product approval.
- Architecture direction: approved and recorded in `../architecture/decisions/ADR-0001-stack-and-deployment.md`: Go backend; React Native + Expo + TypeScript; AWS; PostgreSQL; private S3 media; managed queues/workers where justified; production-grade delivery, security, observability and recovery controls. Multi-region database is deferred pending evidence and budget.
- Workstream 2: first-pass UK legal/regulatory readiness desk review is recorded in `2026-10-22-workstream-2-uk-legal-regulatory-readiness.md`, including an additional source-verified addendum. It is not legal sign-off.
- Current next step: convert Workstream 2 into an owned compliance evidence checklist and then proceed to remaining research gaps in sequence: trust/safety operational capacity; lifecycle/edge-case requirements; data/API/event contracts; matching specification; launch-market and unit economics; accessibility; vendor proof-of-concept; final design-readiness review.
- Mobile UI design remains deferred until the research and requirements readiness gate is passed.
