# Pinchd research: privacy, UK GDPR, age assurance and data governance
**Research date:** 21 October 2026  
**Scope:** UK-first dating service; product, engineering, operations and launch requirements.  
**Status:** Desk research, not legal advice. Obtain qualified UK privacy and online-safety counsel before launch. Mobile UI design remains paused.

## Executive summary
Dating products process unusually intimate personal information. Pinchd should build a privacy and data-governance system before implementation: purpose-specific data mapping, lawful-basis and special-category analysis, a Data Protection Impact Assessment (DPIA), clear retention and deletion schedules, privacy-protective defaults, strong access controls, vendor due diligence, and a proportionate age-assurance strategy. Age assurance and identity verification are related but distinct: proving someone is over 18 does not prove their name, photos, intentions or safety.

## 1. Current UK regulatory context

### UK GDPR and Data Protection Act
The UK GDPR's principles include lawfulness, fairness and transparency; purpose limitation; data minimisation; accuracy; storage limitation; integrity and confidentiality; and accountability. ICO guidance says privacy must be integrated from the planning stage and throughout the lifecycle, and that a DPIA is required where processing is likely to result in high risk to people's rights and freedoms. A DPIA is good practice even where the threshold is uncertain.
- ICO, Data protection by design and by default (updated 5 February 2026): https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/accountability-and-governance/guide-to-accountability-and-governance/data-protection-by-design-and-by-default/
- ICO, Guide to the data protection principles (updated 23 March 2026): https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/data-protection-principles/a-guide-to-the-data-protection-principles/

The Data (Use and Access) Act 2025 has changed parts of the UK data-protection framework; ICO guidance is being updated. Pinchd should check the law and current guidance at each major design/launch gate rather than rely on an old compliance checklist.

### Special-category data
Sexual orientation and sex life are special-category personal data under UK GDPR. Dating profiles may also reveal religious beliefs, health information, political opinions or other sensitive characteristics. Processing special-category data requires both an Article 6 lawful basis and an applicable Article 9 condition, plus any relevant UK law safeguards. Consent is not a universal cure: it must be freely given, specific, informed, unambiguous and withdrawable, and other legal conditions may be necessary depending on the processing. A privacy lawyer should assess the precise feature and purpose.
- ICO, special category data guidance: https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/lawful-basis/special-category-data/

**Product implication:** Treat relationship goals, orientation and intimate profile fields as high-sensitivity data even if the legal classification of every field varies by context. Do not collect a field “just in case”; explain why it is needed and who can see it. Never infer orientation, sexual behaviour, health, religion or other sensitive traits from photos, names, location or messages.

### Online Safety Act and age assurance
Ofcom's age-assurance guidance applies to services in scope of the Online Safety Act, depending on service design and content/risk. Ofcom says age assurance should be highly effective where required and assessed for technical accuracy, robustness, reliability and fairness, with privacy obligations considered. Its July 2026 report reviewed early implementation across sectors including online dating and stresses that no single method eliminates circumvention; layered approaches and vendor due diligence matter.
- Ofcom, Age assurance duties under the Online Safety Act (updated 2 September 2026): https://www.ofcom.org.uk/online-safety/protecting-children/age-assurance
- Ofcom, Use of Age Assurance Report 2026 (published 15 July 2026): https://www.ofcom.org.uk/online-safety/protecting-children/use-of-age-assurance-report-2026
- GOV.UK, Ofcom report on the use of age assurance: https://www.gov.uk/government/publications/ofcom-report-on-the-use-of-age-assurance

Do not assume that a self-declared birth date alone satisfies a highly-effective-age-assurance duty. The correct method and threshold depend on the service's legal scope and current Ofcom guidance; obtain a documented scope assessment.

### Biometrics and identity verification
A face image is personal data; it is not automatically special-category biometric data in every context. Biometric data becomes special-category data when processed through specific technical means allowing or confirming unique identification. Facial age estimation, selfie-to-ID matching, liveness checks and face recognition can have different purposes, risks and legal analyses. Pinchd should not conflate age estimation, liveness, identity proofing and identity uniqueness into one “verified” claim.
- ICO, Biometric data guidance: https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/lawful-basis/special-category-data/what-is-special-category-data/biometric-data/

**Product implication:** Prefer a specialist age-assurance provider that can return a minimal result (e.g. “over 18: yes/no/needs review”) without Pinchd retaining ID documents or facial templates, where legally and technically appropriate. Verify the vendor's method, error rates, accessibility, bias testing, security, deletion behaviour, subprocessors, data locations and incident process. A vendor does not transfer Pinchd's accountability away.

## 2. Data map and proposed handling rules

