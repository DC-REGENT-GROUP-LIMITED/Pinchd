# Workstream 2 — UK legal and regulatory readiness for Pinchd

- **Research date:** 2026-10-22
- **Scope:** UK-first adult dating and social-discovery service; iOS/Android, profiles and photos, likes/comments, matching, private messaging, blocking/reporting, age assurance, optional profile verification and paid Premium subscription.
- **Status:** Research baseline, not legal advice. Counsel must validate applicability, commencement, contracts, operational implementation and launch obligations against the final product and launch date.
- **Approach:** Distinguish statutory duties and regulator guidance from prudent product controls. Do not assume age-gating alone removes all online-safety or child-privacy obligations.

## Executive decision

Pinchd should treat UK legal readiness as a launch gate, not post-launch cleanup. Highest priority: Online Safety Act scope and operational controls; UK GDPR/DPA 2018 including special-category data and DPIA; age assurance and under-18 access assessment; subscription/consumer law; equality/accessibility; vendor, security, retention and incident governance.

Ofcom published specific dating/social-discovery guidance on 18 May 2026. It describes user-to-user services as including services where users share profiles, images, messages or comments and meet through matching systems; it highlights fraud, intimate-image abuse, harassment and stalking risks. This makes a formal OSA scope assessment and documented safety programme a priority. [S1]

## 1. Online Safety Act 2023 — critical launch gate

### Applicability

Pinchd is expected to enable user-to-user interaction through profiles, photos, comments and messages. On this feature description it likely warrants a full OSA scope assessment, but counsel must document the final determination using Ofcom's current tools and guidance. Do not state categorically that every duty applies without classifying the actual service and content features. [S1][S2]

### Required work

- Complete and retain the service scope assessment.
- For an in-scope service, complete the illegal-content risk assessment using Ofcom's applicable categories and guidance; keep it current and reassess before significant service changes.
- Assess whether children are likely to access the service and, if applicable, complete the children's risk assessment and implement relevant duties.
- Create clear, accessible terms and policies; a reporting/complaints route; a process to review and remove illegal content when required; record-keeping and review processes; named operational ownership.
- Map dating-specific abuse cases: romance fraud, impersonation, harassment, stalking and ban evasion, threats/extortion, intimate-image abuse, coercive control, grooming, illegal sexual content and misuse of reports.
- Provide trained moderation and escalation workflows, evidence preservation rules, urgent safety triage, appeal/review paths and documented law-enforcement escalation criteria.
- Treat block enforcement as a safety-critical, server-authoritative action across discovery, likes, messages, notifications, caches, queues and account re-entry.
- Review Ofcom guidance for practical protections for women and girls and integrate applicable measures into safety design.
- If Pinchd permits pornographic content, assess Ofcom's highly effective age assurance requirement for content accessible in profiles and private interactions. Recommended product default pending legal review: prohibit user-generated pornography. [S1][S2]

### Evidence to retain

Scope determination; risk assessments; policy versions; moderation and complaints procedures; risk-control mapping; training and escalation records; incident/review records; significant-change reviews; evidence that controls work in production.

## 2. UK GDPR, Data Protection Act 2018 and Data (Use and Access) Act changes — critical

### Data map and lawful basis

Create a data inventory and map flows for account/contact data, profile content, sexual orientation/sex-life information, relationship preferences, location, messages, reports, ID/age-assurance results, optional verification, payment/entitlement data, analytics, logs and support/moderation evidence.

For every purpose, record controller/processor roles and purpose; Article 6 lawful basis; where special-category data is processed, an applicable Article 9 condition and required DPA 2018 safeguards/policy documentation; minimisation, access, retention, deletion, sharing and user-rights handling; and whether data is inferred or used to make decisions about individuals.

Sexual orientation and data concerning sex life are special-category data; data revealing political opinions, religious/philosophical beliefs, health and biometric data used for identification can also be special-category data. Inferences may qualify depending on intended inference or differential treatment. Do not use sensitive traits for matching/ranking without a documented purpose, lawful basis, Article 9 condition, fairness review and legal sign-off. [S3]

