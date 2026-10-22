# Pinchd Workstream 5 — Open Architecture Decisions and Recommended Defaults
**Date:** 22 October 2026  
**Status:** Proposed defaults for review; not approved product policy, final schema, legal advice or implementation sign-off. Mobile UI design remains paused.

## Purpose
The lifecycle requirements expose decisions that must be resolved before the physical database schema and precise API contracts are frozen. This register proposes conservative defaults to keep engineering moving without silently converting open questions into approved policy. Each item names the recommendation, why it is recommended, and what must be validated.

## 1. Pause behavior
**Recommended default:** Pause stops profile exposure in new discovery and prevents new discovery likes. It does not automatically unmatch or delete existing conversations. Existing matched conversations remain accessible unless the user separately blocks/unmatches or an enforcement restriction applies. Existing outgoing likes should be frozen from new presentation while paused and re-evaluated on resume; do not silently withdraw them.
**Why:** Separates profile visibility from relationship state and avoids unexpected destructive behavior.
**Validate:** User comprehension, queue consistency and abuse/safety implications. Publish exact treatment of previously sent likes before launch.

## 2. Pass reversibility
**Recommended default:** A pass is not reversible in the initial launch version. Provide a deliberate confirmation only if research shows accidental passes are common; do not add an undo feature that secretly re-inserts people into discovery or bypasses pair suppression.
**Why:** Simple, predictable eligibility and queue semantics; avoids a complex second-chance system at launch.
**Validate:** Test with users; ensure pass and block remain distinct and that a pass is not treated as a safety block.

## 3. Like expiry
**Recommended default:** Do not silently expire a valid outgoing like during the initial version. Re-evaluate eligibility whenever it is displayed or acted upon. If a later expiry policy is justified, publish it, define notification behavior, and make expiry idempotent.
**Why:** Avoids introducing hidden timers and user confusion before there is evidence that expiry is needed.
**Validate:** Queue volume, stale-like behavior, consent/eligibility changes and user research. A valid like still cannot create a match after a block, deletion, restriction or hard eligibility change.

## 4. Incoming-like queue
**Approved baseline:** Free users see five eligible incoming likes unblurred at a time in a rolling queue; like-back/pass advances the queue. Remaining eligible likes are blurred while queued. Premium users see the full eligible queue.
**Recommended implementation:** Treat the five-like limit as a presentation entitlement, not a limit on stored incoming interest. Queue items are re-filtered against current account state, eligibility, blocks and pair suppression every time they are read or acted upon. A stale item never creates a match.
**Validate:** Five is an approved launch hypothesis, not a proven optimum. Instrument queue progression, pass/like-back behavior, fairness and user comprehension without fabricating activity or urgency.

## 5. Messaging and attachment retention
**Recommended default:** Retention periods must be defined by data class rather than one global duration. Separate user-visible conversation history, safety evidence, security logs, and backups. After unmatch, active messaging closes immediately; the approved read-only past-match history remains visible for 30 days, with report/block available and no messaging. This 30-day display window does not itself define backend evidence retention.
**Why:** Preserves the approved product experience while avoiding indefinite retention of all content.
**Validate:** UK counsel/privacy review, lawful basis, report/appeal needs, deletion rights, attachment processing, vendor capabilities and backup expiry. Do not promise complete immediate erasure where a documented lawful exception applies.

## 6. Pair exclusion after unmatch and account deletion
**Approved baseline:** A pair is permanently excluded from matching after unmatch, separately from blocking.
**Recommended implementation:** Store the minimum pseudonymous pair-suppression record needed to enforce this policy, separate from public profiles and ordinary analytics. It must not reveal who initiated unmatch. When an account is deleted, do not retain broad profile or message data solely for this feature.
**OPEN LEGAL/PRIVACY GATE:** Determine lawful basis, retention period, deletion transparency, reversibility and whether a minimal suppression token can be retained after account deletion. Do not implement indefinite retention without explicit review.

## 7. Age-assurance outage and appeal
**Recommended default:** Unknown, pending, failed or expired age-assurance state does not grant age-gated access. Show a neutral status and retry/support route; do not expose vendor fraud signals. Appeals must be accessible without requiring the user to enter the restricted product.
**Validate:** Didit configuration and evidence pack, provider callback authenticity, duplicate/out-of-order callbacks, false rejects, manual review capacity, retention/deletion and privacy/legal review.

## 8. Existing matches during pause or non-safety restriction
**Recommended default:** Pause alone does not close existing matches. A safety restriction is capability-scoped and can suspend messaging where necessary. A full suspension disables normal use. Account deletion immediately stops new activity and starts the documented deletion workflow.
**Why:** Avoids conflating a voluntary visibility setting with enforcement or deletion.
**Validate:** Publish clear behavior and test every transition against message send, delivery, push notifications and queued work.

## 9. Premium entitlement refresh
**Recommended default:** Server-authoritative entitlement derived from verified platform purchase/renewal/refund/expiry state. Reconcile platform notifications idempotently and periodically; never trust client-only purchase state. When Premium expires, apply free entitlements on the next authoritative refresh. Existing incoming likes are retained but the visible queue follows the current tier.
**Validate:** Apple/Google billing rules, restore purchases, refunds, grace/billing-retry states, webhook delay, taxes and consumer subscription rules. Safety features remain free regardless of billing state.

## 10. Recovery objectives and resilience
**Recommended default:** Define recovery point objective (RPO) and recovery time objective (RTO) for each critical system before selecting database/backups and cloud services. Start with one AWS region, managed PostgreSQL with Multi-AZ production resilience where feasible, automated backups and tested restoration. Avoid global database deployment until measured needs justify it.
**Why:** Reliability commitments must match actual budget, staffing and tested recovery capability.
**Validate:** Current regional quote, realistic restore drills, backup retention, encryption/key recovery, deletion-tombstone replay and operational ownership. Do not advertise a recovery SLA before it is proven.

## 11. API and data contract defaults
- Version public APIs and use response DTOs/allowlists; never return database records wholesale.
- Derive actor identity from authentication, not request body fields.
- Enforce object-level authorization for every read/write and separate admin audience/permissions.
- Use idempotency keys for likes, messages, match-triggering actions, deletion requests and billing callbacks where retries can duplicate side effects.
- Use a transactional outbox for database changes that must trigger asynchronous effects, with idempotent consumers.
- Keep blocks, eligibility, deletion and message authorization authoritative in PostgreSQL/policy services, never only in caches.
- Prefer short-lived scoped media URLs; verify object ownership and moderation state before serving media.
- Use structured safe error codes; never disclose private preferences, reporter identity, internal risk signals or another user's block state.
- Set service-level objectives after load and pilot testing, not by aspirational guess.

## 12. Architecture decision workflow
For each item, track:
- decision ID and title;
- status: proposed / approved / rejected / deferred;
- decision owner;
- evidence required;
- privacy/legal/vendor dependencies;
- implementation impact;
- date and rationale;
- tests proving the chosen behavior.

Do not treat the recommendations in this document as approved product decisions merely because they are written down. The user-approved product policies remain authoritative, and legal/security gates remain open until supported by evidence.

## Exit gate
Before physical schema/API freeze: confirm decisions on pause and outgoing likes, pass reversibility, like expiry, messaging retention, pair suppression after deletion, age-assurance outage/appeal, restrictions on existing matches, Premium entitlement refresh and recovery objectives. Resolve each explicitly or assign an owner and block dependent implementation.