| Data category | Examples | Recommended handling |
|---|---|---|
| Account and contact | Account ID, email, phone | Separate account identifiers from profile data; verify only as needed; rate-limit and protect recovery flows. |
| Age assurance | Result, method, timestamp, vendor reference | Prefer minimal proof/result and audit metadata; avoid storing ID images unless necessary and approved. Set short retention for transient evidence. |
| Profile and preferences | Photos, bio, age, distance, relationship goals, orientation where volunteered | Field-level visibility and purpose; let users edit/delete; do not reuse for advertising or unrelated inference by default. |
| Location | Approximate location, search radius, GPS if explicitly needed | Prefer coarse location and foreground/one-time access; never display precise home/work location; avoid long-term location trails. |
| Messages and media | Text, images, attachments, safety reports | Treat as highly sensitive; restrict staff access; define retention, abuse-review and legal-hold rules; avoid training models on private messages by default. |
| Matching and analytics | Impressions, likes, match events, ranking reasons | Collect event metadata needed to operate and measure quality; pseudonymise where practical; avoid sensitive profiling. |
| Safety and moderation | Reports, block events, suspected fraud signals, review outcomes | Strict role-based access, audit logs, clear purpose and retention; distinguish allegation from substantiated finding; support appeal and correction. |
| Billing | Subscription status, transaction references | Use payment provider tokens; do not store card details; keep billing records only as required for service, accounting and law. |
| Device/security signals | Login history, rate-limit signals, device/IP risk indicators | Use proportionately for security; don't automatically treat shared IP/device as proof of abuse; document purpose and retention. |

## 3. Lawful processing and consent design
- Maintain a processing register for every purpose: data, purpose, lawful basis, special-category condition (if applicable), recipient, retention, international transfer, and user rights route.
- Separate necessary service processing from optional features such as AI coaching, optional personalisation, research studies and marketing.
- Do not bundle consent to sensitive data processing with unrelated terms or make optional processing a condition of basic use without legal justification.
- Record consent version, timestamp, scope and withdrawal; ensure withdrawal is as easy as giving consent.
- Make privacy notices layered and readable at the point of collection, not only in a long policy.
- If legitimate interests is considered for a purpose, document the purpose, necessity and balancing test; do not assume it can replace the Article 9 condition for special-category processing.
- Complete a DPIA before high-risk profiling, novel AI, systematic monitoring, biometrics or large-scale sensitive-data processing; reassess when the product changes.

## 4. Privacy-by-design requirements

### Collection and visibility
- **PRIV-01:** Every data field has a documented purpose, lawful basis, visibility rule and retention rule before it ships.
- **PRIV-02:** Collect the minimum precision necessary; do not require exact GPS if coarse location suffices.
- **PRIV-03:** Sensitive profile fields are optional where feasible and have explicit audience/visibility controls.
- **PRIV-04:** Private contact details are not shown to matches by default.
- **PRIV-05:** No sensitive-trait inference from photos, messages, names or location for matching, advertising or personalisation.
- **PRIV-06:** Do not sell personal data or permit ad targeting based on intimate profile fields.
- **PRIV-07:** Analytics events avoid raw message text, image contents and unnecessary direct identifiers.

### Retention, deletion and rights
- **PRIV-08:** A retention schedule specifies the period and deletion trigger for each data class, including logs, backups, reports, vendor copies and analytics.
- **PRIV-09:** Account deletion removes or de-identifies data from active systems within a documented period; backup expiry is explained and controlled.
- **PRIV-10:** Retain only limited information where necessary for legal obligations, fraud prevention, safety or dispute handling, with a documented basis, restricted access and time limit.
- **PRIV-11:** Provide a process for access, rectification, erasure, restriction, objection, portability where applicable, and complaints.
- **PRIV-12:** Deleting a profile, pausing discovery, unmatching, blocking and deleting an account are distinct actions with clear consequences.
- **PRIV-13:** Avoid retaining message content indefinitely “in case it is useful”; safety and legal retention exceptions must be specific and justified.

### Security and governance
- **PRIV-14:** Encrypt data in transit and at rest; use managed secrets, least privilege, MFA for privileged access and auditable administrative actions.
- **PRIV-15:** Restrict moderator access to message content to cases where it is necessary for a report, safety investigation or legal obligation.
- **PRIV-16:** Keep production data out of developer laptops, test fixtures, analytics demos and AI prompts unless expressly authorised and appropriately protected.
- **PRIV-17:** Use synthetic or properly anonymised data in development and testing.
- **PRIV-18:** Maintain a breach-response playbook, access reviews, vulnerability management, vendor risk register and incident exercises.
- **PRIV-19:** Perform a privacy and security review for every new data use, AI feature, integration or vendor.
- **PRIV-20:** Publish accurate explanations of how matching, verification and moderation work; do not promise absolute anonymity or safety.

