# Pinchd — Full-Fold Execution and Evidence Register
**Date:** 22 October 2026  
**Status:** Working execution tracker. Owners below are accountable roles, not named people. No evidence is assumed to exist unless linked or recorded.  
**Purpose:** Turn the research pack into a controlled sequence of decisions, evidence collection, tests and launch gates without prematurely starting mobile UI design or production implementation.

## 1. Executive position
Research artifacts now cover product policy, UK legal/regulatory readiness, trust and safety, lifecycle, data classification, schema/API concepts, matching, launch-market selection, monetisation, competitors, accessibility, technical PoCs, traceability and Didit due diligence.

The project is **not yet implementation-ready** for production. The immediate blockers are decisions and evidence that cannot be inferred from the documents:
1. Product confirmation of proposed lifecycle defaults.
2. Qualified privacy/legal review of data purposes, sensitive data, retention/deletion, pair suppression and regulatory scope; a counsel briefing pack is now drafted, but advice/sign-off is still outstanding.
3. Current Didit contract/workflow/data-flow evidence and independent integration tests.
4. Trust-and-safety staffing/coverage and escalation ownership.
5. Execution of technical PoCs and measured cost/recovery evidence.
6. Empirical launch-market and pricing validation.
7. Final signed design-readiness review.

Mobile UI design remains paused.

## 2. Status vocabulary
- **OPEN:** work not completed.
- **BLOCKED:** cannot responsibly proceed until a named dependency is resolved.
- **IN PROGRESS:** evidence collection or test is actively underway and has an owner.
- **READY FOR REVIEW:** complete evidence exists, awaiting accountable review.
- **CLOSED:** decision/evidence is documented and approved by the correct role.
- **NOT APPLICABLE:** rationale and approver documented.

A document being drafted is not evidence that a requirement is implemented or tested. No item may be closed by assumption or by a vendor dashboard alone.

## 3. Priority order

| Priority | Work package | Why this order | Exit evidence |
|---|---|---|---|
| P0 | Confirm product lifecycle defaults | Prevents schema/API rework and contradictory behavior | Signed decision record for pause, pass, like expiry, existing matches during restrictions and deletion UX |
| P0 | Privacy/legal scoping | Determines permitted data, retention, age assurance and launch obligations | Written counsel/privacy determinations and reviewed data map/DPIA decision |
| P0 | Didit evidence collection | Age eligibility is a launch-critical gate; integration claims are not yet proven | Vendor pack answered, workflow/version recorded, DPA/regions/retention reviewed |
| P0 | Trust-and-safety operating model | Safety flows need people, coverage, escalation and measurable service targets | Named accountable owner, coverage plan, escalation directory, tabletop evidence |
| P1 | Freeze test fixtures and test environment | Allows safe reproducible PoCs without real user data | Synthetic fixture plan, isolated credentials, spend caps, logging controls |
| P1 | Run POC-01 to POC-04 | Validates core API/database, concurrency, block/unmatch and event delivery | Dated evidence records with test outputs and independent review |
| P1 | Run POC-10 and POC-06 | Validates eligibility boundaries, query behavior and private media | Zero hard-gate violations; access/deletion tests and measured query costs |
| P1 | Run POC-05 and POC-09 | Vendor and deletion tests depend on privacy/vendor evidence | Callback tests and deletion propagation/restoration proof |
| P1 | Run POC-11 and POC-12 | Critical accessibility, recovery, observability and spend controls | Assistive-tech evidence, restore report, alert/cost report |
| P2 | Run POC-08 | Paid features depend on entitlement and legal/store review | Sandbox evidence, verified entitlement state, refund/cancel/restore tests |
| P2 | Market and price experiments | Avoid choosing a city or price based only on intuition | Documented shortlist, real acquisition tests, sensitivity model and decision |
| P3 | Design readiness review | UI should reflect verified requirements, not fill unresolved policy gaps | Signed gate review with no unowned critical blocker |

## 4. Action register

