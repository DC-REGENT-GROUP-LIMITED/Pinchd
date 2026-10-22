# Pinchd Workstream 3 — Trust & Safety Operating Readiness
**Research date:** 22 October 2026  
**Scope:** UK-first dating platform; operational case handling, severity, staffing, escalation, evidence, appeals, incidents, safety metrics and launch gates.  
**Status:** Proposed operating model based on project research and public-source desk research. Not legal advice, not a staffing commitment, and not proof of compliance. Mobile UI design remains paused.

## Executive decision
Pinchd must not launch into an operating environment where users can report abuse but no accountable person or process can respond. Trust and Safety (T&S) is a product capability and an operational service, not just a report button or classifier. The minimum viable operating model needs: an accountable owner; an accessible reporting and blocking path; severity-based queues; urgent escalation; evidence access controls; documented actions and appeals; coverage limits communicated honestly; and measurable capacity with launch-stop conditions.

The model below distinguishes (a) baseline requirements already in the project, (b) recommended operating targets, and (c) legal/vendor questions that require validation. Response targets are internal service objectives, not statutory deadlines.

## 1. Evidence baseline and limitations
- Existing Pinchd research already identifies harassment, coercion, threats, non-consensual intimate imagery (NCII), cyberflashing, fraud, stalking/doxxing, child safety, ban evasion, account takeover, self-harm, discrimination and false positives as relevant threat scenarios.
- Ofcom's Online Safety Act guidance describes risk assessment, proportionate safety measures, swift handling/removal of illegal content, terms, reporting and complaints duties for in-scope services. Pinchd must confirm its exact service classification and applicable duties with UK counsel.
- Ofcom published dating/social-discovery guidance in May 2026. The live regulatory documents and code version in force at launch must be checked; publication of a code does not necessarily mean it is already in force.
- Report Fraud/City of London Police reported 10,784 romance-fraud reports and over £102 million in reported losses for 2025. These are reported figures, not the total incidence of romance fraud and not a forecast for Pinchd.
- Qualitative research on dating-app safety can illuminate user experience but must not be treated as a representative prevalence estimate unless its methodology supports that claim.
- Sources: Ofcom, Illegal content duties under the Online Safety Act, https://www.ofcom.org.uk/online-safety/illegal-and-harmful-content/illegal-content-duties-under-the-online-safety-act ; Ofcom, Dating and social discovery, https://www.ofcom.org.uk/cy/online-safety/illegal-and-harmful-content/dating-and-social-discovery-know-the-online-safety-risks-rules-and-how-to-comply ; Ofcom regulatory documents, https://www.ofcom.org.uk/online-safety/illegal-and-harmful-content/online-safety-regulatory-documents ; Report Fraud, romance-fraud reporting update, https://www.actionfraud.police.uk/ (confirm exact report page/date before external publication).

## 2. Ownership and decision rights
### Required accountable roles
- **Executive accountable owner:** accepts launch risk, approves resourcing and can pause onboarding or a feature.
- **T&S operational owner:** owns policies, queues, training, quality, escalation rota, vendor oversight and weekly risk reporting.
- **On-duty case handler:** triages incoming cases within declared coverage and follows playbooks.
- **Engineering on-call:** handles platform containment, access-control failures, block bypasses, data exposure, notification leaks and incident evidence.
- **Privacy/security lead (named person or contracted role):** advises on lawful handling, access, retention, breach assessment and vendor processing.
- **Legal adviser:** validates statutory obligations, external reporting duties, law-enforcement requests, retention exceptions and consumer-facing wording.
- **Appeal reviewer:** reviews contested high-impact decisions independently of the initial decision where practicable.

In a small beta one person may hold multiple roles, but critical incident response must have a named backup. Record who is on duty, how they can be reached, and which decisions they are authorised to make. Do not advertise 24/7 monitoring unless it actually exists.

## 3. Severity and internal response targets
Severity should be assigned by credible potential harm and immediacy, not by the user's eloquence, social status, payment tier or whether an automated model scored the report highly. A report may be reclassified as evidence emerges. Multiple reports can be a signal but are not proof.

