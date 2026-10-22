# Pinchd research: monetisation, subscription economics and ethical pricing
**Research date:** 21 October 2026  
**Scope:** Revenue model, premium feature boundaries, subscriptions, payments, pricing experiments, unit economics, cancellation and consumer trust.  
**Status:** Desk research and product strategy, not legal or financial advice. Confirm current UK consumer law and app-store rules before launch. Mobile UI design remains paused.

## Executive summary
Pinchd should monetise optional convenience and expression—not access to safety, consent, core communication or honest compatibility. A credible starting model is a useful free core plus a small number of optional premium benefits that are transparent, reversible and easy to cancel. Pricing should be tested against user-perceived value, retention, refunds, support burden and trust, not only short-term conversion. Avoid hidden pay-to-win ranking, fake scarcity, misleading discounts, confusing renewal terms, paywalls around essential safety, and dark patterns that make cancellation harder than subscribing.

## 1. Proposed free/premium boundary

### Keep free as core product and trust infrastructure
Recommended free features:
- Create and maintain a profile and set relationship goals.
- Set and enforce hard dealbreakers.
- Receive a useful baseline of eligible recommendations.
- Like and match within a transparent, non-manipulative model.
- Send and receive messages after mutual connection.
- Block, report, unmatch and access baseline verification/age assurance where required.
- Use essential privacy controls, notification controls, pause/discovery off, account deletion and data-rights routes.
- See honest “why this profile?” explanations and understand when the eligible pool is limited.
- Access core safety guidance and reporting/appeals routes.

Do not charge to reveal a serious safety report, stop contact, block someone, report abuse, protect a user from an account, or use legally required age assurance. Do not sell access to another user's consent or bypass their preferences.

### Potential premium benefits to test
- Additional profile presentation/customisation and optional prompts.
- Optional coaching for profile wording or message drafting, always user-reviewed.
- Convenience controls such as additional discovery organisation, saved preferences or travel mode, provided they do not bypass hard constraints or misrepresent location.
- More detailed but honest compatibility explanations based on user-provided data.
- Extra ways to manage discovery, such as revisiting previously dismissed profiles, if designed without undermining consent or user controls.
- Advanced soft-preference controls, only when the user explicitly chooses the change and the product labels what changed.

Premium features should be optional enhancements rather than repairs for an intentionally crippled free experience. Avoid paying to send unlimited unsolicited messages, paying to evade blocks, hiding abuse reports behind a subscription, or buying a hidden boost to predicted desirability.

## 2. Competitor pricing and product patterns: interpret carefully
Competitor subscriptions show that dating apps commonly package convenience, discovery controls, visibility and profile tools into tiers. Public prices vary by country, age, platform, promotional offer, billing period and account eligibility. A competitor's current price is not a direct benchmark for Pinchd's willingness to pay.
- Tinder subscription and feature information: https://www.help.tinder.com/hc/en-us/categories/115000755686-Subscriptions-Billing
- Hinge subscription help: https://help.hinge.co/hc/en-us/categories/360001641133-Subscriptions-and-Payments
- Bumble subscriptions and payment support: https://support.bumble.com/hc/en-us/categories/28534052341149-Subscriptions-and-Payments

**Implication:** Recheck live UK in-app prices immediately before final pricing decisions. Compare the actual benefit bundle, refund terms, cancellation steps, free-core usefulness and user trust—not only monthly headline prices. Do not copy competitor feature names or assume their conversion rates transfer to Pinchd.

## 3. Ethical pricing principles
- **Transparent total cost:** show price, billing period, trial end date if any, renewal amount and renewal cadence before purchase.
- **No deceptive urgency:** avoid false countdowns, fabricated “last chance” offers or claims that someone is about to disappear unless factually justified.
- **No shame or guilt:** do not imply a user is single because they failed to pay, or that premium is required to be desirable.
- **No hidden compatibility boost:** never claim a paid tier improves romantic compatibility unless the actual mechanism and evidence support that claim; recommended default is not to sell compatibility rank.
- **No safety paywall:** block, report, core messaging and required verification remain free.
- **Symmetric choice architecture:** make subscribing and declining clear; avoid visually hiding the free option.
- **Cancellation parity:** cancellation should be findable and straightforward, with clear effective date and access end date.
- **No forced continuity by confusion:** renewal is explicit; reminders are accurate; refund and support routes are easy to find.
- **Respect pause and exit:** users can pause discovery, mute notifications and delete accounts without punitive messaging.
- **Accessibility:** pricing, renewal and cancellation information must be understandable to people using assistive technologies and varied reading levels.

## 4. UK legal and platform checks

### Consumer protection and subscriptions
The UK Digital Markets, Competition and Consumers Act 2024 includes a regime for unfair commercial practices and new subscription-contract provisions. The implementation timetable and exact requirements can change; as of this research date, Pinchd must verify the commencement and current rules before launching a subscription flow. Do not assume every provision is already in force merely because the Act exists.
- UK legislation, Digital Markets, Competition and Consumers Act 2024: https://www.legislation.gov.uk/ukpga/2024/13/contents
- GOV.UK, Digital Markets, Competition and Consumers Act guidance: https://www.gov.uk/government/collections/digital-markets-competition-and-consumers-act-2024

Core compliance checklist for legal review:
- Clearly disclose price, duration, renewal, cancellation and trial conversion before consent.
- Preserve a record of the terms and price accepted at checkout.
- Send any required subscription reminders and renewal notices at the applicable times.
- Provide cancellation in the required form and explain when the subscription ends.
- Handle statutory cancellation rights, refunds and exceptions correctly.
- Avoid misleading scarcity, fake discounts, hidden fees and obstructive cancellation flows.
- Keep subscription terms consistent across app-store and direct-web payment routes.

### App-store billing rules
Apple and Google have store-specific rules for in-app digital purchases, external links, billing disclosures, subscriptions, account deletion and refunds. Rules differ by storefront, entitlement, product type and jurisdiction and change frequently. Pinchd must review the latest policies for the exact launch distribution model.
- Apple App Review Guidelines: https://developer.apple.com/app-store/review/guidelines/
- Apple, auto-renewable subscriptions: https://developer.apple.com/app-store/subscriptions/
- Google Play, subscriptions policy: https://support.google.com/googleplay/android-developer/answer/9900533?hl=en-GB
- Google Play, payments policy: https://support.google.com/googleplay/android-developer/answer/9858738?hl=en-GB

**Implication:** Decide early whether subscriptions are sold through Apple/Google, the web, or a compliant mix. Model store commissions, payment processing, taxes, refunds, support and reconciliation before comparing net revenue. Do not build a bypass flow without checking current platform rules.

## 5. Revenue model and unit economics
Model contribution margin per paying user and per active user. Include:
- Gross subscription receipts by plan and billing cadence.
- Store/payment fees, VAT and other applicable taxes.
- Refunds, chargebacks, fraud and payment failures.
- Trial conversion, renewal, churn, downgrades and pauses.
- Hosting, media storage, moderation, verification and customer support costs.
- Acquisition cost by channel and cohort.
- Discounts, introductory pricing and promotion expiry.
- Customer lifetime value (LTV) under conservative retention assumptions.
- Contribution margin after variable service and support costs.
- Cash runway and sensitivity to lower conversion/higher churn.