## 5. Age assurance and verification architecture
Treat the following as separate assertions with distinct evidence:
1. **Age assurance:** enough evidence to determine whether the user meets the service's age threshold.
2. **Identity proofing:** evidence that a person matches a claimed identity, if Pinchd chooses to offer it.
3. **Photo/liveness check:** evidence that a submitted selfie is recent or live and matches a submitted profile image under a defined method.
4. **Account integrity:** signals that may indicate duplicate or abusive accounts, which require corroboration and appeal.
5. **Safety status:** operational outcome of review; not a guarantee that a person is safe.

Recommended approach:
- Assess legal scope and risk with counsel; document the conclusion and review date.
- Compare age-assurance methods for accuracy, robustness, reliability, fairness, accessibility, privacy, fraud resistance, cost and conversion impact.
- Prefer a minimal result token over raw documents or face templates where feasible.
- Make vendor contracts cover controller/processor roles, purpose limits, subprocessors, data locations, retention/deletion, security, breach notification, audit evidence and assistance with rights requests.
- Provide a non-discriminatory exception/review route for people who cannot use the default method, without making the process trivially bypassable.
- Measure pass/fail/manual review rates and disparities; monitor false acceptance and false rejection.
- Never label an account “safe” solely because age or identity checks passed.

## 6. AI and profiling guardrails
- AI-generated suggestions must be optional and user-reviewed.
- Do not train general-purpose models on private messages or intimate images by default.
- Separate safety classification from compatibility ranking.
- Explain the data categories used for personalisation; offer correction, reset and opt-out where appropriate.
- Do not use sensitive data to infer attractiveness, desirability or willingness to engage.
- Human review and appeal must exist for high-impact restrictions; do not treat a probabilistic model output as proof.
- Log model/version, purpose, threshold and outcome for safety decisions, while minimising retained content.
- Conduct a DPIA and security review before deploying automated message scanning, facial processing or other high-risk AI.

## 7. Third-party and international data governance
- Maintain a data-flow diagram showing systems, regions, subprocessors and transfers.
- Review every SDK: what it collects, whether it shares data, whether it builds advertising profiles and whether the feature can be implemented without it.
- Use appropriate UK international-transfer mechanisms and transfer-risk assessments where required.
- Keep a vendor register with owner, purpose, data types, risk rating, contract, security evidence, subprocessors, renewal and exit/deletion plan.
- Ensure operational teams can honour data rights across providers and backups.
- Do not allow third-party analytics or advertising SDKs to collect sensitive profile attributes or message content.

## 8. Post-launch privacy and assurance metrics
- Data-subject request volume, response time, completion and exceptions.
- Deletion completion time across primary systems and vendor systems; backup expiry compliance.
- Privileged access to profile/message/report data and unusual-access alerts.
- Security incidents, near misses, vendor incidents and remediation time.
- Age-assurance pass, fail, manual-review and appeal outcomes, with accessibility and fairness analysis.
- Proportion of profile fields with documented purpose, basis, visibility and retention (target: 100% before launch).
- DPIA completion and remediation closure for all high-risk features.
- Percentage of vendors with current due diligence and documented deletion/exit plans.
- User comprehension of privacy explanations and controls, measured with post-launch surveys or support signals.
- Audit of analytics payloads for raw messages, exact location and unnecessary identifiers.

## 9. Launch gates
1. Complete data inventory, data-flow map and processing register.
2. Obtain legal review of Article 6 bases and Article 9 conditions for each sensitive-data purpose.
3. Complete DPIA and threat model for matching/profiling, age assurance, verification, moderation and AI.
4. Determine Online Safety Act scope and relevant age-assurance / children-safety duties.
5. Select vendors only after privacy, security, accessibility, fairness and contractual review.
6. Approve retention schedule, deletion pipeline, rights-request workflow and breach response.
7. Test role-based access, audit logs, backup restoration/deletion, account recovery and abuse scenarios.
8. Review privacy notices, consent journeys and claims about badges and verification.
9. Complete pre-launch security testing and remediate critical/high issues.
10. Reassess after material changes and at a defined periodic interval.

