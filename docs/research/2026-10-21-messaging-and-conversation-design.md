# Pinchd research: messaging and conversation design
**Research date:** 21 October 2026  
**Status:** Desk research; mobile UI design remains paused pending completion of the research and requirements package.

## Executive summary
Pinchd should optimise for conversations that feel authentic, reciprocal, respectful and useful in helping people decide whether to meet—not for message volume, instant replies or time spent in the app. Product features should reduce the friction of starting and sustaining a conversation while protecting user agency, privacy and safety.

## Evidence review

### 1. Profile-specific openings are more promising than generic prompts
Hinge reported that in its 2024 internal data, likes on text prompts were 47% more likely to lead to a date than likes on photos. Hinge later described testing tailored conversation starters attached to a person's prompts and photos to make thoughtful first messages easier. These are company-reported findings and product rationale; they are not independent evidence and may not generalise to Pinchd.
- Source: Hinge, “Prompt Feedback,” 15 January 2025: https://hinge.co/newsroom/prompt-feedback
- Source: Hinge, “Evolving Together: How Daters Helped Shape Hinge in 2025,” 10 November 2025: https://hinge.co/newsroom/hinge-2025-product-evolution

**Implication:** Make it easy to respond to a specific, voluntarily shared profile detail. Suggestions should be optional, editable, varied and clearly anchored to the profile. Avoid generic copy-paste lines, fabricated familiarity or AI-written messages sent automatically. Do not imply that a suggested opener predicts compatibility.

### 2. Authenticity matters more than a “perfect” AI-written message
Hinge's March 2026 guidance describes users using AI as a second opinion or to check how a message may come across, while emphasising that the message should still sound like the sender.
- Source: Hinge, “Hinge’s Guide to Using AI in Dating,” 5 March 2026: https://hinge.co/newsroom/ai-dating-guide

**Implication:** If Pinchd offers AI assistance, position it as optional coaching: brainstorm a question, make wording clearer, or offer a tone check. The user must review and explicitly send the final message. Do not auto-send, impersonate the user, or encourage manipulative scripts. Offer a “keep my wording” path and avoid storing message drafts longer than necessary.

### 3. Communication pressure and dating fatigue are real design risks
Hinge's 2025 updates included chat-specific notification controls, giving users a way to prioritise selected conversations. Its 2025 Gen Z report says respondents wanted deeper emotional intimacy but often felt hesitant to begin deeper conversations; it also reported that follow-up questions and questions about interests or values were commonly welcomed. These are self-selected platform survey findings, not population-wide estimates.
- Source: Hinge, 2025 product evolution: https://hinge.co/newsroom/hinge-2025-product-evolution
- Source: Hinge, “Gen Z D.A.T.E. Report,” 19 November 2025: https://hinge.co/en-gb/newsroom/2025-GenZ-Report

A 2024 qualitative paper on mobile-online-dating fatigue discusses users' experiences of superficial exchanges, ghosting, hurtful communication, violated boundaries and perceived unequal treatment. A 2026 study of 881 Chinese mobile dating app users found dating fatigue was positively associated with ghosting; its survey design and country context mean it should not be treated as causal proof or assumed to transfer directly to UK users.
- Source: “Coping with mobile-online-dating fatigue and the negative self-fulfilling prophecy of digital dating,” SN Social Sciences: https://link.springer.com/article/10.1007/s43545-024-01042-0
- Source: Yin, Zhou & Peng, “Online Dating Paradox: Why Lonely Hearts Choose to Ghost,” Computers in Human Behavior, online 9 October 2026: https://www.sciencedirect.com/science/article/pii/S0747563226003109

**Implication:** No streaks, guilt nudges, “they're waiting for you” pressure, response countdowns or penalties for slow replies. Provide per-conversation notification controls, mute, pause dating, and clear ways to end or unmatch a conversation. Never treat reply speed alone as a measure of desirability or relationship intent.

### 4. Clarity can help people move from a good date to a second date
Hinge's July 2025 survey of more than 3,000 global daters reported that 49% had held back from following up after a good first date due to fear of seeming too eager. It reported that 75% expected a follow-up on the same day or the next day, and 47% preferred a message that expressed interest in meeting again. These are Hinge Labs survey results, not universal etiquette rules.
- Source: Hinge, “The Follow-Through Formula,” 16 July 2025: https://hinge.co/newsroom/follow-through-formula

