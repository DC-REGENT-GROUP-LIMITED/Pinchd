# Pinchd research: matching algorithms, fairness and explainability
**Research date:** 21 October 2026  
**Status:** Desk research; no mobile UI design in this phase.

## Executive summary
Pinchd's matching system should be a reciprocal recommendation system, not a popularity contest. It must enforce hard preferences exactly, treat soft preferences as ranking signals only when users have clearly allowed that, explain recommendations in plain language, and measure opportunity as well as match outcomes. Safety risk must remain separate from romantic desirability. A small local pool should be handled honestly rather than disguised with fake profiles or silent preference relaxation.

## Evidence and competitive signals

### 1. Matching is two-sided
A peer-reviewed field experiment published in *International Economic Review* found that a two-sided matching approach reduced recommendation inequality, improved predictive accuracy and produced more matched couples than competing algorithms in the platform studied. Its central insight is that a recommendation is only successful when both people are interested; optimising only the likelihood that one person clicks can concentrate exposure on a small subset of users.
- Source: Chen et al., “Reducing Recommendation Inequality via Two-Sided Matching: A Field Experiment of Online Dating,” *International Economic Review* (2023): https://onlinelibrary.wiley.com/doi/10.1111/iere.12631

A 2024 working paper formalises fairness of opportunity in reciprocal recommendation and finds that maximising expected matches alone can still produce significant unfairness. Its proposed approach is a useful research direction, but it is a preprint and not a ready-made production algorithm.
- Source: Tomita & Yokoyama, “Fair Reciprocal Recommendation in Matching Markets,” arXiv (2024): https://arxiv.org/abs/2409.00720

**Implication:** Candidate generation and ranking should consider both parties' stated constraints and likely interest, without treating a predicted “yes” as a fact. Track how often eligible users are shown to others, how concentrated exposure becomes, and whether the system is systematically starving parts of the eligible pool.

### 2. Current competitors describe evolving recommendation systems, but claims need context
Hinge says it introduced a deep-learning recommendation system in 2025 intended to predict mutual compatibility and reported a double-digit increase in matches overall. It also encourages some users to loosen preferences such as distance and explains that dealbreakers affect visibility. This is a company-reported result; the public statement does not establish the same effect for Pinchd or prove that all relaxed filters are acceptable to users.
- Source: Hinge, “Evolving Together: How Daters Helped Shape Hinge in 2025” (10 Nov 2025): https://hinge.co/newsroom/hinge-2025-product-evolution
- Source: Hinge, “How We Connect Daters”: https://hinge.co/how-we-connect-daters

Tinder's September 2026 explanation says its recommendations consider activity and overlap in active times, location, profile information and interaction feedback, and says it no longer uses a traditional Elo score. Tinder also offers optional AI-powered recommendations and a hub to review or delete generated insights. These statements demonstrate transparency and controls as product choices, not independent evidence that the ranking approach is fair or optimal.
- Source: Tinder, “Powering Tinder — The Method Behind Our Matching,” updated 1 Sep 2026: https://www.help.tinder.com/hc/en-us/articles/7606685697037-Powering-Tinder-Behind-Our-Matching
- Source: Tinder, “AI-powered matching,” updated 3 Apr 2025: https://www.help.tinder.com/hc/en-us/articles/34723594883213-AI-powered-matching

**Implication:** Do not copy competitors' engagement signals uncritically. In particular, if activity recency is used to improve the chance that a person can respond, it should not become a hidden popularity or always-online advantage. Inactive accounts should be handled transparently and with lifecycle rules.

### 3. Explainability is a trust requirement
Hinge publicly explains that its recommendation system combines stated compatibility settings with activity signals and says users should understand how their feed is shaped. Tinder explains some ranking factors and gives users controls over optional AI-generated insights.
- Sources: https://hinge.co/how-we-connect-daters ; https://www.help.tinder.com/hc/en-us/articles/7606685697037-Powering-Tinder-Behind-Our-Matching ; https://www.help.tinder.com/hc/en-us/articles/34723594883213-AI-powered-matching

**Implication:** Every recommended profile should have a truthful, short “Why this profile?” explanation based on visible, user-provided information and current settings. Explain categories of signals, not exact scores or another person's private preferences. Provide controls to inspect, edit or remove optional inferred interests.

### 4. Avoid popularity feedback loops
The field experiment above notes that recommendation systems optimising individual click prediction can concentrate recommendations, creating congestion and discouraging other users. A profile that receives early attention may collect still more exposure, while newer or less frequently selected users may struggle to receive any.
- Source: https://onlinelibrary.wiley.com/doi/10.1111/iere.12631

