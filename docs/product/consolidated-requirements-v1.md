# Pinchd consolidated product requirements and technical feasibility
**Version:** 1.0 draft  
**Date:** 21 October 2026  
**Status:** Research synthesis / requirements baseline; not a final build specification.  
**Design status:** Mobile UI design remains paused until research and requirements are approved.

## 1. Purpose
Consolidate the research into one actionable requirements baseline for product, engineering, privacy, safety, operations and commercial decisions. It defines what Pinchd must do, why it matters, how it can be tested, and which decisions remain open. It does not prescribe screen layouts. The current approved technical direction is recorded in ADR-0001; exact service/version choices and implementation remain gated by PoCs and legal, privacy, security and cost review.

## 2. Product direction
**Positioning:** “Serious about compatibility, flexible about relationship goals.”

Pinchd is an inclusive adult dating product that supports people seeking long-term relationships, casual dating, short-term connections, open-minded exploration or an undecided path without treating one goal as morally superior.

The product should help people make relevant, reciprocal connections while preserving their preferences, privacy, agency and safety. It must not optimise only for swipes, session length, paid conversion or notification response.

## 3. Product principles
1. **Preference integrity:** never silently override hard dealbreakers.
2. **Reciprocity:** recommendations should consider potential relevance to both people.
3. **Explainability:** give truthful reasons for recommendations without exposing another person's private data.
4. **Safety separated from desirability:** risk signals may trigger protection and review, not romantic ranking.
5. **Useful free core:** matching, mutual messaging, blocking, reporting, essential privacy controls and legally required age assurance are not premium paywalls.
6. **Privacy by design:** collect only what is necessary, with documented purposes, lawful basis, access, retention and deletion.
7. **Honest growth:** no fake profiles, fabricated activity, misleading scarcity or deceptive availability claims.
8. **Healthy use:** users can pause, control notifications, unmatch and leave without guilt.
9. **Operational realism:** launch only within the capacity to moderate, support and respond safely.
10. **Evidence-led decisions:** distinguish external evidence, competitor claims and Pinchd hypotheses; measure outcomes after launch.

## 4. Scope

### MVP candidate scope
- Adult account creation, authentication and account recovery.
- Age eligibility and a documented assurance approach appropriate to legal scope.
- Profile creation/editing, media upload, visibility and deletion controls.
- Relationship-goal declarations and explicit hard/soft preference settings.
- Discovery with hard eligibility filtering, soft-preference ranking, blocks and prior-dismissal rules.
- Reciprocal recommendation baseline and honest “why this profile?” explanations.
- Likes, mutual matches, messaging between mutual matches and conversation controls.
- Notification preferences, per-conversation mute, discovery pause and account deletion.
- Report, block, unmatch, moderation case management, appeal and incident escalation.
- Basic verification claims with precisely scoped meaning and no unsupported safety promise.
- Subscription/billing foundation only after free/premium boundaries and storefront approach are approved.
- Minimal analytics for product quality, safety, fairness, density and unit economics.
- Staff admin console with least-privilege roles and audit trails.
- Data export/correction/deletion workflows and retention policy.
- Monitoring, backup/restore, incident response and operational dashboards.

### Explicitly out of scope until justified
- Public attractiveness/desirability scores.
- Pay-to-win hidden compatibility boosts.
- Automated message sending by AI.
- Internet-wide facial recognition or sensitive-trait inference.
- Fake profiles, synthetic activity or fabricated likes.
- Broad multi-market launch before local density and operational capacity are demonstrated.
- Microservices by default before scale/team needs justify the complexity.
- Detailed mobile visual design in this research phase.

## 5. Functional requirements and acceptance criteria

### Identity and account lifecycle
- **ID-01:** Users can register, authenticate, recover accounts and manage active sessions securely.
- **ID-02:** Age eligibility is assessed through a documented method chosen after legal, privacy, accuracy, fairness and vendor review.
- **ID-03:** Age assurance, identity proofing, photo/liveness checks and account-integrity signals are described as separate claims.
- **ID-04:** Account recovery resists takeover, SIM/email compromise and repeated automated attempts.
- **ID-05:** Users can pause discovery, unmatch, block, delete their account and request applicable data rights.
- **Acceptance:** account lifecycle flows pass automated and manual tests; no user can access another account's data through ID substitution or stale sessions.