Avoid using “LTV:CAC” as a single proof of viability without stating the model, cohort maturity and assumptions. Early cohorts are noisy; retention and payback estimates should be updated as observed data accumulates.

## 6. Pricing experiments
Use a pre-registered test plan:
1. Test benefit concepts before testing exact price: does the feature solve a real user problem?
2. Test willingness-to-pay with transparent, non-binding research or real offers that comply with consumer law and platform rules.
3. Keep all core functions and safety controls identical across pricing experiments.
4. Randomise fairly at account/cohort level and avoid changing prices in ways that create opaque unequal treatment.
5. Define primary outcome (e.g. contribution margin per eligible active user) and guardrails (refunds, complaints, cancellations, trust, free-user outcomes, safety).
6. Use a sufficiently long window to observe renewal, not just initial purchase.
7. Do not infer demand from button clicks if the offer is not actually available at the displayed terms.
8. Provide accessible disclosures and a straightforward decline path.
9. Stop tests if confusion, complaints, refund rates or safety harms increase.
10. Document the decision, uncertainty and date for review.

Potential experiments:
- Optional profile coaching vs no coaching.
- Travel mode vs standard discovery controls.
- Extra profile customisation.
- Enhanced explanation of compatibility based on declared interests.
- Subscription tier structure vs one simple premium tier.

Do not experiment with charging for safety, changing a user's hard preferences without permission, deceptive price anchoring, or obscuring cancellation.

## 7. Pricing architecture to evaluate
Start by modelling three possible structures; do not commit before market and cost evidence:
- **Free + one premium tier:** easiest to explain and operate; likely a good baseline for a new product.
- **Free + two premium tiers:** more segmentation, but higher choice complexity and risk of feature confusion.
- **Free + occasional one-off convenience purchases:** potentially useful for non-recurring tools, but can encourage manipulative prompts if not carefully designed.

A launch recommendation is to test a simple free core plus one optional premium tier first, unless evidence demonstrates that a different model is clearer and more sustainable. Avoid launching several tiers before there is enough usage data to understand demand.

## 8. Subscription lifecycle requirements
- **MON-01:** Every offer shows total price, billing period, renewal terms, trial conversion and cancellation route before purchase.
- **MON-02:** Checkout records plan ID, displayed price, currency, billing cadence, accepted terms version, timestamp, storefront and receipt/reference.
- **MON-03:** Users can find current plan, renewal date, cancellation steps and effective end date.
- **MON-04:** Cancellation is not hidden behind repeated retention screens, guilt copy or unnecessary support contact.
- **MON-05:** Pause, downgrade and cancellation options are presented accurately where available; do not imply a pause stops billing unless it does.
- **MON-06:** Refund and billing-support routes are clear and consistent with the payment provider/store.
- **MON-07:** Failed payments, renewals and subscription status changes are communicated accurately without exposing sensitive dating activity in notifications.
- **MON-08:** The app does not remove essential safety features when a subscription expires.
- **MON-09:** Users can delete their account independently of cancelling a subscription, and the product explains any separate billing step needed.
- **MON-10:** Pricing changes are communicated with legally and contractually appropriate notice; no silent price changes.
- **MON-11:** Entitlements are server-authoritative and reconcile with store/web receipts; client state alone cannot grant paid access.
- **MON-12:** Analytics events for paywalls and purchases do not include message text, sensitive profile attributes or unnecessary identifiers.

## 9. Metrics and guardrails
Commercial metrics:
- Paywall view-to-purchase conversion by cohort and offer.
- Trial start, trial conversion, renewal, churn, pause, downgrade and reactivation.
- Net revenue and contribution margin per eligible active user.
- Net revenue retention and cohort payback with confidence ranges.
- Refund, chargeback and payment-failure rates.
- Support contacts per 1,000 paying users and time to resolution.

Trust and product guardrails:
- Cancellation completion rate and time; cancellation-related support contacts.
- Refund requests citing confusion or unexpected renewal.
- Complaints about deceptive pricing or paywall pressure.
- Free-user outcomes: eligible recommendations, mutual matches, messaging, safety-control availability and satisfaction.
- Preference-integrity violations and safety-report handling remain at zero/target levels regardless of subscription status.
- Differential impact across cohorts, where lawful and privacy-safe.
- Users who pause or leave should not receive shame-based or excessive win-back messaging.

## 10. Launch gates
1. Define free-core commitments and document why each premium feature is optional.
2. Obtain legal review of consumer/subscription law, privacy and app-store distribution.
3. Choose billing architecture and model gross-to-net economics.
4. Draft and usability-test pricing disclosures, renewal notices and cancellation flow.
5. Verify store policy requirements for each launch country and platform.
6. Test purchase, restore, renewal, refund, cancellation, downgrade, account deletion and entitlement reconciliation.
7. Test failed-payment and grace-period states.
8. Ensure safety, reporting, blocking and essential privacy controls remain available to all users.
9. Establish dashboards for revenue, churn, refunds, complaints and free-user quality.
10. Review pricing and feature boundaries after enough cohort evidence, without manipulating the free experience to force conversion.

## Limitations and open decisions
- Current competitor prices are dynamic and vary by platform, offer, location, age and billing period; this report does not set Pinchd's price.
- Consumer subscription-law commencement and detailed requirements must be rechecked against current UK law at implementation.
- App-store rules and commercial terms may differ by storefront and evolve; confirm current policy before shipping.
- Willingness to pay, premium conversion, retention and unit economics cannot be inferred reliably from desk research alone.
- The exact free/premium boundary, price, billing cadence and payment route remain open decisions.
- Mobile UI design remains paused until the research and requirements package is complete.

## Decision summary
**Recommended direction:** a genuinely useful free core with free safety and mutual messaging; one optional premium tier to test first; transparent billing and easy cancellation; no hidden compatibility ranking or deceptive urgency; and pricing decisions based on contribution margin, retention, user trust and free-user outcomes together.

## Source register
1. UK Digital Markets, Competition and Consumers Act 2024: https://www.legislation.gov.uk/ukpga/2024/13/contents
2. GOV.UK, Digital Markets, Competition and Consumers Act guidance: https://www.gov.uk/government/collections/digital-markets-competition-and-consumers-act-2024
3. Apple App Review Guidelines: https://developer.apple.com/app-store/review/guidelines/
4. Apple, auto-renewable subscriptions: https://developer.apple.com/app-store/subscriptions/
5. Google Play, subscriptions policy: https://support.google.com/googleplay/android-developer/answer/9900533?hl=en-GB
6. Google Play, payments policy: https://support.google.com/googleplay/android-developer/answer/9858738?hl=en-GB
7. Tinder subscription and billing help: https://www.help.tinder.com/hc/en-us/categories/115000755686-Subscriptions-Billing
8. Hinge subscription and payment help: https://help.hinge.co/hc/en-us/categories/360001641133-Subscriptions-and-Payments
9. Bumble subscriptions and payment support: https://support.bumble.com/hc/en-us/categories/28534052341149-Subscriptions-and-Payments