## 10. Limitations and open decisions
- The precise Online Safety Act duties depend on Pinchd's service functionality, content, risk profile and legal classification; counsel must confirm applicability.
- ICO guidance is being updated following the Data (Use and Access) Act 2025. Recheck the current legal position before implementation and launch.
- Whether each field is special-category data depends on context and processing; do not make a blanket assumption that all dating data has the same legal status.
- Age assurance methods have trade-offs. A vendor's claim or regulatory listing is not a substitute for Pinchd's own due diligence and monitoring.
- This is a product research baseline, not legal advice or a completed DPIA.
- UI design remains paused pending completion of the research and requirements package.

## Decision summary
**Recommended direction:** privacy by default, minimal sensitive-data collection, clear purpose and lawful-basis records, a DPIA before high-risk processing, strong deletion and rights workflows, vendor accountability, and age assurance that is both effective and privacy-protective. Verification badges must state exactly what was checked and what remains unverified.

## Source register
1. ICO, Data protection by design and by default (updated 5 Feb 2026): https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/accountability-and-governance/guide-to-accountability-and-governance/data-protection-by-design-and-by-default/
2. ICO, Guide to the data protection principles (updated 23 Mar 2026): https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/data-protection-principles/a-guide-to-the-data-protection-principles/
3. ICO, Special category data: https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/lawful-basis/special-category-data/
4. ICO, Biometric data: https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/lawful-basis/special-category-data/what-is-special-category-data/biometric-data/
5. Ofcom, Age assurance duties (updated 2 Sep 2026): https://www.ofcom.org.uk/online-safety/protecting-children/age-assurance
6. Ofcom, Use of Age Assurance Report 2026 (published 15 Jul 2026): https://www.ofcom.org.uk/online-safety/protecting-children/use-of-age-assurance-report-2026
7. GOV.UK, Ofcom report on the use of age assurance: https://www.gov.uk/government/publications/ofcom-report-on-the-use-of-age-assurance


## 9. Photo/selfie verification, biometric privacy and evidence handling (21 October 2026)

### Separate the purposes
Pinchd must distinguish profile-photo authenticity checks, selfie verification, age assurance and legal identity verification. They answer different questions and should not be described interchangeably. A profile photo comparison may indicate that a live capture resembles uploaded photos; it does not establish a legal identity or guarantee safe conduct. Age assurance estimates or verifies age to a specified confidence; it does not necessarily identify the person. Identity verification establishes identity only to the extent of the method and source used.

Hinge's public help page describes a video-selfie process for a visible Selfie Verified badge and notes that some regions require biometric Face Check. Hinge also warns that verification is not a safety guarantee. Source: https://help.hinge.co/hc/en-us/articles/10303221435539-What-is-Selfie-Verification

### UK data-protection implications
A facial image is personal data when a person is identifiable. It becomes special-category biometric data under UK GDPR when processed through specific technical means for the purpose of uniquely identifying a person. The actual legal classification depends on what the supplier does, not only on whether the product calls it "selfie verification". Before implementation, obtain UK privacy/legal review of the processing purpose, lawful basis, Article 9 condition where relevant, transparency, DPIA requirements, controller/processor roles, international transfers, retention, security and data-subject rights.

The ICO explains that biometric recognition systems can process special-category biometric data and require an appropriate lawful basis and special-category condition where that classification applies. Source: https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/lawful-basis/biometric-data-guidance-biometric-recognition/

### Proposed data-minimisation rules
- Prefer a supplier design that returns a narrowly scoped result (for example, verification passed/failed/retry needed) and deletes raw video, templates and derived biometric data promptly when no longer needed.
- Do not retain selfie video or face templates just in case. Define an explicit purpose and retention period for any evidence retained for disputes, fraud investigation or legal obligations.
- Do not reuse verification images for model training, ad targeting, attractiveness scoring, emotion inference or unrelated profiling.
- Provide an accessible alternative or human review route for users unable to complete a selfie workflow because of disability, device limitations, religious/cultural constraints, appearance changes or technical failure, where consistent with the risk and applicable law.
- Do not expose raw verification media or detailed fraud signals to other users. Publicly display only the minimum status needed for informed decisions.
- Restrict staff access, log sensitive access, encrypt data in transit and at rest, and define breach and supplier-incident procedures.
- Ensure account deletion and verification deletion/retention rules are documented separately; legal holds and necessary fraud evidence must be narrow, time-limited and explained.

### Age assurance is a distinct requirement
Ofcom's July 2026 report reviews how regulated services, including online dating, have used age assurance under the Online Safety Act. It reports early evidence and does not claim final conclusions about effectiveness. Pinchd must separately determine the applicable legal duties and suitable age-assurance method before launch; do not assume selfie verification alone satisfies age assurance. Source: https://www.gov.uk/government/publications/ofcom-report-on-the-use-of-age-assurance