### Profile and preferences
- **PROF-01:** Users control profile facts, photos, prompts, relationship goals and visibility.
- **PROF-02:** Every preference is classified as hard or soft, with clear explanation and reversible control.
- **PROF-03:** Hard constraints are applied before ranking.
- **PROF-04:** Soft preference relaxation requires explicit user action and identifies the change.
- **PROF-05:** Users can correct and remove profile data.
- **Acceptance:** hard-preference compliance is 100% in the standard discovery path; negative tests cover boundary values, missing data, cache results and API access.

### Discovery and matching
- **MATCH-01:** Discovery excludes accounts blocked by either party under the product's documented rules, hidden/deleted/suspended accounts and ineligible profiles.
- **MATCH-02:** Candidate generation considers reciprocal relevance, not one-sided swipe likelihood alone.
- **MATCH-03:** Recommendations are not guarantees of attraction, consent or compatibility.
- **MATCH-04:** “Why this profile?” uses visible, user-provided facts and the viewer's own settings.
- **MATCH-05:** Safety risk is separate from romantic ranking.
- **MATCH-06:** Empty/exhausted discovery states explain the limitation and offer wait, explicit soft-preference expansion, distance change or pause.
- **MATCH-07:** The system measures exposure concentration, eligible-user opportunity, cold-start quality and reciprocal outcomes.
- **MATCH-08:** Ranking must not use profile completeness as a desirability signal; only minimum safety, authenticity and functional participation requirements may be required.
- **MATCH-09:** Interests, hobbies and prompt similarity are soft signals only and cannot override hard constraints or be used as proxies to infer sensitive traits.
- **MATCH-10:** Any controlled exploration remains inside hard eligibility constraints and is monitored for relevance, fairness and safety.
- **MATCH-11:** No public desirability score, hidden paid compatibility advantage, fake profile or fabricated activity.
- **MATCH-12:** Relationship-goal compatibility uses a goal + openness model: users state a primary intention and may optionally select other goals they are genuinely open to.
- **MATCH-13:** Evaluate both users' stated goals and hard/soft settings before ranking; a hard-incompatible pair is excluded from standard discovery.
- **MATCH-14:** Missing, undisclosed and “not sure yet” relationship-goal values are distinct; none is automatically treated as open-minded or compatible.
- **MATCH-15:** Low candidate supply never silently relaxes relationship-goal constraints, and premium features cannot bypass them.
- **Acceptance:** automated tests confirm hard constraints, block exclusions, relationship-goal matrix behavior, cache invalidation, explanation grounding and deterministic behaviour for equivalent inputs; offline evaluation uses documented datasets and bias checks. Monitor reciprocal outcomes, eligible exposure concentration, cold-start quality, and safety separately from romantic ranking.

