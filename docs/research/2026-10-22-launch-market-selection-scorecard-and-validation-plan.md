# Pinchd — Launch Market Selection Scorecard and Validation Plan
**Date:** 22 October 2026  
**Status:** Decision framework; no market selected and no demand assumed.  
**Purpose:** Turn the existing launch strategy into a repeatable market comparison, opt-in demand test and evidence-based launch gate. Mobile UI design remains paused.

## 1. Decision to make
Choose one bounded first launch market only after comparing eligible adult supply, realistic reach, audience breadth, trust-and-safety capacity, cost and learning speed. A city name alone is not a market definition: record the geographic boundary, eligible age range, intended radius defaults, target relationship-goal segments, acquisition channels and support coverage.

Do not select London by default because it is large. Do not select a smaller city solely because it is cheaper. Use the same evidence standard for every candidate.

## 2. Candidate market shortlist process
Start with 3–5 candidate markets, not a national launch. Candidate names should be treated as hypotheses until the evidence pack is complete. For each candidate record:
- A clear boundary based on actual geography and practical travel patterns.
- Credible adult population data and source date.
- The eligible audience estimate with transparent assumptions; never infer protected or sensitive traits from names, images or other proxies.
- Local acquisition channels and realistic reach, including whether partners will actually participate.
- The expected cost per *activated, age-eligible* user, not cost per click or email.
- Capacity for support, moderation, urgent escalation and vendor handling during all promised coverage hours.
- Nearby expansion options if the initial market passes.
- Evidence confidence and what would change the decision.

## 3. Weighted market scorecard
Score each candidate from 1 (weak evidence/poor fit) to 5 (strong evidence/good fit). Add a confidence rating separately; a high score supported only by guesses must not win. Weightings below are starting evaluation weights, not validated facts.

| Criterion | Weight | Evidence needed for a high score |
|---|---:|---|
| Eligible supply potential | 20% | Credible local adult population and conservative estimates for the actual hard-preference pool |
| Reachability and acquisition | 15% | Tested channels, realistic partner access and cost per activated eligible user |
| Reciprocal breadth | 15% | Opt-in survey/waitlist shows a plausible mix of declared goals and preferences without inferring sensitive traits |
| Density/geographic coherence | 10% | A workable boundary and plausible supply within users' chosen radii |
| Trust and operational capacity | 15% | Named coverage, support and moderation capacity with escalation arrangements |
| Cost and runway | 10% | Downside/base/upside cash model and spend cap |
| Learning speed | 10% | Can recruit a diverse enough cohort, instrument outcomes and decide without overspending |
| Expansion adjacency | 5% | Credible path to adjacent areas without undermining existing liquidity |

Calculation: weighted score = sum(score / 5 × weight). Also publish a confidence label (high/medium/low) and a short evidence note for every row. The total is a decision aid, not a substitute for launch gates. A market fails regardless of score if trust/safety coverage, age assurance, legal readiness, or product reliability is insufficient.

## 4. Evidence quality ladder
- **Level A — Observed:** completed, consented waitlist/onboarding behaviour, actual campaign spend, partner commitment in writing, measured support load.
- **Level B — Strong external evidence:** dated official statistics, official reach estimates with stated limitations, documented channel benchmarks relevant to the audience.
- **Level C — Stated intent:** voluntary questionnaire answers or survey interest. Useful but not equivalent to verified activation.
- **Level D — Assumption:** untested conversion, reach, cost, retention or eligibility assumptions. Must be labelled and sensitivity-tested.

Do not mix survey interest with active users or national reach with local active supply. Every market estimate should include a low/base/high range and list its assumptions.
## 5. Bottom-up supply and activation model
Model each market with an explicit funnel. Keep variables separate so real data can replace assumptions.

1. Reachable adult audience: deduplicated audience plausibly reachable through selected channels.
2. Landing-page visits: reach × measured visit rate.
3. Waitlist opt-ins: visits × measured opt-in rate.
4. Qualified opt-ins: opt-ins that meet the intended market scope and voluntarily supply enough information for aggregate modelling.
5. Activated accounts: qualified opt-ins × observed onboarding completion rate.
6. Age-eligible active users: activated accounts × measured age-assurance pass rate × active-within-window rate.
7. Reciprocal eligible supply: active users remaining after both-sided hard-preference compatibility and visibility rules.

Never assume all opt-ins will install, pass age assurance, complete profiles, remain active or be mutually compatible. Use ranges for every unmeasured conversion. Do not publish a single forecast without its downside case.