### Decision gate
No biometric or selfie-verification vendor should be selected until Pinchd has compared verification effectiveness, false rejection, accessibility, fraud resistance, privacy architecture, data location/transfers, deletion guarantees, cost, service availability and independent assurance. Record the decision in a dedicated architecture decision record and complete a DPIA where required before production use.


## 10. Full decision research: age assurance and biometric verification (21 October 2026)

### 10.1 Legal and regulatory baseline to verify before launch
Ofcom's age-assurance guidance says methods capable of being highly effective at determining whether someone is 18 include credit-card checks, digital identity services, email-based age estimation, facial age estimation, mobile-network operator checks, open banking and photo-ID matching. It says self-declaration and payment methods that do not require the user to be over 18 are not capable of being highly effective at 18. Method suitability depends on the service and implementation; this list is not a blanket legal sign-off for Pinchd. Source: https://www.ofcom.org.uk/online-safety/protecting-children/age-assurance

Ofcom's July 2026 report evaluates early evidence from the first six months of child-protection duties and includes the online-dating sector. Ofcom says its findings are early, not final conclusions on effectiveness, and calls on services to follow its guidance, conduct ongoing vendor due diligence and comply with privacy obligations. The regulated service remains responsible when a supplier performs the check. Sources: https://www.gov.uk/government/publications/ofcom-report-on-the-use-of-age-assurance and https://www.ofcom.org.uk/online-safety/protecting-children/use-of-age-assurance-report-2026

The ICO states that processing special-category data requires both an Article 6 lawful basis and an Article 9 condition, and that a DPIA is required for processing likely to be high risk. Its biometric-recognition guidance says organisations must take a data-protection-by-design approach and complete a DPIA before using a biometric recognition system. Legal classification depends on the processing, including whether technical processing is used to uniquely identify a person; a provider's marketing label is not determinative. Sources: https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/lawful-basis/a-guide-to-lawful-basis/special-category-data and https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/lawful-basis/biometric-data-guidance-biometric-recognition/how-do-we-demonstrate-our-compliance-with-our-data-protection-obligations/

### 10.2 Distinguish the questions being answered
- **Age assurance:** determines or estimates whether a person is over a threshold, such as 18. It does not necessarily establish legal identity.
- **Identity verification:** checks whether a person is the named identity on an accepted document or data source. It may collect more information than age-only checks.
- **Selfie/liveness verification:** checks whether a live person is present and/or whether a capture resembles a reference photo. It does not prove that the profile's claims are true or that the person will behave safely.
- **Profile authenticity review:** looks for impersonation, stolen images, scam behaviour or misleading representation. This may use reports and contextual evidence and should not be reduced to a biometric match score.
- **Safety screening:** assesses conduct and risk signals; it is not a substitute for age or identity assurance and must not become a desirability judgement.

Do not display one generic “verified” badge unless the scope is explicitly explained. The status shown to other users should disclose no more than necessary.

### 10.3 Method comparison for a UK adult dating service
**A. Facial age estimation**
- User experience: usually a short face capture; may be fast and not require an identity document.
- Strength: can estimate whether the person is likely above a threshold; Ofcom lists it as capable of highly effective age assurance when implemented appropriately.
- Limits: it estimates age, not identity; uncertain cases need a fallback. Performance can vary with age range, capture conditions and demographic representation. Require independent evaluation on relevant populations, confidence thresholds and a documented fallback.
- Privacy: face images are personal data; if used for unique identification, biometric special-category obligations may apply. Prefer processing that returns only the age-threshold result and deletes captures promptly, subject to the agreed legal basis and retention need.
- Access: consider users whose appearance, disability, assistive technology or camera/device makes capture difficult; provide a genuine alternative.

**B. Photo-ID matching**
- User experience: users provide an identity document and usually a selfie/liveness capture.
- Strength: can support age verification and identity checks, depending on document validation and implementation.
- Limits: creates greater friction; documents can be unavailable, expired, damaged or not supported. Documents can be stolen or forged, and the check does not guarantee conduct or account control indefinitely.
- Privacy: usually the highest data-minimisation burden among the common options. Avoid retaining a full ID image or document number if a verified threshold/result is sufficient. Define strict retention and access rules.
- Access: offer alternatives for people without accepted documents or with document/name mismatch; provide manual review without discriminatory assumptions.

**C. Digital identity or reusable age token**
- User experience: can reduce repeated document uploads where a trusted identity provider or wallet is available.
- Strength: may share a minimal attribute such as “over 18” instead of full identity data.
- Limits: adoption and interoperability vary; users may not have a compatible credential. Validate assurance level, revocation, replay prevention, provenance and the token's actual claims.
- Privacy: potentially strong minimisation if the relying service receives only an age result and a non-linkable or appropriately scoped proof; confirm real implementation and metadata leakage.
- Access: retain a fallback for people without a digital identity credential.