### Likes, matches and messaging
- **LIKE-01:** A like records interest only; it is not a match and does not grant messaging access.
- **LIKE-02:** A match is created only from valid reciprocal interest while both accounts remain eligible and no block, suspension, deletion or other mandatory exclusion applies.
- **LIKE-03:** Before a match, users can withdraw an outgoing like; withdrawal removes its pending effect from future discovery and notification surfaces where technically feasible, and must not generate a misleading “match” state.
- **LIKE-04:** A user may attach an optional, user-written comment to a like on a specific profile photo or prompt. The comment is delivered only in the recipient's incoming-like experience; it does not open a conversation or permit further contact unless the recipient reciprocates with a like and a valid match is created.
- **LIKE-05:** Free users see five eligible incoming likes unblurred at a time in a rolling queue. When they like back or pass on one, the next eligible like moves into the visible set; other eligible likes remain blurred while queued. Premium subscribers can see the full eligible incoming-like queue unblurred and review it incrementally. Five is an initial hypothesis to validate, not a proven optimum. Eligibility, hard preferences, blocks, suspensions, unmatched-pair exclusions and safety restrictions apply before display and cannot be bypassed by payment. Never fabricate likes, activity, urgency or scarcity; do not manipulate ordering to pressure upgrades.
- **LIKE-06:** Thoughtful, relevant comments on a specific photo or prompt may receive a modest, transparent ranking advantage in the recipient's incoming-like ordering, provided the sender is otherwise eligible. This is not a guaranteed boost or guaranteed visibility. Ranking must not override recipient preferences or safety controls, or favour comment length, AI use or payment. Monitor for spam, harassment, low-quality copy-paste comments, inequitable exposure and unwanted pressure; offer reporting, rate limits and recipient controls.
- **LIKE-07:** A match and its conversation are created idempotently: duplicate requests, retries or concurrent reciprocal likes must not create duplicate matches or conversations.
- **LIKE-08:** If eligibility changes before reciprocal interest resolves, re-check current hard constraints and mandatory exclusions before creating a match.
- **LIKE-09:** Users can unmatch without giving a reason. Unmatching takes effect immediately, closes messaging and removes the connection from both users' active match lists, and does not disclose which user initiated it. The former match remains visible in a separate read-only past-matches/history list for 30 days after unmatching, then disappears from that user-facing list. Neither user can send or receive messages through the former connection; the history entry provides report and block actions while visible and never reopens contact. Unmatching does not automatically block either user. Retention of underlying records follows documented, proportionate safety, legal and operational requirements and is distinct from the 30-day user-facing visibility period.
- **LIKE-10:** An unmatched pair is permanently excluded from matching each other again, including through future discovery, likes, retries or account-state refreshes. This pair exclusion is separate from blocking and does not automatically block either user.
- **LIKE-11:** Blocking takes precedence over likes and matches, immediately prevents further contact, and propagates across discovery, likes, messaging, notifications, caches and queued jobs.
- **LIKE-12:** Repeated likes, duplicate actions and stale client state must not bypass rate limits, blocks, unmatched-pair exclusions, preference rules or server-side authorisation.
- **Acceptance:** automated and integration tests cover comment attachment and delivery, recipient preference/block enforcement, comment ranking guardrails, simultaneous reciprocal likes, retries, withdrawal races, preference changes, account suspension/deletion, block-vs-match races, unmatch-vs-message races, permanent unmatched-pair exclusion, 30-day history-list expiry, duplicate notifications and premium bypass attempts.

- **MSG-01:** Mutual-match messaging is available without premium subscription.
- **MSG-02:** Users can send ordinary messages without AI or generated prompts.
- **MSG-03:** Optional openers are profile-specific, editable and user-sent.
- **MSG-04:** AI coaching is opt-in and never sends a message automatically.
- **MSG-05:** Users can mute a conversation, control notifications, unmatch, block and report.
- **MSG-06:** No streaks, reply countdowns or guilt-based response pressure.
- **MSG-07:** Message access, staff review and retention are governed by a documented privacy/safety policy.
- **MSG-08:** Optional conversation starters reference a specific, visible profile detail; suggestions are editable and never sent without explicit user action.
- **MSG-09:** Any AI writing assistance is opt-in, user-reviewed and user-sent; it must not invent facts, impersonate the user or send automatically.
- **MSG-10:** Users have global notification controls and per-conversation mute controls; no streaks, reply countdowns, guilt nudges or punitive reply-speed scoring.
- **MSG-11:** Users can unmatch, block and report from conversation surfaces; block effects propagate across discovery, likes, messaging, notifications, caches and queues.
- **MSG-12:** Read receipts and online status are off by default unless a documented user benefit, privacy review and user control justify them.
- **MSG-13:** Optional post-date follow-up or closure suggestions are editable and never sent automatically.
- **MSG-14:** Any message scanning or contextual safety warning has a defined purpose, proportional intervention, transparency, retention limits and appeal route; automated signals do not establish intent or determine romantic desirability.
- **Acceptance:** only mutual matches can message under MVP rules; block/unmatch immediately prevents future contact as designed across APIs, queues, notifications and cached surfaces. Test suggestion grounding, AI opt-in, notification controls, block propagation, false-positive handling and message-retention rules.

### Trust, safety and moderation
- **SAFE-01:** Reporting is available from profile, message/media, match and relevant account surfaces.
- **SAFE-02:** Blocking works immediately and is free.
- **SAFE-03:** Reports create traceable cases with restricted evidence and an immutable action history.
- **SAFE-04:** Severity levels, triage targets, escalation, emergency and child-safety procedures are documented and staffed.
- **SAFE-05:** Significant restrictions have rationale and an appropriate appeal route.
- **SAFE-06:** IP/device/image signals are corroborating evidence, not sole proof for automatic permanent bans.
- **SAFE-07:** Moderator access is least-privilege and audited.
- **SAFE-08:** Safety policies and response promises match operational capacity.
- **Acceptance:** tabletop exercises cover credible threat, fraud ring, intimate-image abuse, child-safety concern, mass false reporting and vendor outage; block/report end-to-end tests pass before launch.

