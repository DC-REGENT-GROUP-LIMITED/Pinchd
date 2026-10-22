# Pinchd research: trust-and-safety operations and incident response
**Research date:** 21 October 2026  
**Scope:** UK-first dating service; reporting, moderation, escalation, appeals, account integrity, incident response and operational readiness.  
**Status:** Desk research, not legal advice. Confirm current legal duties and service scope with UK counsel. Mobile UI design remains paused.

## Executive summary
A dating app's safety promise is only credible if it is backed by a functioning operational system. Pinchd needs a joined-up process from report intake through triage, evidence preservation, proportionate action, user communication, appeal, and learning. Blocking must work throughout the product; reports must be easy to submit; urgent threats require escalation paths; and moderation capacity must be planned before launch. Do not treat an allegation as proof, a verification badge as a guarantee, or a classifier as a final judge.

## 1. UK regulatory baseline

### Online Safety Act: assess applicability, then build the operating system
Ofcom's current guidance says providers of in-scope user-to-user services must assess illegal-content risks, implement proportionate protections, keep records and review risk assessments. Duties include proportionate steps to prevent users encountering illegal content, mitigate risks, remove illegal content swiftly once aware, explain relevant policies in terms, provide easy reporting and operate a complaints procedure. Ofcom says risk assessments should be kept current and recommends annual review; a significant service change requires reassessment before the change. Dating services are included among examples of user-to-user services, but the precise obligations depend on Pinchd's features and legal scope.
- Ofcom, Illegal content duties under the Online Safety Act (published 18 Feb 2026; updated 25 Jun 2026): https://www.ofcom.org.uk/online-safety/illegal-and-harmful-content/illegal-content-duties-under-the-online-safety-act
- Ofcom, Online safety regulatory documents (updated 12 Oct 2026): https://www.ofcom.org.uk/online-safety/illegal-and-harmful-content/online-safety-regulatory-documents
- Ofcom, Online Safety Act guide for services: https://www.ofcom.org.uk/online-safety/illegal-and-harmful-content/guide-for-services

**Important timing detail:** Ofcom's regulatory-documents page lists updated user-to-user illegal-content Codes of Practice issued on 12 October 2026 as upcoming amendments which take effect 21 days after issue. Do not describe these amendments as already in force on 21 October 2026. Confirm the operative version at implementation and launch.

### Reporting, complaints and affected people
Section 20 of the Online Safety Act addresses content-reporting systems for regulated user-to-user services, including reporting illegal content and, where relevant, content harmful to children. Ofcom's guidance says users should be able to report easily and providers must operate a complaints procedure. The details of applicability and implementation need to be assessed against the service.
- Legislation, Online Safety Act 2023, section 20 (version at 30 Sep 2026): https://www.legislation.gov.uk/ukpga/2023/50/body/2026-09-30
- Ofcom, illegal content duties: https://www.ofcom.org.uk/online-safety/illegal-and-harmful-content/illegal-content-duties-under-the-online-safety-act

### App-store policies
Apple's App Review Guidelines require UGC services to filter objectionable material, provide a way to report offensive content with timely responses, allow users to block abusive users and publish contact information. Google Play's UGC policy calls for robust, effective and ongoing moderation, clear terms, in-app reporting and blocking.
- Apple App Review Guidelines, section 1.2: https://developer.apple.com/app-store/review/guidelines/
- Google Play, User-generated content policy: https://support.google.com/googleplay/android-developer/answer/9876937?hl=en-GB

**Implication:** Store review requirements are a minimum launch gate, not a complete safety strategy. Policies must match real operational capability; do not promise response times the team cannot meet.