**D. Mobile-network operator or open-banking checks**
- User experience: may be low-friction for eligible users, but relies on the user having a supported account or relationship.
- Strength: Ofcom lists both as methods capable of highly effective age assurance at 18 in suitable implementations.
- Limits: data availability, account ownership, network and provider coverage vary. A bank account or phone contract is not automatically proof of the current user's age unless the exact method establishes that fact.
- Privacy: avoid collecting transaction details or unrelated account information. Require the supplier to document what is queried, what Pinchd receives, retention and secondary use.
- Access: not suitable as the only route where coverage is incomplete or exclusionary.

**E. Email/phone age estimation or inference**
- Ofcom lists email-based age estimation among methods capable of highly effective age assurance at 18 when appropriately implemented; not all generic email/phone heuristics meet that bar. Ofcom's 2026 report calls for services relying on age inference for child-protection duties to move to listed highly effective methods unless they can prove their method is highly effective with reliable, compelling evidence. Use only after method-specific legal review and independent validation; do not equate an email address or phone number alone with age verification. Source: https://www.ofcom.org.uk/online-safety/protecting-children/use-of-age-assurance-report-2026

**F. Self-declared date of birth**
- Strength: lowest friction and useful as an initial declaration.
- Limit: Ofcom says self-declaration is not capable of being highly effective at 18 for relevant Online Safety Act duties. It must not be the sole age-assurance control where highly effective assurance is required.
- Privacy: minimal additional collection, but a false declaration is easy to make.

### 10.4 Vendor diligence: what to demand, not a premature winner
Public vendor documentation reviewed includes Yoti's facial age-estimation white paper and Persona's age-assurance/privacy documentation. Yoti publishes accuracy/bias discussion and states that performance depends on data representation; vendor-reported results need independent evaluation on Pinchd's intended population. Persona documents multiple methods, including selfie age estimation, government-ID checks and reusable age tokens, and states that customers control important retention choices. Vendor claims are not independent assurance. Sources: https://www.yoti.com/wp-content/uploads/2024/09/Yoti-Age-Estimation-White-Paper-September-2024-PUBLIC.pdf and https://help.withpersona.com/articles/5I4utwPPHF4TilUvZTJCsh/ and https://withpersona.com/legal/privacy-policy

Require each shortlisted supplier to provide:
- Method-specific independent testing and evidence of compliance with Ofcom's highly effective age-assurance guidance for the exact use case.
- Error rates and confidence intervals near the 18-year threshold, false acceptance and false rejection, subgroup performance, test population and date of testing.
- Liveness/spoof resistance and red-team results, including replayed images/video, deepfakes, virtual cameras and stolen documents where relevant.
- Data flow diagram, controller/processor roles, subprocessors, hosting regions, international transfers, encryption and access logging.
- Exact retention/deletion defaults for raw images, video, document images, face templates, logs and derived attributes; whether deletion can be verified contractually and technically.
- Contractual prohibition on advertising, unrelated profiling and model training with Pinchd user data unless separately and lawfully agreed (default should be no reuse).
- DPIA support, security certifications/audit reports, incident notification SLAs, business continuity, support SLAs, penetration-test summaries and vulnerability handling.
- Accessibility and fallback options, user completion/failure data, supported documents/devices, languages and support for edge cases.
- Price structure, minimum commitments, per-check and retry charges, volume tiers, geographic coverage and exit/export/deletion assistance.
- Clear evidence of how the service meets the relevant regulatory method criteria; do not accept a generic “compliant” statement as sufficient.

### 10.5 Recommended launch posture
**Provisional recommendation, subject to UK legal review:** treat age assurance as a mandatory pre-launch workstream, separate from the optional public selfie-verification badge. First establish the exact age-assurance obligation for Pinchd's service model and intended audience. Shortlist at least two providers and compare a privacy-minimising age-only method (for example facial age estimation or an age token) against a document-backed fallback. Do not choose a supplier or mandate biometric selfie checks for every account until the DPIA, accessibility assessment, vendor due diligence, cost analysis and test results are complete.

If a highly effective age check is required, self-declaration alone is insufficient. Provide a retry and alternative route for uncertain/failed checks, with a documented manual review process. Do not let an unverified or under-age account access adult discovery while a required check is unresolved; ensure enforcement is proportionate and clear.