| Level | Example / threshold | Initial handling target | Required action |
|---|---|---|---|
| **P0 — Critical** | Credible imminent physical threat; suspected child sexual exploitation/abuse material; active exposure or distribution of intimate imagery; severe live security/privacy incident | Immediate escalation as soon as detected during declared coverage; acknowledge/triage within minutes as an operational aim | Protect affected users, restrict content/account or feature as proportionate, preserve necessary evidence securely, notify incident owner and obtain specialist/legal guidance |
| **P1 — High** | Threats, stalking, sexual coercion, repeat targeted harassment, credible fraud campaign, severe impersonation | Human triage within 1 hour during declared staffed coverage as a proposed target | Risk-based protective action, linked-case review, evidence preservation and supervisor escalation |
| **P2 — Standard** | Non-imminent harassment, suspected scam without immediate threat, impersonation, policy violations needing investigation | Triage within 24 hours as a proposed target | Investigate, apply proportionate action, communicate outcome where safe and appropriate |
| **P3 — Support / low urgency** | General product concern, non-urgent appeal, education or policy question | First response within 3 business days as a proposed target | Route to support or appeal workflow; do not let it displace P0/P1 work |

These are proposed internal targets, not a representation of legal deadlines. They must be tested against actual staffing, operating hours, vendor SLAs and forecast volume before launch. If coverage is limited, make that clear and give a safe urgent-report route; do not imply that a queue is monitored when it is not.

## 4. End-to-end case lifecycle
1. **Intake:** accept reports about a profile, message, photo, prompt, suspected scam, safety incident or account. Capture category and a short optional description; minimise free-text collection. Preserve the reported object and relevant identifiers without forcing a user to keep viewing harmful material.
2. **Immediate user protection:** allow the reporter to block or mute immediately, without waiting for investigation. Explain that reporting and blocking are different actions. The reported person must not be told who reported them unless law requires disclosure.
3. **Triage:** check severity, imminence, vulnerability, repeated reports, linked accounts, content type, jurisdiction and whether a specialist playbook applies. Do not treat an accusation, verification status, or single model output as conclusive.
4. **Containment:** where justified, remove or restrict content, limit contact, suspend an account, prevent further exposure, or temporarily disable a feature. Prefer proportionate and reversible interim measures when facts are uncertain and risk permits.
5. **Evidence and investigation:** preserve only evidence needed for safety, enforcement, legal obligations or appeal. Record who accessed it, why, what was reviewed and what decision followed. Restrict intimate imagery and highly sensitive evidence to specially authorised staff.
6. **Decision:** apply the published rules consistently. Record the policy provision, evidence basis, uncertainty, action, duration, decision maker and any linked cases. Use high-impact enforcement review where practicable.
7. **Communication:** provide a concise outcome or status where safe and lawful. Avoid revealing reporter identity, sensitive investigative methods or information that could increase risk.
8. **Appeal:** provide an accessible route to challenge significant enforcement decisions. Review new evidence and policy application; separate the reviewer from the original decision where practicable. Urgent protective restrictions need not be lifted merely because an appeal is filed.
9. **Closure and learning:** close only when immediate risk is addressed, required actions and communications are complete, evidence access is controlled, and follow-up tasks have owners. Feed anonymised trends into policy, product and engineering changes.

## 5. Non-negotiable product/system controls
- **Block is authoritative server-side state.** Every discovery result, incoming/outgoing like, match creation, message send/delivery, push notification, queued job, cache response and re-entry path must re-check current account and block state. A stale cache must never override a block.
- **Report and block remain available to free users.** Safety actions must not be paywalled, hidden behind a premium upsell, or delayed to increase engagement.
- **Report is not an automatic guilt finding.** Use risk-based temporary safeguards where needed, but distinguish intake from final enforcement.
- **No reporter disclosure by default.** Internal access and external disclosure are governed by least privilege, law and documented process.
- **No silent safety failure.** If reporting, blocking, age eligibility or urgent escalation is unavailable, raise an operational incident and consider stopping new onboarding or affected functionality.
- **Do not use paid status, attractiveness, match count or revenue value to set safety priority.**
- **Avoid unsafe evidence previews.** Users and staff should not be forced to reopen explicit or threatening content to report it; use safe summaries/thumbnails and controlled evidence access.
- **Age assurance and optional profile verification are separate.** A profile badge is not proof that a person is safe, truthful or legally eligible. Didit integration, vendor due diligence, error/appeal handling, retention and deletion must be evidenced before relying on it.

