# Pinchd — Lifecycle Defaults and Privacy Decision Log
**Date:** 22 October 2026  
**Status:** Consolidated decision proposals and open gates. This is not a record of new user approval.  
**Purpose:** Make unresolved decisions explicit before schema/API freeze, identify which choices can be product-confirmed and which require legal, vendor, technical or empirical evidence.

## 1. Decision rule
The explicit user-approved product baseline remains authoritative. This log does not silently promote recommendations into approved policy. A decision is closed only when the named owner records the choice, rationale, date, dependencies, user-facing wording where applicable and acceptance tests. Legal/privacy questions require qualified review and cannot be closed by product preference alone.

## 2. Product defaults needing confirmation

| ID | Topic | Recommended default | Why | Consequence if accepted | Confirmation needed |
|---|---|---|---|---|---|
| LD-01 | Pause and new discovery | Paused profile is hidden from new discovery and cannot receive/create new discovery likes | Clear boundary between visibility and relationship history | Candidate queries, incoming queue and cached feeds must filter paused accounts | Product owner |
| LD-02 | Existing matches while paused | Pause alone does not unmatch or close messaging | Avoids destructive surprise and treats pause as visibility control | Existing matches remain accessible unless separate block/unmatch/restriction applies | Product owner + T&S |
| LD-03 | Outgoing likes during pause | Freeze new presentation; do not silently withdraw; re-evaluate on resume and before match | Preserves user action while preventing stale eligibility from becoming a match | Likes remain stored but cannot bypass current eligibility | Product owner |
| LD-04 | Pass reversibility | Pass is non-reversible in launch version | Simplifies candidate state and avoids hidden re-entry behavior | No undo or second-chance queue at launch; pass remains distinct from block | Product owner; validate in user research |
| LD-05 | Outgoing-like expiry | No silent time-based expiry at launch | Avoids hidden timers and unexpected loss of interest | Likes remain until withdrawn or invalidated by current eligibility/safety state | Product owner; monitor stale queue behavior |
| LD-06 | Removed prompt/photo referenced by like comment | Show current profile or neutral unavailable-content label; never fabricate removed content | Maintains truthful context while limiting unnecessary evidence retention | Need stable content reference and privacy-reviewed evidence policy | Product owner + privacy/T&S |
| LD-07 | Premium entitlement refresh | Server-authoritative; reconcile verified store state; apply free queue on expiry/refund after authoritative refresh | Prevents client-forged entitlement and preserves entitlement integrity | Billing sandbox, restore, grace, refunds and delayed notifications must be tested | Product + billing owner; counsel/store review |
| LD-08 | Restriction vs existing match | Capability-scoped safety restriction may disable messaging; full suspension disables normal use | Supports proportionate response while preserving safety reporting/appeals where appropriate | Requires explicit enforcement capability matrix | Product + T&S/legal |
| LD-09 | Account deletion behavior | Stop new activity and hide from discovery immediately; durable deletion workflow with truthful progress and documented exceptions | Prevents further interaction while deletion tasks propagate | Requires clear re-auth/grace rules, deletion status and vendor/backup orchestration | Product + privacy/legal |
| LD-10 | Pass and pair suppression | Pass does not create permanent pair exclusion; unmatch does | Keeps ordinary discovery choice separate from safety and relationship state | Separate pass history, block records and unmatch suppression | Product owner |

## 3. Decisions already approved — do not reopen without new evidence
- Hard eligibility constraints precede ranking; never silently relax them.
- Legal minimum age is a separate gate from user-configured discovery age range.
- Age-assurance pending, unknown, failed or expired never grants gated access.
- Like is interest, not match or messaging permission; a match requires valid reciprocal interest and current eligibility.
- Block is server-authoritative and propagates across discovery, likes, messages, notifications, caches and queues.
- Unmatch closes active messaging immediately, removes active lists, does not reveal initiator, preserves read-only past-match history for 30 days, allows report/block while visible and permits no messaging.
- A pair is permanently excluded from matching after unmatch, separately from block.
- Free users see five eligible incoming likes unblurred at a time in a rolling queue; Premium users see the full eligible queue. Five remains a launch hypothesis to validate, not a proven optimum.
- Essential matching, mutual messaging, blocking, reporting, legally required age assurance and deletion remain free.
- No fake profiles/activity, fabricated likes, deceptive urgency/scarcity, pay-to-win compatibility or guaranteed outcomes.
- Optional profile verification is separate from age assurance.
- Mobile UI design remains paused until research and requirements are concluded.