**Implication:** Evaluate exposure concentration and distribution, not only clicks, likes or matches. Consider freshness and limited exploration among eligible profiles, while respecting hard constraints and avoiding random exposure that users perceive as irrelevant. Never publicly rank people by desirability.

### 5. Preference integrity must be explicit
Competitor products sometimes suggest loosening filters to find more people. Pinchd should distinguish hard constraints from soft preferences as a first-class product rule. A small pool is a product reality, not permission to override user intent.
- This requirement follows Pinchd's product principles; it is not a claim that external research has established a single best filter policy.

**Implication:** Hard constraints (for example, age range or relationship intention when explicitly marked as non-negotiable) must be enforced in the normal matching path. Soft preferences may influence ranking, but relaxation must be explicitly opted into and the result clearly labelled. If the pool is exhausted, offer clear choices: wait, broaden a named soft preference, extend distance, or pause discovery.

## Proposed matching model (conceptual, not a final algorithm)
1. **Eligibility gate:** apply age, location/legal eligibility, visibility status, relationship-goal compatibility where declared hard, blocks, prior dismissals and other hard constraints.
2. **Reciprocal candidate score:** estimate whether both people may find the other relevant using declared preferences and user-provided profile details. Treat estimates as uncertain and never as consent.
3. **Ranking:** balance reciprocal relevance, freshness, variety, and fair opportunity across eligible users.
4. **Exploration:** reserve a controlled, measured share of exposure for eligible candidates with limited evidence, without bypassing hard preferences.
5. **Explanation:** generate a human-readable reason from auditable factors, such as a shared interest or compatible stated relationship goals.
6. **Feedback and control:** let users correct interests, change preferences, dismiss profiles and disable optional personalisation.
7. **Monitoring:** assess quality, fairness, preference compliance, safety and local pool health continuously.

Do not use safety-risk scores as romantic desirability signals. Do not infer sensitive traits from photos or message contents. Do not use facial recognition or other highly intrusive signals to infer identity or attractiveness.

## Proposed requirements and acceptance criteria

### Preference integrity
- **MATCH-01:** Every hard preference is represented explicitly and applied before ranking.
- **MATCH-02:** In standard discovery, hard-preference compliance is 100% for eligible results, excluding clearly documented data correction or moderation edge cases.
- **MATCH-03:** The system never silently relaxes hard preferences.
- **MATCH-04:** Any soft-preference expansion requires explicit user choice, identifies which preference changed and is reversible.
- **MATCH-05:** When no candidates remain, the product explains the reason at an appropriate level and offers wait/broaden/pause choices without guilt.

### Reciprocal matching and fairness
- **MATCH-06:** Ranking evaluates potential relevance to both people, not just one-sided swipe likelihood.
- **MATCH-07:** A prediction is never presented as a guarantee of interest, compatibility or consent.
- **MATCH-08:** Dashboards measure eligible-user exposure, exposure concentration, distinct mutual matches and outcomes by cohort where lawful and privacy-safe.
- **MATCH-09:** New and returning users have a defined cold-start path that does not require long-term behavioural tracking to receive useful recommendations.
- **MATCH-10:** Active status may be used only for a documented purpose such as avoiding stale recommendations; it must not create an unbounded always-online advantage.
- **MATCH-11:** No public desirability score, hidden shadow ranking disclosed as fact, or premium purchase that secretly boosts romantic compatibility.
- **MATCH-12:** Safety risk is held in a separate pipeline and can trigger verification, rate limits, review or restrictions, never a lower romantic desirability score.

### Explainability and data control
- **MATCH-13:** Each recommendation can provide a truthful explanation using visible profile facts and the viewer's settings.
- **MATCH-14:** Explanations do not reveal another user's private preferences, inferred sensitive traits or internal risk flags.
- **MATCH-15:** Users can review and correct their own profile attributes and declared preferences.
- **MATCH-16:** Optional AI personalisation is clearly opt-in, explains its data inputs, and offers a way to inspect/delete generated insights.
- **MATCH-17:** Recommendation logging is minimised, access-controlled and governed by documented retention periods.

