# Research review: Hinge “Making Confident Moves” redesign
**Reviewed:** 21 October 2026  
**Source:** Irene Montalban, “Hinge App Redesign: Making Confident Moves” (12 December 2019) — https://medium.com/@irenexmontalban/hinge-app-redesign-making-confident-moves-3917f6e71204  
**Related source:** Designer’s Dribbble project — https://dribbble.com/shots/14037228-Hinge-App-Redesign

## Executive recommendation
**Explore private saved profiles; defer a lightweight music prompt; do not implement public reviews/ratings.** This is a designer case study from 2019, not evidence that the feature improved real-world relationship outcomes at scale. The author reports qualitative usability feedback, including preference for eight bookmark slots and a 24-hour window. Treat those details as prototype findings to consider, not validated product requirements.

## What the designer explored
The case study identified a decision-making problem: people want to make thoughtful choices but may feel constrained to like a profile element or comment immediately. Three solution areas were explored:
1. **Time-sensitive bookmarking:** save profiles to revisit and compose a more thoughtful message later. The prototype used eight slots and a 24-hour expiry.
2. **Music:** a fun potential conversation starter, but with integration/content-bank effort and possible fit issues.
3. **Reviews:** extra social information, but highly subjective and potentially harmful.

## Pinchd feature assessment

### 1. Private shortlist / saved profiles — EXPLORE
**Potential benefits**
- Lets users return when they have enough time and attention.
- Reduces accidental loss of an interesting profile while browsing.
- Supports intentional dating without requiring an immediate like or message.
- Can complement profile-specific likes and comments rather than replacing them.

**Risks and design constraints**
- A permanent shortlist can become a backlog and create decision fatigue.
- A timer, countdown or expiring saved profile can manufacture urgency and pressure.
- Saving must not notify the saved person or expose the list to other users.
- Saved profiles must obey hard preferences, blocks, visibility changes, account deletion and moderation restrictions.
- Avoid a paid-only save feature; core connection should not depend on premium access.

**Provisional MVP shape**
- Private “Saved” list accessible from discovery and profile.
- Save/unsave action available on a profile.
- No expiry timer and no notification to the other person.
- Revalidate eligibility and block state when opening saved profiles; suppress inaccessible profiles with a clear, privacy-safe explanation.
- Keep likes and messages as explicit separate actions; saving must never imply romantic interest or send anything.
- Consider a cap only if usability or abuse data justifies one. If used, explain it plainly and avoid guilt-based prompts.
- Add optional user-controlled reminder only if later evidence shows it helps; reminders must be quiet, configurable and easy to disable.

**Decision:** Candidate for MVP or early beta, subject to checking scope against the consolidated requirements. High product fit; relatively low technical complexity, with meaningful privacy and state-consistency edge cases.

### 2. Music feature — DEFER / TEST LIGHTWEIGHT VERSION
A full music integration adds third-party dependencies, API availability, licensing/content rights, moderation and maintenance concerns. Start, if desired, with an optional text-based “song I’d share with you” prompt or a user-entered favourite song. Do not autoplay audio. Assess accessibility, privacy, abuse risk and whether it creates substantive conversations before investing in an integration.

**Decision:** Not an MVP dependency. Consider a later experiment only after core profile and messaging flows are healthy.

### 3. Public user reviews or ratings — DO NOT BUILD
Public reviews are a poor fit for a dating product where reports can be subjective, intimate and retaliatory. Risks include harassment, defamation disputes, coercion, fake reviews, outing, brigading and permanent reputational harm. A public score could become a desirability/ranking mechanism and damage trust.

**Safer alternative:** If post-date feedback is ever explored, keep it private, structured, optional and purpose-limited. Do not publish reviews or use feedback as a public popularity score. Do not use sensitive or unverified claims to change someone's romantic visibility. Any internal use should have clear notice, minimisation, access controls, retention limits, bias assessment and an appeal path for material consequences.

**Decision:** Exclude from MVP and roadmap unless a substantially different, safety-reviewed use case emerges.

