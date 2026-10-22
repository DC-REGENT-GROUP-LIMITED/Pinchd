# Pinchd — Matching Engine Specification and Fairness Test Plan (Draft)
**Date:** 22 October 2026  
**Status:** Policy-to-engine translation for review; not validated production weights or an empirical claim about dating outcomes.  
**Scope:** Hard eligibility, reciprocal compatibility, ranking, missing values, deterministic tie-breaks, exploration, explanations and synthetic tests.

## 1. Non-negotiable principles
1. **Eligibility is a gate, not a score.** A candidate who fails an active hard constraint must not be returned or reintroduced through exploration.
2. **Evaluate both directions.** Candidate A must be eligible for B and B for A. One-sided interest is not a match.
3. **Hard and soft preferences remain distinct.** Soft preferences affect ordering; hard constraints cannot be relaxed to fill a feed.
4. **Missing is not a positive answer.** Unknown, undisclosed, stale or unverified values do not mean compatible, open-minded or “no preference.”
5. **Safety and romantic ranking are separate.** Age assurance, blocks, restrictions, reports and consent are gates or safety workflows, not score features.
6. **Explanations must be truthful.** Reasons must come from current, permitted data; never expose sensitive inferences or enforcement signals.
7. **Profile completeness is not desirability.** More photos, longer bios, more prompts or verification status are not general romantic-worth scores.
8. **No paid ranking advantage.** Premium cannot improve hidden ranking or bypass constraints.
9. **Version and measure policy.** Candidate decisions should be reproducible from a versioned policy and privacy-approved feature snapshot.

## 2. Inputs and canonical definitions
Inputs may include account lifecycle, current age assurance, visibility and restrictions, profile/media gates, both users' age/gender/radius preferences, coarse distance freshness, per-field strictness, relationship goals, mutual likes, matches, blocks, pair exclusions, approved soft signals, exposure history, policy version and correlation ID.

Do not pass identity documents, biometric templates, private safety evidence, inferred sexuality, protected-trait proxies or unapproved sensitive fields into ranking.

### Relationship goals
Represent long-term, short-term, casual, open-minded/exploring, not-sure-yet and undisclosed as distinct values. Store goal separately from openness and from the other person's desired goals.
- Explicit incompatibility excludes only where a user deliberately set a hard constraint.
- Soft goal fit may influence ordering when both values are known and the approved matrix supports it.
- “Not-sure-yet” and “open-minded/exploring” are not interchangeable.
- Missing/undisclosed never implies openness to casual dating.
- Long-term preference is not consent to casual dating.
- Unsupported combinations yield neutral/unknown, not fabricated compatibility.

### Age and distance
- Age range is always a hard discovery boundary; legal minimum age is an independent gate.
- Radius is hard by default and uses an approved coarse location method. No exact-location disclosure or triangulation.
- If a hard gate cannot be checked reliably, withhold the candidate until safely refreshed; never silently widen age or radius.
- Location freshness, uncertainty and travel mode are explicit inputs; travel mode does not override another person's hard boundaries.

## 3. Eligibility pseudocode
Illustrative only; production code must run server-side and be tested at discovery, like, match and message boundaries.

```text
eligible(A, B, now):
    if A.id == B.id: return false
    if not account_allows_discovery(A, now): return false
    if not account_allows_discovery(B, now): return false
    if not current_age_assurance_pass(A, now): return false
    if not current_age_assurance_pass(B, now): return false
    if not profile_and_media_gates_pass(A, B): return false
    if has_discovery_restriction(A, now) or has_discovery_restriction(B, now): return false
    if block_exists(A, B) or block_exists(B, A): return false
    if pair_exclusion_exists(A, B) or active_match_exists(A, B): return false
    if not age_preference_pass(A, B) or not age_preference_pass(B, A): return false
    if not gender_preference_pass(A, B) or not gender_preference_pass(B, A): return false
    if not radius_preference_pass(A, B, now) or not radius_preference_pass(B, A, now): return false
    if not all_hard_attribute_preferences_pass(A, B): return false
    if not all_hard_attribute_preferences_pass(B, A): return false
    return true
```

- Keep internal reason codes for testing/operations; do not reveal sensitive block, enforcement or fraud details to users.
- Use the same eligibility contract in both directions; any intentional asymmetry must be specified.
- Re-check eligibility inside mutation transactions. A discovery result is not permission to like or message.
- An unset preference imposes no constraint; an unknown value for a configured hard constraint does not pass.
- Age assurance, enforcement and block state must be authoritative and fresh enough for the operation.