## 4. Legal/privacy decisions that cannot be closed by product approval

| ID | Open gate | Required decision/evidence | Accountable role | Blocks |
|---|---|---|---|---|
| PV-01 | Lawful basis for each processing purpose | Article 6 mapping and any special-category condition/safeguards for sensitive data; purpose and minimisation | Privacy/legal | Public launch; affected data use |
| PV-02 | DPIA and high-risk processing | Documented screening and DPIA where required; risks, mitigations and residual-risk approval | Privacy/legal + security | High-risk profiling/biometrics/novel AI |
| PV-03 | Age-assurance vendor and data flow | Current contract/DPA, workflow, fields, regions, subprocessors, accuracy/error, appeal, retention/deletion and callback proof | Privacy/legal + vendor owner | Age-assurance implementation/launch |
| PV-04 | Message/attachment retention | Separate retention trigger/period for visible history, message bodies, attachments, reports, audit logs and backups | Privacy/legal + T&S | Schema freeze and public launch |
| PV-05 | Pair suppression after deletion | Whether minimal pseudonymous suppression may persist, lawful basis, duration, re-identification risk, disclosure and rights handling | Privacy/legal + product | Post-deletion suppression implementation |
| PV-06 | Safety evidence and lawful holds | Case-specific access, purpose, retention, legal hold, external disclosure and deletion exception policy | Privacy/legal + T&S | Public launch |
| PV-07 | Account/contact retention and deletion | Grace/security policy, deletion timeline, vendor propagation, backup expiry, exceptions and completion wording | Privacy/legal + engineering | Public launch |
| PV-08 | International transfers/subprocessors | Processor register, data-flow regions, support access, contracts and transfer assessment | Privacy/legal + vendor owner | Vendor production use |
| PV-09 | Subscription law and store rules | Launch-date applicability, cancellation/refund/reminders/disclosures and current platform rules | Counsel + billing owner | Paid launch |
| PV-10 | OSA and child-access scope | Written scope determination, applicable risk assessments, controls and evidence using current rules | Counsel + T&S lead | Public launch |

## 5. Technical and operational decisions before production
- Define formal RPO/RTO for authentication, matching, messaging, safety/reporting, billing and deletion workflows; prove backup restoration and deletion-tombstone replay.
- Select exact AWS services, regions, database version/extensions, queues, logging/metrics and secrets tooling after PoCs and cost/security review.
- Confirm idempotency behavior for likes, matches, messages, deletion and provider callbacks.
- Define moderation/support staffing and coverage; internal severity targets must be resourced and tested, not presented as statutory deadlines.
- Define accessibility device/assistive-tech test matrix and close severe defects in critical safety flows.
- Confirm alert ownership and escalation for block propagation, age-gate anomalies, deletion backlog, outbox backlog, billing divergence and unusual spend.
- Record exact implementation versions and run reproducible tests before marking any PoC passed.

## 6. Closure record template
For each decision, record:
- Decision ID and exact question.
- Status before/after: proposed, approved, rejected, deferred or blocked.
- Decision owner and reviewer.
- Decision date, rationale and alternatives considered.
- Evidence reviewed and links/file references.
- Legal/privacy/vendor dependencies and any residual risk.
- User-facing notice/help wording affected.
- Data model/API/event changes affected.
- Acceptance tests and result.
- Review trigger and next review date.

## 7. Immediate sequence
1. Product owner confirms or rejects LD-01 through LD-10; where no decision is made, preserve them as open and do not freeze dependent contracts.
2. Privacy/legal owner closes PV-01 through PV-10 only with appropriate advice and evidence; no invented retention periods or lawful bases.
3. Vendor owner completes the Didit evidence request pack and obtains current documents.
4. Engineering executes non-vendor PoCs first, then runs POC-05 and POC-09 only when the relevant evidence/policy is sufficient.
5. Update the PostgreSQL blueprint and API contracts from the signed decision record; do not implement schema fields that imply unapproved indefinite retention.
6. Update the requirements traceability matrix with evidence links and actual statuses.
7. Conduct a design-readiness review only after research, risk ownership and acceptance gates are complete.

**Current status:** LD-01 through LD-10 remain proposed recommendations pending product confirmation. PV-01 through PV-10 remain open gates. No legal sign-off, Didit integration pass or technical PoC pass is claimed by this document.