## 2. Threat model for a dating service
Design operational playbooks for foreseeable categories:
- Harassment, coercion, threats and repeated unwanted contact.
- Sexual harassment, unsolicited explicit images, cyberflashing and non-consensual intimate imagery.
- Impersonation, fake profiles, romance fraud, financial solicitation and extortion.
- Stalking, doxxing, threats to disclose sensitive information, location inference and account surveillance.
- Child-safety concerns, underage users, grooming and suspected child sexual abuse material (CSAM).
- Ban evasion, coordinated abuse, mass reporting and malicious attempts to weaponise moderation.
- Account takeover, SIM/email compromise, recovery abuse and suspicious access.
- Self-harm threats, credible imminent threats to a person, and emergencies.
- Discrimination, hate, targeted abuse and threats tied to protected characteristics.
- False positives, context-dependent language, reclaimed terms, satire and consensual adult discussion that may be misclassified.

Police.uk describes digital abuse as including monitoring, stalking, harassment, threats, impersonation and non-consensual posting of images. Pinchd's safety guidance should make clear that digital abuse is not the victim's fault and provide calm, practical next steps.
- Police.uk, Digital domestic abuse: https://www.police.uk/advice/advice-and-information/daa/domestic-abuse/digital-domestic-abuse/

## 3. End-to-end case lifecycle
1. **Intake:** Report from profile, message, image, match, or support contact. Capture the relevant target and reason without forcing the reporter to compose a long narrative.
2. **Immediate protection:** Offer block/mute and safety guidance immediately; a user should not have to wait for a case decision to stop contact.
3. **Triage:** Apply severity and confidence rubric; identify imminent danger, child-safety concerns, intimate image abuse, credible threats and repeat patterns.
4. **Evidence preservation:** Preserve relevant content, timestamps, account IDs, message IDs and action history under restricted access and a documented retention/legal-hold policy. Avoid collecting unrelated private conversations.
5. **Investigation:** Review context and corroborating signals. Distinguish allegation, evidence, finding and enforcement outcome.
6. **Action:** Apply proportionate measures such as content removal, warning, feature restriction, temporary suspension, verification, permanent ban or escalation to appropriate authorities when required or justified.
7. **Communication:** Tell the reporter what can safely be shared, explain next steps and avoid promising a particular outcome before review.
8. **Appeal/review:** Provide a route for users affected by significant restrictions to request human review, subject to legal and child-safety constraints.
9. **Closure:** Record the reason, policy version, evidence category, action, reviewer and timestamps.
10. **Learning:** Examine recurrence, missed signals, false positives, time to action and product/design changes that could prevent similar harm.

## 4. Severity and service-level framework
The following is a proposed internal operational target, not a legal deadline or promise to users. Calibrate to launch geography, staffing, evidence and counsel advice.

| Tier | Examples | Initial handling target | Immediate controls |
|---|---|---|---|
| P0 Critical | Credible imminent threat to life; immediate child-safety danger; suspected active exploitation; credible extortion or intimate-image threat with urgent risk | Immediate automated escalation/queue; on-call human acknowledgement target measured in minutes during staffed coverage | Preserve essential evidence; block contact; restrict account/content when justified; follow emergency/child-safety protocol |
| P1 High | Threats, stalking, doxxing, non-consensual intimate imagery, cyberflashing, credible fraud/extortion, repeat targeted harassment | Rapid human triage; proposed target within 1 hour during staffed coverage | Prevent continued contact; preserve evidence; escalate specialist review |
| P2 Standard | Harassment, impersonation, suspicious solicitation, repeated unwanted contact, abusive content | Proposed target within 24 hours | Block/mute; queue for review; investigate pattern |
| P3 Low / support | Profile quality, non-urgent disputes, general policy questions, non-urgent appeals | Proposed target within 3 business days | Provide support route; avoid implying an urgent safety investigation |

Operational rules:
- Publish only response commitments that are actually staffed and monitored.
- Measure both first acknowledgement and meaningful action; an auto-reply is not a completed review.
- If Pinchd cannot staff urgent reports reliably, narrow launch scope or add a qualified moderation partner before opening access.
- Define 24/7 escalation coverage or clearly communicate coverage limits and direct users to emergency services when appropriate.
- Keep safety actions proportional to confidence and severity; uncertainty may justify temporary protection while facts are reviewed, but not careless permanent conclusions.

