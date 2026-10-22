# Pinchd — Premium Pricing and Unit Economics Model
**Date:** 22 October 2026  
**Status:** Planning model; no subscription price or launch budget is approved.  
**Product direction:** One optional Premium tier, monthly-first, no free trial; essential matching, messaging, safety and legally required age assurance remain free.

## 1. Decision summary
Pinchd should test willingness to pay only after the free product promise is clear and there is enough local eligible supply to deliver value. Premium should improve convenience and control, not buy romantic ranking advantage or safety. Do not select a price by copying a competitor's headline price; local prices, promotions, renewal periods, taxes and store terms vary.

Candidate benefits approved for further evaluation:
- Full eligible incoming-like queue (free users see five eligible incoming likes unblurred at a time in a rolling queue).
- Profile customisation.
- Optional, user-reviewed profile coaching.
- Saved/organised discovery lists.
- Additional discovery controls.
- Truthful, user-controlled travel mode with privacy and anti-stalking safeguards.
- Detailed compatibility explanations grounded in user-provided data.

Not permitted: hidden ranking advantage, paid comment boost, guaranteed matches/dates, unsolicited messaging privileges, or paid access to blocking/reporting/consent.

## 2. Pricing decision framework
No exact price is set. Evaluate a small, coherent set of monthly price points in the UK, then consider annual billing only if user research and cancellation/renewal economics justify it. Do not add several tiers at launch unless evidence shows a clear user need.

For each tested price, collect:
- Comprehension: what the member believes Premium includes.
- Stated intent and actual purchase conversion, kept separate.
- Refund, cancellation, complaint and chargeback rates.
- Renewal retention and voluntary cancellation reasons.
- Free-to-paid conversion by acquisition cohort and market.
- Whether the free product still provides a meaningful dating experience.
- Whether value is concentrated in a feature that should remain free for safety, consent or fairness.

Do not use deceptive dark patterns, fake countdowns, false scarcity, shame-based cancellation flows, or subscription prompts that interrupt safety/reporting. Cancellation should be easy to find and complete. No free trial is approved for launch.

## 3. Current platform and tax facts to model
These are sourced facts as of the research date; they are not a final billing decision.

- UK standard VAT is 20% for most goods/services, subject to the correct tax treatment and seller/merchant arrangement. Source: HMRC, https://www.gov.uk/vat-rates
- Apple states its standard App Store digital goods/services commission is generally 30%, with 15% for qualifying subscriptions and certain programmes; eligibility and contract terms matter. Source: Apple Developer Program membership, https://developer.apple.com/programs/whats-included
- Apple's Small Business Programme can offer 15% commission to approved eligible developers under its programme terms. Source: Apple programme terms, https://developer.apple.com/support/downloads/terms/schedules/Schedule-2-and-3-English-UK.pdf
- Google Play's official service-fee page describes a UK/EEA/US rollout effective 30 June 2026 with subscription fees that can comprise a 10% service fee plus a 5% billing fee for qualifying transactions, and other fee structures depending on install status, programme and billing route. Do not reduce this to a universal 15% rule. Source: https://support.google.com/googleplay/android-developer/answer/112622?hl=en-GB
- Store policies and fees can change; verify current developer account terms, billing route, regional rollout and eligibility before implementation. External payment paths and terms need specialist legal/store review.

For a consumer price that includes VAT, a simplified initial model may use net-of-VAT price = consumer price / 1.20, but only if standard VAT applies and the business is responsible for accounting for it in that way. Confirm actual VAT treatment with an accountant.

## 4. Unit economics formulas
Use one model per platform/channel because store fees, VAT handling, refunds and tax reporting can differ.

Let:
- P = consumer price per paid period, including applicable tax.
- V = applicable VAT fraction if included in P.
- F = store/payment fee fraction and any fixed fee.
- R = refund/chargeback allowance as a share of billed value.
- S = variable service cost per paying account per period (support, moderation, verification and infrastructure allocation).
- A = acquisition cost allocated to the cohort.
- M = active paying subscribers in the period.
- C = monthly cancellation/churn rate for the paid cohort.

