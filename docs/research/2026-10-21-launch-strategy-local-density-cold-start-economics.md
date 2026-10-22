# Pinchd research: launch strategy, local density and cold-start economics
**Research date:** 21 October 2026  
**Scope:** UK-first launch planning, market selection, network effects, local liquidity, acquisition, retention and measurement.  
**Status:** Desk research and planning framework. No market or launch location is selected by this note. Mobile UI design remains paused.

## Executive summary
Dating products have a local, two-sided liquidity problem: the app is only useful when a user can discover enough eligible, active people they might genuinely want to meet—and those people also have a reasonable chance of seeing relevant candidates. Total registrations and downloads can hide a weak market. Pinchd should launch in one deliberately bounded market or cohort, measure the active eligible pool by preference segment, and expand only when quality, safety and operational readiness pass defined gates. Never manufacture demand through fake profiles, misleading “active now” claims, false scarcity or spam.

## 1. Why local density matters
Dating is a network product, but unlike a global content feed, the usefulness of a recommendation often depends on geography, distance preferences, age eligibility, relationship intentions and reciprocal preferences. The relevant metric is therefore not “number of sign-ups”; it is the number of real, active, eligible candidates who could plausibly be shown to each user.

A 2023 peer-reviewed field experiment in online dating found that a two-sided matching system could improve matching outcomes while reducing recommendation inequality in the studied platform. It reinforces that platform value depends on mutual suitability and how opportunity is allocated, not just total activity.
- Source: Chen et al., “Reducing Recommendation Inequality via Two-Sided Matching: A Field Experiment of Online Dating,” *International Economic Review* (2023): https://onlinelibrary.wiley.com/doi/10.1111/iere.12631

Network effects can be nonlinear: a small increase in eligible local supply may substantially improve a user's options, while a larger total user count outside their relevant geography or preferences may add little value. This is a product hypothesis for Pinchd to measure, not a universal fixed threshold.

**Implication:** Report density as a distribution across active users and segments, not a single global average. Monitor how many eligible profiles each user has in a chosen radius and how many of those are recently active, profile-complete and able to receive recommendations.

## 2. Competitor and market evidence: use cautiously
Ofcom's UK Online Nation 2025 reporting provides third-party context on reach for major dating services. Reach is not the same as active users, paying subscribers, local availability, user quality or revenue, and should not be used as a direct estimate of Pinchd's obtainable market.
- Source: Ofcom, Online Nation 2025 report and data: https://www.ofcom.org.uk/media-use-and-attitudes/online-habits/online-nation-2025
- Source: Ofcom, Online Nation 2025 interactive report: https://www.ofcom.org.uk/media-use-and-attitudes/online-habits/online-nation-2025-interactive-report

Competitors have used different approaches to reduce cold start: focused markets, invite/referral loops, profile completion, curated introductions, event-based discovery and broader recommendation systems. These are examples, not proof that any single tactic will work for Pinchd. The product must not rely on scraping competitor users, importing contact books without clear permission, or making deceptive claims about local activity.

**Implication:** Build a bottom-up local market model using credible public population data, opt-in waitlist data, acquisition experiments and conservative activation assumptions. Do not extrapolate national app reach into local addressable users without a documented method.

## 3. Define a launch market by evidence, not intuition
Score candidate launch markets using:
- **Eligible population:** likely adult audience within the intended age, distance and relationship-goal range, based on lawful, non-sensitive public data and explicit user opt-ins.
- **Reachability:** cost and realistic ability to reach relevant people through local communities, creators, partnerships, PR and paid acquisition.
- **Geographic coherence:** a boundary that corresponds to real travel patterns and users' chosen distance settings, not an arbitrary marketing map.
- **Audience balance and breadth:** sufficient diversity of relationship goals and preferences to support reciprocal matching; avoid over-optimising for total headcount.
- **Operational support:** ability to provide moderation, support, verification and incident response for the launch region.
- **Trust context:** availability of credible local partners and channels for safe launch education.
- **Cost and learning speed:** ability to run a bounded experiment without overspending or creating an unmanageable moderation queue.
- **Expansion adjacency:** whether a successful market can grow into nearby areas without immediately diluting density.

Do not assume London is automatically the best first market because it is large. It may provide reach and diversity but also high competition and acquisition cost. A smaller, coherent city or community could offer faster learning if the eligible audience is reachable and sufficient.