| ID | Action | Accountable role | Required output/evidence | Dependency | Status |
|---|---|---|---|---|---|
| ACT-01 | Confirm lifecycle defaults LD-01–LD-10 | Product owner | Dated accepted/rejected/deferred choices and user-facing behavior | Review decision log | OPEN |
| ACT-02 | Confirm minimum-age policy and age-assurance purpose | Product + privacy/legal | Written age policy, lawful purpose, approved workflow and fallback | UK regulatory review + Didit evidence | BLOCKED |
| ACT-03 | Complete OSA scope and child-access assessment | UK counsel/compliance | Written scope determination, current regulator tool outputs and required risk assessments | Final product/content scope | OPEN |
| ACT-04 | Complete data purpose/lawful-basis register | Privacy/legal | Purpose-by-purpose Article 6 basis and Article 9 condition/safeguards where applicable | Data map and actual feature scope | OPEN |
| ACT-05 | Complete DPIA screening and DPIA where required | Privacy lead + security | Screening rationale or approved DPIA, mitigations and residual-risk decisions | Data flow, matching/location/age-assurance details | OPEN |
| ACT-06 | Approve retention/deletion schedule | Privacy/legal + T&S | Retention trigger/period, exceptions, owner and deletion method per data class | Vendor terms, safety needs, legal review | BLOCKED |
| ACT-07 | Decide post-deletion pair suppression | Privacy/legal + product | Documented lawful basis, minimum token design, duration, disclosure and rights handling | Approved pair-exclusion policy and privacy review | BLOCKED |
| ACT-08 | Obtain Didit evidence | Vendor owner | Completed vendor request, contract/DPA, subprocessors, regions, workflow, retention, accuracy, callback docs | Vendor response | BLOCKED |
| ACT-09 | Confirm Didit callback/state design | Backend + security | Threat model, state machine, signature/replay/account binding tests | ACT-08 | BLOCKED |
| ACT-10 | Set T&S coverage model | T&S lead + operations | Hours, on-call, capacity model, escalation directory, urgent-case coverage and training | Risk assessment and budget | OPEN |
| ACT-11 | Approve synthetic test-data plan | Engineering + privacy | Fixture catalog with no real sensitive data, retention and cleanup | Test scope | READY FOR REVIEW |
| ACT-12 | Establish isolated PoC environment | Platform lead | Least-privilege credentials, spend caps, CI, safe logs and teardown | ACT-11 | IN PROGRESS — isolated local cluster and app role created; Engineering/Security/Privacy review, teardown demonstration and full environment approval remain open |
| ACT-13 | Execute POC-01 | Backend/platform | Reproducible Go/Postgres baseline, migration and transaction evidence | ACT-12; isolated database/user, network boundaries and teardown approval | IN PROGRESS — SQL transaction baseline plus minimal Go API unit/vet, isolated DB integration and loopback HTTP smoke checks passed; migrations/CI, broader API behavior, hardening and independent review remain open; see `2026-10-22-poc01-postgres-baseline-record.md` and `2026-10-22-poc01-go-api-baseline-record.md` |
| ACT-14 | Execute POC-02 and POC-04 | Backend/platform + QA | Concurrency/idempotency and outbox reliability results | POC-01 | NOT STARTED |
| ACT-15 | Execute POC-03 and POC-07 | Backend/messaging + T&S | Block/unmatch races and notification suppression evidence | POC-01, POC-04 | NOT STARTED |
| ACT-16 | Execute POC-10 | Matching/data + QA | Hard-gate property tests, query plan and privacy checks | Product policy confirmed | NOT STARTED |
| ACT-17 | Execute POC-06 | Media/platform + security | Private object, moderation, metadata stripping and deletion tests | Storage/media policy selected | NOT STARTED |
| ACT-18 | Execute POC-05 | Identity/backend + privacy/legal | End-to-end Didit callback, failure, appeal and deletion evidence | ACT-08, ACT-09 | BLOCKED |
| ACT-19 | Execute POC-09 | Privacy/platform + security | Durable deletion/retry/backup restore evidence | ACT-06, ACT-07 | BLOCKED |
| ACT-20 | Execute POC-11 | Accessibility + QA | VoiceOver/TalkBack, large text, keyboard/switch and vendor SDK test results | Critical flows defined | NOT STARTED |
| ACT-21 | Execute POC-12 | Platform/operations + finance | Cost attribution, alerts, backup restore and approved RPO/RTO evidence | Test environment, workload assumptions | NOT STARTED |
| ACT-22 | Execute POC-08 | Billing + backend + counsel | Store sandbox purchase/renewal/cancel/refund/restore results | Paid scope, platform accounts and legal review | BLOCKED |
| ACT-23 | Validate launch market | Product/growth + T&S | Evidence-backed 3–5 market shortlist and capped channel tests | Budget, eligible-supply assumptions | OPEN |
| ACT-24 | Validate Premium price and economics | Product/finance + counsel | Sensitivity analysis, price test plan, fee/tax and churn assumptions | Vendor/cloud cost inputs and legal review | OPEN |
| ACT-25 | Conduct design-readiness review | Product + engineering + privacy/legal + T&S | Signed gate review, remaining risks, accepted evidence and UI authorization | P0/P1 gates closed or explicitly scoped | BLOCKED |

## 5. Evidence register — required records
Create or link one record per evidence item. Keep source files in the repository with dates and owners; do not put secrets or raw identity/biometric data in the repository.