### DPIA and privacy by design

Start a DPIA before building high-risk profiling or biometric/age-assurance flows. Include matching/ranking, location, sensitive preference fields, safety signals, AI coaching, moderation evidence, data sharing and children-access risks. Evaluate discrimination, physical harm, stalking, re-identification, loss of confidentiality and exclusion; document mitigations and residual risk. The ICO says a DPIA is required where processing is likely to result in high risk; doing it early is prudent even before the final trigger is settled. [S4][S5]

Controls:
- Plain-language privacy notice, layered at collection points.
- Purpose-specific collection and opt-in for genuinely optional features.
- No exact location in profile display; no triangulation; avoid storing precise coordinates longer than needed.
- Keep age-assurance payloads and ID images with the vendor where possible; receive a minimal pass/fail, age-band or token result only when sufficient. Verify actual Didit data flow and retention.
- Distinguish age assurance from optional profile-photo verification.
- No raw messages, IDs, exact coordinates, access tokens or sensitive content in analytics/logs.
- Access controls, MFA, audit trails, encryption, retention/deletion and vendor propagation.
- Rights process for access, correction, objection/restriction and erasure; explain lawful exceptions and backup/record retention rather than promising instant deletion from every immutable backup.
- Data breach assessment and notification runbook.
- Confirm whether Pinchd must appoint a DPO or representative and whether an ICO data protection fee is payable; check current rules before launch. [S6]

### International transfers and vendors

Create a subprocessor/vendor register for AWS, age assurance, email/SMS, push, analytics, moderation, support and payments. Review processing terms, subprocessor changes, data location, international transfers, access support, retention/deletion, security assurance, incident notification and exit/export. UK region selection alone does not establish compliance if vendors or support access involve other countries.

## 3. Age assurance and child access — critical

- Define a clear 18+ policy and age assurance workflow; method must be risk-based and proportionate, resistant to trivial circumvention, accessible and designed to minimise data.
- Document a children's access assessment: whether children are likely to access Pinchd in reality, considering appeal, distribution, user behavior, marketing and actual evidence. An adults-only label alone does not settle the Children's Code question.
- ICO guidance says the Children's Code applies to information society services likely to be accessed by children, not only services aimed at children. [S7][S8]
- If children are likely to access the service, assess Code standards and the appropriate age-assurance/protection approach. If it is not appropriate for children, focus on effective access restriction and document why the approach is proportionate. [S7][S8]
- Ofcom's OSA age-assurance duties are feature/content-dependent. Dating services that permit pornographic content are specifically called out for highly effective age assurance; confirm relevant rules and current guidance with counsel. [S1][S2]
- Design appeals for false rejection and accessibility/discrimination risks. Limit age-check data to age assurance; prohibit reuse for unrelated profiling.
- Test underage detection, retry/failure, manual review, appeals, vendor outage, account recovery and deletion.

## 4. Subscriptions, consumer protection and app stores — critical before charging

Approved product direction: one optional Premium tier, monthly-first, no free trial, simple cancellation, no paid safety/consent privileges, and no pay-to-win romantic compatibility. This product policy does not replace legal compliance.

The DMCC Act 2024 subscription-contract chapter and Consumer Subscription Contracts (Cancellation and Information) Regulations 2026 require a launch-date-specific commencement and applicability check. Government published dedicated subscription-contract guidance on 19 October 2026, so this needs review before billing implementation. Do not assume every subscription provision is already in force merely because the Act exists; the legislation page identifies some provisions as prospective. [S9][S10]

Pre-launch actions:
- Obtain legal review of applicable subscription rules and effective dates at intended launch.
- Display price, billing interval, renewal, key restrictions, cancellation method and material terms clearly before purchase.
- Use affirmative consent and durable confirmation/receipts as applicable.
- Provide simple cancellation and any legally required reminders, cooling-off rights/refunds and notices if applicable.
- Avoid misleading claims about likes, profile visibility, compatibility, success, scarcity or guaranteed outcomes.
- Make free versus Premium entitlements clear; no hidden ranking advantage, fabricated likes or deceptive urgency.
- Implement server-side entitlement state, receipt validation, renewal/expiry/refund/revocation handling, restore purchases and support.
- Check current Apple App Store and Google Play payment, subscription, cancellation, refund, account deletion and UGC policies before release.
- Review VAT, corporate/tax and accounting obligations with a qualified adviser.

