# Pinchd — Synthetic Test-Data and Fixture Specification
**Date:** 22 October 2026  
**Status:** Proposed engineering/privacy review artifact. No fixtures have yet been generated or tests executed.  
**Purpose:** Define safe, deterministic test inputs for POC-01–POC-12 and lifecycle requirements without copying or reconstructing real people.

## 1. Safety boundary
All test records must be generated from fixed fictional values or seeded generators. Never import production, beta, vendor, support, analytics or social-media data; never use real identity documents, faces, phone numbers, email addresses, precise locations, messages or report content. Do not create synthetic data by perturbing a real record. No test fixture may be sent to a live age-assurance or payment workflow.

Use reserved domains (for example, `example.invalid`) and non-routable phone placeholders. Use fictional names that are not selected by searching for real-person matches. Coordinates must be omitted or use documented test coordinates for a fictional grid; avoid mapping coordinates to residences. Images should be generated geometric/color assets with no faces, metadata or embedded text, or use a deterministic placeholder byte sequence where image processing is not under test.

## 2. Fixture catalogue
| Fixture family | Minimum cases | Data boundary | Used by |
|---|---|---|---|
| Account lifecycle | invited, active, paused, restricted, deletion-pending, deleted, recovery-pending | UUIDs from fixed test namespace; synthetic timestamps | LC-01–LC-36, POC-01/09 |
| Age assurance | pending, pass, fail, inconclusive, expired, unavailable, malformed, replayed, wrong-session | Synthetic signed callback keys only; no identity artifacts | POC-05; eligibility gates |
| Preferences/matching | reciprocal eligible pair, one-sided preference, hard-gate mismatch, unknown optional field, tie, blocked pair | Fictional enum values and seeded preference combinations | POC-10; matching tests |
| Social graph | like/pass/match/unmatch/block, simultaneous block and message, stale candidate cache | Synthetic account IDs; no real contact graph | POC-02–04/07 |
| Messaging/safety | benign placeholder, reportable category labels, duplicate report, urgent-severity routing marker | No realistic traumatic narrative or intimate imagery; use category codes | POC-03/07; T&S flow |
| Media | valid placeholder, oversized, wrong MIME, metadata-bearing test file created solely for stripping test, rejected malware fixture from approved test corpus | Never use real photos or harmful content | POC-06 |
| Billing | sandbox product IDs, renewal/cancel/refund/restore events, duplicate/out-of-order webhook | Provider sandbox or local contract doubles only | POC-08 |
| Deletion/restore | pending job, retry, dead-letter, tombstone replay, late vendor callback | Synthetic identifiers; verify no resurrection | POC-09/12 |
| Accessibility | semantic labels, focus order, text scaling and reduced-motion states | No user research participant data in fixtures | POC-11 |

## 3. Determinism and reproducibility
- Pin generator version, seed, fixture schema version, timezone (UTC), locale and clock source.
- Each test names the fixture set/version and expected invariant; random cases record seed on failure.
- Keep the minimal fixture set in source control only after privacy/security review; use no secrets or generated credentials.
- Make fixtures idempotent and namespace-isolated per run; teardown must remove rows, objects, queues and local artifacts.
- Contract tests use local fakes or vendor-documented sandbox payloads with synthetic identifiers. Never fabricate a payload and label it as vendor evidence.
- Keep test logs redacted; assert absence of tokens, message bodies, identity fields and sensitive profile data.

## 4. Required invariant coverage
1. Unknown/pending/failed/expired age state denies age-gated access.
2. Eligibility is recalculated server-side; client flags cannot grant access.
3. Blocks take effect synchronously for reads, writes, queued notifications and stale caches.
4. Deleted or deletion-pending accounts cannot be reactivated by retry, restore or late callback.
5. Duplicate commands/events do not create duplicate likes, messages, charges or deletion effects.
6. Hard matching constraints are never softened by ranking, exploration, Premium or cold-start fallback.
7. Untrusted media is private until validated/moderated; public projections use allowlists.
8. Synthetic event payloads omit raw sensitive fields and secrets.
9. Out-of-order event handling preserves a valid monotonic lifecycle state.
10. Accessibility fixtures verify interaction semantics but do not substitute for human testing.

## 5. Handling, retention and acceptance
Store fixtures only in the approved isolated PoC environment. Keep generated artifacts to the minimum test needs; cleanup at run end, with any retained failing fixture reviewed and stripped of unnecessary content. No exact retention period is assigned here: environment owner must define run-artifact expiry within the approved environment policy.

Acceptance for ACT-11 / EV-09 requires: catalog reviewed by engineering and privacy; generator and schema version recorded; seeded reproducibility demonstrated; a scan/inspection confirms no real-user or sensitive identity data; cleanup procedure documented; and approval recorded by accountable roles. This document alone does not close ACT-11 or EV-09.

## 6. Linked work
Supports `2026-10-22-requirements-to-test-traceability-matrix.md`, `2026-10-22-technical-proof-of-concept-plan.md`, and ACT-11/EV-09 in `2026-10-22-full-fold-execution-and-evidence-register.md`. All related tests remain NOT RUN until dated evidence is attached.