### Privacy, data rights and security
- **PRIV-01:** Each data field has a purpose, lawful basis, visibility, retention and deletion rule.
- **PRIV-02:** Sensitive data is minimised; no sensitive-trait inference from photos, names, location or message contents.
- **PRIV-03:** A DPIA is completed before high-risk profiling, biometrics or novel AI, and legal scope is documented.
- **PRIV-04:** Data rights and deletion cover relevant primary systems, processors and backup expiry.
- **PRIV-05:** Logs and analytics exclude message text, exact location, identity documents, secrets and unnecessary direct identifiers.
- **PRIV-06:** Encryption, MFA, least privilege, managed secrets, audit trails, backups and incident response are implemented.
- **PRIV-07:** Vendors undergo privacy, security, accessibility, retention and exit review.
- **Acceptance:** privacy/security sign-off includes data-flow map, processing register, DPIA decisions, access review, deletion/export test, restore test and critical/high vulnerability remediation.

### Monetisation and billing
- **BILL-01:** Core matching, mutual messaging, blocking, reporting, unmatching, essential privacy controls and legally required age assurance remain free. Incoming-like visibility is governed by the product's transparent free-versus-paid display policy below; users must never be misled about whether likes exist.
- **BILL-02:** Premium benefits are optional and cannot override another user's hard preferences, consent, blocks or unmatched-pair exclusions.
- **BILL-03:** Subscription architecture is an open decision: compare Free + one paid tier against Free + two paid tiers. Prefer one paid tier at launch unless a second tier has a clearly distinct benefit bundle and evidence that added choice improves user value without confusion.
- **BILL-04:** Incoming-like model hypothesis: free users see the first five eligible incoming likes unblurred and can like back or pass on them; further eligible incoming likes may be blurred until they move into the visible set or the user subscribes. Subscribers can see the full eligible incoming-like queue unblurred and review it incrementally, liking back to create a match or passing/dismissing to move on. Five is a testable starting hypothesis, not a validated optimum.
- **BILL-05:** Incoming-like ordering must not be manipulated to create artificial scarcity. Show accurate counts only when the count is reliable, explain the blur/visibility rule clearly, and never fabricate likes, activity or urgency. Hard preferences, blocks, suspensions, unmatched-pair exclusions and safety controls apply before display and cannot be bypassed by payment.
- **BILL-06:** Candidate premium benefits for validation are full unblurred incoming-like queue with incremental review, profile customisation tools, optional user-reviewed profile/message coaching, saved or organised discovery lists, additional discovery organisation controls, travel mode that truthfully reflects location, and more detailed compatibility explanations based only on user-provided data.
- **BILL-07:** Do not sell hidden romantic ranking advantages, paid comment boosts, guaranteed likes/matches/dates, unsolicited messaging privileges, safety controls or access to another user's consent. Commented-like ranking is independent of payment.
- **BILL-08:** Any advanced soft-preference controls require explicit user action, clearly state what changes, and cannot silently relax hard constraints.
- **BILL-09:** Price, cadence, renewal, trial conversion, cancellation and refund terms are clear before purchase; cancellation is straightforward and not obstructed by guilt or repeated retention screens.
- **BILL-10:** Entitlements are server-authoritative and reconciled with the selected payment/store system.
- **BILL-11:** Cancellation and account deletion are clearly distinguished, with billing consequences explained.
- **BILL-12:** Refund, payment failure, restore purchase and renewal states are tested; essential safety functionality is unaffected by expiry or failed payment.
- **BILL-13:** Measure contribution margin and retention alongside refunds, complaints, trust, free-user outcomes and safety; do not optimise solely for conversion or revenue.
- **Acceptance:** end-to-end purchase lifecycle tests pass on each launch platform; safety functionality is unaffected by expiry or failed payment; every premium benefit has a documented user problem, free/premium boundary, guardrail and validation plan.