## 4. Soft ranking model
Generate candidates only from the eligible set. Validate retrieved candidates with canonical eligibility before returning them. Candidate generation must not use sensitive data or hidden desirability proxies.

| Component | Role | Initial handling |
|---|---|---|
| Relationship-goal fit | Soft alignment where known | Neutral if unknown; matrix requires research |
| Mutual stated lifestyle fit | User-selected soft preferences | Score approved fields only |
| Shared interests/prompts | Optional relevance | Modest weight; no sensitive inference |
| Intent recency | Potentially stale field | Use only if editable and freshness is handled |
| Exposure balancing | Avoid repetitive overexposure | Bounded adjustment; fairness monitoring |
| Exploration | Controlled variety | Sample only from eligible candidates |
| Profile completeness | Not a desirability signal | Do not score as romantic worth |
| Premium/payment/comment length | Prohibited ranking inputs | Excluded |
| Safety flags | Not ranking inputs | Eligibility/enforcement workflow |

No numeric weights are approved. A simple additive score may be considered only after each feature, missing-value behavior and weight are documented and validated. Do not present scores as probabilities of attraction, compatibility or relationship success.
## 5. Unknown values, tie-breaks and exploration
- Unknown feature contribution is neutral or omitted with documented score normalization; it is never coerced into agreement or disagreement.
- Do not reward users for revealing sensitive data. Optional missing fields should not automatically penalize a profile.
- Stale intent is omitted or handled under a transparent freshness rule, not silently treated as current.
- If all soft features are unknown, eligible candidates may still be ordered using stable variety/exposure rules.
- For near-equal scores: bounded exposure balancing, then a deterministic per-viewer/per-policy-version seed for a feed window, then stable ID as final fallback.
- Do not use paid status, protected traits, inferred vulnerability, comment length or arbitrary account age as tie-breakers.
- Bind cursors to policy version and feed window. Material eligibility changes require revalidation; stale cursors restart safely.
- Exploration samples only from the already eligible set. It never widens age, gender, radius or hard preferences.
- Set exploration rate only as an experiment parameter after instrumentation exists. Add a rollback switch and compare with a stable baseline.
- Cap repeat impressions and monitor eligible candidates who are never surfaced. Do not use prior popularity as an unexamined desirability multiplier.

## 6. Reciprocal compatibility and explanations
The feed is a recommendation, not a guarantee. Any reciprocal compatibility estimate must use documented fields and versioned rules, show unknown when coverage is insufficient, avoid unvalidated percentages, and never imply a like exists unless a valid reciprocal like exists.

Approved explanation patterns may include:
- “You both selected [shared interest].”
- “Their stated relationship goal aligns with yours.” (only when both values are known and the approved matrix supports it)
- “This profile fits the preferences you set.” (generic fallback without disclosing hidden settings)
- “You both answered [prompt topic] in a similar way.” (only if the similarity method is explainable and approved)

Never expose another user's exact distance, hidden preferences, age-check status, enforcement status, report history or reasons for absence. Never invent a shared interest. If no grounded reason exists, omit the explanation. Validate projection versions so stale data cannot create false claims.

## 7. Abuse, privacy and security
- Rate-limit and detect scripted likes, account farming, repeated create/withdraw cycles and scraping without protected-trait proxies.
- Keep anti-abuse features separate from romantic ranking; restrict staff access.
- Block/unmatch must invalidate feeds, queues, caches and notifications; re-check canonical state at action time.
- Avoid exposing exact candidate counts, queue totals or hidden-block state where they enable enumeration.
- Do not train/evaluate ranking on private messages, safety reports or sensitive fields without separate purpose, lawful basis, transparency and review.
- Reports route to safety operations; they are not merely a score penalty that leaves someone circulating.

## 8. Synthetic acceptance tests
Use deterministic unit/integration tests with seeded accounts, known policy version and expected eligibility. Add property-based and concurrency tests as implementation matures.

