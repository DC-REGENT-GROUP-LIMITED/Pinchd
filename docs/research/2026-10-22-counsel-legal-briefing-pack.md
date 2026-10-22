# Pinchd — Counsel Legal and Privacy Briefing Pack
**Date:** 22 October 2026  
**Status:** Draft instruction pack for qualified UK counsel and privacy review. Not legal advice, not a legal opinion and not an approval.  
**Client/project:** DC Regent Group Limited / Pinchd (confirm legal entity and controller before reliance).  
**Proposed scope:** UK-first adult dating/social-discovery service; final launch jurisdictions, feature scope and operator details must be confirmed.

## 1. Purpose and requested deliverables
Please provide a written, dated assessment against the actual product and launch configuration, identify assumptions and missing facts, cite the law/guidance version used, and separate mandatory duties from prudent controls. Please flag matters requiring specialist counsel or regulator engagement. Do not treat this research pack as a substitute for facts or legal instructions.

Requested outputs:
1. Applicability memo and launch-blocking obligations with accountable owner and evidence.
2. Reviewed processing inventory, Article 6 basis and Article 9 condition/safeguards per purpose where applicable; DPIA screening and DPIA recommendation/completion.
3. Approved retention/deletion schedule, legal-hold criteria, backup/restore and post-deletion pair-suppression advice.
4. Age-assurance purpose, threshold, method, necessity/proportionality, vendor roles/transfers, data minimisation, false-reject/appeal and notice advice.
5. UGC/Online Safety Act scope, risk assessments, reporting/complaints, recordkeeping and operational requirements applicable to this configuration.
6. Review of eligibility/discovery rules, protected-characteristic risks, accessibility and age-check disparate impact.
7. Premium/subscription, consumer information, cancellation/refund and platform-billing advice for intended launch date and channels.
8. Terms, privacy notice, community standards, marketing claims, complaints/contact, IP/trademark and corporate/tax matters in or out of scope.
9. A clear readiness opinion with unresolved assumptions; no generic “compliant” conclusion absent evidence.

## 2. Facts supplied for confirmation (not counsel-verified)
- Product direction: “Serious about compatibility, flexible about relationship goals.”
- Proposed UK-first adult dating service with user profiles, preference-based discovery, reciprocal likes/matches, messaging, reporting/blocking, optional profile-photo verification and intended age assurance.
- Proposed stack: Go, React Native/Expo, AWS and PostgreSQL; these do not establish legal compliance.
- Sensitive data may arise from relationship goals, preferences, messages, safety reports, age assurance and inferred attributes. Exact fields and purposes remain subject to confirmation.
- Proposed controls include hard eligibility gates, private media, server-authoritative age state, block suppression, deletion orchestration and optional AI coaching; confirm actual MVP scope and whether each will launch.
- Existing research documents are working drafts. No counsel sign-off, vendor proof, user testing or technical PoC completion is evidenced in this pack.
- Mobile UI design remains paused pending research/requirements and design-readiness gates.

## 3. Priority questions requiring written advice
### A. Entity, geography and service classification
1. Confirm contracting entity, controller/processor roles, establishment, launch territories, age range and target audience.
2. Does the actual service meet the applicable Online Safety Act user-to-user/service definitions? Which Ofcom codes, duties, risk assessments, records and review cadence apply?
3. What child-access assessment and child-protection measures are required given the service's intended adult-only policy and likely access?
4. Which current Ofcom tools/documents and effective dates must the company retain as evidence? Please cite exact versions and outputs.

### B. Data protection, sensitive data and retention
5. For each actual field and purpose (profile, preferences, matching, messages, safety, age assurance, analytics, AI, support, billing), identify Article 6 basis and Article 9 condition/safeguards where applicable; address inferences and data sharing.
6. Is a DPIA required for the combined processing? If not, document screening rationale, residual risks and reassessment triggers.
7. Which fields are special-category data or criminal-offence data in this design? How should consent, explicit consent or another condition be assessed without assuming a single basis for all uses?
8. What retention periods/triggers are justified for account data, likes/passes, pair suppression, messages, safety evidence, banned-account records, age results, vendor copies, audit logs, analytics and backups?
9. Can minimum pair-suppression survive account deletion, on what basis/design/duration/disclosure, and how are rights and objections handled?
10. What deletion/erasure, access, restriction, portability and automated-decision transparency processes are required, including vendor/backup propagation and lawful holds?