### Launch and operations
- **OPS-01:** Launch geography is chosen using a documented local eligible-supply and acquisition model.
- **OPS-02:** Staged beta and expansion gates cover density, safety, support capacity, preference compliance and unit economics.
- **OPS-03:** No fake profiles, fabricated likes or misleading activity.
- **OPS-04:** Safety queues, support ownership, service targets, incident playbooks and escalation coverage are proven before launch.
- **OPS-05:** Dashboards track quality, safety, fairness, fatigue, trust, commercial outcomes and operations.
- **Acceptance:** a go/no-go review records evidence, open risks, owners and explicit approval; launch cannot proceed with unresolved critical security or safety issues.

## 6. Non-functional requirements
Targets below are draft placeholders to be finalised against platform, budget, threat model and load tests.

### Security and privacy
- Server-side authorisation for every profile, media, conversation, report, admin action and paid entitlement.
- TLS for network traffic; appropriate encryption at rest and key management.
- MFA for privileged staff and least-privilege access.
- Audited administrative and evidence access.
- Secret scanning, dependency scanning, code review and vulnerability handling in CI/CD.
- No production personal data in local development fixtures.
- Documented retention and deletion for logs, media, messages, reports, backups and vendors.
- Threat model for scraping, stalking, location triangulation, credential stuffing, account takeover, insider access and ban evasion.

### Reliability and recovery
- Define availability and recovery objectives before public beta.
- Use automated backups and regularly test restoration.
- Queue/report failures must be observable and alert to an owner.
- Critical safety events and block operations should fail safely; degraded infrastructure must not silently permit blocked contact.
- Maintain incident runbooks, rollback procedures and change-control gates.

### Performance and scalability
- Establish baseline latency budgets for onboarding, discovery, likes, match creation, messaging and blocking.
- Load-test expected beta peak plus a documented headroom factor.
- Use pagination and rate limits; avoid returning entire candidate pools or sensitive profile data.
- Use asynchronous work for non-blocking media processing and notifications, with retry/idempotency rules.
- Avoid premature microservices; start with clear modular boundaries and split only when justified.

### Accessibility and usability
- Target WCAG 2.2 AA for web surfaces where applicable and use platform accessibility guidance for mobile.
- Support screen readers, dynamic text, sufficient contrast, reduced motion and accessible error states.
- Ensure safety/reporting and cancellation flows are accessible and not hidden behind gestures alone.
- Test with automated tooling and manual assistive-technology review.

### Observability
- Structured logs with sensitive-data redaction.
- Metrics, traces, dashboards and alert ownership.
- Correlation IDs without exposing personal content.
- Product analytics events documented in a versioned schema.
- Distinguish product events from security/audit records and apply separate access/retention rules.

## 7. Provisional domain architecture
Proposed modular domains:
1. Identity and account lifecycle.
2. Profiles and media.
3. Preferences and eligibility.
4. Discovery and reciprocal matching.
5. Likes and matches.
6. Messaging and notifications.
7. Trust, safety and moderation.
8. Verification and age assurance.
9. Billing and entitlements.
10. Analytics and experimentation.
11. Admin and support.
12. Privacy rights and data lifecycle.

Keep domain ownership and API contracts explicit. Enforce critical rules on the server, not only in mobile clients. The matching system should consume eligible candidates and return explainable results; it must not have authority to override hard preferences or safety blocks.

## 8. Provisional data entities
- UserAccount, AuthenticationSession, RecoveryEvent.
- Profile, ProfileMedia, ProfilePrompt, RelationshipGoal.
- Preference, PreferenceType (hard/soft), PreferenceChange.
- AgeAssuranceResult, VerificationClaim, VerificationProviderEvent.
- CandidateImpression, RecommendationReason, Like, Match.
- Conversation, Message, ConversationState, NotificationPreference.
- Block, Report, ModerationCase, EvidenceReference, ModerationAction, Appeal.
- Subscription, PurchaseReceipt, Entitlement, RefundEvent.
- ConsentRecord, PrivacyRequest, RetentionRule, DeletionJob.
- AuditEvent, SecurityEvent, ExperimentAssignment, AnalyticsEvent.

These are conceptual entities; normalisation, exact fields, identifiers, retention and access policies require data modelling and DPIA review.