## 5. Consumer marketing and unfair commercial practices — critical

The DMCC Act's unfair-commercial-practices regime applies to practices from 6 April 2025 according to current government guidance. Claims and interface design must not mislead users or omit material information. [S11]

- Do not fabricate likes, matches, activity, scarcity or other users.
- Do not suggest payment guarantees dates, compatibility or success.
- Explain material limits and renewal terms in context.
- Make cancellation and account deletion discoverable, not deliberately obstructive.
- Ensure testimonials, influencer promotions, endorsements and referral incentives are accurately disclosed and substantiated.
- Substantiate safety, verification and compatibility claims; verification does not guarantee identity, good intent or safety.
- Keep pricing and promotion tests under review for unfair or discriminatory outcomes.

## 6. Equality, discrimination and accessibility — high priority / counsel review

Have UK counsel assess the Equality Act 2010 and relevant service-provider duties against final eligibility, gender identity, discovery preferences, safety and verification policies. Avoid simplistic conclusions about lawful exceptions in dating: service type, intimate nature, protected characteristics and exact operation matter.

- Keep gender identity distinct from discovery preferences.
- Preserve the approved rule that women includes cis and trans women and do not default to cis/trans exclusion.
- Audit age assurance, photo verification, ranking, moderation and support outcomes for disparate failure rates; provide accessible appeal/review.
- Avoid inferring protected traits from photos, names, location or behavior.
- Provide accessible mobile flows, readable terms, screen-reader support, sufficient contrast, alternatives to voice-only interactions and accessible reporting.
- Review any sex-based eligibility, filtering or exclusion with counsel before launch; don't infer legal permission from competitor behavior.

## 7. Location, stalking and user safety — high priority

- Collect the least precise location needed; do not show home/work coordinates or exact distance if it could enable triangulation.
- Obtain explicit user action before travel mode; show clear expiry and visibility; avoid revealing movement history or home location.
- Block and report must be available from profiles and conversations, with clear effects and no notification to the blocked person that reveals the reporter.
- Prevent recontact through new accounts as far as proportionate and lawful; define appeal and false-positive handling.
- Avoid push notification content that reveals sensitive relationship or message information on lock screen by default.
- Define urgent threat, stalking, intimate-image abuse and imminent physical-harm escalation procedures, including evidence handling and staff safety.

## 8. Security, cyber resilience and operational governance — high priority

- Threat model and asset/data-flow inventory.
- MFA, least privilege, role separation and audited staff access.
- Secure SDLC, peer review, dependency/secret scanning, CI tests and environment separation.
- Encryption in transit and at rest; managed secrets, rotation and least-privilege service identities.
- Rate limits, abuse detection and secure account recovery.
- Private media, short-lived access, validated uploads and malware/content checks as appropriate.
- Backup restoration tests, incident response, breach notification assessment and vendor outage plans.
- Independent penetration test or proportionate security assessment before public launch and after material changes.
- Documented retention for messages, reports, banned-account evidence, audit logs, backups and legal holds.

## 9. AI and automated decision features — review before release

AI coaching is optional, user-reviewed and never auto-published or auto-sent. Before release:
- Document purpose, data inputs, vendor retention/training terms, user disclosure, opt-out and deletion.
- Never send sensitive profile fields or private messages to an AI vendor by default.
- Prevent generation of harassment, manipulation, discriminatory advice or unsafe prompts.
- Ensure suggestions are editable and user-controlled.
- Keep AI outputs out of authoritative safety decisions and romantic ranking unless a separate documented assessment approves a narrowly defined use.
- Assess profiling/automated-decision rights, transparency, discrimination and DPIA implications with counsel.