## Research quality and limitations
- The article is a student/designer case study published in 2019, not a controlled production evaluation.
- The article describes interviews with young professionals aged 22–27 and iterative usability tests, but does not provide enough detail to infer representative demand or long-term outcomes.
- Reported user preference for eight slots and 24 hours should not be treated as a universal optimal limit.
- The feature's value should be judged by user-reported usefulness, thoughtful reciprocal interactions and downstream conversations/dates—not by saved-profile count or time spent in the app.
- Hinge's broader product principles support defining desired outcomes and avoiding features that add engagement without improving users' lives; this is useful context, not proof this exact concept will work for Pinchd.

## Suggested validation after launch
Instrument only minimal events: save, unsave, open saved profile, like/comment from saved profile, and save-to-action conversion. Do not log profile text or message content in analytics. Compare a cohort with access to saved profiles against a suitable baseline or staged rollout, and assess:
- Whether users find it useful.
- Whether it increases meaningful profile-specific likes/comments and reciprocal conversations.
- Whether saved profiles become a stale backlog.
- Whether it affects reporting, blocking, or privacy complaints.
- Whether benefits differ across user groups or market-density levels.

Avoid optimising for saves as an end in themselves. Provide user controls and a clear way to clear the list.

## Decision summary
| Concept | Recommendation | Timing |
|---|---|---|
| Private saved profiles | Explore; no forced expiry | MVP candidate / early beta |
| Music prompt | Lightweight text prompt only if useful | Later experiment |
| Public reviews/ratings | Reject | Exclude |
| Private post-date feedback | Only with safeguards and a clear purpose | Future research |

## Sources
- Irene Montalban, Medium case study (2019): https://medium.com/@irenexmontalban/hinge-app-redesign-making-confident-moves-3917f6e71204
- Irene Montalban, Dribbble project: https://dribbble.com/shots/14037228-Hinge-App-Redesign
- Hinge product design principles: https://hinge.co/it-it/newsroom/product-design-principles


## 8. Discovery and recommendation system research (21 October 2026)

### Officially documented Hinge behaviour
Hinge describes its current recommendation system as a collection of algorithms that combines stated preferences with activity signals. Its 2025 update says a deep-learning system is intended to predict mutual compatibility: not only who a person may like, but who is likely to like them back. Hinge says recommendations take account of both daters' preferences, past likes/skips/matches and mutual dealbreakers. Source: https://hinge.co/how-we-connect-daters

Hinge's privacy/help explanation says it uses information supplied by members and their use of the service—including age, gender, location, preferences, dealbreakers, likes, skips, matches and, in its wording, people with whom users may exchange phone numbers—to recommend people to each user and recommend each user to others. It also says users can change discovery settings to filter profiles. Source: https://help.hinge.co/hc/en-us/articles/360010956733-Automated-Decision-Making-and-Profiling-at-Hinge

Hinge's published explanation has described some preferences as flexible when not treated as dealbreakers. Its Help Centre page, updated in 2026, explains that when a preference is not strict, the system prioritises profiles that fit but may show others. In July 2026, Reddit users reported a UI change from the old dealbreaker toggle to a more explicit Flexible Filters choice. The Reddit post is useful for usability feedback, but is not authoritative evidence of universal rollout or defaults. Source: https://www.reddit.com/r/hingeapp/comments/1v8q389/dealbreakers_is_now_flexible_filters/

Hinge's official 2025 product update says it encouraged users to be open-minded about certain filters, including distance, and to understand how dealbreakers and "Prefer Not to Say" affect visibility. It reports a double-digit increase in matches following global recommendation-system changes; this is a company-reported result, and the public page does not provide a full independent evaluation design. Source: https://hinge.co/newsroom/hinge-2025-product-evolution


### User-reported pain points: treat as signals, not prevalence estimates
Public Reddit discussions include recurring complaints about:
- People appearing outside stated age or distance ranges when users thought their filters were strict.
- Confusion between a preference and a dealbreaker, including uncertainty about whether a setting affects profiles shown, profiles shown to others, or both.
- Low trust in filters after a user sees an apparent exception, whether due to flexible settings, location representation, stale state, product changes or a genuine defect.
- Relationship intention and relationship type being important enough that users want clearer filtering; some older threads complain that such filters were paid or incomplete.
- Missing or hidden answers on family-plan fields making compatibility uncertain even when a user has selected a strict preference.

