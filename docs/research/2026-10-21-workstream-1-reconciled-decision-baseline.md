# Workstream 1 — reconciled decision baseline
**Date:** 21 October 2026  
**Status:** Working baseline; product-owner approval still required for proposed items.  
**Scope:** Research and requirements reconciliation only. No application code or mobile UI changed.

## Purpose
Use this alongside `2026-10-21-research-gap-audit-and-product-decision-register.md`. The audit is complete, but the product decision register is only complete when each item has an explicit status. Research recommendations are not automatically approved product decisions.

## Confirmed direction
- Product positioning: “Serious about compatibility, flexible about relationship goals.”
- Research and requirements must be completed before mobile UI design begins.
- Relationship intentions remain distinct and non-judgmental: long-term, short-term, casual, open-minded/exploring and not-sure-yet.
- Hard constraints are applied before ranking and are never silently relaxed. Soft changes require explicit user choice.
- User-selected age range is a hard discovery boundary; legal minimum-age eligibility is separate.
- User-selected distance radius is hard by default; location display is approximate and must not enable triangulation.
- Gender identity and discovery preferences are distinct; selected discovery categories are hard constraints. Implementation requires privacy, accessibility and legal review.
- Existing children and future family plans are distinct; relevant preferences are user-controlled. Lifestyle preferences default soft with an explicit hard option.
- Sensitive beliefs and traits are optional/minimised and not used for matching without documented purpose and review.
- Matching is reciprocal and explainable. Safety risk is not romantic desirability.
- Blocks, reports and essential safety controls remain free. Payment cannot bypass consent, blocks or hard preferences.
- No fake profiles/activity, fabricated likes, deceptive scarcity, hidden pay-to-win compatibility or promises of romantic outcomes.
- No free trial at launch.
- Treat the compliance requirements the owner confirmed in Didit as workflow-planning requirements. Do not remove required checks merely to reduce cost or clear a dashboard indicator.
- Optional profile verification and age assurance are separate. A badge must describe only checks actually performed. Didit liveness is configured; profile-photo matching and end-to-end integration are not complete.

## Proposed or unresolved — not approved by this register
| Item | Status | What must happen before commitment |
|---|---|---|
| Five unblurred incoming likes for free users | Hypothesis | Validate trust, fairness, usability and conversion; do not call five optimal. |
| Full unblurred incoming-like queue for Premium | Proposed | Confirm the free experience remains genuinely useful and transparent. |
| One Premium tier and monthly-first billing | Proposed | Compare against a free-only beta; model costs and storefront rules. |
| Commented likes get modest ordering advantage | Proposed | Assess abuse, fairness and recipient agency; no payment-based advantage. |
| Single-profile discovery feed | Undecided | Compare scanability, accessibility and alternatives before design. |
| Voice, polls, AI coaching, saved lists and travel mode | Later candidates | Keep outside MVP unless approved based on user value and cost. |
| Launch city, radius defaults, subscription price | Undecided | Use density, acquisition, cost and willingness-to-pay evidence. |
| Matching weights, thresholds and exploration | Undecided | Specify a transparent baseline, offline tests and fairness guardrails. |
| Stack, hosting and providers | Proposed | Complete proof-of-concept, privacy/security and vendor reviews. |
| Permanent pair exclusion after unmatch and 30-day user-facing history | Documented requirement; confirm before implementation | Test account recreation, reporting, privacy and retention edge cases. |

## Working MVP proposal
The MVP candidate includes account creation/recovery; age eligibility and an approved assurance approach; profiles, photos and text prompts; relationship intentions and preferences; eligibility-filtered discovery; reciprocal recommendations; likes and optional contextual comments; mutual matching and messaging; report, block, unmatch, moderation and appeals; notification controls and discovery pause; account deletion/data-rights workflows; restricted staff tooling, audit logs, monitoring, backups and minimal analytics.

Subscription implementation should wait until free/Premium boundaries, pricing and payment/store approach are approved. Do not assume voice, polls, AI coaching, saved collections or a single-profile feed are MVP requirements. No UI design in this workstream.

## Contradictions to resolve
1. **Free core vs blurred incoming likes:** define a transparent boundary that preserves meaningful free connection and never manufactures scarcity.
2. **Unmatch history vs deletion:** distinguish user-visible 30-day history from evidence retention, legal holds and deletion obligations.
3. **Verification vs age assurance:** liveness alone does not establish age or match a selfie to a profile photo.
4. **Unknown preference values:** specify per-field behavior; missing or undisclosed must not silently count as compatible with a hard rule.
5. **Policy vs algorithm:** the relationship-goal matrix is a policy baseline, not validated ranking weights.
6. **Draft-status drift:** requirements, preference policy and ADRs have different statuses; this document and the audit register should be linked from each relevant source, without pretending all details are approved.

## Required status vocabulary
- **Confirmed:** explicitly approved by the product owner.
- **Proposed:** recommended but not approved.
- **Hypothesis:** requires evidence or testing.
- **Legal review required:** legal applicability or interpretation is not settled.
- **Vendor confirmation required:** capability, contract or workflow details need supplier evidence.
- **Deferred:** intentionally out of current scope.
- **Blocked:** cannot proceed until a named dependency is resolved.

## Reconciliation protocol
1. Preserve confirmed decisions; do not ask the owner to decide them again.
2. For every proposed item, record rationale, trade-offs, acceptance criteria and validation plan.
3. Ask only for genuine product choices that research cannot settle.
4. Keep legal and vendor-dependent items visible rather than resolving them by assumption.
5. Update source documents only when status is clear, with a link to this baseline.
6. Close Workstream 1 when MVP scope and exclusions are approved, contradictions are resolved, and critical unknowns have an owner and validation gate.

## Readiness outcome
This baseline consolidates the current state but does not itself constitute final MVP approval. The next workstream is legal and regulatory readiness after Workstream 1 decisions are closed. Mobile UI design remains paused.


## Approval recorded — 2026-10-22

The user explicitly approved all previously proposed/unresolved Workstream 1 product decisions. They are now approved product direction, with empirical hypotheses, legal checks and implementation gates preserved where stated. Do not treat the earlier “approval pending” language in this file as current.

The user also approved the engineering direction recorded in ADR-0001: Go backend; React Native + Expo + TypeScript mobile; AWS; PostgreSQL; private S3 media; managed queues/workers when useful; production-grade CI/CD, tests, infrastructure as code, secret management, monitoring, budget alerts and recovery plans. Use a single primary AWS region initially, with Multi-AZ managed database resilience in production where feasible; multi-region database deployment is a future option, not a launch requirement. Exact service choices and versions remain subject to technical proof-of-concept, security/privacy/legal review and current cost validation.

Cost figures discussed are planning estimates only: development/testing USD $25–$100/month; lean production USD $150–$500+/month; growth USD $500–$2,000+/month. Excludes labour, legal/compliance, age-assurance charges, paid vendors and app-store fees. Do not present these as quotes or guaranteed costs.

Next workstream: UK legal and regulatory readiness, recorded in docs/research/2026-10-22-workstream-2-uk-legal-regulatory-readiness.md. Mobile UI design remains deferred until research and requirements are concluded.