## 9. API and event principles
- Use versioned API contracts and schema validation.
- Idempotency for match creation, payment webhooks, moderation actions and deletion jobs.
- Server-side validation for preference changes, blocks, entitlements and access permissions.
- Avoid exposing exact coordinates, internal risk scores, private preferences or moderation evidence through public APIs.
- Webhooks must be authenticated, replay-protected and safely retried.
- Event taxonomy must define owner, purpose, schema, retention and whether an event can contain personal data.
- Make privacy deletion and block propagation reliable across caches, search indexes, notifications, analytics and third parties.

## 10. Quality and testing strategy
- **Unit tests:** eligibility, hard/soft preferences, ranking constraints, entitlements, retention and moderation policy logic.
- **Property-based tests:** no candidate violates hard constraints; blocks suppress contact; premium status cannot bypass safety.
- **Contract tests:** API schemas, billing providers, age-assurance provider, notification provider and moderation vendor.
- **Integration tests:** registration, age assurance, profile, discovery, like/match, messaging, report/block, deletion and restore.
- **End-to-end tests:** user lifecycle and critical safety flows across iOS, Android and admin.
- **Security tests:** IDOR/BOLA, session expiry, account recovery, rate limits, upload validation, webhook replay, privilege escalation and scraping.
- **Privacy tests:** analytics/log payload inspection, deletion propagation, access-control review and retention expiry.
- **Accessibility tests:** screen-reader and dynamic-text checks, contrast, keyboard support for web admin, accessible report/cancel flows.
- **Load tests:** discovery, messaging, notification queues, report bursts and launch cohort peaks.
- **Operational exercises:** safety incident tabletop, backup restoration, provider outage and rollback.

## 11. Core analytics and decision metrics
### Product value
- Onboarding completion and time to first eligible recommendation.
- Eligible candidate pool distribution and discovery exhaustion.
- Reciprocal matches per eligible cohort and time to first mutual match.
- Conversation initiation and two-way conversation measures.
- Optional self-reported “met” / “would meet again” outcomes.

### Preference and fairness
- Hard-preference violation rate (target zero).
- Soft preference expansion and reversal rates.
- Eligible exposure distribution and concentration.
- Cold-start performance and user-reported relevance.
- Recommendation explanation usefulness and correction rates.

### Safety and trust
- Reports per active user/conversation by category/severity.
- Human triage/action time, critical backlog, repeat harm and recurrence.
- Block effectiveness, appeal reversal and QA false-positive/negative indicators.
- Account integrity and scam outcomes.
- User trust/safety feedback and moderation capacity.

### Wellbeing and commercial sustainability
- Pause, notification control, healthy exit and self-reported fatigue where volunteered.
- Free-user outcomes and core feature access.
- Conversion, renewal, churn, refunds, chargebacks and contribution margin.
- Acquisition cost per activated eligible user.
- Support cost, moderation cost and operational capacity.

Never use session length, swipes, messages sent or paid conversion as the sole success metric.

## 12. Dependencies and critical path
1. Legal/service-scope assessment: age threshold, Online Safety Act, privacy and subscriptions.
2. Product decisions: hard vs soft preference taxonomy, relationship-goal compatibility, blocks, message rules and verification claims.
3. Launch market and cohort: local density model, acquisition budget and support coverage.
4. Privacy architecture: data inventory, lawful basis, DPIA, retention, deletion and vendor review.
5. Trust-and-safety operations: policies, staff coverage, escalation, appeals and evidence handling.
6. Technical stack decision: client, backend, database, storage, queues, hosting, observability and payment architecture.
7. Data and API design: domain models, contracts, event taxonomy and migrations.
8. Build/test: end-to-end safety, preference, deletion and billing tests.
9. Beta gates: density, reliability, safety operations, security and economics.
10. Mobile design only after product requirements, architecture decisions and primary user journeys are approved.

## 13. Open decisions register
| Decision | Current status | Decision owner / evidence needed |
|---|---|---|
| Initial launch geography | Open | Compare eligible supply, acquisition, cost and safety capacity |
| Age assurance provider/method | Open | Legal scope, accuracy, fairness, privacy, accessibility and vendor review |
| Identity/photo verification | Open | Define precise claim, user value, DPIA and false-positive process |
| Hard vs soft preference taxonomy | Partly defined | Product policy and boundary-case tests |
| Relationship-goal matching rules | Open | Compatibility matrix and user-control policy |
| Message/media retention | Open | Privacy, safety, legal and operational requirements |
| Moderation coverage model | Open | Launch volume, severity targets, budget and on-call capacity |
| Technical stack and hosting | Proposed, not final | ADR-0001 evaluation and team capability |
| Billing route and premium plan | Monthly-first candidate; one optional Premium tier; no free trial at launch; price and final bundle open | Store rules, UK law, net economics and trust tests |
| AI features and provider | Optional / open | Privacy, safety, data use, review and deletion controls |
| Success thresholds and launch gates | Draft | Pre-register targets against beta assumptions |
| Public verification badge wording | Open | Exact proof scope and comprehension testing |

