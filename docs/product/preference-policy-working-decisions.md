# Preference policy — working decision log
**Status:** In progress; only explicitly confirmed decisions are marked decided.  
**Last updated:** 21 October 2026  
**Related:** `docs/product/consolidated-requirements-v1.md`

## Confirmed decisions

### 1. Relationship intention is user-controlled
**Decision:** Users choose whether relationship intention is a hard dealbreaker or a soft preference.

**Required behavior**
- The control is explicit, understandable and reversible.
- If marked hard, incompatible stated intentions are excluded before ranking.
- If marked soft, intention may affect ranking and must be represented honestly; it cannot override any other hard preference.
- Never infer someone's relationship intention from their profile, messages, photos or behavior if they have not declared it.
- Do not silently relax this preference when the candidate pool is small. Offer an explicit change and explain the likely effect.
- Do not reveal one user's private preference settings to another user.

### 2. Age range is always a hard boundary in discovery
**Decision:** The age range a user sets is a hard eligibility boundary. Profiles outside that range must not appear in discovery just because the eligible pool is small.

**Required behavior**
- Apply the age boundary before ranking, recommendations, fallback logic or premium features.
- Never widen an age range silently or automatically.
- Users may change their own age range explicitly in settings; any change must be clear and reversible.
- Legal minimum-age eligibility is a separate mandatory gate and can never be overridden by a user's chosen age range.
- Test exact lower and upper boundary values, missing/unverified age data, cached discovery results and preference changes.
- Do not expose another person's hidden preferences or private age-assurance evidence.

### 3. Distance is a hard boundary by default; displayed location is approximate
**Decision:** The selected distance radius is a hard discovery limit by default. Profile location display is vague and privacy-preserving, inspired by the broad approach used by Hinge, without copying proprietary UI.

**Required behavior**
- Apply distance eligibility before ranking, recommendations, fallback logic or premium features.
- Never expand the radius silently or automatically.
- Users may explicitly widen or narrow their radius; changes must be clear and reversible.
- When the pool is too small, explain the limitation and offer a choice to expand the radius, wait, or adjust another preference.
- Display only an approximate area, such as a broad neighbourhood, borough or town, where appropriate; do not show street addresses or precise map pins.
- Distance labels should be rounded or broad enough not to enable triangulation. Do not expose exact coordinates, live movement, home/work location or raw GPS data to other users.
- Request only the location precision needed for the feature. Consider approximate or user-selected location if it can satisfy the product need.
- Define the distance calculation and handling of stale/approximate location during implementation; do not present precision that the underlying location data cannot support.
- Test boundary distances, missing location, location-permission denial, stale location, cache invalidation and a user changing their radius.
- Legal age eligibility, blocks and safety restrictions remain mandatory exclusions regardless of distance.

### 4. Gender identity and discovery preferences follow the Hinge-style model
**Decision:** Separate a user's gender identity from the genders they want to see in discovery. Users explicitly select who they want to see; those selections act as hard discovery constraints. Offer inclusive gender identity options and let users control how their gender is displayed. This adopts the approach described in Hinge's public gender-diversity guidance, without copying its interface or proprietary implementation.

**Required behavior**
- Ask separately about a user's gender identity and who they want to see.
- Let users select all genders they want to be shown, using an inclusive and clearly explained set of options.
- Apply selected discovery categories as hard eligibility constraints before ranking. Do not show profiles outside selected categories because of low supply, engagement prediction or premium status.
- Gender identity must be self-described; never infer it from photos, names, appearance, voice or behavior.
- “Women” includes cisgender women and transgender women; do not define “women” as cisgender women only.
- Do not introduce a cis/trans exclusion filter by default. Any future additional compatibility controls require separate user research, privacy/legal review, anti-discrimination and abuse assessment, and clear voluntary disclosure controls.
- Let users control whether and how their gender identity appears on their profile, while ensuring the system still has the minimum data needed to apply their chosen discovery preferences.
- Do not reveal private gender identity fields or internal matching settings to other users.
- Treat gender identity and sexual orientation as sensitive personal data: minimise collection, limit access, document lawful basis and retention, and do not use them for unrelated ranking or advertising.
- Provide respectful reporting and moderation against harassment or discriminatory abuse.