## 5. Emergency, child-safety and intimate-image playbooks
### Imminent danger
- Provide clear guidance to contact UK emergency services (999) where someone is in immediate danger.
- Do not promise that Pinchd can dispatch emergency help unless an operationally verified capability exists.
- Maintain a named escalation owner, decision tree, secure incident channel and post-incident review process.

### Child-safety concern / suspected CSAM
- Restrict access to the material and account as appropriate under the documented legal and safety process.
- Do not ask staff to download, forward, or casually inspect suspected illegal imagery.
- Use trained specialist handling, access controls, evidence-preservation rules and legal reporting/escalation procedures confirmed by counsel.
- Never expose the reporter or potential victim to the reported account through unnecessary notifications or workflows.
- Maintain age-assurance exceptions and appeals without creating an easy bypass.

### Non-consensual intimate imagery and cyberflashing
- Offer a clear report route with immediate blocking and minimal repeated exposure to the content.
- Limit copying and staff access; follow current legal requirements and specialist evidence handling.
- Consider platform-level matching or hashing only after a legal, privacy, accuracy and human-review assessment; do not rely on automated matches alone.
- Ensure takedown, account action, appeal and recurrence monitoring are coordinated.
- Do not require a victim to repeatedly view or upload the content to make a report.

## 6. Ban evasion, fraud and coordinated abuse
- Link enforcement records to the account and event, preserving the rationale and appeal history.
- Use device, IP, image reuse and behavioural patterns only as corroborating risk signals; shared networks, household devices and VPNs can create false positives.
- Avoid automatic permanent bans based on a single weak signal.
- Apply rate limits and friction to high-risk actions: mass likes/messages, rapid account creation, repeated external-contact solicitation and repeated account recovery attempts.
- Keep the signals used for abuse prevention separate from romantic compatibility ranking.
- Detect report brigading and coordinated abuse, but do not penalise good-faith reporters solely because a report is unsubstantiated.
- Provide a secure account recovery process with re-verification and notification of security changes.
- Maintain a human review and appeal path for severe enforcement, with a fast route for urgent safety concerns.

## 7. Reporting and blocking product requirements
- **SAFEOPS-01:** Report entry points exist on profile, message, image/media, match and support surfaces.
- **SAFEOPS-02:** The user can choose a reason and optionally add context; the form is short and accessible.
- **SAFEOPS-03:** Reports create a traceable case with immutable event history and role-restricted evidence.
- **SAFEOPS-04:** The reporter can block immediately without waiting for moderation.
- **SAFEOPS-05:** Blocking suppresses future discovery, likes, notifications, message requests and cached profile surfaces; test each path.
- **SAFEOPS-06:** The blocked person is not told the reporter's identity or given unnecessary information that facilitates retaliation.
- **SAFEOPS-07:** Users can report on behalf of someone affected by content where the law/policy requires or permits it.
- **SAFEOPS-08:** Case status communication is privacy-preserving and does not disclose another user's confidential enforcement details.
- **SAFEOPS-09:** Significant restrictions have a documented rationale and an appropriate appeal path.
- **SAFEOPS-10:** Appeals are reviewed by a person with authority to overturn the original decision, where feasible.
- **SAFEOPS-11:** Report categories, help text and terms match the actual moderation policy and legal duties.
- **SAFEOPS-12:** Publish accurate support/contact details and operational hours.
- **SAFEOPS-13:** Maintain a policy version and decision reason for each action.
- **SAFEOPS-14:** Safety operations can apply a temporary protective restriction while preserving review and appeal.
- **SAFEOPS-15:** Provide user-facing safety guidance without blaming victims or making guarantees of safety.