## 4. Pre-launch validation without interviews
Pinchd does not need to treat interviews as a prerequisite. It can reduce uncertainty with privacy-preserving, non-deceptive tests:
1. A landing page that accurately explains the product proposition and asks for an explicit opt-in waitlist registration.
2. Separate acquisition experiments across candidate markets and channels with clear privacy notices.
3. A preference questionnaire limited to what is needed to estimate eligible-pool structure; do not request unnecessary sensitive details.
4. A transparent waitlist that makes no claim that matches or launch dates are guaranteed.
5. A referral test where invitations are voluntary, rate-limited and not uploaded from contacts without permission.
6. A readiness survey asking users what launch radius and relationship goals they would choose, with responses clearly labelled as stated intent rather than verified demand.
7. A market model with low/base/high activation assumptions and explicit uncertainty ranges.

Do not create fake profiles, simulate active users, manufacture likes or claim a user is likely to find matches when supply is not yet there.

## 5. Local liquidity metrics
Define metrics before opening the market:
- **Eligible active supply:** count of real accounts active within a defined rolling window, with complete enough profiles, that satisfy a user's hard constraints and visibility rules.
- **Candidate availability distribution:** median, 10th percentile and 90th percentile of eligible candidates per active user, segmented by chosen radius and relevant declared preferences.
- **Reciprocal candidate availability:** eligible candidates where declared hard constraints are compatible in both directions.
- **Time to first relevant profile:** time from onboarding completion to first eligible, non-repeated recommendation.
- **Time to first mutual match:** cohort distribution, not a single average; interpret cautiously because users differ in goals and behaviour.
- **Discovery exhaustion rate:** share of sessions/users who reach no new eligible profiles within their current settings.
- **Profile freshness:** share of shown candidates active within a documented window, while avoiding a punitive always-online advantage.
- **Activation:** proportion of opt-in waitlist accounts that complete onboarding, verify required age threshold and use discovery within a defined period.
- **Retention by cohort:** D1/D7/D30 return, but paired with satisfaction, pause and healthy exit measures rather than treated as the only outcome.
- **Quality and safety:** report rate, block rate, preference compliance, complaints, moderation backlog and support capacity.
- **Acquisition efficiency:** cost per activated, eligible user, not cost per email address or download alone.

Segmenting metrics by sensitive traits requires careful legal and privacy review. Prefer user-declared preferences and aggregated analysis; do not infer sensitive traits from names, photos or messages.

## 6. Proposed launch gates
These are suggested decision rules to calibrate in the market model, not universal benchmarks:
- **Gate A — Demand evidence:** waitlist and acquisition tests show credible interest from a sufficiently broad local pool, not only high click-through.
- **Gate B — Density:** the lower tail of eligible active supply meets a pre-agreed minimum for the launch cohort; do not rely on the average alone.
- **Gate C — Preference integrity:** automated and end-to-end tests demonstrate zero normal-flow hard-preference violations.
- **Gate D — Safety operations:** moderation, report/block flows, urgent escalation, age assurance and incident response are staffed and tested.
- **Gate E — Product readiness:** onboarding, profile, discovery, matching, messaging, account recovery, deletion and support work end-to-end.
- **Gate F — Economics:** base-case acquisition and support costs are sustainable, with downside runway modelled.
- **Gate G — Learning plan:** primary outcomes, guardrails, event definitions, experiment design and decision owners are documented.
- **Gate H — Expansion:** expand only after the current market meets quality, safety, retention and density targets for multiple cohorts.

If a gate fails, delay launch, reduce scope or continue building the waitlist. Do not relax hard preferences to make the numbers look better.

## 7. Launch sequencing
### Phase 0 — Research and instrumentation
- Complete product requirements, privacy review, safety operations plan and event taxonomy.
- Build a bottom-up market model with assumptions and sensitivity analysis.
- Select candidate markets only after comparing them consistently.

### Phase 1 — Waitlist and market experiment
- Test value proposition and acquisition channels.
- Measure opt-in waitlist quality, declared preferences, activation intent and cost.
- Publish clear expectations about launch timing and data use.
- Avoid implying that sign-up guarantees a match.