### C. Age assurance and vendors
11. What age-assurance standard/method and evidence are proportionate for this service and intended age threshold? Separate assurance from optional identity/photo verification.
12. What must be assessed in Didit's specific workflow, DPA, subprocessors, data regions, support access, transfers, accuracy, accessibility, retention, deletion and security terms?
13. Are any age-check inputs/results biometric or special-category data in this context? What safeguards, notices and alternative/appeal route are required?
14. What evidence is necessary before relying on vendor outcome, and what fallback is lawful/safe during outage, inconclusive result, false rejection or service termination?

### D. Safety, equality, accessibility and consumer terms
15. What reporting, complaint, illegal-content, urgent-harm, evidence preservation, appeals and staffing requirements apply? Does the proposed operating model meet duties in practice?
16. Review profile eligibility, gender identity, discovery preferences, visibility, ranking and verification for Equality Act/service-provider risks. What objective documentation or alternative is required?
17. What accessibility standard and reasonable-adjustment obligations apply to signup, age assurance, discovery, reporting, blocking, cancellation and deletion?
18. For intended paid channels and launch date, which subscription/consumer rules are in force; what notices, confirmation, reminders, cancellation, cooling-off, refund and store terms apply?
19. Review claims about compatibility, verification, safety, scarcity, outcomes, pricing and Premium advantages for substantiation and unfair-practice risk.

## 4. Documents and evidence to provide
- Product requirements, approved and proposed decision logs, lifecycle matrix, matching rules and user journey descriptions.
- Data classification/retention map, architecture/schema blueprint, data-flow diagrams, event/API contracts and proposed deletion process.
- Didit due-diligence pack plus actual vendor responses, contract/DPA, workflow/version, sample redacted payloads, region/subprocessor and deletion evidence when received.
- Trust-and-safety operating model, risk taxonomy, coverage/staffing plan, incident/escalation/complaints runbooks.
- Accessibility acceptance criteria and future participant/device test plan.
- Subscription economics and proposed billing channels/terms; no final price or launch date assumed.
- Security threat model, access model, incident/breach plan and future PoC/test evidence as it becomes available.
- Existing legal desk review: `2026-10-22-workstream-2-uk-legal-regulatory-readiness.md`; status dashboard: `2026-10-22-full-fold-research-status-and-launch-readiness.md`.

## 5. Decision log template
| Ref | Question | Counsel conclusion / authority | Facts assumed | Required action/evidence | Owner role | Due/review trigger | Status |
|---|---|---|---|---|---|---|---|
| LEG-01 | OSA scope and duties |  |  |  |  |  | OPEN |
| LEG-02 | Child access / age assurance |  |  |  |  |  | OPEN |
| LEG-03 | Article 6/9 basis and DPIA |  |  |  |  |  | OPEN |
| LEG-04 | Retention, deletion, pair suppression |  |  |  |  |  | OPEN |
| LEG-05 | Didit contract, data flow and transfers |  |  |  |  |  | OPEN |
| LEG-06 | Safety operations and records |  |  |  |  |  | OPEN |
| LEG-07 | Equality/accessibility |  |  |  |  |  | OPEN |
| LEG-08 | Subscriptions and consumer journey |  |  |  |  |  | OPEN |
| LEG-09 | Terms, notices, claims, IP |  |  |  |  |  | OPEN |

## 6. Current source pack and limits
Use the legal workstream's dated source list as a research starting point, then independently verify authoritative sources and effective/commencement dates on the advice date. Key existing official references include Ofcom's Online Safety Act dating/social-discovery guidance and age-assurance guidance; ICO special-category and DPIA guidance; applicable UK legislation and commencement instruments; GOV.UK/CMA consumer guidance; and current Apple/Google developer policies. See the linked workstream for URLs and the research limitations.

The existing desk review is a first-pass research artifact, not a legal determination. The intended launch date, territories, final feature scope, controller identity, vendor workflow and actual data flows are unresolved or unverified. Counsel should state which conclusions change if any assumption changes.

## 7. Completion and hand-off rule
Counsel answers and signed/dated determinations should be stored as separate evidence records; do not overwrite this question pack with inferred conclusions. Map each answer to ACT-02–08, ACT-25 and EV-02–07/16 in the execution register. Until reviewed evidence is received, corresponding items remain OPEN or BLOCKED. This pack does not close a legal gate or authorize public registration, paid billing, vendor production use or mobile UI design.