## 8. Moderation workforce and governance
- Assign a named accountable owner for trust and safety, even at a small launch.
- Create a written policy matrix mapping each report type to severity, evidence standard, action range, escalation and appeal.
- Train reviewers in trauma-informed communication, harassment, coercive control, romance fraud, intimate image abuse, child safety, bias and evidence handling.
- Use role-based access and least privilege; record who viewed or exported sensitive evidence.
- Provide reviewer support, workload limits, rotation and wellbeing safeguards for exposure to distressing content.
- Maintain quality assurance: sample decisions, dual-review critical cases, calibrate disagreements and document policy changes.
- Separate moderation performance from simplistic “number of bans” targets that encourage over-enforcement.
- Use vendors only with clear service levels, security/privacy due diligence, escalation procedures, audit rights, training requirements and exit plans.
- For a small launch, a trained specialist partner may be safer than pretending one founder can cover urgent reports around the clock.

## 9. Appeals, errors and procedural fairness
- Explain the policy category and action in language the user can understand, while withholding details that could enable evasion or endanger reporters.
- Allow users to submit relevant context and correct mistaken identity/account-linkage.
- Review severe actions with human judgement; do not let the original automated classifier be the sole appeal reviewer.
- Track reversals and false positives by policy category and, where lawful and proportionate, affected cohort.
- Avoid treating a report as a proven fact; assess evidence and context.
- Avoid requiring the reporter to confront the accused or negotiate directly.
- Provide a clear distinction between content removal, account restriction, suspension and permanent ban.
- Preserve the ability to take immediate protective action where risk is urgent, with prompt review afterward.

## 10. Operational metrics and dashboard
- Reports per 1,000 active users and per 1,000 active conversations, broken down by category and severity.
- Median and 90th/95th percentile time to first human review, action and closure; separate automated acknowledgement from human review.
- P0/P1 queue age, unresolved critical cases, on-call acknowledgement and escalation completion.
- Repeat offending and recurrence within 7/30/90 days, using documented definitions.
- Block effectiveness failures, including reappearance in discovery/likes/notifications.
- Appeal rate, appeal response time, reversal rate and reasons for reversal.
- False-positive and false-negative indicators from QA sampling and post-incident reviews.
- Reporter abandonment rate and accessibility issues in reporting flow.
- Case backlog, reviewer workload, wellbeing signals and coverage gaps.
- Fraud/impersonation verification outcomes and recurrence.
- Time to remove/restrict reported illegal content where applicable.
- User trust/safety perception and whether users know how to report.
- Safety performance by launch cohort and user segment where lawful, proportionate and privacy-safe.

No single metric proves safety. More reports may reflect rising harm or easier reporting; fewer reports may reflect improvement or loss of trust. Interpret metrics alongside qualitative case review and product context.

## 11. Incident response and learning loop
Maintain an incident playbook with:
- Trigger and severity classification.
- Incident commander and backup.
- Legal/privacy/security contacts and decision rights.
- Secure communication channel and evidence access rules.
- Containment actions and rollback criteria.
- User communication and regulator/law-enforcement escalation decision tree.
- Vendor escalation and evidence preservation.
- Recovery criteria and monitoring.
- Blameless post-incident review within a defined period, including what failed, what worked, and assigned remediation owners/dates.
- A change-control rule: significant product changes trigger a safety risk review before release.

Run tabletop exercises before launch for at least: romance-fraud ring, credible threat, non-consensual intimate image, child-safety report, account takeover, false-positive mass report and moderation vendor outage.

## 12. Launch gates
1. Confirm Online Safety Act scope and applicable Ofcom codes with qualified counsel.
2. Complete service-specific illegal-content risk assessment and children's access assessment as applicable; document decisions.
3. Publish terms, moderation rules, reporting route, complaints procedure and contact information.
4. Staff and test report queues, escalation paths, emergency playbook and appeal operations.
5. Test block effectiveness end-to-end, including cached results, notifications, likes and search/discovery.
6. Establish secure evidence handling, retention, access logging and legal-hold procedures.
7. Confirm child-safety and intimate-image abuse procedures with specialist legal/operational advice.
8. Set realistic internal service targets and prove capacity against expected launch volume.
9. Conduct tabletop exercises and resolve critical findings.
10. Establish dashboard, QA sampling, governance review and periodic risk reassessment.
11. Review app-store policies immediately before submission.
12. Reassess risks before material changes and at least annually where required/recommended.