### 10.6 Data retention and privacy controls
- Collect the minimum result needed: normally an over-18 threshold outcome, method/version, timestamp and a narrowly scoped transaction reference—not the ID document itself unless necessary.
- Prefer immediate deletion of raw capture data after result determination unless a documented, necessary and lawful reason justifies a defined retention period.
- Avoid storing reusable face templates centrally where a vendor can complete the check and return a signed result instead.
- Separate public badge status from private verification evidence; only a small, authorised trust/safety group may access sensitive evidence for defined cases.
- Document deletion timelines, backups, dispute holds and supplier deletion confirmation. User account deletion must trigger the relevant verification-data deletion workflow unless a specific lawful retention requirement applies.
- Do not use age-assurance images for matching, recommendations, profile coaching, advertising, identity graphing or unrelated analytics.
- Complete a DPIA before any biometric recognition system is used; record alternatives considered, necessity/proportionality, risks, mitigations and residual risk acceptance.

### 10.7 Acceptance criteria and launch gates
- The legal minimum age and applicable age-assurance duty are documented and reviewed before launch.
- The chosen method has evidence that it meets the applicable Ofcom standard for the specific purpose; self-declaration alone is not treated as sufficient where highly effective assurance is required.
- At least one workable fallback exists for failed, uncertain or inaccessible checks, without granting access before a required check is satisfied.
- Supplier due diligence, data-flow mapping, processor terms, retention/deletion controls, security review and DPIA are completed before production use where required.
- Pinchd stores only the minimum result and audit evidence necessary; raw media is deleted according to the approved policy.
- Public badges describe the check precisely and do not claim a person is safe or truthful.
- False acceptance, false rejection, retry, abandonment, appeal and accessibility metrics are monitored and reviewed by subgroup where lawful and privacy-safe.
- No vendor is selected solely on headline accuracy or price; evidence, privacy, accessibility, resilience, support and exit options are compared.

### 10.8 Open decisions
- Confirm service classification and exact legal age-assurance requirements for the planned UK launch.
- Decide whether all users need a public selfie badge, whether it is voluntary, and what additional verification triggers are justified.
- Select the acceptable age-assurance method and fallback after supplier evaluation.
- Set evidence retention, appeal, manual-review and deletion policies.
- Decide whether a vendor's output can be stored as a minimal signed result and whether an independent privacy/security review is required.

### Sources
1. Ofcom, Age assurance duties under the Online Safety Act: https://www.ofcom.org.uk/online-safety/protecting-children/age-assurance
2. Ofcom, Use of Age Assurance Report 2026: https://www.ofcom.org.uk/online-safety/protecting-children/use-of-age-assurance-report-2026
3. GOV.UK, Ofcom report on the use of age assurance (16 July 2026): https://www.gov.uk/government/publications/ofcom-report-on-the-use-of-age-assurance
4. ICO, Special category data: https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/lawful-basis/a-guide-to-lawful-basis/special-category-data
5. ICO, Demonstrating compliance with biometric recognition obligations: https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/lawful-basis/biometric-data-guidance-biometric-recognition/how-do-we-demonstrate-our-compliance-with-our-data-protection-obligations/
6. Yoti, Age Estimation White Paper (September 2024): https://www.yoti.com/wp-content/uploads/2024/09/Yoti-Age-Estimation-White-Paper-September-2024-PUBLIC.pdf
7. Persona, Age assurance methods: https://help.withpersona.com/articles/5I4utwPPHF4TilUvZTJCsh/
8. Persona, Processor Privacy Policy (9 April 2026): https://withpersona.com/legal/privacy-policy


## 11. Startup-friendly supplier shortlist and cost strategy (21 October 2026)

### Context
DC REGENT GROUP LIMITED is an early-stage company and the first three identity/age vendors reviewed appeared too expensive. The target is a low-commitment, usage-based supplier that can support testing and a small beta without a large setup fee or monthly minimum. Do not select solely on lowest headline price: Pinchd must confirm the exact age-assurance method meets the applicable Ofcom standard, supplier privacy/security, accessible fallback, data retention/deletion and service continuity.

### Publicly advertised options worth checking
**Stripe Identity** — UK pricing page currently advertises document + selfie verification at £1.25 per completed verification and the first 50 verifications free. It is straightforward pay-as-you-go and has prebuilt UI, but the public pricing describes identity-document verification, not a standalone facial age-estimation product. It may be useful for an ID fallback or a small pilot; confirm whether it meets Pinchd's exact age-assurance use case before treating it as the primary age gate. Source: https://stripe.com/gb/identity