## 14. Definition of research/requirements completion
The research phase can be considered ready for a design/build decision when:
- Core product principles and MVP scope are approved.
- Functional and non-functional requirements have owners and acceptance criteria.
- Hard/soft preferences and relationship-goal rules are documented.
- Legal scope, DPIA plan, age assurance and data governance have named owners.
- Trust-and-safety policies, staffing, incident response and appeals are credible for the intended launch.
- Launch market assumptions and density gates are modelled.
- Free/premium boundary and billing approach are approved in principle.
- Stack ADR and provisional domain/data model are reviewed.
- Analytics definitions, guardrails and experiment plan are agreed.
- Critical unknowns are either resolved or explicitly accepted as hypotheses with a validation plan.
- No unresolved critical safety or privacy risks are hidden behind “we will fix it later.”

## 15. Decision summary
**Recommendation:** approve this as the consolidated draft requirements baseline, then resolve the open-decision register and technical ADR before screen design or production implementation. Use controlled beta measurement to validate assumptions, while keeping hard preference integrity, safety, privacy and user agency as non-negotiable constraints.

## Source documents in this repository
- `docs/research/2026-10-21-messaging-and-conversation-design.md`
- `docs/research/2026-10-21-matching-algorithms-fairness-explainability.md`
- `docs/research/2026-10-21-privacy-uk-gdpr-age-assurance-data-governance.md`
- `docs/research/2026-10-21-trust-safety-operations-incident-response.md`
- `docs/research/2026-10-21-launch-strategy-local-density-cold-start-economics.md`
- `docs/research/2026-10-21-monetisation-subscription-economics-ethical-pricing.md`
- `docs/architecture/decisions/ADR-0001-stack-and-deployment.md`
- `docs/security/privacy-foundations.md`
- `docs/trust-and-safety/operating-model.md`


## 16. Verification and age-assurance decision record (provisional; research complete, decisions open)

### Proposed requirements
- **AGE-01 — Mandatory age assurance:** determine the legal minimum age and applicable UK age-assurance obligations for Pinchd's service before launch. Use a method evidenced to meet the applicable Ofcom standard where highly effective assurance is required; self-declaration alone is insufficient for that purpose.
- **AGE-02 — Separate controls:** age assurance, legal identity verification, selfie/profile verification and safety review are distinct. A result in one category must not be described as proof of another.
- **VER-01 — Public badge scope:** any public verification badge must describe the exact check performed and must not imply that the user is safe, truthful or legally identified unless that specific claim is actually established.
- **VER-02 — Voluntary badge hypothesis:** public selfie verification should remain optional at launch unless product research, risk assessment and legal review establish a proportionate reason to require it. This does not change any mandatory age-assurance duty.
- **VER-03 — Risk-triggered review:** enhanced verification may be requested for documented, proportionate reasons such as credible impersonation reports or suspected account takeover. A trigger is not proof of wrongdoing.
- **VER-04 — Fair failure handling:** failed, low-confidence or inaccessible checks must have a retry and alternative/manual review path where feasible. Failure alone is not proof of malicious conduct. Required age checks must be completed before access to features for which they are a legal prerequisite.
- **VER-05 — Appeals:** meaningful account restrictions based on verification or automated risk signals must have a documented review and appeal process, except for temporary protective actions needed to address urgent risk.
- **PRIV-VER-01 — Data minimisation:** store the minimum result and audit evidence needed. Prefer prompt deletion of raw selfie/video/document captures and avoid storing face templates where a vendor can return a scoped result. Retention exceptions must be necessary, lawful, defined and documented.
- **PRIV-VER-02 — No secondary use:** verification data must not be used for attractiveness scoring, recommendations, advertising or unrelated model training.
- **PRIV-VER-03 — Vendor gate:** no production vendor is selected until method-specific effectiveness evidence, subgroup error rates, spoof resistance, privacy architecture, data location/transfers, deletion controls, accessibility, cost, support and exit options are compared and reviewed.
- **PRIV-VER-04 — DPIA:** complete a DPIA before using biometric recognition where required and complete privacy/legal review of the full data flow and supplier contract before production.
- **SAFETY-VER-01 — Impersonation reporting:** users can report suspected impersonation, stolen images, scams and suspicious account behaviour from profile or conversation surfaces without contacting the reported person.