| ID | Setup | Expected result |
|---|---|---|
| MATCH-01 | A's age preference excludes B | B not returned regardless of score |
| MATCH-02 | B's age preference excludes A | B not returned to A |
| MATCH-03 | Either age assurance pending, expired or unknown | Not eligible |
| MATCH-04 | Hard radius exceeded | Exclude; no silent widening |
| MATCH-05 | Candidate not in selected gender preference | Exclude; no identity inference |
| MATCH-06 | Hard attribute fails either direction | Exclude before ranking |
| MATCH-07 | Soft preference mismatch, hard gates pass | Remains eligible; may rank lower |
| MATCH-08 | Soft field undisclosed | Neutral/omitted contribution |
| MATCH-09 | Relationship goal undisclosed | Never assume openness to casual |
| MATCH-10 | Only one side likes | No match or messaging permission |
| MATCH-11 | Reciprocal likes arrive concurrently | One match and one conversation |
| MATCH-12 | Block occurs during feed paging | Action-time like rejected |
| MATCH-13 | Pair exclusion exists | Excluded from rediscovery |
| MATCH-14 | More photos/prompts, no better fit | Completeness alone does not raise score |
| MATCH-15 | Candidate has Premium | No rank advantage or gate bypass |
| MATCH-16 | Long comment on a like | Comment length does not affect ranking |
| MATCH-17 | Explanation source is stale | Omit or recompute; never fabricate |
| MATCH-18 | All soft features unknown | Stable, non-fabricated ordering |
| MATCH-19 | Candidate restricted from discovery | Exclude, do not merely penalize |
| MATCH-20 | Stale cursor after policy change | Safe restart/error; no ineligible leak |
| MATCH-21 | Location too stale/uncertain | Apply safe policy; no precise-location leak |
| MATCH-22 | Exploration chooses a candidate | All hard gates still pass |
| MATCH-23 | Block/unmatch with queued push | Worker rechecks and suppresses delivery |
| MATCH-24 | Gender identity differs from discovery preferences | Keep fields separate; apply approved rules only |
## 9. Property-based invariants
- Every returned candidate passes eligibility in both directions.
- Changing only a soft feature never changes hard eligibility.
- Lowering a score never makes an ineligible candidate eligible.
- Premium changes neither eligibility nor rank score.
- Profile completeness alone does not improve romantic-fit score.
- A block or pair exclusion suppresses a candidate regardless of score.
- Same inputs, policy version and seed produce the same order.
- Missing values never become positive agreement by coercion.
- Concurrent reciprocal likes never create duplicate active matches or conversations.

## 10. Metrics and fairness review
Measure by market and time window using lawful, privacy-safe cohort analysis:
- Eligible pool size, zero-result rate and candidate supply.
- Impression distribution, repeat exposure and unseen eligible share.
- Like, reciprocal-like, match and conversation-start rates as funnel diagnostics, not personal worth.
- Exposure by onboarding cohort and profile completeness, without scoring completeness as desirability.
- Hard-gate error audits using synthetic and consented test accounts.
- Explanation accuracy and stale/missing explanation rates.
- Report/block rates and post-block message/notification incidents.
- Feature coverage and missingness across relevant cohorts.
- Premium versus free outcomes as a fairness audit, never a rationale for paid ranking advantage.

Do not optimize only for swipes, likes or session time. Define success around relevant, consensual conversations, user-reported fit and safety outcomes, with guardrails against harassment, unwanted contact, exclusion and manipulation. Observational data alone does not establish causation.
## 11. Launch gates and unresolved research
1. Approve a relationship-goal compatibility matrix through user research, including unknown behavior.
2. Specify the hard/soft field catalogue and per-field strictness semantics.
3. Validate candidate retrieval, geo precision, stale-location behavior and policy versioning in a proof-of-concept.
4. Define deterministic tie-breaks, cursor behavior, exploration experiment and rollback.
5. Complete a threat model for enumeration, scraping, stalking and inference from results.
6. Complete privacy review of features, metrics, cohort reporting and retention.
7. Build synthetic fixtures for gender combinations, relationship goals, unknowns, location boundaries, blocks, restrictions, stale states and concurrency.
8. Run property-based tests and adversarial review before controlled beta.
9. Establish baseline metrics and alert thresholds before ranking experiments.
10. Keep numeric weights, score thresholds and exploration rates unapproved until supported by evidence.

**Next step:** review this policy translation against the approved decision baseline, then implement pure eligibility/ranking functions and synthetic tests only after open semantics are resolved. This document does not authorize production ranking or mobile UI design; UI work remains paused.