## 6. Specialist playbooks to complete before launch
### Imminent threat, stalking or coercive control
Prioritise immediate safety, restrict contact and visibility as appropriate, preserve relevant evidence, escalate to the on-duty lead, and give neutral safety guidance. Do not promise that Pinchd can dispatch emergency services. Verify jurisdiction-specific escalation and law-enforcement procedures with counsel.

### NCII, sexual exploitation and cyberflashing
Stop further exposure promptly where legally and technically appropriate; restrict access to evidence; avoid downloading, copying or redistributing imagery unnecessarily; record identifiers and action history; use a specialist escalation path. Suspected child sexual abuse material must never be circulated internally as ordinary case evidence. Obtain counsel-approved handling, reporting and preservation procedures before launch.

### Romance fraud, extortion and impersonation
Look for linked accounts, repeated scripts, suspicious contact patterns and off-platform money requests, while avoiding unsupported accusations. Protect users from further contact, preserve relevant message/account evidence, apply proportionate enforcement, and provide practical reporting guidance. Never guarantee reimbursement or recovery of funds.

### Ban evasion and coordinated abuse
Use a risk-based combination of account, device, behavioural and verified signals only after privacy/legal review. Avoid treating shared IP, household devices or accessibility tools as conclusive evidence. Use graduated restrictions and review paths; monitor false positives and disparate impact.

### Account takeover and platform compromise
Revoke sessions or credentials where appropriate, protect recovery channels, preserve security logs, assess exposure, notify security/privacy leads and follow the incident and breach-assessment plan. Avoid sending sensitive account details to an email address that may itself be compromised.

### Self-harm or immediate wellbeing concern
Use a compassionate, non-clinical response and jurisdiction-appropriate resources. Do not present moderators as therapists or imply emergency monitoring. Define escalation boundaries and staff wellbeing support.

## 7. Evidence, privacy and access controls
- Use role-based access, least privilege, strong authentication and audit logs for case records and evidence.
- Separate ordinary support access from sensitive T&S evidence access.
- Record purpose, case ID, access reason, export/download activity and disclosure.
- Define retention by evidence category and legal purpose; avoid indefinite retention by default. Deletion must propagate to derived stores and vendors subject to documented lawful holds.
- Establish a lawful request workflow for police/courts, including identity/authority validation, scope minimisation, approval, secure transfer and audit trail.
- Do not use case evidence to train models or for unrelated analytics without a separately assessed lawful basis, transparency and controls.
- Build a privacy/security incident path for accidental evidence disclosure, incorrect access permissions, vendor incidents and lost devices.

## 8. Staffing and capacity model
Do not infer staffing from MAU alone. Estimate demand by report/contact rate, severity mix, duplicate/link rate, average handling time, investigation complexity, appeals, QA, training, leave, incidents and vendor escalation delays.

A first planning equation is:
**Required case-handler hours = expected cases × average handling minutes ÷ 60**, then add explicitly modelled time for quality review, appeals, training, supervision, incident reserve and leave. This is a workload estimate, not a staffing recommendation until real or pilot data exists.

### Stage gates
- **Closed test:** invite-only and capped; named T&S owner and backup; defined operating hours; all test reports routed; urgent escalation rehearsed; block/report and deletion flows tested; staff trained on evidence handling.
- **Local beta:** monitored queue with documented rota; daily backlog and severity review; weekly sample-based QA; appeal reviewer; incident simulation; measured handling-time distribution and peak-load assumptions.
- **Growth:** planned coverage across hours/time zones; specialist queue ownership; vendor contingency; regular policy calibration; staff wellbeing and exposure controls; capacity forecasts tied to acquisition plans.
- **Before each expansion:** prove that moderation capacity, escalation, vendor support and incident response scale with the expected user cohort. Do not widen invitations simply because acquisition demand is strong.

## 9. Metrics and governance
Review weekly during beta and more frequently during incidents. Segment by severity, category, language/region where justified, and product surface without exposing unnecessary sensitive attributes.