A simplified contribution per paid period is:
`Contribution ≈ (P / (1 + V)) × (1 - F - R) - S`

This is only a rough sensitivity formula. The true calculation must model whether the store fee applies before/after taxes, fee on refunded transactions, fixed billing fees, currency conversion, platform-specific terms and whether the business or store is merchant of record. Do not treat this expression as an accounting rule.

Useful metrics:
- **Net revenue per subscriber:** recognised subscription revenue less taxes, store/payment fees, refunds and chargebacks.
- **Contribution margin:** net revenue less variable infrastructure, moderation, support, verification and other variable service costs.
- **Gross margin:** define consistently with accounting policy; don't mix it with contribution margin.
- **Paid conversion:** new paying accounts / eligible active free accounts exposed to a clear offer.
- **Paid retention:** paying cohort still active and entitled after N renewal periods / initial paying cohort.
- **Refund rate:** refunded transactions / billed transactions (also report value-weighted rate).
- **Cost per activated eligible user:** acquisition spend / users who complete onboarding, pass required age assurance and are active in the defined window.
- **Payback period:** acquisition cost / monthly contribution, using a cohort model and only where contribution is positive.
- **LTV:** cohort-based discounted contribution over observed/defensible retention, not a made-up multiple of monthly revenue.

Never infer LTV from a short beta or use stated purchase intent as a substitute for conversion.
## 5. Cost inventory

### Fixed and semi-fixed
- Product engineering and maintenance labour.
- Legal counsel, privacy/DPIA work, consumer subscription review and contract review.
- Trust-and-safety operations, training, on-call coverage and incident exercises.
- Customer support tools and staffing.
- Age-assurance vendor minimums, implementation and audit support.
- Security testing, monitoring, backups, incident response and compliance evidence.
- Developer accounts, release management, accounting and insurance.

### Variable or usage-based
- Age-assurance checks, retries, appeals and fraud review.
- Image/video storage, transformations, scanning and delivery.
- Database, API compute, queues, logs, metrics, alerting and backups.
- Push/SMS/email, authentication and customer support contacts.
- Moderation vendor usage and human case time.
- Store/payment fees, taxes, refunds and chargebacks.
- Acquisition, creator/partner payments and referral incentives.
- AI inference if optional coaching ships; include opt-in rate, input/output size, abuse controls and provider data terms.

Keep verification and safety costs visible; don't bury them in a generic cloud line item. Model report volume and high-severity incident spikes, not only average usage.

## 6. Scenario model
Build Low / Base / High cases, with inputs and sources visible. A useful initial model should include:
- Waitlist acquisition by channel and cost.
- Activation and age-assurance completion.
- Monthly active eligible users and retention.
- Share of active users exposed to the Premium offer.
- Conversion by price point and platform.
- Refunds, chargebacks, churn and renewal.
- Store fee route and VAT assumptions.
- Cost per age check and retry rate.
- Moderation/report rate, average handling time and escalation cost.
- Infrastructure cost by active user and media usage.
- Labour, legal, vendor minimums and contingency.
- Cash runway and downside spend cap.

Never assume that higher engagement automatically means better economics. Model a case where eligible supply is weak and paid conversion falls; do not compensate by degrading the free experience.
## 7. Ethical free/Premium boundary
| Capability | Free/Premium baseline | Rationale |
|---|---|---|
| Age assurance required by law/policy | Free | Eligibility/compliance gate |
| Core profile and discovery | Free | Basic service value |
| Hard preference controls | Free | Preference integrity and consent |
| Like/pass and mutual match | Free | Core reciprocal interaction |
| Messaging after mutual match | Free | Core consent-based connection |
| Block, report, safety support and appeal | Free | Safety must not be monetised |
| Account pause and deletion | Free | User control and exit |
| Required account/security controls | Free | Security and lawful access |
| Full incoming-like queue | Candidate Premium benefit | Convenience; free rolling queue still useful |
| Profile customisation | Candidate Premium benefit | Cosmetic convenience only |
| Saved/organised discovery lists | Candidate Premium benefit | Organisation, not hidden ranking |
| Extra discovery controls | Candidate Premium benefit | Must not override another user's hard preferences |
| Travel mode | Candidate Premium benefit | Only with truthful location and anti-stalking safeguards |
| Detailed compatibility explanations | Candidate Premium benefit candidate | Basic explanation and preference integrity must remain truthful; never sell fabricated certainty |
| AI profile coaching | Optional candidate benefit | Opt-in, user-reviewed, accessible non-AI path |
| Hidden rank boosts or guaranteed outcomes | Prohibited | Unfair and misleading |
| Unsolicited messaging or safety-control unlocks | Prohibited | Violates consent/safety boundary |