### Phase 2 — Controlled beta
- Open to a limited, geographically coherent cohort.
- Monitor supply, reciprocal availability, user support, safety queues and onboarding failures daily.
- Use staged invitations to avoid a burst of sign-ups overwhelming moderation.
- Keep an accessible feedback and reporting route.

### Phase 3 — Stability and optimisation
- Tune onboarding, relevance, preference controls and notification defaults.
- Run controlled experiments with guardrails.
- Evaluate cohorts over enough time to distinguish novelty from sustainable use.
- Track both successful connections and healthy exits/pauses.

### Phase 4 — Adjacent expansion
- Expand radius or geography deliberately, not all at once.
- Recalculate eligible pool and safety capacity before each expansion.
- Preserve local density by avoiding unnecessary fragmentation into too many tiny markets.
- Pause acquisition if support or moderation capacity is exceeded.

## 8. Acquisition channels and trust
Potential channels to test:
- Local creators whose audiences align with the product's respectful, relationship-goal-inclusive positioning.
- Community organisations, social groups, alumni networks and local events, with no implied endorsement unless formally agreed.
- Referral loops with consent, rate limits and anti-spam protections.
- Local PR and founder-led explanation of product differentiation.
- Paid acquisition with transparent creative and cohort-level cost measurement.
- Educational content on safer dating, preference integrity and healthy use.

Guardrails:
- No fake profiles, paid actors presented as ordinary members, fabricated likes or misleading “people are waiting” messages.
- No scraping, unsolicited bulk messages or uploading address books without informed permission.
- No targeting based on sensitive inferred traits.
- Disclose incentives and sponsorships.
- Do not promise relationship outcomes or guaranteed matches.
- Do not let premium status buy hidden compatibility or bypass another user's preferences.

## 9. Cold-start product policies
- Start with enough onboarding/profile quality to make recommendations useful, but do not create unnecessary form fatigue.
- Use profile-completion guidance that explains value without coercion.
- Use explicit, opt-in preference expansion for soft constraints only.
- Explain when the current pool is exhausted and provide choices: wait, broaden a named soft preference, extend distance or pause.
- Use controlled exploration only among eligible profiles and monitor fairness.
- Give new users a fair chance at exposure without granting permanent advantage to new accounts.
- Apply clear lifecycle rules to stale accounts; never label an account “active now” without a defensible basis.
- Avoid infinite swiping as the primary engagement mechanic; finite discovery can support deliberation, but test whether it improves user outcomes.
- Provide honest empty states rather than fabricated activity.

## 10. Unit economics and downside planning
Build a scenario model with explicit inputs:
- Reachable adult audience in the market.
- Landing-page conversion to waitlist.
- Waitlist-to-activated-user rate.
- Cost per waitlist lead and cost per activated user by channel.
- Profile completion and age-assurance completion rates.
- Eligible local pool after hard constraints.
- Reciprocal recommendation and match rates.
- 30/90-day active cohort retention and healthy pause/exit.
- Free-to-paid conversion, churn, refunds and payment fees.
- Verification, moderation, support, cloud, storage, messaging and vendor costs.
- Staffing/on-call coverage and contingency for safety incidents.
- Cash runway and acquisition spend caps.

Model low/base/high cases and one or more severe downside cases. Sensitivity-test activation, local eligibility, acquisition cost, retention and moderation cost. Do not use revenue forecasts that assume every registration becomes an active user or that retention grows automatically with more notifications.

## 11. Decision framework
For each candidate market, record:
- Evidence date and source.
- Population and audience assumptions.
- Acquisition channels and estimated cost range.
- Waitlist quality and stated preferences.
- Candidate supply estimate by relevant segment.
- Safety and moderation coverage.
- Risks and uncertainties.
- Launch-gate status.
- Decision owner, review date and trigger for proceeding or stopping.

**Decision rule:** choose the market where Pinchd can most credibly create a sufficiently dense, safe and relevant local network within available budget and operational capacity—not simply the market with the biggest headline population.

## Limitations and open questions
- Ofcom online reach data is useful market context but is not local active dating-app supply, revenue or market share.
- Local pool thresholds depend on geography, user preferences, activity and matching design; there is no universal minimum.
- Waitlist sign-ups and stated intent are imperfect proxies for activated use.
- Retention metrics alone do not establish relationship quality or wellbeing.
- The launch geography, budget, marketing mix, age range and operating capacity remain undecided.
- No mobile UI design should begin until research and requirements are complete.