## Limitations and open questions
- The Online Safety Act duties depend on the precise service design and scope; legal review is essential.
- Ofcom's 12 October 2026 issued amendments are listed as upcoming and become effective 21 days after issue; verify the current operative version before relying on them.
- Service-level targets above are proposed internal targets, not statutory deadlines or public promises.
- Exact reporting obligations for child-safety or intimate-image incidents depend on facts and applicable law; specialist counsel must confirm procedures.
- The optimal moderation staffing model depends on launch geography, user volume, coverage expectations and budget.
- Mobile UI design remains paused pending completion of the research and requirements package.

## Decision summary
**Recommended direction:** implement an end-to-end safety operation before launch, with easy reporting and immediate blocking, severity-based human review, evidence protection, clear escalation and appeals, trained moderation capacity, and measurable learning loops. Start small enough to operate safely; do not launch a user base or feature set beyond the team's ability to respond.

## Source register
1. Ofcom, Illegal content duties under the Online Safety Act (18 Feb 2026; updated 25 Jun 2026): https://www.ofcom.org.uk/online-safety/illegal-and-harmful-content/illegal-content-duties-under-the-online-safety-act
2. Ofcom, Online safety regulatory documents (updated 12 Oct 2026): https://www.ofcom.org.uk/online-safety/illegal-and-harmful-content/online-safety-regulatory-documents
3. Ofcom, Online Safety Act guide for services: https://www.ofcom.org.uk/online-safety/illegal-and-harmful-content/guide-for-services
4. Online Safety Act 2023, section 20 (version at 30 Sep 2026): https://www.legislation.gov.uk/ukpga/2023/50/body/2026-09-30
5. Apple App Review Guidelines, section 1.2: https://developer.apple.com/app-store/review/guidelines/
6. Google Play, User-generated content policy: https://support.google.com/googleplay/android-developer/answer/9876937?hl=en-GB
7. Police.uk, Digital domestic abuse: https://www.police.uk/advice/advice-and-information/daa/domestic-abuse/digital-domestic-abuse/


## 10. Profile authenticity, verification and impersonation controls (21 October 2026)

### Risk basis
Ofcom identifies fake personas used for romance fraud, intimate-image abuse/sextortion, harassment and stalking as risks for dating and social-discovery services. It reports that 29% of adult internet users had experienced romance or dating scams and 6% had encountered fraud on a dating website or app; these are broad survey measures, not platform-specific fake-profile prevalence. Source: https://www.ofcom.org.uk/cy/online-safety/illegal-and-harmful-content/dating-and-social-discovery-know-the-online-safety-risks-rules-and-how-to-comply

UK Trading Standards warns that scammers may use stolen photos and fabricated stories to build trust before seeking money, gifts or personal information. Source: https://www.economy-ni.gov.uk/news/protect-yourself-online-dating-scams

Hinge's voluntary Selfie Verification uses a video selfie and visible badge, and Hinge cautions that a badge does not guarantee safety. Some regions use a required biometric Face Check. Source: https://help.hinge.co/hc/en-us/articles/10303221435539-What-is-Selfie-Verification