**Implication:** Offer optional, low-pressure post-date message templates that express interest honestly, ask about mutual interest, or politely close the loop. Never send these without user review and consent. Avoid implying that a follow-up is owed, that one timing rule fits everyone, or that non-response means bad intent.

### 5. Respectful messaging requires both prevention and accountable review
Hinge reports using pre-send reflection prompts for potentially offensive content, recipient feedback on flagged messages, custom comment filters and a combination of automated detection and trained human moderation.
- Source: Hinge, 2025 product evolution: https://hinge.co/newsroom/hinge-2025-product-evolution
- Source: Hinge, “How are interactions with other members processed?” updated 23 December 2025: https://help.hinge.co/hc/en-us/articles/360060794853-How-are-interactions-with-other-members-processed-by-Hinge

**Implication:** Build a proportionate safety layer: user-controlled filters, report/block access within the conversation, contextual warnings for clearly risky content, and human review for serious or ambiguous cases. Warnings should encourage reconsideration rather than shame. Automated systems should not claim to know intent, and a report should not be treated as proof by itself. Keep blocking effective across discovery, likes, notifications and cached content.

### 6. Follow-through signals need careful interpretation
Hinge's Signals feature describes recent behaviours such as thoughtful participation, responding to interest and follow-through, while explicitly stating that the badge does not guarantee responsiveness, compatibility or relationship outcomes.
- Source: Hinge Help Centre, “What is Signals?”, updated 29 April 2026: https://hinge.co/help.hinge.co/hc/en-us/articles/49858544145555-What-is-Signals

**Implication:** Pinchd should not launch a public “good dater” score based on message frequency or response speed. If future trust or reliability signals are considered, publish what they measure, avoid making promises they cannot support, allow correction/appeal where appropriate, and assess bias against people with different schedules, disabilities, communication styles or caregiving responsibilities. Do not infer relationship quality from app activity alone.

## Proposed Pinchd requirements

### Conversation start
- **MSG-01:** A user can start a conversation by referencing a specific, visible profile prompt, photo caption or shared interest.
- **MSG-02:** Suggested openers are optional, editable, and never sent without explicit user action.
- **MSG-03:** Suggestions must not invent shared experiences, personal facts or emotional intimacy.
- **MSG-04:** Offer varied opening styles (curious, playful, direct) without stereotyping by gender, sexuality, ethnicity or age.
- **MSG-05:** Users can send an ordinary message without using AI or a suggested template.

### Conversation quality and agency
- **MSG-06:** Notification settings include global controls and per-conversation mute/prioritisation.
- **MSG-07:** No streaks, response countdowns, punitive inactivity scores or guilt-based nudges.
- **MSG-08:** Users can unmatch, block and report from the conversation with clear, understandable consequences.
- **MSG-09:** Users can pause discovery and control notifications without deleting their account.
- **MSG-10:** Any AI writing assistance is opt-in, user-reviewed and user-sent; it is not presented as a compatibility prediction.
- **MSG-11:** Any post-date follow-up or closure templates are optional and editable; no automated sending.
- **MSG-12:** Do not expose read receipts or online status by default without a documented user benefit, privacy review and user control.

### Safety, privacy and moderation
- **MSG-13:** Report and block controls are available at every stage of the conversation.
- **MSG-14:** A block suppresses the blocked account from future discovery, likes, notifications and conversation surfaces, subject to legally necessary record retention.
- **MSG-15:** Message scanning, if used, has a defined safety purpose, data-minimisation rules, retention limits and a transparent privacy explanation.
- **MSG-16:** Safety classifiers trigger proportionate interventions; they do not independently determine romantic desirability or compatibility.
- **MSG-17:** Users can understand why content was warned, filtered or removed and have an appropriate appeal path.
- **MSG-18:** Moderation operations define severity levels, triage targets, escalation routes, evidence-access controls and staff wellbeing safeguards.