### 5. Children and family plans are user-controlled preferences
**Decision:** Users choose whether their preferences about children and family plans are hard dealbreakers or soft preferences, where those fields are offered.

**Required behavior**
- Use clear, non-judgmental options (for example, wants children, doesn't want children, has children, open to children, unsure, or prefer not to say where appropriate). Final labels and available combinations need product and privacy review.
- Distinguish whether someone already has children from whether they want children in the future; these are different questions.
- Let users decide separately whether each relevant field is a hard dealbreaker or a soft preference, rather than bundling all family-related answers into one rule.
- A hard mismatch excludes the profile before ranking. A soft mismatch can influence ranking but never overrides any other hard rule.
- Never infer parenthood or family plans from photos, age, messages or other behavior.
- Do not expose hidden preference settings to another user or imply that any one family choice is better.
- If a relevant value is unknown or undisclosed, follow an explicit policy; do not silently treat it as compatible.
- Keep the fields optional unless a specific, reviewed product requirement justifies otherwise.
- Ensure family-related data is collected for a defined matching purpose, protected appropriately, and covered by retention/deletion rules.

### 6. Smoking, alcohol and lifestyle preferences
**Decision:** Use a hybrid model: users can mark clearly defined lifestyle preferences as hard dealbreakers or soft preferences where the field is offered. Default to **soft preference** to avoid unnecessarily narrowing the pool, except where the user explicitly chooses a hard boundary.

**Required behavior**
- Keep each field independent; do not bundle smoking, alcohol and other lifestyle choices into one global lifestyle score.
- Provide clear, neutral options such as smoking, alcohol use, and other lifestyle topics only when there is a defined matching purpose and a privacy review.
- Let the user explicitly mark each offered field as a dealbreaker. A hard mismatch excludes the profile before ranking; a soft mismatch may influence ranking only.
- Never infer a lifestyle habit from a photo, appearance, venue, social circle, messages or behavior.
- Do not label users as better or worse based on lifestyle answers; do not use these answers for unrelated ranking or advertising.
- Keep answers optional, support “prefer not to say” where appropriate, and define how missing/unknown values behave before launch.
- Do not silently change a user's hard/soft setting when the pool is small; explain the limitation and let the user decide whether to change it.
- Treat health-related or otherwise sensitive lifestyle data with extra care; avoid collecting detail that is not needed for matching.
- Ensure premium features cannot bypass hard lifestyle preferences.

### 7. Religion, political views and other sensitive preferences
**Decision in principle:** Keep these fields optional, minimise collection, and do not use them in matching until a specific purpose, privacy impact and legal basis have been reviewed.

**Required behavior**
- Do not require users to disclose religion, political views or other sensitive beliefs to create or use a profile.
- Do not infer these traits from names, photos, location, social connections, messages or behavior.
- Do not use sensitive fields for advertising, engagement optimisation or unrelated ranking.
- Before any sensitive field is used for eligibility or ranking, document the user benefit, exact purpose, lawful basis and any additional conditions required for special-category data, retention/deletion rules, access controls, user-facing explanation and abuse risks.
- Keep profile disclosure and matching use as separate choices where feasible; displaying a field on a profile must not automatically make it a matching signal.
- If review does not establish a clear, lawful and proportionate use, omit the field from matching. Do not let premium status change this rule.

### 8. Interests and prompt similarity
**Decision:** Shared interests, hobbies and prompt similarity are soft matching signals only.

**Required behavior**
- These signals may improve ordering or explain why a recommendation could be relevant, but must never override hard eligibility rules.
- Do not equate shared interests with guaranteed compatibility or use them as a proxy to infer sensitive traits.
- Avoid presenting a single opaque compatibility score as a definitive measure of relationship potential.

### 9. Profile completeness
**Decision:** Profile completeness is not a desirability score. Require only the minimum information needed for safe, functional participation.

**Required behavior**
- Do not rank people as more desirable simply because they have more photos, longer biographies, more prompts or more polished profiles.
- Define minimum profile requirements around functionality, authenticity and safety—not popularity or attractiveness.
- Offer optional prompts and profile fields without implying that users must disclose more personal information to be seen.
- Do not let paid features bypass minimum safety or authenticity requirements.

### 10. Matching and recommendation fairness
**Decision:** Pinchd should use reciprocal recommendations with hard eligibility gates, soft-signal ranking, fair-opportunity monitoring and truthful explanations. The production algorithm and ranking weights remain open until technical, privacy and post-launch validation work is complete.

**Required behavior**
- Apply all hard constraints before ranking; standard discovery must not return profiles that violate a hard constraint.
- Evaluate potential relevance to both people rather than optimising only one person's likelihood of liking a profile. Predicted interest is uncertain and never implies consent.
- Use only explicitly declared preferences and permitted profile information; do not infer sensitive traits from photos, messages, names or location.
- Keep safety risk and moderation decisions separate from romantic desirability and compatibility ranking.
- Monitor exposure opportunity and concentration, including whether eligible new or less frequently active users are effectively excluded from discovery.
- Any controlled exploration must remain within hard constraints and be measured for relevance, fairness and safety.
- Provide a truthful, concise “Why this profile?” explanation based on visible profile information and the viewer's own settings. Never expose hidden preferences, sensitive inferences or internal safety flags.
- If the eligible pool is small, be honest and offer user-controlled choices to wait, broaden a named soft preference, expand distance, or pause. Never use fake profiles, fabricated activity or silent filter relaxation.
- Do not sell a hidden compatibility advantage or public desirability score.
- Tune and validate ranking using controlled experiments with pre-defined primary and guardrail metrics; do not equate engagement with relationship quality.

**Initial acceptance criteria**
- Hard-preference violation rate in standard discovery: zero, with exceptions limited to documented data-correction or moderation edge cases.
- Any soft-preference expansion requires explicit user action, identifies what changed and is reversible.
- Recommendations can explain their main reasons without disclosing private data.
- Safety decisions do not reduce a person's romantic desirability score; they may trigger separate verification, rate limits, review or restrictions.
- Monitoring covers reciprocal match outcomes, unique eligible profiles receiving exposure, exposure concentration, cold-start performance, user complaints and safety outcomes separately.

### 11. Relationship-goal compatibility — decided policy
**Decision:** Use a goal + openness model. Users state a primary relationship intention and may optionally select other goals they are genuinely open to. The compatibility matrix below is a policy baseline; implementation details remain subject to privacy/accessibility review and post-launch validation.

**Core rules**
- Relationship goals are not morally ranked. Long-term, short-term, casual and undecided goals are all valid.
- A user's hard relationship-goal setting is an eligibility constraint. A soft setting may affect ranking but never overrides any hard constraint.
- Evaluate both people's stated goals and hard/soft settings before ranking. If either person's hard rule excludes the pair, do not show the profile in standard discovery.
- Keep long-term, short-term, casual, open-minded/exploring and not-sure-yet states distinct. “Prefer not to say” and missing data are also distinct from “open-minded.”
- Do not assume someone seeking long-term is open to casual dating, or vice versa, unless they explicitly say so.
- If the pool is small, never silently relax relationship-goal constraints. Offer a clear choice to change a soft preference, wait or pause.
- Premium features cannot bypass relationship-goal dealbreakers.
- Changes to goals or preference strength must invalidate affected cached recommendations.

**Compatibility matrix — adopted policy baseline**
| Viewer goal | Potentially compatible candidate goals | Caution |
|---|---|---|
| Long-term | Long-term; open-minded if candidate explicitly includes long-term; unsure if viewer allows uncertainty | Short-term/casual are excluded when the viewer marks long-term as a hard requirement and the candidate has not declared openness to long-term |
| Short-term | Short-term; casual if meanings align; open-minded if candidate explicitly includes short-term; unsure if viewer allows uncertainty | Do not assume a long-term-only candidate is open to short-term |
| Casual | Casual; short-term if meanings align; open-minded if candidate explicitly includes casual; unsure if viewer allows uncertainty | Do not assume a long-term-only candidate is open to casual |
| Open-minded / exploring | Goals the person explicitly says they are open to; another open-minded person; unsure if uncertainty is acceptable | This label does not imply willingness to accept every goal |
| Not sure yet | Any goal only when the other person's stated openness and hard/soft settings allow it | Uncertainty is not automatically compatible |

“Potentially compatible” means eligible for ranking only after both users' hard constraints and all other eligibility checks pass. It never implies attraction or consent. This matrix is the policy baseline and must be validated against real outcomes after launch; it is not a claim that every user with a given label will want every other label.

**Acceptance criteria**
- Relationship-goal compatibility is evaluated for both people before ranking.
- Hard-incompatible pairs never appear in standard discovery.
- Missing, undisclosed and unsure values have explicit, distinct behavior.
- Tests cover each matrix combination, both users' hard/soft settings, cache invalidation and premium bypass attempts.

## Global rules proposed for confirmation
1. **Hard constraints precede ranking.** No ranking, premium feature, cold-start fallback or engagement objective can override a hard constraint.
2. **No silent relaxation.** When discovery is limited, explain the limitation and let the user choose which preferences to change.
3. **Blocks and safety restrictions are never preferences.** They are mandatory exclusions and cannot be relaxed.
4. **User-stated data only.** Do not infer sensitive traits or preferences from photos, names, messages or precise location.
5. **Unknown is not automatically compatible.** If a hard rule requires a value and the candidate's value is missing or unknown, apply a documented fail-closed policy for that rule, or ask the user whether unknown values may be shown. Do not silently assume a match.
6. **Explain without leaking.** Explain recommendations using the viewer's own settings and the candidate's visible or explicitly shared information, without exposing another person's hidden settings.
7. **Audit and test.** Preference changes should take effect predictably across discovery, likes, cached results and recommendation APIs. Record policy version/change metadata without logging unnecessary sensitive values.
8. **Free integrity.** Payment cannot bypass another user's hard preferences, blocks or consent.
9. **Incoming likes use a transparent rolling queue.** Free users see five eligible incoming likes unblurred at a time. When they like back or pass on one, the next eligible like moves into the visible set; further likes may remain blurred while queued. Subscribers can see the full eligible incoming-like queue unblurred and review it incrementally. Five is an initial hypothesis to validate, not a proven optimum. Never fabricate likes, activity, urgency or scarcity, or manipulate ordering to pressure an upgrade. Apply hard preferences, blocks, suspensions, unmatched-pair exclusions and safety restrictions before display; payment cannot bypass them.
10. **Unmatching is quiet and immediate.** Remove the connection from both active match lists, close messaging, and do not disclose who initiated the unmatch. Keep the former match in a separate read-only history area for both users for 30 days after unmatching; then remove it from the user-facing history list. The history entry offers report and block actions while visible, but never allows renewed contact. An unmatched pair is permanently excluded from matching each other again. Unmatching does not automatically block; blocking and reporting remain separate actions. Retain underlying records only as justified by documented safety, legal or operational needs, independently of the 30-day display period.
11. **Commented likes may receive modest visibility preference.** Users may attach a user-written comment to a specific profile photo or prompt. The comment does not permit contact unless interest is reciprocated. Relevant comments may receive a transparent, modest ordering advantage in the recipient's incoming likes, but no guarantee of visibility; payment, AI use and comment length are not ranking advantages, and hard preferences, blocks, safety controls and recipient agency always take precedence.
12. **Ethical premium baseline.** Keep core matching, incoming likes, mutual messaging, blocking, reporting, unmatching, essential privacy controls and legally required age assurance free. Start with one optional premium tier as a hypothesis. Candidate benefits for validation: profile customisation, optional user-reviewed profile/message coaching, saved or organised discovery lists, extra discovery organisation, truthful travel mode, and more detailed compatibility explanations based on user-provided data. Do not sell hidden ranking advantages, paid comment boosts, guaranteed romantic outcomes, safety or access to consent. Price and feature demand remain unvalidated; legal, privacy and app-store review are launch gates.

## Preference categories to settle
| Category | Working proposal | Status |
|---|---|---|
| Age range | Hard exclusion | **Decided** |
| Gender identity and genders wanted in discovery | Separate fields; user-selected gender categories are hard constraints | **Decided in principle; implementation/privacy review required** |
| Sexual-orientation compatibility | Use only explicitly provided settings; no inference | Proposed; privacy review needed |
| Distance | Hard exclusion by default; explicit user-controlled expansion; approximate displayed location | **Decided** |
| Relationship intention | User chooses hard or soft | **Decided** |
| Children/family plans | User chooses hard or soft per relevant field | **Decided** |
| Smoking/alcohol/lifestyle preferences | Soft by default; user can explicitly make each a hard dealbreaker | **Decided** |
| Religion/politics/other sensitive profile fields | Optional; collect only when necessary; no inference; do not use for matching until reviewed | **Decided in principle; privacy/legal review required before launch** |
| Interests, hobbies and prompt similarity | Soft ranking signals only; never override hard constraints | **Decided** |
| Profile completeness | Not a desirability score; only minimum information needed for safe, functional participation | **Decided** |
| Blocks, suspensions, age ineligibility and account deletion | Mandatory exclusion, not configurable preference | Non-negotiable product principle |

## Edge cases to test
- Candidate is outside age range by one boundary value.
- Candidate's age, relationship intention or other hard-rule field is missing.
- Candidate's gender identity is hidden or optional profile display is disabled, but eligibility still requires matching data.
- User changes a gender selection while discovery results are cached.
- Candidate's existing parenthood and future family plans differ; each preference must be evaluated separately.
- User marks one family field hard and another soft.
- Candidate has not disclosed family plans; the system must follow the explicit unknown-value rule.
- User marks a lifestyle field as a hard dealbreaker, and the candidate's answer is missing or undisclosed.
- One person blocks another while a like, message notification or queued recommendation is pending.
- User changes a preference while discovery results are cached, then narrows it again.
- Candidate changes relationship intention or gender display settings after appearing in discovery.
- Paid user tries to override a hard preference.
- Empty pool: product explains low availability without silently widening filters.
- A recommendation explanation would reveal another user's private preferences.
- User pauses, hides or deletes an account; stale profiles disappear from discovery and downstream surfaces.
- Distance boundary, stale/approximate location, permission denial and location change invalidation.
- Approximate profile location and rounded distance labels do not enable reliable triangulation.

## Decision protocol
For each remaining category, record:
- User-facing wording and available values.
- Whether the user may mark it hard or soft.
- Behavior when the candidate's value is unknown.
- Whether it is used for eligibility, ranking, or both.
- Whether any sensitive-data or legal review is needed.
- How the choice is explained, changed and tested.

Do not treat the proposed categories as final policy until explicitly approved.

## Source
- Hinge, “Supporting Gender Diversity in the Hinge Community”: https://help.hinge.co/hc/en-us/articles/360053562713-Supporting-Gender-Diversity-in-the-Hinge-Community
- Hinge, “How Hinge decided which gender options to include”: https://help.hinge.co/hc/en-us/articles/4407404339603-How-did-Hinge-decide-which-gender-options-to-include
- Hinge, “How We Connect Daters”: https://hinge.co/how-we-connect-daters
- Hinge, “Match Note” / gender-specific preferences and Gender Toggle: https://hinge.co/newsroom/match-note


## Status and reconciliation note — 21 October 2026

This file records working policy decisions; its overall status remains **In progress** until reconciled with the product decision baseline. Refer to `../research/2026-10-21-workstream-1-reconciled-decision-baseline.md`. Confirmed policy principles should be preserved; any remaining unknown-value behavior, field-level legal/privacy review, and implementation detail must remain explicit rather than inferred from a draft requirement.