Core metrics:
- Reports per 1,000 active users, by category and surface.
- Time to first human triage and time to protective action, median and high percentiles.
- Open backlog and oldest-case age by severity.
- Repeat reports, repeat victimisation signals and linked-account clusters.
- Appeal rate, overturn rate and reason.
- False-positive/false-negative estimates from reviewed samples, with uncertainty stated.
- Block bypass, unwanted contact after block, and notification leakage tests.
- Report abandonment and completion rate.
- Case-handler workload, overtime, leave coverage and exposure-related wellbeing indicators.
- Vendor uptime, escalation latency, and data deletion/retention exceptions.
- Safety outcomes after action, not just volume of removals or bans.

Metrics must not reward indiscriminate bans, suppressed reporting or premature closure. Review examples qualitatively and include user impact, fairness and unintended consequences.

## 10. Incident command and launch-stop conditions
For a severe incident, name an incident commander, preserve a timeline, contain the risk, establish one internal communication channel, track decisions and owners, and conduct a blameless post-incident review with corrective actions. Determine legal, regulatory, user and vendor notification obligations with counsel and the privacy/security lead; do not guess deadlines.

**Pause new onboarding or affected functionality** when any of the following is true:
- No named owner or reachable escalation path for critical reports.
- Block/report controls fail or a block can be bypassed by messaging, notifications, discovery or queues.
- A critical queue is unattended beyond the declared service objective.
- Age eligibility or required reporting mechanisms are unavailable.
- A serious privacy/security incident creates uncontrolled exposure.
- Case volume exceeds safe handling capacity and there is no credible mitigation.
- A required vendor is unavailable and the fallback process is untested.
- Evidence access is uncontrolled or staff are using unapproved channels to share sensitive material.

Resumption requires a recorded decision, tested fix or mitigation, residual-risk acceptance by the accountable owner, and any required legal/privacy sign-off.

## 11. Pre-launch acceptance checklist
- [ ] Named accountable executive, T&S owner, backup and on-call engineering contact.
- [ ] Published community standards and reporting categories approved by counsel.
- [ ] Report and block paths tested across profile, message, match, notification and queued-delivery surfaces.
- [ ] P0/P1 escalation rota and declared coverage tested with tabletop exercises.
- [ ] Playbooks approved for threats/stalking, NCII/cyberflashing, suspected child exploitation, fraud/extortion, account takeover and privacy incidents.
- [ ] Evidence access, audit logging, retention, deletion, lawful holds and police-request process reviewed.
- [ ] Appeal process and high-impact enforcement review documented.
- [ ] Vendor contracts, data flows, SLAs, subprocessors, security evidence and contingency plans reviewed.
- [ ] Forecast workload and staffing assumptions tested against a capped pilot.
- [ ] Metrics dashboard and launch-stop authority defined.
- [ ] Online Safety Act scope, risk assessment, terms, reporting/complaints and code version reviewed by UK counsel.
- [ ] Store requirements for user-generated content, blocking/reporting and in-app account deletion verified against current Apple and Google policies.
- [ ] Didit age-assurance evidence pack and end-to-end test results reviewed; unresolved integration gaps treated as launch blockers where they affect legal eligibility or required controls.
- [ ] No UI design work begins until research/requirements and the design-readiness gate are complete.

## 12. Open questions requiring evidence, not assumption
1. Which support/T&S hours can Pinchd genuinely staff at closed test, beta and launch?
2. What case rate and severity distribution does the pilot show, and what average handling time follows?
3. Which specialist escalation providers or counsel are available for P0/child-safety/NCII cases?
4. What exact UK reporting, preservation and external-disclosure obligations apply to Pinchd's eventual service design?
5. Which vendors/processors touch reports, evidence, identity checks and messages, and in which regions?
6. What retention periods are necessary for each evidence category, and which exceptions are legally justified?
7. What are the tested fallback behaviours when reporting, age assurance, messaging, push delivery or the T&S console is unavailable?
8. Which measures detect repeat abuse without over-relying on shared device/IP signals or creating discriminatory outcomes?

## Completion status
This workstream supplies a proposed operating model and acceptance checklist. It is not complete until owners, coverage, playbooks, vendor evidence, pilot capacity data and legal review are attached to the project. The existing research note remains the supporting background; this file is the operational-readiness synthesis.