## 11. Pricing research update — 21 October 2026

This section adds a current UK storefront snapshot and clarifies the research-first decision. It does not establish a launch price.

### UK competitor snapshot (observed 21 October 2026)

Public UK iOS listings expose multiple in-app purchase price points, but often omit the billing duration or show offers for different cohorts. Treat these as observed listing examples, not clean monthly equivalents.

| App | Observed UK App Store price examples | What the official feature descriptions establish |
|---|---|---|
| Tinder | Plus £3.99; Gold examples £7.49, £11.99 and £14.99 | Tinder Gold includes seeing who liked the user; Plus includes conveniences such as unlimited likes, rewind and Passport. |
| Bumble | Boost examples £7.49, £8.99 and £10.99; Premium examples £13.99, £14.99 and £16.99 | Bumble Premium includes seeing the “Liked You” queue; Boost and Premium bundles differ. |
| Hinge | Hinge+ examples £12.99 and £14.99; other listed subscription entries include £14.99 and £29.49 | The UK App Store listing exposes several prices, but does not make all billing durations and plan mappings clear. Hinge's official help centre should be checked for feature details. |

Sources: UK App Store listings for Tinder (https://apps.apple.com/gb/app/tinder-dating-app-meet-date/id547702041), Bumble (https://apps.apple.com/gb/app/bumble-dating-app-meet-people/id930441707), Hinge (https://apps.apple.com/gb/app/hinge-dating-app-match-date/id595287172); official Tinder subscription help (https://www.help.tinder.com/hc/en-gb/articles/115004487406-Tinder-subscriptions); official Bumble subscription guide (https://support.bumble.com/hc/en-us/articles/32668790872733-Understanding-Bumble-s-paid-features-and-subscription-plans).

**Interpretation:** Seeing incoming likes is an established paid benefit among large competitors. It is a plausible value proposition for Pinchd Premium, but competitor pricing alone cannot tell us what Pinchd users will pay or what price would maximise long-term value. Do not compare a weekly price to a monthly price, or an App Store listing example to a standard web price, without confirming duration and eligibility.

### Willingness to pay: evidence and limits

- A 2026 DatingScout UK overview reports 18.9% preferring premium features, 34.3% preferring affordable options and 46.9% preferring free-with-ads. Its public page does not provide enough methodological detail here to treat these numbers as a forecast for Pinchd; use as directional context only. Source: https://www.datingscout.co.uk/online-dating-statistics
- A 2026 Minds UK Gen Z study about willingness to pay for advanced safety features uses a simulated persona panel rather than a conventional recruited representative sample. Do not treat its willingness-to-pay estimates as population estimates. Its directional lesson is consistent with Pinchd's principles: users may expect safety to be a baseline responsibility, so safety should not be a premium paywall. Source: https://getminds.ai/studies/dating-apps-safety-feature-valuation-uk-2026
- General freemium research identifies perceived functional, emotional, social and economic value as relevant to upgrade and retention intentions, but findings from another digital service cannot directly predict dating-app conversion. Source: https://onlinelibrary.wiley.com/doi/full/10.1111/isj.12262

**Conclusion:** No sufficiently strong public evidence located in this pass can determine Pinchd's actual willingness to pay, conversion rate, churn or price elasticity. Validate through direct qualitative research, a transparent non-binding price-sensitivity study, and later properly designed live offers once legal and product readiness permit. Recruit across relationship goals, ages, genders, regions, income constraints and free/paid dating-app experience; avoid treating one segment as representative of all users.

### Sustainable unit economics: model gross-to-net revenue

A useful first model should be parameterised rather than built around an assumed price. For each candidate price and billing period, calculate:

1. Consumer price inclusive of applicable tax.
2. Less VAT or other applicable indirect taxes, using accountant-reviewed treatment.
3. Less store/payment fees under the actual distribution and billing terms.
4. Less refunds, chargebacks, fraud and failed-payment leakage.
5. Less variable costs attributable to paying users (support, moderation, media, verification and infrastructure where applicable).
6. Resulting contribution per paying subscriber and contribution per eligible active user.
7. Compare against acquisition cost, retention, renewal and payback; model low/base/high scenarios and uncertainty.

Illustrative store-fee caution (verify again at implementation): Google Play's current UK/EEA/US service-fee documentation describes new subscription fee terms from 30 June 2026, including a 10% service fee plus a 5% billing fee for qualifying transactions under the stated conditions. Eligibility, install cohort and programme terms matter. Source: https://support.google.com/googleplay/android-developer/answer/112622?hl=en-GB. Apple commission terms vary by programme, storefront and subscription tenure; do not hardcode a universal rate. Sources: https://developer.apple.com/help/app-store-connect/distributing-apps-in-the-european-union/commissions-fees-and-taxes and https://www.apple.com/uk/newsroom/2020/11/apple-announces-app-store-small-business-program/.

The model must separately account for the cost of maintaining a healthy free user base: moderation and safety operations, user support, discovery quality, and infrastructure. Do not optimise conversion by deliberately degrading free matching, safety or consent. A product that extracts subscription revenue but loses local network density or trust may be commercially unsustainable.

### Research-first decision and next steps

**Decision recorded:** Research competitor pricing, user willingness to pay and sustainable unit economics before setting Pinchd Premium pricing. No price has been approved.

Recommended sequence:
1. Confirm competitor feature bundles and billing periods directly from official help pages and storefront listings immediately before benchmarking.
2. Interview prospective users about the problem solved by seeing all incoming likes and other candidate premium benefits; separate interest from actual willingness to pay.
3. Test non-binding price concepts and package comprehension, including whether the five-like rolling queue feels fair and useful.
4. Build gross-to-net scenarios with explicit assumptions for tax, storefront fees, conversion, churn, refunds, support and variable service cost.
5. Define launch decision thresholds and guardrails before any live price test: contribution margin per eligible active user, renewal/retention, refund and complaint rates, trust, free-user outcomes and safety.
6. Choose and test a price only after the evidence and financial model are reviewed. Keep one optional Premium tier as the baseline hypothesis unless research supports a clearer alternative.

Mobile app design remains paused until research and requirements work is complete.


## 12. User value research: which paid benefits matter? (21 October 2026)

### What external evidence supports

**1. Some users do pay for dating-app features, but payment is not universal.** Pew Research Center's 2023 US survey found 35% of adults who had used online dating had ever paid to use a site or app. The share was 22% among users under 30 and 41% among users aged 30 and older. This is historical US evidence, not a UK forecast or a measure of willingness to pay for Pinchd specifically. Source: https://www.pewresearch.org/internet/2023/02/02/the-who-where-and-why-of-online-dating-in-the-u-s/

**2. Seeing incoming likes is a familiar paid benefit.** Current product help and market reviews describe full incoming-like visibility as a common premium feature. That makes it an understandable value proposition, but not proof that users consider any particular price fair or that the feature alone will convert them. Competitor pricing and bundles are documented in section 11.

**3. User discussions repeatedly describe price/value disappointment.** Recent Reddit discussions include complaints that subscriptions cost too much, paid boosts do not lead to more matches, and a paid queue can contain few suitable profiles. These are self-selected anecdotes, not representative survey results; use them to generate interview questions, not estimate market percentages. Examples: https://www.reddit.com/r/DatingApps/comments/1vaj3f5/ranting_about_dating_app_premium_prices/ and https://www.reddit.com/r/OnlineDating/comments/1i9nmny/

**4. Trust can be damaged when paid features appear to withhold potential connections.** App Store reviews and public user discussions include complaints about profiles being gated behind paid discovery features or upgrade prompts implying urgency. Reviews are anecdotal and cannot establish how common this experience is, but the risk is strategically relevant to Pinchd's positioning. Example: https://apps.apple.com/gb/app/hinge-dating-app-match-date/id595287172?see-all=reviews

**5. Safety should remain a baseline rather than a premium hook.** A 2026 UK Gen Z safety-feature study reports low simulated willingness to pay for advanced safety add-ons, but uses a target-audience simulation rather than a conventional representative survey. Do not generalise its numerical result. Its methodology limitation aside, it does not justify paywalling safety. Source: https://getminds.ai/studies/dating-apps-safety-feature-valuation-uk-2026

### Implications for Pinchd — hypotheses, not facts

- **Full incoming-like visibility:** likely the clearest initial Premium proposition because users understand the benefit immediately. Validate whether five visible likes in a rolling queue feels useful and fair before treating it as the final free allowance.
- **More profile/discovery organisation:** potentially useful to people with a busy queue or limited time, but likely less compelling before Pinchd has sufficient local user density.
- **Profile and message coaching:** may help users who feel unsure what to write; it must remain optional, editable, and user-reviewed, and must not auto-send.
- **Detailed compatibility explanations:** could support Pinchd's trust and explainability promise if the reasons are grounded in declared data. Do not sell an inflated or opaque compatibility score.
- **Travel mode and extra discovery controls:** lower-priority until there is evidence that users need them and they can be delivered without bypassing hard preferences or misrepresenting location.
- **Paid boosts / hidden ranking advantage:** exclude from the proposed package. Users may judge them as untrustworthy when outcomes are uncertain, and they conflict with Pinchd's fairness principles.

### Proposed research method

**Phase A — discovery interviews (qualitative, not representative).**
Recruit a varied set of UK dating-app users, including people who pay and do not pay, and people with different relationship goals and levels of dating-app fatigue. Ask about their last real subscription decision rather than asking only what they hypothetically might buy:
- What prompted the last upgrade, cancellation or refusal to pay?
- What did they expect to improve, and what actually improved?
- How useful was seeing incoming likes, and how relevant were those likes to their stated preferences?
- Did they feel the free product was useful in its own right or deliberately impaired?
- Which benefits would they miss if removed? Which felt like clutter?
- What would make them trust a subscription offer and cancellation process?

**Phase B — concept comparison.**
Show balanced, non-binding concept cards with equal visual prominence and plain language:
A. Full incoming-like queue unblurred.
B. Full incoming-like queue plus profile customisation.
C. Full incoming-like queue plus optional profile/message coaching.
D. Full incoming-like queue plus detailed declared-data compatibility explanations.
E. Free-only / no subscription.
Ask participants to explain what each package does, rank usefulness, identify concerns and select a plausible price range. Randomise presentation order where practical to reduce order effects. Do not use fake scarcity or imply the offers are live if they are not.

**Phase C — price sensitivity.**
After the benefit is understood, use a recognised price-sensitivity method or structured discrete-choice exercise with clearly stated billing periods and currency. Include a genuine no-purchase option. Distinguish stated willingness to pay from actual purchase behaviour; report ranges and uncertainty rather than a single precise number. Ensure any eventual live test uses transparent offers and legally reviewed disclosures.

**Phase D — live validation after launch readiness.**
Pre-register the test and use account/cohort-level assignment. Primary metric should be contribution margin per eligible active user over an appropriate period, not raw conversion alone. Guardrails should include renewal and churn, refunds, complaints, free-user matching quality, queue relevance, user trust, and safety. Never withhold essential safety, messaging, privacy, consent or hard-preference integrity for experimental purposes.

### Research decision
Continue with **Free + one optional Premium tier** as the baseline hypothesis. Treat full incoming-like visibility as the lead candidate, and profile customisation, coaching and compatibility explanations as secondary candidates to compare. Do not set price or lock the exact bundle until qualitative and quantitative validation plus gross-to-net modelling are complete.

### Evidence limitations
Public evidence located in this research pass does not establish UK-specific willingness to pay for Pinchd's exact feature bundle, expected conversion, optimal free-like count, price elasticity, retention, or causal effects of premium on match/date outcomes. Reddit and app reviews are self-selected; US survey results are not UK estimates; simulated audience studies are not equivalent to representative probability samples. These sources inform hypotheses only.


## 13. Public online discussion scan — no participant recruitment (21 October 2026)

### Scope and method
This pass uses public Reddit threads, a UK newspaper pricing comparison and a UK paid-dating usage survey listing. Reddit posts are self-selected anecdotes and the search results are not a representative sample; they are used to identify recurring objections, unmet needs and testable hypotheses. No claims about prevalence should be made from comment counts or votes.

### Themes found in public discussion

**A. Price feels unjustified when outcomes do not improve.** In a July/August 2026 r/DatingApps discussion, users complain that subscriptions can cost $15–$20+ per month or even per week, and that boosts did not lead to matches. Others disagree, saying value depends on location, profile and the alternative app. Source: https://www.reddit.com/r/DatingApps/comments/1vaj3f5/ranting_about_dating_app_premium_prices/ (published July 2026; viewed 21 October 2026).

**B. Users ask whether Premium creates more matches, not just more controls.** A recent thread asking whether dating-app subscriptions are necessary shows users seeking first-hand evidence before paying; responses include recommendations to use free options and to pay only when there is a sufficiently large queue and time to review it. Source: https://www.reddit.com/r/DatingApps/comments/1wvqb67/is_it_necessary_to_pay_a_subscription_on/ (October 2026).

**C. Value depends on local supply and user context.** In Hinge and cross-app threads, commenters say extra likes or filters can be useful for some people, but are less valuable when the local pool is small, few relevant likes arrive, or the user already gets enough matches. One thread also describes the incoming-like queue becoming a chore rather than a benefit when the queue is very large. Sources: https://www.reddit.com/r/hingeapp/comments/p03xl1/ and https://www.reddit.com/r/DatingApps/comments/1e93ln4/ (2021 and 2024; older evidence, useful as qualitative themes only).

**D. Users distrust fake or irrelevant likes.** A 2026 r/OnlineDating discussion about paid memberships raises concerns about receiving likes before a profile is complete and about paid queues that do not contain profiles matching the user's preferences. Source: https://www.reddit.com/r/OnlineDating/comments/1sp8ftu/ (April 2026). This directly supports Pinchd's requirement that eligibility checks run before any incoming-like queue or paywall presentation.

**E. Price transparency and subscription duration matter.** The Guardian's UK consumer explainer (9 November 2024) documents large price differences by dating service and subscription duration, and reports that some providers promote a lower effective monthly price for longer commitments. It is not a current 2026 price list, but illustrates why comparisons must distinguish weekly, monthly and multi-month total costs. Source: https://www.theguardian.com/lifeandstyle/2024/nov/09/the-price-of-love-how-much-does-dating-cost-pay-bill-apps.

### Quantitative context with caveats
Statista's UK fee-based online dating usage page reports a September 2025 survey of 6,177 UK respondents aged 18–64, asking whether they had paid for online dating in the past 12 months. The accessible page does not expose a usable share in the result, so this scan does not quote a percentage. Source: https://www.statista.com/forecasts/1466296/fee-based-online-dating-usage-in-the-uk/.

### Implications for Pinchd
1. Sell a clearly described convenience benefit, not a promise of more matches or dates.
2. The full incoming-like queue should be useful only because it contains eligible, real interest. Apply hard preferences, blocks, account restrictions and safety exclusions before display and before subscription entitlement checks.
3. The five-like rolling queue needs validation: five may be enough for some users and too few or too many for others. Monitor how many free users actually exhaust the visible queue, whether newly revealed likes are relevant, pass/reciprocate rates, complaints and free-user match outcomes.
4. Avoid subscription urgency, misleading countdowns, “someone is waiting” pressure, fabricated like counts and promises that payment improves compatibility.
5. Offer clear billing-period and total-cost comparisons. If multiple durations are offered later, show both the total charged and the equivalent monthly cost without obscuring commitment length.
6. Keep one Premium tier as the baseline, but do not assume seeing likes alone is enough. Compare that benefit against coaching, profile customisation and explainable compatibility with balanced concept descriptions.
7. For commercial viability, segment hypotheses by local supply/density and queue size, but do not use sensitive traits to target prices or apply opaque personalised pricing.

### Updated evidence-based conclusion
The online scan supports incoming-like visibility as an understandable premium benefit, but shows strong reasons to avoid promising outcomes: users question whether subscriptions increase matches, dislike paying for irrelevant queues, and see value as dependent on local supply and personal context. Public discussion cannot establish the optimal price, feature bundle, five-like allowance or expected conversion. Continue to research and model these rather than setting a price now.


## 14. Blurred incoming-like model: trust and fairness scan (21 October 2026)

### Scope
Reviewed public Reddit discussions about Hinge's blurred incoming-like queue, free vs paid visibility, and paid features that users perceive as manipulating urgency or outcomes. These are self-selected reports and opinions, not representative evidence. Individual allegations about platforms hiding, fabricating, or reordering likes are unverified and must not be repeated as established fact.

### Public discussion themes

**1. The blurred queue is understood, but can create confusion about what is free.**
In a September 2026 r/hingeapp thread, users distinguish between being able to review likes one at a time for free and paying to see the full unblurred queue at once. Some find the incremental flow acceptable; others dislike not being able to compare multiple profiles. Source: https://www.reddit.com/r/hingeapp/comments/1wmmy94/is_everything_locked_behind_paywalls/ (21–22 September 2026).

**2. Users can interpret blurring as withholding a potential connection.**
A February 2026 r/OnlineDating discussion describes the blurred-photo upsell as frustrating and asks whether connection is being subordinated to monetisation. A June 2024 r/dating post similarly describes disappointment on seeing a blurred like and a subscription prompt. These posts reflect individual perceptions, not proof of intentional manipulation. Sources: https://www.reddit.com/r/OnlineDating/comments/1r29cdx/why_are_dating_apps_designed_this_way/ and https://www.reddit.com/r/dating/comments/1dn5rbg.

**3. Claims of fake or strategically withheld likes damage trust, even when unverified.**
Users in a March 2025 r/DatingApps post and October 2026 r/HingeStories/r/SwipeHelper posts speculate that apps manipulate the timing or quality of likes to induce payment. The posts do not establish that the allegations are true. The product lesson is reputational: Pinchd should be able to explain that displayed counts represent real, eligible incoming interest and should not use fake likes, fake notifications, hidden queues of ineligible profiles, or fabricated urgency. Sources: https://www.reddit.com/r/DatingApps/comments/1j8lrrk, https://www.reddit.com/r/HingeStories/comments/1wtjx3a/hinge_bursts/, https://www.reddit.com/r/SwipeHelper/comments/1wxu3ae/so_its_basically_an_open_secret_that_hinge_and/.

**4. Users worry about likes expiring or being used as pressure.**
An April 2026 r/OnlineDating discussion criticises expiring likes and interprets the change as a potential subscription pressure tactic. This is opinion about a product change, not evidence of intent. Avoid adding expiry timers or loss-aversion mechanics to the incoming-like queue unless there is a clear user-benefit case; default to no artificial expiry. Source: https://www.reddit.com/r/OnlineDating/comments/1socwog/hinge_update_expiring_likes_signals/.

**5. Queue relevance and ability to compare profiles matter.**
In the September 2026 discussion, users describe reviewing incoming likes incrementally and others say they would prefer to compare several profiles at once. Another participant says a high volume of likes can make it difficult to compare people or manage conversations. This suggests users differ: a one-at-a-time queue can reduce overload, while a larger unblurred view can help comparison. Source: https://www.reddit.com/r/hingeapp/comments/1wmmy94/is_everything_locked_behind_paywalls/.

### Design and policy implications (requirements, not UI design)

- The free rolling queue should show five eligible profiles unblurred simultaneously; users can make an informed choice to reciprocate or pass. The next eligible like then enters the visible set.
- Blurred items must not be misleading: do not reveal identity-sensitive information through thumbnails, metadata, names, comments or notifications beyond the explicitly approved product policy. Do not display an exact count if the count is stale, incomplete or includes ineligible accounts.
- The same hard eligibility gates apply to free and paid users. Premium reveals more eligible likes; it does not change who is eligible or the order for romantic ranking.
- Do not expire incoming likes just to encourage a subscription. Do not fabricate counts, activity, “someone is waiting” prompts, countdowns or claims that a like will disappear.
- Do not change queue contents or apparent quality in response to subscription purchase, subscription expiry, user vulnerability, profile desirability, or predicted willingness to pay. The queue should reflect real eligible interest and documented ordering rules.
- Explain the distinction between a like and a match: an incoming like is interest, not a match or a promise of compatibility; a match still requires reciprocal interest.
- Keep pass and like-back controls equally clear. Passing must not trigger guilt copy or repeated retention prompts. Passing should reveal the next eligible like in the rolling queue.
- Monitor whether free users exhaust the five visible slots, time to action, eligible-like relevance, pass/reciprocate rates, complaints about misleading counts, paid conversion, refunds, and free-user match outcomes. Segment analyses only where lawful and privacy-safe.
- Provide clear subscription terms and a straightforward cancellation route; never suggest payment guarantees dates, attraction or better people.

### Decision
**Retain the rolling-queue hypothesis, with explicit anti-manipulation safeguards.** The public scan does not establish that five is the optimal free allowance or that a blur model increases trust or sustainable revenue. It shows a real trust risk when blurring is perceived as withholding connection or when counts/ordering are suspected of being manipulated. Validate the queue rules and five-like allowance against transparency, relevance, free-user outcomes and complaints before locking them for launch.

### Evidence quality note
Reddit discussions are useful for discovering language, frustrations and edge cases. They cannot establish prevalence, causality, a competitor's internal intent, or Pinchd's optimal pricing. Allegations about fake/withheld likes are recorded only as perceptions and reputational risks, not facts.


## 15. Subscription duration, renewals and cancellation: public research scan (21 October 2026)

### Scope and evidence quality
This scan compares publicly listed UK dating-app subscription offers, public Reddit discussions about pricing/renewals, and current UK government guidance. App Store in-app purchase menus show multiple prices but do not always identify each price's duration or eligibility; do not treat each listed price as a confirmed monthly plan. Reddit reports are individual experiences and do not establish how often a problem occurs.

### Competitor patterns observed

- **Tinder:** The UK App Store listing shows multiple Tinder Gold price points (£7.49, £11.99, £13.99, £14.99), plus Tinder Plus at £3.99 and paid boosts. It states subscriptions automatically renew unless auto-renew is turned off, and that users manage cancellation through Apple account settings. The listing does not reliably map every displayed price to a billing period. Source: https://apps.apple.com/gb/app/tinder-dating-app-meet-date/id547702041.
- **Hinge:** The UK App Store lists Hinge+ at £12.99 and £14.99, as well as other Hinge subscription prices. It describes automatic renewal and account-settings cancellation. The public listing does not clearly map every price to a specific billing period or offer. Source: https://apps.apple.com/gb/app/hinge-dating-app-match-date/id595287172.
- **Bumble:** The UK App Store lists Boost at £7.49, £8.99 and £10.99; Premium at £13.99, £14.99 and £16.99; and a 7-day Premium offer at £9.99 (another listed 7-day price is £24.99). The difference between offers illustrates why duration and package details must be confirmed before comparing. Source: https://apps.apple.com/gb/app/bumble-dating-app-meet-people/id930441707.

**Interpretation:** Competitors offer multiple durations and price points, but the public storefront snapshot is not sufficient to infer an optimal duration for Pinchd. Weekly access may suit a short, time-limited need but risks feeling expensive per week; longer commitments can lower the effective monthly cost while asking users to pay more upfront and commit longer.

### Public discussion themes

**1. Long-term discounts can obscure the actual commitment.** A Reddit explanation of Hinge pricing describes a “starting at” monthly equivalent based on a six-month upfront payment versus a higher one-month price. Commenters disagree about whether this is normal pricing or deceptive framing; one concern is that users may not notice the total upfront charge. This is a 2023 US discussion, so it is a qualitative example, not current UK pricing evidence. Source: https://www.reddit.com/r/hingeapp/comments/11uztmb/how_hinge_premium_pricing_works/.

**2. Users report confusion when cancellation routes differ by payment channel.** A January 2026 Bumble Reddit thread describes a user unable to find cancellation because the purchase channel appeared to differ from what they expected; replies suggest the user may have bought through Apple. This is an individual unverified support experience, but highlights the need to identify the billing provider clearly and direct users to the correct cancellation route. Source: https://www.reddit.com/r/Bumble/comments/1q44wb4/cant_cancel_my_subscription/.

**3. Users may view longer contracts as risky for dating.** The Hinge pricing thread includes the view that paying upfront for several months is unattractive because a user may meet someone quickly, while other commenters say longer terms may be worth it if they improve the experience. This suggests flexibility has value in dating, where subscription need may end suddenly.

### Current UK regulatory note

GOV.UK published official guidance on subscription contracts on 19 October 2026 explaining the regime under the Digital Markets, Competition and Consumers Act 2024 and the Consumer Subscription Contracts (Cancellation and Information) Regulations 2026. Source: https://www.gov.uk/government/publications/subscription-contracts.

The government also describes requirements intended to reduce subscription traps, including clearer information, straightforward cancellation, reminders before certain renewals or trial conversions, and new 14-day cooling-off rights in specified circumstances. Sources: https://www.gov.uk/government/news/consumers-to-save-around-400-million-every-year-from-government-crackdown-on-costly-subscription-traps and https://www.gov.uk/government/consultations/consultation-on-the-implementation-of-the-new-subscription-contracts-regime/outcome/government-response-to-the-consultation-on-the-implementation-of-the-new-subscription-contracts-regime-web-accessible-version.

**Launch gate:** Legal counsel must confirm commencement, scope and exact obligations for Pinchd's specific product, platform and payment route. Do not treat this research summary as a complete compliance checklist.

### Recommended Pinchd billing hypothesis

- Start with a **monthly auto-renewing subscription** as the primary option to minimise upfront commitment. Make renewal, price, billing period, cancellation and access end date prominent before purchase.
- Consider an optional **three-month plan** only if research shows meaningful demand and the total upfront charge is clearly disclosed alongside the equivalent monthly cost. Do not default to six- or twelve-month commitments before retention and user trust are understood.
- Do not launch weekly billing by default. If later offered for a genuine short-term use case, show its total cost and renewal cadence clearly; do not use it as a high-pressure entry offer.
- No free trial or discounted trial should convert without explicit, clear consent to the future charge and an understandable reminder/cancellation path. Trial terms and legal duties must be reviewed.
- Cancellation should be easy to find and as straightforward as subscribing. State whether access continues to the paid-through date and whether cancellation stops only future renewal. Never add guilt, fake scarcity or repeated obstructive retention screens.
- Show the actual total payable for the selected term. If showing an “equivalent monthly” figure, show the full upfront charge and commitment duration at least as prominently.
- If payment is made through Apple or Google, explain that cancellation may need to occur through the relevant store. For direct web billing, provide a direct account cancellation route and a confirmation record.
- Send accurate receipts, renewal reminders where required, and cancellation confirmations; never make notifications expose sensitive dating activity.

### Economic trade-offs to model

For each term, model upfront cash receipts separately from earned revenue and retention. Longer terms may reduce billing churn and payment processing frequency, but can increase refund/support complexity and user regret; a monthly plan may lower commitment friction but expose the business to more frequent churn. Weekly plans may yield higher short-term revenue per week but could harm trust if they appear expensive or exploit urgency. These are hypotheses to model, not observed Pinchd outcomes.

### Decision status
**Provisional recommendation:** monthly-first billing with transparent renewal and simple cancellation; consider three-month upfront pricing only after validation. Do not set the subscription price or introduce multiple Premium tiers based on competitor listings alone.

### Evidence limitations
The current UK App Store pages list several in-app purchase prices but do not reliably expose duration, account eligibility, promotion rules or billing-channel differences for every price. Reddit discussions are self-selected and often non-UK. Current legal requirements need exact legal review before launch. No public source located here establishes the optimal Pinchd billing cadence, conversion, renewal rate or willingness to pay.


## 16. Free-trial decision — confirmed (21 October 2026)

**Decision:** Pinchd will not offer a free trial for Premium at launch.

### Requirements
- Do not use an automatic free-trial-to-paid conversion flow at launch.
- Present the Premium price, billing period, renewal terms and cancellation route clearly before the user subscribes.
- Keep the free tier useful and honest; do not degrade core matching, safety, mutual messaging, privacy controls or hard-preference integrity to pressure purchase.
- Users can choose to subscribe when they understand the offer; do not imply a trial is required to experience the app's core value.
- Revisit this decision only if later evidence establishes a clear user benefit and a compliant, transparent trial design. Any reconsideration requires an explicit product decision; it is not a default roadmap item.

### Rationale
The current strategy favours a clear, low-confusion subscription model. A trial introduces extra terms, renewal expectations, reminder/cancellation obligations and the risk of unexpected charges. No evidence gathered so far establishes that a free trial is necessary to validate demand for Pinchd Premium. Research should first test the value of the paid feature bundle and monthly billing without relying on trial conversion.

**Status:** Confirmed product decision; launch pricing and final Premium bundle remain open.


## 17. Cancellation and retention research scan (21 October 2026)

### Evidence reviewed
This scan combines public Reddit discussions, UK consumer research, and industry benchmark reporting. Reddit posts are anecdotal and self-selected; they are evidence of possible user perceptions, not representative estimates or proof of platform intent.

### Why users say they cancel or decline to pay

**1. The paid feature did not improve perceived outcomes.** In a March 2026 r/OnlineDating discussion, users asking whether Hinge+ was worth paying for received mixed responses, including advice that paying does not guarantee matches and that profile quality or real-world dating may matter more. In an October 2026 r/hingeapp thread, a user asked whether Hinge+ / HingeX and boosts produce better results, illustrating uncertainty about the value of the paid bundle. Sources: https://www.reddit.com/r/OnlineDating/comments/1rk1pho/is_paying_for_hinge_really_worth_it/ and https://www.reddit.com/r/hingeapp/comments/1wv7i5z/anyone_see_real_results_from_paying_for_hinge/.

**2. The visible likes did not justify the cost.** A September 2026 r/Bumble post says the author cancelled after paying to inspect a set of incoming likes and deciding the result was a reason not to continue. This is an individual anecdote, but it illustrates the risk that unblurred access disappoints if queue relevance is poor. Source: https://www.reddit.com/r/Bumble/comments/1wk5evs/canceled_subscription_today/.

**3. Users suspect the app changes visibility around subscription status.** Reddit threads describe users seeing a burst of likes after cancellation or suspecting profile visibility changed after unsubscribing. The reports are not verified; comments also suggest alternative explanations such as filters changing when paid access ends. Record these as trust perceptions, not facts about competitor algorithms. Sources: https://www.reddit.com/r/Bumble/comments/1rpy9v8/cancelled_subscription_received_100_new_likes/ and https://www.reddit.com/r/Bumble/comments/1pwcnfg/my_subscription_lapses_and_all_of_a_sudden_i_have/.

**4. Local relevance and preference integrity influence whether users stay.** In a February 2026 Bumble thread, a user said they were leaving the paid plan after repeated frustration with out-of-area or travel-mode profiles and then described a more relevant local experience elsewhere. This is one user's account, but directly reinforces Pinchd's hard distance boundary and transparent travel-mode principles. Source: https://www.reddit.com/r/Bumble/comments/1rc0yml/premium_expires/.

**5. Unwanted renewal and cancellation friction can deter sign-up.** Citizens Advice reported on 8 March 2024 that 26% of UK adults had accidentally taken out a subscription in the previous 12 months; among those affected, 40% cited automatic renewal without their knowledge and 39% cited forgetting to cancel a free trial. It also reported that 54% of people had decided not to sign up for a subscription because they were concerned cancellation would be difficult. These figures concern subscriptions generally, not dating apps specifically, and are dated 2024. Source: https://www.citizensadvice.org.uk/about-us/media-centre/press-releases/consumers-spend-688-million-on-unused-subscriptions-in-the-last-year/.

**6. Industry retention data suggests dating-app subscriptions are volatile.** Business of Apps' 2025 dating-app benchmarks report says that among 2024 cohorts, over 95% of monthly subscribers were inactive 12 months later, and over three-quarters of annual subscribers were inactive at that point. Its definition is inactivity, not necessarily formal cancellation, and the figures are category-level estimates rather than a Pinchd forecast. Source: https://www.businessofapps.com/data/dating-app-benchmarks/.

### Retention lessons for Pinchd

- **Sell control and convenience, not romantic outcomes.** The user should understand exactly what Premium unlocks: the full eligible incoming-like queue, plus any additional benefits that later prove valuable.
- **Queue relevance is part of the value.** Hard preferences, blocks, account restrictions, unmatched-pair exclusions and safety rules must apply before incoming likes enter the queue. Payment cannot reveal ineligible profiles.
- **Do not manipulate visibility around renewal or cancellation.** Do not increase apparent likes, change recommendation quality, manufacture urgency, or send "someone is waiting" prompts because a subscription is about to expire or a user has cancelled.
- **Let users leave cleanly.** Cancellation should be straightforward, discoverable and confirmed. Explain when paid access ends and that deleting the app/account may not cancel a store-billed subscription where applicable.
- **Respect successful dating and pauses.** If a user pauses dating, meets someone, or no longer needs the service, make it easy to stop renewal or pause discovery. Do not shame or guilt them into staying.
- **Make the free tier credible.** Core matching, mutual messaging, safety and privacy must not degrade when Premium expires. The user can remain on Free without punitive discovery changes.
- **Measure value beyond short-term conversion.** Track cohort renewal, churn, refunds, chargebacks, complaints, cancellation completion, queue relevance, eligible matches and free-user outcomes. Do not optimise only for initial purchase.
- **Offer optional feedback at cancellation.** Ask one neutral, skippable question (for example, price, not enough relevant likes, feature not useful, taking a break, met someone, technical issue, other). Do not require a reason or obstruct cancellation. Treat answers as directional self-reports.
- **Support billing-channel clarity.** Tell users whether Apple, Google or Pinchd's direct billing handles payment and cancellation, and link them to the correct route. Keep subscription entitlements server-authoritative and reconcile receipts.

### Product decisions already aligned with this evidence
- No free trial at launch.
- One optional Premium tier as the initial hypothesis.
- Monthly-first billing as the provisional candidate; price remains open.
- No paid ranking advantage or guaranteed romantic outcome.
- No fake likes, fabricated activity, misleading counts or subscription-linked manipulation of queue quality.
- Easy cancellation and transparent renewal terms are requirements, subject to final legal/platform review.

### Evidence limitations
Public discussions reveal failure modes and user language but cannot estimate prevalence. Industry retention statistics depend on dataset and definitions; inactivity is not identical to cancellation. UK consumer subscription figures cover many categories, not only dating apps. None of these sources establishes Pinchd's expected churn or lifetime value.

### Research conclusion
The strongest retention hypothesis is that users stay when Premium repeatedly delivers a clear, relevant convenience benefit and trust that the free experience is not being manipulated. For Pinchd, this means genuine eligible likes, preference integrity, transparent subscription terms, no false promises, and easy exit. Validate this through post-launch cohort data rather than adopting competitor retention assumptions.


## 18. Premium value, local density and unit-economics implications (21 October 2026)

### External evidence reviewed
**Dating-app engagement and acquisition remain difficult.** Adjust's 29 January 2026 industry review reports that global dating-app installs fell 4% and sessions fell 7% in 2025. Reported day-30 retention improved from 5% in 2024 to 6% in 2025, while average session length fell from 13.21 to 11.49 minutes. Adjust also reports cost per install rising from $1.46 to $2.76 and paid/organic acquisition ratio rising from 1.65 to 2.14. These are aggregated mobile-marketing benchmarks, not UK-only data and not forecasts for Pinchd. Source: https://www.adjust.com/blog/state-of-dating-apps/

**London is a plausible candidate for a focused launch, but market size is not product-market fit.** Global Dating Insights, reporting on 4 August 2026 and citing industry estimates, says YouGov research found 9% of London adults were actively using a dating app in mid-2025 and 31% had used one previously. The article also reports Ofcom estimated 4.9 million UK adults used an online dating service during 2024. These figures describe the overall category, not Pinchd's reachable audience, age-eligible pool, relationship-goal mix or likely acquisition rate. Treat the cited market-value forecasts as estimates rather than audited market totals. Source: https://www.globaldatinginsights.com/news/london-dating-apps-may-create-a-750m-569m-subscription-business/

**Subscription retention is weak in published category benchmarks.** Business of Apps' 2025 dating-app benchmarks reports more than 95% of monthly subscribers inactive after 12 months, with over three-quarters of annual subscribers inactive at that point. "Inactive" should not be treated as identical to cancellation, and methodology may differ from Pinchd's eventual analytics. Source: https://www.businessofapps.com/data/dating-app-benchmarks/

### Why density changes Premium value
Pinchd's proposed Premium lead benefit is seeing the full eligible incoming-like queue unblurred. The feature is only valuable when a user has enough **real, eligible, relevant incoming likes** to make review worthwhile. Therefore:
- A small or newly launched local pool can mean too few eligible likes to justify a subscription, even if the feature works correctly.
- High raw user counts do not prove adequate density for a particular person's age range, distance boundary, relationship goals and discovery preferences.
- A user can have a legitimate quiet period; do not imply that a low like count means a defect or that paying will fix it.
- Never widen age, distance, relationship-goal or other hard preferences to improve the apparent queue or create monetisation opportunities.
- Do not sell a guaranteed number of likes, matches, replies or dates.
- Measure eligible incoming-like availability and reciprocal connection outcomes across launch areas and relevant preference cohorts, while protecting privacy and avoiding small-group disclosure.
- If supply is thin, focus on acquisition and healthy density, not on selling visibility into an empty or irrelevant queue.

### Premium value hypothesis to validate
**Hypothesis:** A single optional Premium tier can earn willingness to pay when it saves time and gives transparent control over reviewing genuine, eligible incoming interest—without withholding core connection or safety.

Candidate bundle, still subject to validation:
1. Full eligible incoming-like queue unblurred (incremental review, not pressure-based expiry).
2. Better organisation of discovery and saved profiles, only if research shows users need it.
3. Optional profile coaching that the user reviews and edits; never auto-publish or auto-send.
4. More detailed compatibility explanations grounded in user-provided information, with no sensitive inference and no false precision.
5. Truthful travel mode only if implementation can keep distance and location expectations clear.

Do not add benefits merely to make a bundle appear larger. Validate each candidate against real user problems, accessibility, privacy, engineering cost and support burden. Do not introduce a second paid tier until the first bundle has evidence of demand and sustainable economics.

### Lightweight unit-economics model: required inputs, not invented estimates
Do not set a price from competitor screenshots alone. Build a model with explicit assumptions and sensitivity ranges.

Core monthly calculations:
- **Gross subscription revenue** = active paid subscribers × displayed monthly price (or monthly-equivalent revenue recognised under the chosen plan).
- **Net subscription revenue** = gross revenue − store/payment fees − refunds/chargebacks − applicable taxes and other variable billing costs.
- **Contribution margin** = net subscription revenue − variable service, moderation, customer support and other per-subscriber costs.
- **Paid subscriber acquisition cost (CAC)** = attributable acquisition spend ÷ new paying subscribers attributed to that spend.
- **Subscriber lifetime value (LTV)** = expected monthly contribution × expected paid lifetime, adjusted for discounting and variable retention where material.
- **Payback period** = CAC ÷ monthly contribution per acquired subscriber, using a consistent cohort definition.
- **Break-even active payers** = monthly fixed operating costs ÷ contribution per active paying subscriber, if contribution is positive and the simplified assumptions hold.

Model low/base/high cases for:
- price and store/payment fee;
- free-to-paid conversion;
- monthly renewal/churn;
- refunds, chargebacks and billing support;
- acquisition mix (organic, referral, partnership and paid);
- service, moderation, age-assurance and support costs;
- local launch pool and eligible incoming-like availability;
- proportion of Premium value that users can experience without paying.

The model must separate **app installs**, registered users, verified/age-eligible users, monthly active users, eligible discovery supply, incoming-like volume, reciprocal matches, paid subscribers and retained paid subscribers. Do not assume an install equals an active or eligible dater.

### Decision gates before launch pricing
1. **Demand gate:** User research supports the specific Premium jobs-to-be-done; price sensitivity is tested without implying guaranteed dating outcomes.
2. **Density gate:** A defined launch area has sufficient eligible, active supply for the intended discovery experience; thresholds must be set from product goals and beta data rather than copied from competitors.
3. **Economics gate:** Conservative unit-economics scenarios show a credible route to covering payment, infrastructure, safety, support and acquisition costs. Do not count speculative lifetime value as current cash.
4. **Trust gate:** Users understand what is free, what Premium changes, price, billing interval, renewal and cancellation before purchase.
5. **Safety and fairness gate:** Monetisation does not bypass hard preferences, blocks, consent, age assurance, safety processes or privacy controls.
6. **Measurement gate:** Instrument a pre-registered beta plan and report cohort results; do not optimise only for paywall conversion.

### Research conclusion
A large dating market and strong category subscription revenue do not by themselves establish Pinchd's viable price or Premium demand. The product should prove a dense, trustworthy local experience first. Monetisation should be evaluated against eligible local supply, actual feature usage, cohort retention, net revenue and safety/support costs. Pricing remains open; no numeric price or conversion target is approved by this research.