Examples include a 2026 thread describing profiles well outside an age/distance range (https://www.reddit.com/r/hingeapp/comments/1sk1bv9/do_filters_not_work/), a 2025 report of profiles apparently outside enabled dealbreakers (https://www.reddit.com/r/hingeapp/comments/1jx5slj), and a July–October 2026 discussion about the Flexible Filters wording change and confusion around whether preferences were relaxed (https://www.reddit.com/r/hingeapp/comments/1v8q389/dealbreakers_is_now_flexible_filters/). These are self-selected reports, cannot establish how often the problem occurs, and may describe different versions or settings. They do establish that the distinction between strict and flexible preferences must be understandable and testable.

### Implications for Pinchd
1. Separate **eligibility** from **ranking**. Apply legal age gate, account/safety eligibility, blocks, permanent unmatched-pair exclusions and every active hard preference before recommendation ranking.
2. Make every preference's effect explicit: "Hard boundary — never show profiles outside this setting" versus "Soft preference — prioritise matches, but allow other eligible profiles." Avoid relying on a hidden toggle or unexplained terminology.
3. Do not silently relax a hard preference to fill the feed, improve match counts, boost engagement or support Premium conversion. If a user opts into a flexible setting, confirm exactly which field becomes flexible and label any resulting profile accordingly.
4. Use reciprocal ranking only after hard eligibility gates: estimate relevance to the viewer and likelihood of reciprocal interest, then rank among eligible candidates. Do not treat predicted popularity as desirability or as a safety score.
5. Be cautious with behavioural learning. Likes, passes and matches may inform soft ranking, but users should be able to correct mistaken inferences, reset or limit personalisation, and understand at a high level what signals influence recommendations.
6. Missing or "prefer not to say" data is unknown, not a match. For hard preferences, do not silently assume an unknown answer satisfies the requirement; make the outcome explicit and minimise unnecessary exclusion by offering users a clear choice where appropriate.
7. Keep core preference integrity free. Premium may offer extra convenience filters only if justified, but payment must not be required for users to enforce core boundaries around age, distance, relationship intention or other agreed hard constraints.
8. Test preference enforcement end-to-end: candidate retrieval, ranking, caches, incoming likes, notifications, search refresh, location changes, profile edits and stale recommendation queues. A hard preference must be enforced at the server-authorised eligibility layer, not only in the interface.

### Proposed acceptance criteria to carry into requirements
- A profile outside any active hard preference is not returned, liked, recommended, notified about or surfaced through an alternate queue.
- Every soft preference is labelled as soft; if a user changes it to hard, the new rule applies immediately to subsequent discovery and incoming-like eligibility, with stale results invalidated.
- The user can see a concise explanation for why a profile is recommended, based on non-sensitive, user-provided compatibility factors, without exposing another person's private preferences.
- A recommendation may be ranked highly only if it passes all mandatory eligibility checks; recommendation score cannot override a block, report-related restriction, suspension, deletion, age gate or unmatched-pair exclusion.
- Learning from skips or likes adjusts only permitted soft ranking signals and never changes explicit hard boundaries.
- Unknown or withheld data is handled according to a documented field-specific rule; no unknown answer is silently treated as satisfying a hard preference.
- Preference changes are audited and tested across feed, likes, notifications, caches and reciprocal recommendation paths.

### Research conclusion
Hinge's useful model is a two-sided recommendation system that combines stated preferences and behavioural signals, with the product encouraging some users to widen preferences. Public user reports show that the distinction between soft preferences and strict dealbreakers can be confusing and that perceived exceptions can erode trust. Pinchd should adopt two-sided ranking and explainability, but differentiate itself by making hard boundaries reliable, visible and free. This is a product policy recommendation, not a claim that Hinge's current system always violates its published rules.


## 9. Full research: profile prompts, authenticity and conversation design (21 October 2026)

### 9.1 Hinge's documented profile model
Hinge's profile design uses photos and structured Prompts to give viewers more than a swipeable image. Its public materials describe text prompts, voice prompts, polls, and optional profile coaching. Its April 2025 profile-refresh guide recommends showing different sides of a person's life, using current photos, avoiding generic or clichéd answers, and using voice or polls when they help express personality. Hinge describes Voice Prompts as up to 30 seconds and Prompt Polls as a way to present three options and invite interaction. Source: https://hinge.co/newsroom/profile-refresh-guide

Hinge introduced its “Your World” prompt collection in June 2025 with Esther Perel. The prompts invite users to talk about their social role, people and experiences that shape them, what they enjoy, and small personal rituals. Hinge also retired older prompts as it refreshed the collection. This demonstrates an active editorial strategy: prompt libraries need curation and refresh, not simply an ever-growing catalogue. Source: https://hinge.co/newsroom/your-world-prompts-with-esther-perel

Hinge's January 2025 Prompt Feedback launch offers three levels of feedback: keep an answer, try a small change, or go deeper. Feedback is private and optional; the user chooses whether to change their answer. Hinge reported that in its own 2024 data, likes on text prompts were 47% more likely to lead to a date than likes on photos, and that 63% of surveyed Hinge daters found it difficult to know what to include on their profiles. These are company-reported figures, not independently verified causal estimates; users who like a prompt may differ from those who like a photo. Source: https://hinge.co/newsroom/prompt-feedback

Hinge's November 2025 explanation of Prompt Feedback says nearly 40% of Hinge users reported finding prompts difficult to navigate. It frames the tool around transparency, authenticity and equity, and recognises that short answers can sometimes be intentional. This figure differs from the 63% figure above because the questions and/or survey timing may differ; do not treat them as contradictory without the underlying survey instruments. Source: https://hinge.co/en-gb/newsroom/about-prompt-feedback

Hinge launched Convo Starters in December 2025, later announcing UK availability on 23 February 2026. The feature gives three personalised ideas tied to a particular photo or prompt, encouraging a question, exchange of views or deeper discussion. It does not write a ready-to-send message; users compose their own words. Hinge reported that 72% of surveyed daters were more likely to consider someone whose like included a message and that likes with comments were twice as likely to lead to a date; these are company-reported findings and do not prove that comments alone cause dates. Source: https://hinge.co/newsroom/convo-starters


### 9.2 Public user discussions: recurring themes and limitations
A December 2025 Reddit discussion ranking Hinge prompts shows strong opinions about prompt choice and tone. Commenters criticised generic statements such as wanting kindness or good communication when no specific detail makes the answer distinctive; they also criticised negative, ultimatum-like prompts and prompts that put the burden on the viewer to start a conversation. Other commenters disagreed about how long or short answers should be, illustrating that there is no single answer length or prompt that works for everyone. Source: https://www.reddit.com/r/hingeapp/comments/1pg27hn/a_ruthless_ranking_of_the_2025_hinge_prompts/

A 2023 Reddit thread on common prompt mistakes similarly argues that one-word, generic and cliché responses can feel low-effort. It is a community guide rather than a controlled study, and its advice is partly subjective. Source: https://www.reddit.com/r/hingeapp/comments/143yedd

Profile-review threads in 2025–2026 show users asking for help with answers that feel bland, struggling to receive interest, or reporting that matches do not ask questions and conversations fade. Such threads are selected for people seeking help and are not representative samples. They nevertheless point to product needs: users want help expressing themselves, want profile content to give others something concrete to respond to, and cannot be promised that better prompts will produce matches. Example: https://www.reddit.com/r/hingeapp/comments/1x4jrw5/ and https://www.reddit.com/r/hingeapp/comments/1x97ifx/

Another recurring discussion is that profile advice can become formulaic: users imitate whatever prompt is said to perform well, making profiles less individual. Community advice is inconsistent—some people advocate very short answers, others detailed stories; some prioritise photos above all, others value text. This disagreement is a reason not to turn popularity anecdotes into rigid profile scoring rules. Source: https://www.reddit.com/r/hingeapp/comments/1pg27hn/a_ruthless_ranking_of_the_2025_hinge_prompts/

### 9.3 Research beyond Hinge
A 2006 peer-reviewed interview study of 34 online daters found that people manage a tension between presenting an ideal self and wanting to appear authentic, using small cues to infer whether others' claims are credible. This is an older study of a large online dating site, not current Hinge users, but the tension remains directly relevant to profile design. Source: Ellison, Heino and Gibbs, “Managing Impressions Online: Self-Presentation Processes in the Online Dating Environment,” Journal of Computer-Mediated Communication: https://onlinelibrary.wiley.com/doi/10.1111/j.1083-6101.2006.00020.x

A 2025 preprint based on semi-structured interviews with 20 Chinese college-educated young adult dating-app users reports that participants sometimes selected idealised photos and presented traits they believed others wanted, while also hiding identifying details to reduce privacy risk. This sample is small and culturally specific, so it cannot establish prevalence in the UK; it highlights the tension between expressive profiles, accurate representation and privacy. Source: https://arxiv.org/abs/2502.11430

Research on deceptive self-presentation has found that profile inaccuracies can occur, though a study of 80 online daters found deviations in physical attributes were generally small and people reported relationship information as more accurate than photos. This is not a basis for assuming users lie; it supports avoiding interfaces that reward exaggerated claims or make polished performance feel necessary. Source: https://pubmed.ncbi.nlm.nih.gov/18593866/

A 2024 qualitative study of 18 users of a voice-based dating app in China found that voice-based interaction changes impression management and how users perceive each other. Its small, culturally specific sample does not prove voice prompts improve match quality, but it supports treating voice as a distinct expressive option rather than a universal requirement. Source: https://arxiv.org/abs/2402.19328


### 9.4 Recommended Pinchd profile principles
1. **Invite specific self-expression, not a sales pitch.** Prompt topics should help users share a small story, preference, ritual, opinion, relationship value or future activity. Avoid prompts that reward generic claims such as “I am kind” without context.
2. **Use a balanced profile.** Encourage a mix of self-description, lived experience, lightness/playfulness and a prompt that is easy to respond to. Do not require every profile to be witty, vulnerable, highly verbal or extroverted.
3. **Make prompts accessible across relationship goals.** Include prompts for people seeking long-term relationships, casual dating, short-term dating, exploring, or unsure. Avoid implying that one goal is more mature or worthy than another. Do not pressure users to disclose private relationship history, trauma, sexual experiences or sensitive beliefs.
4. **Offer optional profile coaching.** Coaching should point out generic phrasing and suggest ways to add a concrete detail, but must not rewrite a user into a generic persona. Show suggestions privately, explain why they may help, allow rejection, and never publish changes without explicit approval.
5. **Conversation suggestions should be profile-grounded.** Offer optional question ideas tied to a specific photo or prompt, but don't auto-send, impersonate the user, or present AI-written language as the user's own. Let the user edit everything and disable the feature.
6. **Avoid engagement manipulation.** Do not claim a particular prompt guarantees matches or dates. Do not penalise short, quiet, neurodivergent, disabled or non-native-speaker communication styles simply because they differ from popular examples.
7. **Multimodality should be optional.** Text should be sufficient to create a complete profile. Voice/video/polls can add expression, but must not become paid prerequisites or create disadvantage for users who cannot or do not want to record them.
8. **Truthfulness and recency.** Give users control to update profile content and encourage current, representative photos without demanding highly identifying locations, workplace, family details or other information that creates privacy risk.
9. **Prompts should create reply opportunities.** Prefer a prompt answer that includes a specific hook or an invitation to share a view over a prompt that merely asks others to do all the work. But don't require every answer to end in a question.
10. **Keep profile quality separate from desirability.** Completion status can help users understand setup, but must not become a hidden attractiveness score or a reason to reduce discovery exposure unfairly.

### 9.5 Candidate profile-content framework (not final UI)
A complete profile could combine: (a) representative photos; (b) a prompt showing everyday personality or interests; (c) a prompt showing values or what a good connection looks like to the user; and (d) a conversation-friendly prompt that reveals a story, choice, opinion or activity. Exact counts, editing flow, media limits, prompt library and whether voice/polls launch in version one remain open pending product, accessibility, safety and technical review.

Prompt categories to research and author in a later content-design pass:
- Everyday life and routines.
- Interests and enthusiasms, with a concrete example.
- Humour and playful choices without humiliation or adversarial bait.
- Values and how someone treats people, framed as real behaviours rather than virtue-signalling labels.
- What a good date or shared activity could look like.
- Relationship intentions and what the user is genuinely open to, without forcing disclosure beyond the separate intention fields.
- Personal growth and what the user is learning, without pressuring disclosure of trauma.
- Community, friendship and meaningful experiences, while avoiding identifying third parties without consent.

### 9.6 Candidate requirements and acceptance criteria
- Users can build a usable profile with text and photos; optional voice/video/poll elements are not required to participate.
- Prompt library supports diverse relationship goals, communication styles and accessibility needs; no goal is framed as superior.
- Prompt answers are editable and deletable; preview matches the content that other users see.
- Profile coaching is optional, private by default, gives suggestions rather than a replacement identity, and requires explicit user approval before edits publish.
- Any AI conversation inspiration is tied to content the recipient chose to show, clearly labelled as suggestions, editable, disableable, and never auto-sent.
- Do not infer sensitive traits from prompt content or use prompts as covert psychological or desirability scoring.
- Do not penalise blank/short answers with hidden ranking penalties beyond transparent minimum functional profile requirements.
- Report/block and moderation controls are accessible from profile surfaces; prompt responses cannot contain threats, harassment, discriminatory abuse or contact details used to evade safeguards.
- Measure prompt engagement, likes with comments, match-to-conversation starts, conversation quality proxies and user-reported usefulness; interpret observational correlations cautiously and monitor inequity by broad, privacy-safe cohorts.

### 9.7 Research conclusion
Hinge's profile strategy has evolved from text prompts into a broader system of photos, text, voice, polls, optional coaching and profile-grounded conversation suggestions. Its stated aim is to help users show specific, authentic details and make first messages easier. Company-reported metrics suggest prompts and comments are associated with better downstream outcomes, but public sources do not establish causal effects independently. Reddit feedback highlights cliché, generic answers, negative tone and the danger of formulaic advice, while research on online self-presentation shows the tension between authenticity, impression management and privacy. Pinchd should adopt optional, specific, conversation-friendly prompts and user-controlled coaching, while avoiding a universal ideal profile, forced vulnerability, AI impersonation and any claim that profile polish guarantees romantic outcomes.


## 10. Profile authenticity, photo integrity and verification (21 October 2026)

### Hinge's documented verification approach
Hinge's Help Centre says Selfie Verification is voluntary in most contexts: users complete a short video-selfie process and receive a visible badge. Hinge explicitly warns that verification does not guarantee a person's safety. In some regions, a biometric Face Check is required; the process and terms differ by region. Source: https://help.hinge.co/hc/en-us/articles/10303221435539-What-is-Selfie-Verification

The practical lesson is that verification is a signal, not a guarantee. A badge can indicate that a profile passed a particular check at a particular time; it cannot prove that every profile statement is true, that the account remains under the same person's control, or that the person will behave safely.

### Threat landscape and external evidence
Ofcom's current guidance for dating and social-discovery services identifies romance fraud, fake personas, intimate-image abuse/sextortion, harassment and stalking as relevant risks. It reports that 29% of adult internet users had experienced romance or dating scams and 6% had encountered fraud on a dating website or app, based on Ofcom research. These are broad survey measures, not rates of fake profiles on any specific app. Source: https://www.ofcom.org.uk/cy/online-safety/illegal-and-harmful-content/dating-and-social-discovery-know-the-online-safety-risks-rules-and-how-to-comply

McAfee's February 2026 UK survey release says one in five respondents had encountered a fake profile or AI-generated bot online and 34% had spotted AI-generated or modified photos on dating or social apps. This is vendor-sponsored survey research; the public release should be treated as a warning signal, not an independently validated prevalence estimate. Source: https://www.mcafee.com/en-gb/consumer-corporate/newsroom/press-releases/2026/20260210.html

UK Trading Standards' February 2025 warning says romance scammers use stolen photos and fabricated stories to build trust and then seek money, gifts or personal information. Report Fraud has also published UK cases involving offenders who used dating apps to build relationships and steal money. Sources: https://www.economy-ni.gov.uk/news/protect-yourself-online-dating-scams and https://www.reportfraud.police.uk/romance-scammer-jailed-after-deceiving-on-dating/


### Research on self-presentation and photo accuracy
Academic research cautions against treating all profile curation as deception. A 2008 study of 80 online daters found discrepancies in physical attributes were common but generally small; participants reported their photographs as less accurate than relationship information. Source: https://pubmed.ncbi.nlm.nih.gov/18593866/

A 2009 study of 54 dating profile photographs found independent judges rated around one-third as inaccurate, often due to age of photo, retouching or differences in appearance; the sample was small and from an earlier online dating context. Source: https://academic.oup.com/joc/article-abstract/59/2/367/4098407

A 2025 study of 180 online-dating profile owners and 196 judges examined profile-based personality impressions and self-presentation. It reinforces that profile impressions are not perfect readouts of someone's actual personality; a profile should be treated as an introduction, not a psychological diagnosis. Source: https://journals.sagepub.com/doi/10.1177/01461672251386488

A 2025 interview study of 20 Chinese college-educated young adult dating-app users found that participants sometimes chose idealised photos and presented desired traits, while also hiding identifying details for privacy. Its small, culturally specific sample is not representative of UK users, but it illustrates that authenticity and privacy can be in tension. Source: https://arxiv.org/abs/2502.11430

These studies are not a reason to police ordinary flattering photos, makeup, styling, selective presentation or personal privacy. Product policy should distinguish reasonable self-presentation from impersonation, stolen identity, materially misleading current appearance, fraud, harassment or attempts to evade a block or suspension.


## 11. Pinchd verification workflow: voluntary profile verification vs age assurance (21 October 2026)

### Product decision
Public profile/selfie verification is voluntary for ordinary Pinchd users. It must not be confused with any age-assurance or age-gating duty that may apply in a specific jurisdiction and service context. If legal review determines age assurance is required, implement it as a separate, clearly explained workflow and gate only the access/activity that the law requires.

### Hinge reference
Hinge's published help article (updated 17 December 2025) describes voluntary Selfie Verification using a video selfie, liveness check and 3D face authentication against profile photos. Hinge says the badge is not a guarantee of safety. Hinge also notes that in some regions a Face Check may be required, with different biometric retention conditions. This is evidence of Hinge's public product behaviour, not evidence that Pinchd should copy its biometric retention or regional mandatory-check model. Source: https://help.hinge.co/hc/en-us/articles/10303221435539-What-is-Selfie-Verification

### Proposed voluntary badge workflow
1. User explicitly initiates “Verify my profile” after a plain-language explanation of purpose, biometric processing, provider, retention/deletion, badge meaning, failure handling and available alternative.
2. Run one liveness check to establish that the submitted capture is live rather than a static image/replay.
3. Run one 1:1 face match between the live selfie and the user's current eligible profile photo(s). Do not use face search/1:N gallery search for ordinary verification.
4. Return verified / unable to verify / failed / needs review as distinct states. A failed or uncertain result should not silently imply fraud. Offer a limited reasonable retry and accessible support/review path.
5. On pass, issue a badge that explains the narrow claim (“selfie checked against profile photos”), not a guarantee of identity, honesty or safety.
6. Re-evaluate verification when profile photos materially change; remove or suspend the badge if the evidence no longer supports it, with user notification and a retry path.
7. Store the minimum result/audit data needed. Confirm provider retention, subprocessors, transfers and deletion contractually; do not retain raw selfie or biometric templates longer than necessary. Complete UK GDPR/DPIA/legal review before production biometric processing.

### Didit workflow review from supplied screenshots
The screenshot shows: Face Match -> Liveness -> duplicate Face Match -> Device & IP Analysis -> AML Screening -> branching by US state -> Age Estimation -> final decision. The advertised workflow cost shown is $0.63–$0.73.
- Remove duplicate Face Match; one face match is sufficient for the intended 1:1 comparison unless Didit documents a specific separate purpose.
- Order liveness before the single face match where technically supported.
- Exclude AML screening from voluntary profile verification. Sanctions/PEP screening is not the same purpose as confirming a live person matches their profile photos; no business need for dating users has been established.
- Do not make Device & IP Analysis a default prerequisite for the badge. Consider it only for a documented anti-fraud/security purpose, with notice, proportionality, retention and false-positive review.
- Remove age estimation and state branching from the voluntary badge workflow. If needed, create a separate age-assurance workflow whose requirements are based on legal scope and the launch jurisdiction, not merely state labels.
- Configure final decisions with separate pass, retry, review and fail outcomes; do not treat a technical error or inconclusive score as proof of deception.
- Confirm the actual module selected in Didit, since its public price list currently distinguishes passive liveness ($0.10) from active liveness ($0.15); screenshot pricing may reflect a different module/configuration.

### Privacy, accessibility and fairness guardrails
Biometric face matching can involve special-category biometric data under UK GDPR depending on purpose and processing. Before production, obtain legal/privacy review, identify lawful basis and Article 9 condition where applicable, perform a DPIA where required, provide an accessible non-biometric alternative or explain the limitations, test demographic performance and false rejection rates, and define a human review/appeal route. A badge is optional; don't penalise non-participation by reducing ordinary discovery exposure or compatibility ranking.

### Age assurance is separate
Ofcom's guidance updated 2 September 2026 lists methods capable of being highly effective at determining whether a user is 18+ in applicable contexts, including facial age estimation and photo-ID matching, but whether Pinchd has a duty and where to place an age gate depends on service scope and content/risk. Selfie profile matching is not by itself proof that a user is over 18. Source: https://www.ofcom.org.uk/online-safety/protecting-children/age-assurance


### Correction: Didit node type for profile-photo matching (21 October 2026)
The generic Face Match node visible in the supplied Didit workflow screenshot is document-oriented and should not be assumed to compare a selfie against a dating profile photo. Didit's Face Match 1:1 product describes live-selfie-to-document-portrait matching. Didit separately documents a **Biometric Authentication Workflow**: liveness, then face matching between the `portrait_image` supplied when creating the session and the current selfie, with optional IP analysis. Source: https://docs.didit.me/reference/workflows
Therefore, do not tell implementation to use the generic document Face Match node for Pinchd's Hinge-style badge. First confirm that the Didit account/workflow builder supports the Biometric Authentication Workflow and that Pinchd can securely pass the user's existing profile photo as `portrait_image`. If this is only available through an API/session-creation integration, configure the flow there rather than forcing a document-based workflow. No identity document should be required for voluntary profile verification.


### Didit workflow documentation re-check — correction and final configuration (21 October 2026)
Official Didit Workflows documentation identifies four workflow types. The relevant documented type is **Biometric Authentication Workflow**: (1) Liveness detection, (2) face matching between the `portrait_image` sent when creating the session and the current selfie, and (3) optional IP analysis. Source: https://docs.didit.me/reference/workflows
The standard **KYC Workflow** instead starts with ID Verification and can optionally add Liveness and Face Match between the document portrait and selfie. Source: https://docs.didit.me/reference/workflows and https://docs.didit.me/reference/face-match-11
Important limitation: the documentation describes Biometric Authentication as returning-user authentication and does not explicitly guarantee that every account/workflow configuration supports an arbitrary dating-profile photo as the `portrait_image`. Before building production verification around profile photos, confirm support with Didit and test session creation. If supported, use the Biometric Authentication Workflow, pass the user's selected current profile photo as `portrait_image` via server-side session creation, retain Liveness as the first step, leave IP Analysis off initially, and configure final outcomes. Do not use the KYC document Face Match node, do not add duplicate Face Match, AML, or age estimation to this optional badge flow. If arbitrary profile-photo reference is not supported, ask Didit whether Face Match API can compare a live capture with an uploaded profile photo, or select a provider that explicitly supports that use case.
Cost note: Didit's current public price page lists Biometric Authentication at $0.10/auth; standalone passive liveness is $0.10, Face Match 1:1 is $0.05, IP Analysis $0.03, and AML $0.20. Do not add these prices together for a Biometric Authentication module unless the console bills separate modules in the configured flow; verify the actual account quote and billing semantics. Source: https://didit.me/pricing/