## 10. Corporate, intellectual property and launch housekeeping

Before launch, obtain advice on company contracting and liability, app terms, privacy/cookie notices for web properties, intellectual property and asset licences, brand/trademark clearance for Pinchd, employment/contractor IP assignments, insurance, tax/VAT and complaints handling. These topics were not exhaustively researched in this pass and remain open.

## Prioritised action register

| Priority | Action | Evidence required |
|---|---|---|
| P0 | OSA scope and feature/content classification | Signed scope memo; Ofcom tool result; duty checklist |
| P0 | Illegal-content and children-access assessments as applicable | Approved risk assessments and review cadence |
| P0 | DPIA and Article 6/9 lawful-basis schedule | DPIA, data map, processing register, special-category condition |
| P0 | Age assurance design/vendor review | Data-flow, accuracy/failure/appeal tests, retention and contract review |
| P0 | Moderation/report/block/urgent-harm operations | Runbooks, staffing plan, test evidence, escalation ownership |
| P0 | Subscription legal review | Commencement check, consumer journey review, cancellation/refund specification |
| P0 | Security/privacy readiness | Threat model, access model, retention/deletion tests, incident plan |
| P1 | Equality/accessibility review | Counsel memo, accessibility test report, bias/appeal checks |
| P1 | Vendor and transfer assessment | Subprocessor register, DPAs, transfer assessment where applicable |
| P1 | Marketing claims and brand/IP clearance | Claim substantiation, trademark and asset licence checks |
| P1 | AI feature review | AI DPIA delta, vendor terms, safety tests, opt-in/opt-out design |
| P1 | ICO fee and governance check | Fee-check result, DPO/representative determination |

## Explicit unresolved questions

1. Will launch be UK-only, or will EEA/other countries be served at launch?
2. Will user-generated nudity or pornography be prohibited? Recommended initial position: prohibit, with policy and technical controls, pending counsel.
3. What exact age-assurance data does the configured Didit workflow return and retain? Verify live configuration; liveness being configured does not prove photo matching or end-to-end compliance is complete.
4. What legal entity will operate the service and act as controller?
5. What is the actual moderation coverage model, response SLA, escalation availability and budget?
6. What billing routes will be used in iOS, Android and any web checkout?
7. What is the retention schedule for messages, reports, banned-account evidence, audit logs and backups?
8. What lawful basis and Article 9 condition apply to each sensitive field and inference?
9. Which Ofcom OSA duties apply to the final service configuration and content types?
10. What provisions of the new subscription regime are in force on the planned launch date?

## Sources

[S1] Ofcom, Dating and social discovery: know the online safety risks, rules and how to comply, published 18 May 2026: https://www.ofcom.org.uk/online-safety/illegal-and-harmful-content/dating-and-social-discovery-know-the-online-safety-risks-rules-and-how-to-comply

[S2] Ofcom, Age assurance duties under the Online Safety Act, updated 2 September 2026: https://www.ofcom.org.uk/online-safety/protecting-children/age-assurance

[S3] ICO, Special category data: https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/lawful-basis/a-guide-to-lawful-basis/special-category-data

[S4] ICO, How do we do a DPIA?: https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/accountability-and-governance/data-protection-impact-assessments-dpias/how-do-we-do-a-dpia/

[S5] Data Protection Act 2018, section 123: https://www.legislation.gov.uk/ukpga/2018/12/body/2026-09-30

[S6] GOV.UK, Pay the data protection fee: https://www.gov.uk/data-protection-register-notify-ico-personal-data

[S7] ICO, Services covered by this code: https://ico.org.uk/for-organisations/uk-gdpr-guidance-and-resources/childrens-information/childrens-code-guidance-and-resources/age-appropriate-design-a-code-of-practice-for-online-services/services-covered-by-this-code/

[S8] ICO, Age assurance for the Children's code: https://ico.org.uk/about-the-ico/what-we-do/information-commissions-opinions/age-assurance-for-the-children-s-code/

[S9] GOV.UK, Subscription contracts guidance, published 19 October 2026: https://www.gov.uk/government/publications/subscription-contracts