| Evidence ID | Evidence | Minimum acceptance | Current state |
|---|---|---|---|
| EV-01 | Product decision record | LD-01–LD-10 status, owner, rationale and acceptance tests | OPEN |
| EV-02 | OSA and child-access determination | Counsel/compliance sign-off and current official assessment outputs | OPEN |
| EV-03 | Data flow and data inventory | All purposes, fields, processors, regions, recipients and lifecycle paths mapped | DRAFT EXISTS; REVIEW OPEN |
| EV-04 | Lawful-basis and special-category register | Every purpose has documented legal assessment; no assumed basis | OPEN |
| EV-05 | DPIA screening/DPIA | Signed screening and DPIA if required, with residual risks | OPEN |
| EV-06 | Retention/deletion schedule | All classes, triggers, periods, exceptions, owners and proof of deletion | OPEN |
| EV-07 | Didit vendor evidence pack | Current contract, DPA, workflow, data, region, accuracy, retention, callback and support evidence | REQUEST PACK DRAFTED; VENDOR EVIDENCE OPEN |
| EV-08 | T&S operating plan | Coverage, staffing, urgent escalation, incident tabletop and capacity | RESEARCH EXISTS; OPERATING DECISIONS OPEN |
| EV-09 | Test environment controls | Synthetic fixtures, least privilege, spend caps, logging and cleanup | FIXTURE JSON/REFERENCE CHECK AND 6 LOCAL GO TESTS PASS; GO 1.25.6/GIT 2.52.0; DEDICATED POSTGRESQL 17.10 CLUSTER ON 127.0.0.1:55433 WITH SCRAM, ACL-RESTRICTED SECRETS AND LEAST-PRIVILEGE ROLE; ONLY THE NEW PoC DATABASE WAS ACCESSED/MODIFIED, EXISTING SERVICE/DATA UNTOUCHED; ENGINEERING/SECURITY/PRIVACY REVIEW, TEARDOWN DEMONSTRATION AND ENVIRONMENT APPROVAL OPEN |
| EV-10 | POC evidence set | All POC-01–POC-12 with actual outputs and independent review | POC-01 SQL transaction baseline plus minimal Go API/readiness smoke CONDITIONAL PASS (no independent review; migrations/CI/security hardening/full POC-01 incomplete); POC-02–POC-12 NOT RUN/BLOCKED per dependencies |
| EV-11 | Accessibility evidence | Critical flows tested on representative devices/assistive technology | CRITERIA EXIST; TESTING OPEN |
| EV-12 | Recovery report | Backup restore, deletion tombstone replay, approved RPO/RTO and evidence | OPEN |
| EV-13 | Cost evidence | Current regional quote, measured PoC cost, vendor charges and workload assumptions | OPEN |
| EV-14 | Market evidence | Eligible-supply estimate, channel tests, demand signals and shortlist rationale | OPEN |
| EV-15 | Pricing/unit economics evidence | Sensitivity analysis with verified platform fees, taxes, conversion and churn assumptions | MODEL EXISTS; INPUTS OPEN |
| EV-16 | Final launch/design gate | Owners sign scope, risks, tests and permitted next stage | BLOCKED |
| EV-17 | Counsel briefing pack | Focused questions, source hand-off, evidence checklist and answer log | DRAFT EXISTS; COUNSEL REVIEW NOT RECEIVED |

## 6. Guardrails for execution
- Keep research, synthetic PoC, controlled beta and production environments separate.
- No production user data or real identity/biometric material in a PoC by default.
- Do not provision paid infrastructure until cost owner approves a bounded budget and alerts.
- Do not select a data retention period just to simplify schema design.
- Do not mark age assurance complete from liveness configuration alone.
- Do not relax hard matching constraints to solve low density.
- Do not open public registration without staffed reporting/blocking/urgent escalation.
- Do not let Premium affect safety, consent, hard eligibility or hidden romantic ranking.
- Do not call a technical test a legal compliance certification.
- Every exception must state scope, risk, accountable approver, compensating control and expiry date; safety-critical invariants cannot be waived by convenience.

## 7. What can be done without further user or counsel input
- Prepare synthetic fixtures and isolated PoC scaffolding once the environment/budget guardrails are approved.
- Execute architecture-independent tests for idempotency, outbox semantics, safe errors and block/unmatch invariants in a disposable environment.
- Request Didit documentation and complete the vendor evidence pack from actual responses.
- Prepare a counsel briefing containing focused questions, source references and decision points.
- Develop market and price experiment protocols without choosing a launch city or price prematurely.
- Maintain traceability links and record actual test results without marking unexecuted items passed.

## 8. Current next move
Start with **ACT-01, ACT-03–ACT-08, ACT-10 and ACT-11** in parallel where owners are available. In the absence of named owners, keep status OPEN rather than imply assignment. Engineering can then prepare the isolated synthetic test harness and begin POC-01, with subsequent tests blocked only by genuine dependencies.

**No named owner, vendor response, legal sign-off, test pass, city, price or launch date is invented in this register.**