**Didit** — public age-verification documentation advertises 500 free verifications per month and then age estimation at US$0.10/check, with an ID-verification fallback advertised at US$0.15/check when triggered. This is a potentially startup-friendly price structure. These are vendor-published claims; confirm UK availability, full terms, tax/currency, the exact method and independent evidence of meeting the relevant Ofcom standard, biometric-data handling, retention, false-accept/false-reject performance and fallback behaviour in writing. Source: https://updates.didit.me/id/solutions/age-verification/

**WeVerify** — public product page advertises browser-based age estimation with liveness at €0.20/check and says it returns an over/under-threshold result without retaining personal data after the check. This is a promising low-friction price indication, but treat its compliance statements as vendor claims pending evidence and contractual review. Confirm UK service availability, processor/subprocessor terms, retention claims, audit evidence and the exact Ofcom standard before use. Source: https://weverify.uk/age.html

**Certivus Verify** — public page advertises a free sandbox with 50 checks/month, a £20 one-off deposit credited against first-month usage, fresh checks at £1.60 and pass reuse at £0.05. It may be useful if its reuse network is relevant, but the low reused-pass price depends on users already having a Certivus Pass. Verify company maturity, service references, assurance evidence, UK data residency claim, contractual terms and service resilience. Source: https://certivus.com/identity-verification

**Proof of Age** — public pricing page lists a free developer tier for demo/testing and paid plans from £99/month for 2,000 proofs. This is less attractive before Pinchd has predictable volume, but may be worth revisiting after beta if its specific assurance methods and economics fit. Source: https://proofofage.app/pricing

**AgeCheck API** — public page advertises €0.10 per verification, no setup fee or monthly minimum and 1,000 free checks for early-access members. Treat this as a vendor's launch offer and verify availability, business maturity, security, legal assurance evidence, data processing and service-level commitments before any production decision. Source: https://www.agecheckapi.com/pricing

### Indicative cost illustration (not a forecast)
For 100 successful checks in a month, using only the advertised headline rates before VAT, retries, fallback checks or engineering/support costs:
- Didit at US$0.10/check would be US$10 if the monthly free allowance does not apply or is exceeded; within its advertised 500 free checks, usage fees could be zero subject to terms.
- WeVerify at €0.20/check would be €20.
- Stripe Identity at £1.25/check would be £125 after the advertised 50 free verifications.
- Certivus fresh checks at £1.60/check would be £160, after its £20 activation deposit is accounted for separately as a credit against usage.

Currencies are not directly comparable, and these examples exclude any tax, FX, retries, manual review, fallback checks and integration cost. They are only useful for shortlisting.

### Recommended startup approach
1. **Do not integrate three providers.** Shortlist Didit and WeVerify for a low-cost age-estimation-first flow, and keep Stripe Identity or another document provider as a fallback candidate only if needed. Add Certivus as a quote/assurance comparison if its terms are attractive.
2. Ask each shortlisted provider to confirm in writing: the specific method used; evidence it meets Ofcom's highly effective age-assurance criteria for Pinchd's use case; false acceptance/rejection rates near 18; subgroup performance; liveness/spoof testing; what data Pinchd and the supplier retain; deletion time; hosting and transfers; subprocessors; DPIA support; accessibility alternatives; and exact all-in pricing including retries and fallback.
3. Ask for a startup/beta arrangement: free sandbox, no minimum commitment, pay-as-you-go, no annual contract, no setup fee, low-volume cap and written notice before pricing changes.
4. Start in sandbox, test normal, failed, borderline, low-light, device/accessibility and retry paths, then run a small controlled pilot. Do not let an unverified user access features for which age assurance is a legal prerequisite.
5. Before production, complete supplier security/privacy review, processor terms, data flow mapping, DPIA where required and legal review. Confirm the company's Hiscox PI and cyber policies with the broker; do not assume any provider's liability or security posture is covered by the policies.
6. Keep mandatory age assurance free to the user and outside Premium. Re-check unit economics as the user base grows.

### Vendor scoring gate
Score each vendor against: (a) Ofcom method evidence and independent assurance; (b) privacy and data minimisation; (c) false acceptance/rejection and accessibility; (d) security and operational maturity; (e) all-in low-volume price; (f) integration effort; (g) contractual liability, breach notification and support; (h) exit and deletion. Legal compliance and adequate security are pass/fail gates, not tradeable scores. The cheapest provider must not win if it cannot demonstrate the required assurance.

### Research conclusion
The most startup-friendly public pricing structures found are Didit's advertised 500 free checks/month with low per-check pricing, WeVerify's advertised €0.20/check, and Stripe Identity's 50 free ID checks followed by £1.25 per completed check. These are shortlist candidates only, not approved suppliers. Pinchd should run a low-cost, usage-based pilot after written assurance evidence and privacy/security review, and should not commit to enterprise minimums while volume is unknown.