## Measurement plan (post-launch; no interviews required as a prerequisite)
Track cohort-level measures and segment responsibly to detect unequal effects:
- **Conversation initiation:** percentage of matches with at least one message within 24 hours and 7 days; report by user cohort and acquisition channel.
- **Reciprocity:** proportion of conversations with meaningful two-way exchange, using carefully defined, privacy-reviewed event metadata rather than analysing intimate message content by default.
- **Conversation-to-date progression:** optional self-reported “met in person” and “would meet again” events, clearly labelled as voluntary and incomplete.
- **User agency:** use of notification controls, mute, pause, unmatch and closure features; user-reported usefulness where available.
- **Safety:** reports per 1,000 active conversations, report reason distribution, time to triage/action, recurrence, appeal reversal rate and block effectiveness.
- **Wellbeing and fatigue proxies:** self-reported fatigue where voluntarily collected, notification disablement, pause usage and post-session satisfaction. Treat proxies as ambiguous, not diagnoses.
- **AI quality:** suggestion adoption/edit/reject rates, user-rated authenticity, safety issues and outcome differences. High acceptance alone is not success.
- **Fairness:** compare outcomes across relevant cohorts where lawful and privacy-safe; do not infer sensitive traits from photos or message content.

## Experiment hypotheses
1. Profile-specific optional openers improve first-message initiation versus a generic blank composer, without increasing reports or reducing perceived authenticity.
2. Per-conversation notification controls improve self-reported control and reduce unwanted notification burden without materially harming reciprocal conversations.
3. Optional, user-reviewed follow-up/closure templates improve clarity and reduce uncertainty without increasing pressure or unwanted contact.
4. A contextual pre-send reflection prompt reduces harmful messages without disproportionately blocking benign communication, humour or reclaimed language.

Use controlled experiments with pre-registered primary outcomes, guardrail metrics and stopping criteria. Do not optimise solely for messages sent, session duration or conversion. Treat low-volume safety outcomes cautiously; supplement experiments with operational review.

## Limitations and open questions
- Most product metrics cited above come from Hinge and may reflect its audience, definitions and commercial incentives.
- Several survey findings are self-reported; observational relationships do not establish causation.
- The 2026 ghosting study used a Chinese sample; UK transferability is uncertain.
- Pinchd-specific effect sizes cannot be known from desk research. They must be measured after launch.
- Messaging content is highly sensitive. Any content analysis or AI assistance needs a separate privacy, security, lawful-basis and retention assessment before implementation.
- Do not begin mobile screen design until the research/requirements package is complete.

## Decision summary
**Recommended direction:** thoughtful profile-specific openers; optional AI coaching with human review; user-controlled notifications; low-pressure follow-up/closure support; accessible block/report controls; proportionate, transparent moderation; and outcome measures centred on reciprocal, safe connection rather than engagement volume.

## Source register
1. Hinge, Prompt Feedback (15 Jan 2025): https://hinge.co/newsroom/prompt-feedback
2. Hinge, Evolving Together (10 Nov 2025): https://hinge.co/newsroom/hinge-2025-product-evolution
3. Hinge, Gen Z D.A.T.E. Report (19 Nov 2025): https://hinge.co/en-gb/newsroom/2025-GenZ-Report
4. Hinge, AI in Dating Guide (5 Mar 2026): https://hinge.co/newsroom/ai-dating-guide
5. Hinge, Follow-Through Formula (16 Jul 2025): https://hinge.co/newsroom/follow-through-formula
6. Hinge, Signals Help Centre (29 Apr 2026): https://hinge.co/help.hinge.co/en-us/articles/49858544145555-What-is-Signals
7. Hinge, interaction processing (updated 23 Dec 2025): https://help.hinge.co/hc/en-us/articles/360060794853-How-are-interactions-with-other-members-processed-by-Hinge
8. “Coping with mobile-online-dating fatigue and the negative self-fulfilling prophecy of digital dating,” SN Social Sciences: https://link.springer.com/article/10.1007/s43545-024-01042-0
9. Yin, Zhou & Peng, “Online Dating Paradox: Why Lonely Hearts Choose to Ghost,” Computers in Human Behavior (online 9 Oct 2026): https://www.sciencedirect.com/science/article/pii/S0747563226003109


## 10. Profile prompts and conversation-support research (21 October 2026)