### Minimum spreadsheet fields
Market; boundary; source/date; reachable audience low/base/high; channel; spend cap; visits; opt-ins; qualified opt-ins; onboarding completion; age-assurance pass; active-window rate; estimated reciprocal pool; cost per opt-in; cost per activated eligible user; support/moderation hours; confidence; owner; next validation date.

## 6. Pre-launch experiment plan
**Experiment A — Proposition test**
- Present the real positioning: “Serious about compatibility, flexible about relationship goals.”
- Describe what exists and what is planned accurately; do not imply a working app or guaranteed launch date if neither is true.
- Use a privacy notice, explicit waitlist opt-in and an easy withdrawal/deletion path.
- Measure qualified opt-in and cost per qualified opt-in, not just click-through.

**Experiment B — Market/channel comparison**
- Run small, capped tests across shortlisted markets with comparable creative, landing-page content and measurement.
- Avoid comparing markets with materially different spend, targeting or campaign duration without normalising and documenting it.
- Do not scrape competitor audiences, upload contacts without permission, buy dubious lists or fabricate testimonials.

**Experiment C — Preference and radius survey**
- Ask only the minimum fields needed to estimate aggregate reciprocal supply.
- Make questions optional where possible; explain purpose and retention.
- Separate identity from discovery preferences. Do not infer sensitive traits.
- Treat answers as stated intent, not proof of active demand or legal eligibility.

**Experiment D — Partner feasibility**
- Contact a small, relevant set of local community organisations, venues, creators or event organisers.
- Record actual responses and commitments; a theoretical list of partners is not evidence of reach.
- Do not promise features, safety outcomes or audience volumes that are not yet supported.

**Experiment E — Waitlist-to-beta conversion**
- Only after product and legal gates permit recruitment, measure actual onboarding, age assurance, profile readiness, activation and early retention.
- Keep the cohort bounded and instrument support/safety load before expanding.
## 7. Metric definitions and decision rules
- **Qualified opt-in rate:** qualified waitlist opt-ins / unique landing-page visitors.
- **Cost per qualified opt-in:** acquisition spend / qualified opt-ins.
- **Cost per activated eligible user:** acquisition spend / accounts that complete required onboarding, pass required age assurance and are active in the agreed window.
- **Reciprocal eligible pool:** candidates who pass both users' hard constraints and all safety/visibility gates; do not substitute raw registrations.
- **Lower-tail density:** distribution of eligible candidate counts per active user; publish median and lower percentile, not only the average.
- **Exhaustion rate:** proportion of users/sessions with no new eligible candidates under current settings.
- **Operational load:** reports, blocks, appeals and support requests per active cohort, with time-to-first-response and backlog age.
- **Healthy exit:** pause, delete or leave without treating it automatically as failure; collect optional reason only with transparent purpose.

Before experiments begin, write down the minimum evidence thresholds, sample limitations, decision owner and stop-loss budget. Do not set universal density or conversion numbers without local evidence. Use pre-agreed gates to avoid moving the goalposts after seeing results.

## 8. Go / hold / stop decision
**GO to a controlled beta only if** demand evidence is credible, lower-tail reciprocal supply is adequate for the declared launch cohort, hard-preference compliance tests pass, age assurance is operational, moderation and support coverage is real, privacy/legal gates are owned, deletion and blocking are tested, and downside cost is affordable.

**HOLD** if sample sizes are too small, the evidence relies mostly on assumptions, eligible supply is uneven, or operational coverage is not yet proven. Continue capped validation or narrow the cohort.

**STOP or redesign the market plan** if acquisition is unaffordable in the downside case, reciprocal supply cannot support the proposition, partners fail to materialise, or safe operations cannot be staffed. Do not respond by widening hard preferences or using fake activity.

## 9. Required evidence pack per candidate market
- One-page boundary and audience rationale.
- Dated public demographic sources and assumptions.
- Weighted scorecard plus confidence ratings.
- Low/base/high funnel and cost model.
- Acquisition test plan, spend limit and actual results.
- Waitlist privacy notice and data minimisation record.
- Aggregate preference/radius findings with small-group suppression.
- Written partner responses/commitments.
- Moderation/support coverage and incident escalation plan.
- Decision log: GO / HOLD / STOP, approver, date and review trigger.

## 10. Current decision and next steps
No city, launch date, radius default, acquisition budget or density threshold is approved by this document. Existing research explicitly leaves the market unselected.

Next actions:
1. Choose a shortlist using credible population and acquisition evidence.
2. Build a transparent low/base/high workbook using the fields above.
3. Validate channel access and partner willingness with small capped experiments.
4. Set launch-gate thresholds before seeing conversion results.
5. Revisit the scorecard after actual waitlist and beta data.
6. Do not start mobile UI design until the broader research and readiness gates are complete.