### Proposed layered controls for Pinchd
1. **Account integrity:** enforce one active account per person where technically and proportionately supportable; detect repeat account creation, ban evasion, suspicious device/account clusters and rapid account rotation without treating shared networks or devices as proof of abuse.
2. **Photo integrity:** compare uploaded profile photos with an in-app capture during optional or risk-triggered verification; check for obvious stolen-image reuse and manipulated media using proportionate tools. Treat automated results as signals for review, not conclusive guilt.
3. **Visible verification status:** distinguish clearly between unverified, selfie-verified and any stronger identity/age-assurance status. Explain what each check verifies and what it does not. Never imply that verification guarantees honesty, safety, age or ongoing account control beyond the actual check.
4. **Report pathways:** make reporting impersonation, stolen images, AI-generated impersonation, scams and suspected underage accounts easy from the profile and conversation. Allow evidence attachment without requiring a victim to confront the suspected account.
5. **Risk-based review:** combine signals such as duplicate photos, scam-related language, high-volume outreach, repeated reports and suspicious account behaviour. Use human review for consequential actions where feasible, with documented thresholds, appeal and audit trails.
6. **No appearance policing:** do not use attractiveness, body type, disability, race, gender presentation, age appearance or cosmetic editing as risk proxies. Normal makeup, filters and flattering photographs are not automatically fraud.
7. **Fraud friction:** educate users not to send money, gift cards, crypto or financial details to someone they have not independently verified; warn on suspicious payment requests and attempts to move conversations off-platform rapidly. Avoid intrusive monitoring of private messages beyond clearly disclosed, proportionate safety purposes and legal review.
8. **Block and re-entry prevention:** a block must immediately stop contact and discovery between the pair. Suspected block evasion through alternate accounts should enter a safety review and cannot be allowed to recreate contact automatically.
9. **Incident response:** maintain a documented escalation route for credible threats, extortion, suspected minors, identity theft and financial fraud; preserve only necessary evidence under a defined retention/legal process and escalate to appropriate authorities where legally required or appropriate.
10. **Fair enforcement:** notify users of meaningful enforcement outcomes where safe, provide an appeal channel, and separate temporary protective measures from final findings. Avoid disclosing detection logic that would enable evasion.

### Verification strategy recommendation
Do not require biometric selfie verification for every user by default without a clear risk assessment, accessibility evaluation, privacy impact assessment, supplier review and UK legal analysis. Consider voluntary verification as an early trust signal and proportionate step-up verification for defined risk scenarios. Any mandatory age assurance is a separate legal/product control and must not be conflated with selfie verification or identity verification.

### Operational metrics
Track report volume and categories, confirmed impersonation/scam cases, time to triage, time to restrict high-risk accounts, repeat-offender rate, appeal overturn rate, false-positive rate where measurable, verification completion/failure by device/accessibility cohort, and user trust feedback. Review metrics for inequitable impacts and avoid optimising solely for the number of accounts banned.

### Acceptance criteria
- Verification status describes the actual check and never claims to guarantee safety.
- Users can report impersonation, scam behaviour and suspicious images directly from profile or chat.
- Verification failure does not automatically imply malicious behaviour; a retry and appeal route exists unless urgent risk justifies a temporary protective restriction.
- Automated detection cannot silently impose permanent enforcement without a documented review standard and appeal path, except where immediate protective action is necessary.
- Blocked pairs cannot contact each other through alternate discovery or messaging paths; suspected evasion is handled by safety operations.
- Verification and evidence collection follow data-minimisation, retention, access-control and privacy-review requirements.


## 11. Verification policy decision research: voluntary badge, risk-based checks and age gate (21 October 2026)

### Recommended policy architecture
Keep three separate layers in the product and operations model:
1. **Mandatory age assurance:** satisfy the legal requirements applicable to Pinchd's UK service and prevent under-age access where required. This is a launch gate, not a paid benefit.
2. **Optional profile selfie-verification badge:** a trust signal about a specific selfie/profile comparison, only if the check can be explained accurately and implemented with acceptable privacy/accessibility safeguards. Do not present it as proof of legal identity or safety.
3. **Risk-triggered enhanced verification:** use only for defined, documented scenarios such as suspected impersonation, credible account takeover, repeated high-risk reports or appeal review. A trigger begins a proportionate review; it is not itself proof of wrongdoing.