## Decision summary
**Recommended direction:** use a focused local launch, transparent waitlist tests, reciprocal eligible-supply metrics, staged beta, safety capacity gates and conservative unit economics. Expand only when a market demonstrates meaningful density, acceptable safety performance and sustainable acquisition—not merely downloads or sign-ups.

## Source register
1. Ofcom, Online Nation 2025 report and data: https://www.ofcom.org.uk/media-use-and-attitudes/online-habits/online-nation-2025
2. Ofcom, Online Nation 2025 interactive report: https://www.ofcom.org.uk/media-use-and-attitudes/online-habits/online-nation-2025-interactive-report
3. Chen et al., “Reducing Recommendation Inequality via Two-Sided Matching: A Field Experiment of Online Dating,” *International Economic Review* (2023): https://onlinelibrary.wiley.com/doi/10.1111/iere.12631


## 12. Hinge-style geographic model — clarified working decision

Pinchd will use a Hinge-inspired user-selected location and maximum-distance radius model, rather than choosing a city-wide launch boundary as the discovery rule. The launch-market analysis remains about where to seed the first user community and acquisition channels; each user's location and distance preferences determine discovery eligibility.


### Hinge reference research and Pinchd interpretation
Hinge's official product explanation says its recommendations use stated preferences alongside activity and recommendation-system signals, and that it uses information to recommend people to each user and vice versa. Source: https://hinge.co/how-we-connect-daters

Hinge's Help Centre describes location and distance as user-controlled discovery preferences, and its public materials describe dealbreakers for age and location. Product options and paywall treatment can vary by market and over time; Pinchd should replicate the useful interaction model, not assume every current Hinge detail is stable or suitable. Sources: https://help.hinge.co/hc/en-us/articles/360011063294-How-do-I-change-my-Preferences and https://help.hinge.co/hc/en-us/articles/360011165853-How-do-I-set-my-location

A July 2026 Reddit discussion reports a Hinge distance-setting map with a circle indicating the selected radius; users noted this made the geographic boundary easier to understand, including around rivers and borders. This is anecdotal and the feature may be in a phased test, so treat it as inspiration to investigate rather than a guaranteed universal Hinge feature. Source: https://www.reddit.com/r/hingeapp/comments/1uz3abi/distance_setting_is_now_updated/

### Geo requirements for Pinchd
- Users can choose or adjust a location, with device location permission optional where a manually selected location is available.
- Users set a maximum-distance radius. The default product rule is a hard boundary: profiles outside the radius are excluded unless the user explicitly opts into a clearly labelled flexible mode.
- Use straight-line radius as the initial matching definition if following Hinge's documented pattern; make clear that radius is not travel time. Evaluate route-based distance only if later research demonstrates user need and engineering/privacy trade-offs are acceptable.
- Where practical, visualise the radius on a map so users can understand coverage, but never expose another person's precise coordinates or exact home/work location.
- Display approximate location only. Do not provide a live map of a person, precise distance that enables triangulation, or repeated location updates that expose movement patterns.
- Store the minimum location precision needed for matching. Separate any precise device coordinates from public profile location; document consent, purpose, retention, access control and deletion behaviour before implementation.
- Changing the selected location or maximum radius triggers a fresh eligibility evaluation and invalidates stale discovery results. It does not silently modify age, relationship-goal or other hard preferences.
- A city chosen for launch is a go-to-market seed area, not a hard product boundary. Users near the edge of that area may discover compatible users outside it if those profiles are inside their own selected radius and otherwise eligible.
- Launch density should be assessed using eligible active users within user-selected radii and preference constraints, not a single city population count. Aggregate metrics should be privacy-preserving and should not reveal small cohorts.

### Updated launch-market approach
The prior candidate list (London, Manchester, Birmingham, Bristol and Leeds) remains a shortlist for acquiring the initial community, not a ranking of discovery geography. Compare them using evidence-backed reachable early adopters, acquisition cost, operating/safety readiness and likely eligible supply. For each candidate, model more than one plausible catchment and radius distribution; do not assume every resident is in every other resident's dating pool.

No launch city, default radius, permitted radius range or map interaction has been selected. These require a product decision informed by accessibility, privacy, technical feasibility, market research and controlled testing. Mobile UI design remains deferred until the research and requirements phase is complete.