## Post-launch measurement plan
- Hard-preference violation rate (target: zero in standard flow).
- Candidate eligibility pool size and proportion with no available candidates.
- Reciprocal match rate, measured per eligible recommendation cohort rather than raw app-wide count.
- Time to first relevant recommendation and time to first mutual match.
- Exposure distribution: share of impressions/likes concentrated in top deciles; Gini or another documented inequality measure; unique eligible profiles receiving exposure.
- New-user cold-start performance versus mature accounts, with safety and relevance guardrails.
- “Why this profile?” usefulness and correction rates.
- Soft-filter expansion frequency, reversal rate, and whether users understood what changed.
- Fairness and quality comparisons across relevant cohorts, only where lawful, proportionate and privacy-preserving.
- Complaints about irrelevant recommendations, preference violations, perceived unfairness and deceptive explanations.
- Safety reports and block/report rates monitored separately from romantic compatibility metrics.

## Experiment hypotheses
1. Reciprocal ranking increases mutual matches per eligible recommendation compared with one-sided ranking, without materially worsening hard-preference compliance or exposure inequality.
2. A controlled exploration policy increases the number of distinct eligible users who receive meaningful exposure without degrading user-reported relevance.
3. Plain-language “Why this profile?” explanations improve trust and perceived control without creating false certainty.
4. Explicit soft-filter expansion increases user understanding and satisfaction compared with generic “broaden your preferences” nudges.
5. A freshness signal can reduce stale recommendations without materially disadvantaging users who are less frequently online.

For each experiment, pre-register primary and guardrail metrics, use a defined analysis window, evaluate segment-level harms, and stop if hard-preference violations or safety outcomes worsen. Avoid concluding that an experiment improves real-world relationship quality from app engagement alone.

## Cold-start and local density
- Start with a limited launch geography or clearly defined cohort so eligible active supply is not fragmented across too many markets.
- Measure the number of active, eligible candidates per user, not just total registrations.
- Use honest explanations for a thin pool; never generate fake profiles, fabricated activity or false scarcity.
- Keep user choice intact: users can wait, change a specific soft preference, expand distance, or pause.
- Avoid showing already-dismissed profiles repeatedly unless the user explicitly chooses to revisit them.
- Define a reactivation and stale-account policy that protects privacy and does not misrepresent who is available.

## Limitations and open questions
- Competitor explanations describe their own systems and may omit details; reported improvements are not independent causal proof.
- The 2023 field experiment is strong evidence that two-sided design can help in the studied platform, but effects may vary with user base, objectives and market conditions.
- The 2024 fairness paper is a preprint and should be treated as a research lead, not a production-ready standard.
- No desk research can identify Pinchd's best ranking weights or estimate local pool density before launch.
- Sensitive-data use, profiling, automated decision-making and fairness audits require a UK GDPR/DPIA and legal review before implementation.
- Mobile UI design remains paused until the research and requirements package is complete.

## Product decision log
The following matching direction has been adopted as the product policy baseline; this does not approve a production algorithm or ranking weights:
- Hard eligibility gates before ranking; zero hard-preference violations in standard discovery.
- Reciprocal relevance rather than one-sided like prediction alone.
- Fair-opportunity and exposure-concentration monitoring.
- Truthful “Why this profile?” explanations using visible information and the viewer's own settings.
- Safety and trust enforcement remain separate from romantic desirability.
- No silent relaxation, fake activity, fake profiles or paid compatibility advantage.
- Controlled post-launch experiments with relevance, fairness, privacy and safety guardrails.

## Decision summary
**Recommended direction:** hard eligibility gates, reciprocal ranking, fair exposure monitoring, transparent recommendation reasons, explicit user control over soft-preference relaxation, privacy-minimised personalisation, and honest handling of cold start. Treat this as the product's matching policy baseline; choose and tune a production algorithm only after technical feasibility, privacy review and controlled post-launch measurement.

## Source register
1. Chen et al., *International Economic Review* (2023): https://onlinelibrary.wiley.com/doi/10.1111/iere.12631
2. Tomita & Yokoyama, arXiv (2024): https://arxiv.org/abs/2409.00720
3. Hinge, “Evolving Together” (10 Nov 2025): https://hinge.co/newsroom/hinge-2025-product-evolution
4. Hinge, “How We Connect Daters”: https://hinge.co/how-we-connect-daters
5. Tinder, “Powering Tinder — The Method Behind Our Matching” (updated 1 Sep 2026): https://www.help.tinder.com/hc/en-us/articles/7606685697037-Powering-Tinder-Behind-Our-Matching
6. Tinder, “AI-powered matching” (updated 3 Apr 2025): https://www.help.tinder.com/hc/en-us/articles/34723594883213-AI-powered-matching
7. Tinder, “Most Compatible” (updated 6 Jul 2026): https://www.help.tinder.com/hc/en-gb/articles/42908926408333-Most-Compatible