### Evidence summary
Hinge's public feature set includes structured text prompts, optional voice prompts, prompt polls, private prompt coaching, and profile-grounded Convo Starters. Convo Starters provides several ideas based on a specific profile photo or prompt; users write the actual message themselves. Hinge announced UK availability in February 2026. Sources: https://hinge.co/newsroom/profile-refresh-guide and https://hinge.co/newsroom/convo-starters

Hinge reports that 72% of surveyed daters were more likely to consider someone whose like includes a message and that likes with comments were twice as likely to lead to a date. These are company-reported associations and should not be treated as independent causal proof. Its 2025 Prompt Feedback materials also report that users struggle to know what to write; numbers differ across publications (63% in its January launch announcement versus nearly 40% in a November explanation), likely reflecting different questions/timing. Do not merge these into one measure without source survey methods. Sources: https://hinge.co/newsroom/prompt-feedback and https://hinge.co/en-gb/newsroom/about-prompt-feedback

Community discussions on Reddit often criticise cliché, generic claims, negative/ultimatum-style prompts and answers that give another person nothing specific to respond to. Other commenters disagree about ideal answer length and the importance of prompts versus photos. These discussions are self-selected anecdotes, not representative evidence; the disagreement argues against rigidly scoring users by a single profile-writing formula. Sources: https://www.reddit.com/r/hingeapp/comments/1pg27hn/a_ruthless_ranking_of_the_2025_hinge_prompts/ and https://www.reddit.com/r/hingeapp/comments/143yedd

Research on online self-presentation describes the tension between presenting an ideal self and trying to be authentic, while newer qualitative work notes that people may want to express personality while protecting identifying information. Voice-based interaction may offer a different sense of presence, but available qualitative studies do not prove voice prompts improve outcomes for UK users. Sources: https://onlinelibrary.wiley.com/doi/10.1111/j.1083-6101.2006.00020.x and https://arxiv.org/abs/2502.11430 and https://arxiv.org/abs/2402.19328

### Requirements recommended for Pinchd
- Prompt topics should elicit concrete details, stories, preferences, values-in-action, everyday routines and shared-activity ideas—not only generic traits or lists of demands.
- Provide a balanced choice of playful, practical, reflective and relationship-goal-inclusive prompts. No user should have to be witty, extroverted, highly verbal or disclose trauma to create a compelling profile.
- Optional coaching should help the user add specificity while preserving their own voice. It must be private, editable, rejectable and never auto-publish.
- Optional conversation inspiration should be tied to the profile content the other user chose to share. It should offer ideas, not ready-to-send impersonation; users write, edit and send the message themselves.
- Support text-only participation. Voice, video and polls can be optional expression modes, never requirements for a complete usable profile or prerequisites for core matching.
- Do not promise that a prompt, AI suggestion or comment will guarantee a match or date. Do not penalise short answers or communication styles through hidden desirability scoring.
- Let users edit, replace and delete prompt answers; ensure profile preview accurately reflects what others see.
- Avoid prompting unnecessary sensitive disclosures, identifying details about third parties, or intimate history that users may later regret sharing.
- Conversation starters should encourage a specific observation, open question, shared perspective or light exchange. Avoid pressure, sexualised assumptions, intrusive questioning or generic copy-and-paste scripts.
- AI suggestions are opt-in, clearly suggestions, grounded in visible profile content, user-reviewed and never auto-sent. Provide a setting to disable them.

### Metrics and safeguards
Measure profile completion friction, prompt answer edits, user-rated usefulness of coaching, likes with comments, match-to-first-message conversion, conversation starts and user-reported relevance. Treat these as product diagnostics, not proof of relationship quality. Monitor whether coaching disproportionately benefits certain writing styles or users, and ensure no ranking penalty is attached to declining AI help, voice or polls. Do not optimise for message volume alone; include reports, blocks, unwanted sexual content, response pressure and user comfort.

### Conclusion
Prompts should make it easier for users to express themselves and easier for others to respond meaningfully, while keeping authenticity and agency with the person. Pinchd should adopt optional, profile-grounded conversation support, but reject forced vulnerability, AI impersonation, generic formulaic profiles and guaranteed-outcome claims.