Premium entitlement is server-authoritative. Store purchase, renewal, refund, cancellation, grace period and expiry must be reconciled with signed/provider notifications and tested for replay, delay and disagreement. A payment failure must never remove safety controls or silently falsify subscription status.

## 8. Experiments and decision gates
1. Test value proposition before testing price: users must understand what the free product does.
2. Use a small set of transparent monthly prices; do not present fake discounts or invented reference prices.
3. Randomise or stage tests carefully; document sample sizes and do not claim statistical certainty from tiny cohorts.
4. Do not make a price test a prerequisite for safety or essential product access.
5. Predefine guardrails: complaints, refunds, chargebacks, cancellation friction, free-user satisfaction, hard-preference violations and support burden.
6. Keep billing UX clear about price, billing period, auto-renewal, cancellation route and what happens on expiry.
7. Test restore purchases, account changes, family/store edge cases as applicable, duplicate webhook, refund/revocation and delayed events.
8. Review consumer-law and platform requirements before selling; upcoming legal commencement dates and applicability require counsel verification.
9. Do not launch paid subscriptions until app-store products, terms, privacy notices, refund/support workflow and server entitlement checks are production-ready.
## 9. Launch economics gates
Before authorising spend or pricing:
- A downside case shows enough runway for legal, vendor, support and safety obligations even if conversion is poor.
- The free service can operate without Premium revenue during the initial validation window.
- Unit cost includes age assurance, moderation and support, not only cloud hosting.
- Premium benefits are testable and delivered; no placeholder benefit is sold.
- Billing state is server-authoritative and audited.
- Cancellation, restore, refund and deletion paths are tested on each platform.
- Store fees and VAT assumptions have been verified for the actual route and entity.
- Contribution margin and cohort retention are measured separately from gross billings.
- A price can be changed without deceptive anchoring or trapping existing subscribers.
- No price or monetisation decision overrides the safety, legal or accessibility gates.

## 10. Open decisions
1. Exact monthly price and whether/when annual billing is appropriate.
2. Which Premium benefits make the final launch scope and which remain later candidates.
3. Platform-specific billing route, commission eligibility, VAT treatment and merchant-of-record responsibilities.
4. Store cancellation/refund and new subscription regulation applicability at the actual launch date.
5. Age-assurance and moderation vendor price quotes, minimum commitments and support SLAs.
6. Whether a full incoming-like queue meaningfully improves user value without harming the free experience.
7. Price-test method, minimum sample size, experiment budget and decision owner.
8. Cohort retention evidence needed before LTV is reported.
9. How refunds, chargebacks and billing disputes affect entitlement and safety review.
10. The cost allocation method for support, moderation, compliance and shared infrastructure.

## 11. Sources and review caveat
- HMRC, VAT rates: https://www.gov.uk/vat-rates
- Apple Developer Program membership/fees: https://developer.apple.com/programs/whats-included
- Apple Schedule 2 and 3 (UK terms PDF): https://developer.apple.com/support/downloads/terms/schedules/Schedule-2-and-3-English-UK.pdf
- Google Play Console Help, Service fees (UK/EEA/US rollout details): https://support.google.com/googleplay/android-developer/answer/112622?hl=en-GB
- Apple, App Review Guidelines: https://developer.apple.com/app-store/review/guidelines/

Recheck source pages and the developer's actual contract before external publication or billing implementation. This model is not financial, tax or legal advice; no exact price or revenue forecast is approved.