[S10] Digital Markets, Competition and Consumers Act 2024, Part 4 (live legislation and commencement annotations): https://www.legislation.gov.uk/ukpga/2024/13/part/4/2026-09-30

[S11] GOV.UK / CMA, Unfair commercial practices, updated 18 November 2025: https://www.gov.uk/government/publications/unfair-commercial-practices-cma207/unfair-commercial-practices

## Research limitations

This is a focused first-pass desk review, not a complete legal opinion. It does not replace a full review of the Online Safety Act, UK GDPR/DPA 2018 as amended, current ICO/Ofcom guidance, app-store contracts, vendor terms, tax obligations or the exact launch feature set. Laws and commencement provisions may change. Revalidate all conclusions with UK-qualified counsel immediately before launch and after significant feature or jurisdiction changes.


## Addendum — source verification and implementation detail (2026-10-22)

### A. Online Safety Act: specific operating obligations

Ofcom's current guidance says an in-scope user-to-user service must complete a suitable and sufficient illegal-content risk assessment, take proportionate steps to prevent users encountering illegal content and mitigate risks of priority offences, swiftly remove illegal content when required, explain protections in terms of service, provide easy reporting and complaints, and keep written records. New services or services newly brought in scope generally have a three-month risk-assessment window; a significant service change requires a new assessment before the change, and Ofcom recommends annual review. Pinchd should not use the outer legal deadline as a reason to launch without a completed pre-launch risk assessment. [S12][S13]

Ofcom's regulatory documents page lists updated user-to-user Illegal Content Codes of Practice issued 9 September 2026, replacing the prior version that was in force until 29 September 2026. The legal/compliance owner must use the current version and record which measures apply, rather than copying an older checklist. [S14]

Build controls around the specific risk model: registration and impersonation, fraud/scams, sexual coercion, intimate-image abuse, cyberflashing, threats, stalking, harassment, illegal sexual content and child sexual exploitation/abuse. Risk controls should include content/report pathways, prompt triage, evidence-preservation rules, repeat-offender controls, ban-evasion response, staff access restrictions, and an audit trail showing action and rationale. A report button alone is not a compliance programme.

### B. Age assurance: vendor diligence is Pinchd's responsibility

Ofcom's July 2026 age-assurance report says services using age assurance should follow its HEAA guidance in full, carry out regular due diligence on vendors, and meet privacy/data-protection obligations; outsourcing does not remove the service provider's responsibility for effectiveness. [S15]

For the Didit integration, retain a dated evidence pack: configured workflow/version; exact data fields returned; whether images or biometric templates are stored and by whom; vendor subprocessors and regions; false acceptance/rejection and fallback handling; accessibility and appeal route; security incident terms; retention/deletion behavior; and test results. A dashboard status or configured liveness step is not evidence that the whole age-assurance flow is production-ready. Separate age assurance from optional profile-photo verification.

### C. Subscription law: commencement clarified

The Consumer Subscription Contracts (Cancellation and Information) Regulations 2026 were laid on 19 October 2026 and come into force on **1 January 2027**. The Regulations apply only to subscription contracts entered into on or after the relevant DMCCA subscription chapter comes into force. The DMCCA Part 4 Chapter 2 page still labels that chapter prospective in the version reviewed. Pinchd must re-check the actual commencement instruments and the applicable launch date with counsel before activating Premium; do not assume the whole new regime applies to a subscription started before commencement, nor assume existing consumer-law duties do not apply. [S9][S10][S16]

Engineering should support compliant changes without a rewrite: versioned offer/terms, timestamped affirmative consent, durable purchase confirmation, reminders where required, cancellation and cooling-off state machine, refund/renewal/expiry events, proof of notice delivery, and an auditable entitlement history. Confirm whether purchases are in-app, web or both, since the consumer journey and app-store rules differ.

### D. App-store distribution is a product risk, not final paperwork

Apple's current App Review Guidelines explicitly say the dating category is established and new apps need a meaningfully different or improved experience; Apple also requires UGC apps to provide objectionable-content filtering, reporting, timely responses, blocking and published contact information. Apps supporting account creation must provide in-app account deletion. [S17][S18]