### Launch gates
1. Confirm applicable legal age and age-assurance duties with UK counsel.
2. Document the purpose and precise claim of each verification method.
3. Compare at least two suppliers or explain why a viable alternative is unavailable.
4. Complete method-specific due diligence, accessibility review, privacy review and DPIA where required.
5. Test false acceptance/rejection, retry, abandonment, appeal outcomes and performance across relevant cohorts.
6. Approve retention/deletion, incident response, staff access, user communications and supplier exit procedures.
7. Keep age assurance and any optional public verification badge outside Premium paywalls.

These are proposed requirements based on research, not a declaration that legal review or supplier selection is complete. No provider is selected and no biometric method is approved for production.


## 17. Startup-friendly age-assurance supplier requirements

- **VEND-AGE-01:** Prefer a pay-as-you-go supplier with a free sandbox or meaningful startup allowance, no annual commitment, and no material minimum monthly spend during beta.
- **VEND-AGE-02:** Compare low-cost candidates (including Didit, WeVerify and Stripe Identity, with Certivus as an additional quote candidate) on the exact method, Ofcom evidence, false acceptance/rejection, privacy, security, accessibility, all-in pricing and contract terms. These are candidates only; no provider is approved.
- **VEND-AGE-03:** No vendor may be used in production for a legally required age gate until the specific method's effectiveness and applicable regulatory fit are documented and approved.
- **VEND-AGE-04:** Prefer minimal result sharing and prompt deletion of raw capture data; confirm processing roles, subprocessors, hosting, transfers, retention, incident notification and deletion in writing.
- **VEND-AGE-05:** The pilot must estimate total cost per completed eligible account, including retries, failed checks, fallback checks, manual review, FX/VAT and support—not only advertised successful-check price.
- **VEND-AGE-06:** Keep mandatory age assurance free to users and separate from Premium.


## 18. Voluntary profile verification workflow

- VER-01: Public selfie/profile verification is voluntary by default and user-initiated. Non-participation must not reduce baseline eligibility, compatibility ranking, or ordinary discovery exposure.
- VER-02: Explain the exact purpose, checks, provider, biometric/data processing, retention/deletion, badge meaning, retries, alternatives and support before consent/activation.
- VER-03: Baseline flow is one liveness check followed by one 1:1 face match against the user's current profile photo(s); no duplicate match step and no 1:N face search for ordinary badge verification.
- VER-04: A successful check may issue a badge whose wording describes only the check performed. The badge must not imply that a person is safe, truthful, criminal-record checked, or verified against government identity unless those checks actually occurred.
- VER-05: Result states must distinguish passed, failed, inconclusive/needs review, technical error and retry. An inconclusive result is not a finding of fraud.
- VER-06: Provide reasonable retries, accessible fallback/review and a clear support path. Test false acceptance/rejection and demographic performance; review any thresholds before production.
- VER-07: Reassess badge validity when the relevant profile photos change; notify the user and provide a path to re-verify if the badge is suspended or removed.
- VER-08: Minimise biometric data and retention. Before launch, complete privacy/legal review, lawful-basis analysis, Article 9 condition analysis where applicable, DPIA where required, supplier due diligence, retention/deletion and international-transfer review.
- VER-09: Device/IP analysis is not a default prerequisite for voluntary profile verification. Use only where a documented and proportionate security purpose supports it.
- VER-10: AML/sanctions screening is excluded from ordinary dating-profile verification unless a separate, evidenced legal/business requirement is established.
- VER-11: Age assurance is a separate workflow and legal question. Optional profile verification must not substitute for legally required age assurance; age assurance must not be made mandatory without documented legal/product justification.
- VER-12: State-specific branching must reflect validated legal requirements and current jurisdictional status. Do not select an age-assurance method solely because a workflow template lists a state law.