Ofcom's guidance lists multiple methods capable of highly effective age assurance at 18, including facial age estimation, digital identity, photo-ID matching, mobile-network operator checks, open banking, credit-card checks and email-based age estimation. It does not treat self-declaration alone as highly effective at 18. The method must be assessed for the specific implementation and legal duty. Source: https://www.ofcom.org.uk/online-safety/protecting-children/age-assurance

Ofcom's July 2026 review says the evidence covers the first six months of child-protection duties and is not a final verdict on every method. It urges regulated services to follow the guidance, perform continuing supplier due diligence and meet privacy obligations. Pinchd remains responsible for its compliance even if a vendor runs the check. Source: https://www.ofcom.org.uk/online-safety/protecting-children/use-of-age-assurance-report-2026

### Should the selfie badge be mandatory?
**Recommendation: do not require a public selfie-verification badge for every user by default at launch.** First validate whether the badge meaningfully reduces impersonation and improves user trust, while measuring drop-off, accessibility failures, false rejection and support cost. A public badge can be voluntary; risk-triggered verification may be required under clearly disclosed rules when justified. This is a product recommendation, not a conclusion that voluntary checks alone meet age-assurance obligations.

Reasons:
- A selfie/profile check does not prove legal identity, truthful profile claims or safe conduct.
- Mandatory biometric capture increases privacy, accessibility and onboarding burden.
- Verification failures may result from poor lighting, device quality, disability, appearance changes or provider error; these are not proof of fraud.
- The safety benefit must be evidenced against less intrusive options and proportionate fraud controls.

### Risk-triggered verification policy
Potential triggers for step-up checks include a credible impersonation report, repeated photo reuse across unrelated accounts, suspicious account recovery, or a documented pattern of high-volume scam-like behaviour. Do not trigger solely because of attractiveness, race, skin tone, gender presentation, disability, age appearance, nationality, LGBTQ status or other protected/sensitive characteristics. Use multiple independent signals where possible and review bias/false positives.

Where immediate protective action is necessary, temporarily limit risky functions while a human review is pending. Explain the restriction at a useful level, avoid exposing detection logic, preserve a safe appeal route, and reinstate access promptly when evidence does not support the concern. Permanent enforcement requires a documented standard of evidence and a meaningful appeal path except where law or urgent safety needs require otherwise.

### Required operating workflow
1. Trigger is logged with reason category and signal source.
2. User receives a clear request describing what must be verified, why, what data is collected and what happens if they do not complete it.
3. Provide accessible capture guidance, retries and an alternative route where feasible.
4. Automated outcomes are checked against confidence and quality thresholds; ambiguous outcomes go to human review rather than automatic permanent bans.
5. Record only the minimum outcome and case evidence required; restrict access to authorised staff.
6. Communicate decision and appeal path, except where disclosure would create a specific safety or legal risk.
7. Apply deletion/retention rules to raw images, documents and case records; periodically audit provider and internal handling.

### Metrics and review cadence
Track verification request rate, completion, abandonment, retry, failure and appeal overturn; confirmed impersonation and scam incidence; time to resolve; repeat offender rate; false positive indicators; and the effect on user activation. Break down by device/accessibility cohorts and demographic groups only where lawful, necessary, statistically meaningful and privacy-protective. Avoid small-cell reporting. Reassess thresholds and vendor performance at a defined cadence and after material incidents or provider/model updates.

### Decision gates
- No public selfie badge ships until its exact meaning, user-facing explanation, data flow, accessibility and privacy impact are approved.
- No risk-triggered check ships without documented triggers, review standards, appeal, retention and access-control rules.
- Age assurance is selected separately and must meet the applicable legal standard; a public selfie badge cannot substitute for it unless the exact method independently meets that requirement.
- Vendor contracts must address subprocessors, security, incident notification, data retention/deletion, model training/secondary use, audit evidence, service continuity and exit assistance.