This is a major strategic risk for Pinchd: a well-built generic swipe/like app may be rejected even if technically sound. Prepare a clear product differentiation statement and evidence for the app review notes: relationship-goal compatibility model, user-controlled hard constraints, explainable reciprocal recommendations, and safety design. Do not rely only on a tagline or visual styling to prove meaningful differentiation.

Store policy checks must be re-run immediately before submission. Confirm Google Play's latest user-generated content, dating/safety, subscription, account deletion and billing requirements from current official Play policy pages. The exact applicable payment route depends on the app, jurisdiction and current programme terms; do not hardcode a payment architecture before that review.

### E. Privacy and data architecture consequences

- Keep profile attributes and preference fields purpose-tagged so optional sensitive fields are not automatically collected, exposed, ranked or sent to vendors.
- Separate raw vendor response/evidence from the minimum product decision (e.g. age threshold passed) and limit access to the smallest authorised set of staff/services.
- Maintain a data-flow map that covers primary DB, object storage, backups, logs, analytics, search/cache, queues, customer support, payment processors and age-assurance vendor.
- Deletion must be an orchestrated lifecycle: user-facing status, immediate access revocation, active data deletion or anonymisation, vendor propagation, cache/index/queue handling, and documented lawful retention exceptions. Define backup expiry separately and avoid claiming that every backup is immediately erased if the actual process does not do that.
- Make block/report state authoritative and synchronously effective for new access decisions; asynchronous propagation may clean secondary stores but must not permit an interim re-contact window.
- Retain a legal hold mechanism with restricted access and recorded reason/expiry; do not use indefinite retention as the default.

### F. Launch decision gates

Do not open public registration until owners can demonstrate:
1. Written OSA applicability memo and current Ofcom assessment artifacts.
2. Completed DPIA and documented Article 6/Article 9 basis for sensitive-data processing, approved by privacy counsel/DPO where applicable.
3. Age-assurance test results, vendor due diligence, appeal/failure handling and data-minimisation evidence.
4. Moderation staffing and escalation coverage matched to actual operating hours and likely report volume; tested urgent-harm playbooks.
5. End-to-end block, report, deletion, data export and account recovery tests.
6. Subscription consumer journey and commencement review signed off before first paid transaction.
7. App-store differentiation, UGC moderation, published support contact and account deletion demonstrated in the submitted build.
8. Accessibility checks for signup, age assurance, discovery, reporting, blocking, cancellation and deletion.
9. Security review, penetration test or proportionate independent assessment, backup restoration test and incident tabletop.
10. Approved privacy notice, terms, community standards, moderation policy, retention schedule and complaints process.

## Additional authoritative sources

[S12] Ofcom, Check how to comply with the illegal content rules (includes risk-assessment timing, records, removal, terms, reporting and complaints): https://www.ofcom.org.uk/Online-Safety/check-how-to-comply

[S13] Ofcom, Illegal content duties under the Online Safety Act, updated 25 June 2026: https://www.ofcom.org.uk/online-safety/illegal-and-harmful-content/illegal-content-duties-under-the-online-safety-act

[S14] Ofcom, Online safety regulatory documents and guidance (current code versions): https://www.ofcom.org.uk/online-safety/illegal-and-harmful-content/online-safety-regulatory-documents

[S15] Ofcom, Use of Age Assurance Report 2026, updated 27 July 2026: https://www.ofcom.org.uk/online-safety/protecting-children/use-of-age-assurance-report-2026

[S16] UK Parliament, Consumer Subscription Contracts (Cancellation and Information) Regulations 2026, commencement and application: https://statutoryinstruments.parliament.uk/instrument/OvlNv3Jx

[S17] Apple, App Review Guidelines, including dating-app differentiation and UGC controls: https://developer.apple.com/app-store/review/guidelines/

[S18] Apple, Offering account deletion in your app: https://developer.apple.com/support/offering-account-deletion-in-your-app/
