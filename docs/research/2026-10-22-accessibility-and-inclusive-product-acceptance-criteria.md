# Pinchd — Accessibility and Inclusive Product Acceptance Criteria
**Date:** 22 October 2026  
**Status:** Requirements baseline for later design and engineering; not a legal compliance certification.  
**Scope:** Native iOS/Android accessibility, inclusive language, profile and matching inclusivity, safety/reporting, media, age assurance, and test strategy. Mobile UI design remains paused.

## 1. Standards and approach
Use the four WCAG principles—perceivable, operable, understandable and robust—as an engineering framework. WCAG 2.2 is a W3C Recommendation; its newer criteria include focus visibility, minimum target size, alternatives to dragging, consistent help, avoiding redundant entry and accessible authentication. Native app conformance and legal obligations need a separate applicability review; do not describe a native app as WCAG-conformant solely from this checklist.

Primary references:
- W3C WCAG 2.2: https://www.w3.org/TR/WCAG22/
- W3C WCAG overview: https://www.w3.org/WAI/standards-guidelines/wcag/
- W3C WCAG 2.2 changes: https://www.w3.org/WAI/standards-guidelines/wcag/new-in-22/
- Apple Human Interface Guidelines — Accessibility: https://developer.apple.com/design/human-interface-guidelines/accessibility/
- Apple Human Interface Guidelines — VoiceOver: https://developer.apple.com/design/human-interface-guidelines/voiceover/
- Apple Human Interface Guidelines — Inclusion: https://developer.apple.com/design/human-interface-guidelines/inclusion/
- Apple App Review Guidelines: https://developer.apple.com/app-store/review/guidelines/

## 2. Product-level accessibility principles
- Every essential task must be completable without voice, video, audio, gestures that require dragging, or optional AI.
- Essential actions cannot rely on colour, position, animation, sound or an icon without a programmatic label.
- Support OS accessibility settings rather than overriding them for brand consistency.
- Offer clear text alternatives for media and non-text status. Never require a user to disclose disability to access an alternative.
- Use plain language, consistent terminology, predictable navigation and actionable error messages.
- Respect reduced-motion settings; avoid flashing, rapid transitions and unnecessary parallax.
- Do not treat disability, assistive-technology use, atypical speech, short answers or missing optional media as low desirability.
- Accessibility must include safety and consent flows, not only onboarding and profile screens.
- Test with disabled people and assistive-technology users, compensating them appropriately; automated checks alone are insufficient.

## 3. Measurable acceptance criteria

### A. Visual and text presentation
- Support system text scaling, including large accessibility sizes, without clipping, overlapping, truncating essential controls or hiding safety actions.
- Verify contrast against the applicable target for text, icons, boundaries and states; never use colour as the sole state indicator.
- All functional images and controls have meaningful accessible names; decorative images are ignored by screen readers.
- Profile photo alternatives are handled carefully: provide meaningful labels for controls and content structure without inventing descriptions of a person's traits or exposing private information.
- Text can reflow where platform patterns support it; error messages and button labels remain visible at large text settings.
- Dynamic content changes (new match, error, verification result) are announced without moving focus unexpectedly.
- Dark mode, increased contrast and colour-filter settings remain usable.

### B. Screen readers and semantics
- All core flows are operable with VoiceOver on iOS and TalkBack on Android.
- Focus order follows reading/action order; no keyboard/screen-reader traps.
- Custom controls expose correct role, label, value, selected/disabled state and available actions.
- Photo carousels, prompt responses, incoming-like queues and match lists expose position and navigation controls without requiring swipes that screen readers intercept.
- A like, pass, block, report, unmatch, pause and delete action has an explicit spoken label and a confirmation/result announcement appropriate to its risk.
- Errors identify the affected field and a corrective action. Focus moves to the error summary only when doing so improves recovery.
- Status messages such as “profile hidden” or “message not sent” are programmatically announced and visible.

### C. Motor and alternative input
- Core workflows do not require drag-and-drop; every reorder or media selection has a button-based alternative.
- Touch targets meet platform guidance and are spaced to reduce accidental activation; test WCAG 2.2 target-size principles where applicable.
- Support external keyboard/switch access where platform frameworks provide it; focus indication remains visible and unobscured.
- Time-limited steps provide adequate time and a clear extension/retry route unless a genuine security need is documented.
- High-impact actions such as delete, block and unmatch are not triggered by a single accidental swipe; provide a clear action and recovery/confirmation policy.
- Biometric or age-assurance failure has an accessible alternative or appeal route when permitted by the approved risk model.
### D. Cognitive, language and emotional accessibility
- Use plain English and short, action-oriented labels. Explain “like”, “match”, “pass”, “pause”, “block”, “unmatch” and “delete” consistently.
- Explain relationship-goal options with neutral examples and distinguish “not sure yet” from “open-minded/exploring” and “undisclosed”.
- Do not force users through long, dense onboarding before they understand data use, age assurance and safety controls.
- Preserve entered profile data after recoverable validation errors; avoid redundant re-entry.
- Give clear progress and recovery instructions for slow or failed age checks, uploads, purchases and deletion.
- Do not use shame, countdowns, streaks, guilt, or urgency to pressure likes, replies, subscriptions or profile disclosure.
- Make “no thanks”, skip, close and decline paths as legible as positive actions.
- Allow users to pause or leave without punitive messaging; support healthy exit.

### E. Audio, video and optional AI
- Text remains a complete baseline profile format. Voice, polls and AI coaching are optional enhancements.
- Any meaningful audio/video content must have an accessible text alternative or equivalent content path; do not make speech the sole carrier of essential information.
- Caption instructional or explanatory video; provide transcripts for user-facing instructional audio.
- Voice recording and playback have visible controls, clear recording state and accessible error/permission messaging.
- AI coaching is opt-in, editable and user-reviewed; it never auto-publishes or auto-sends. Users can ignore suggestions without penalty.
- AI must not assume standard grammar, neurotypical expression, cultural familiarity, attractiveness or desirability. Test output quality across language styles and disability-related communication differences.
- Clearly identify generated suggestions and avoid presenting an inference as a fact about the person.

### F. Gender, relationship goals and inclusion
- Store gender identity separately from discovery preferences. Do not make identity fields mandatory beyond documented necessity.
- Avoid default cis/trans exclusions; apply the approved gender preference policy consistently.
- Support relationship-goal combinations without moral judgement; missing values do not imply openness.
- Avoid gendered defaults in onboarding and marketing copy. Use neutral, respectful language where possible.
- Do not infer or expose protected traits from names, photos, location, prompts or behavior.
- Ensure users can understand and correct their own profile representation; provide reporting/support for misrepresentation or targeted abuse.
- Test for different body sizes, skin tones, ages within the adult product scope, gender expressions, accents and communication styles without scoring them as desirability.

### G. Safety and high-impact flows
- Reporting, blocking, muting where offered, unmatching, pausing and account deletion must be reachable without relying on gestures alone.
- Block and report controls remain accessible from relevant profile, message and match contexts, with consistent labels.
- Safety confirmation copy explains consequences plainly without exposing confidential investigation methods.
- Do not require a user to confront, message or notify the reported person in order to report.
- A user who cannot complete a standard age-assurance flow needs accessible vendor support and an appeal route; no inaccessible alternative may create an age-gate bypass.
- Safety notices should not depend on sound or colour alone; urgent instructions use text and clear action.
- Essential safety features are free and cannot be hidden behind inaccessible subscription UI.
## 4. Platform and device test matrix

| Test area | Minimum test coverage | Acceptance evidence |
|---|---|---|
| iOS screen reader | VoiceOver through registration, profile, discovery, likes, messages, block/report, unmatch, pause and deletion | No unlabeled core controls; logical focus order; status/error announcements |
| Android screen reader | TalkBack through the same flows | Equivalent task completion and semantics |
| Text scaling | Default, large accessibility sizes and bold text on both platforms | No critical clipping/overlap; all safety controls available |
| Display and contrast | Light/dark, increased contrast, colour filters, grayscale | States remain distinguishable without colour alone |
| Motion | Reduce Motion enabled; low-power/animation settings where applicable | No essential information depends on motion |
| Motor access | Switch/keyboard/accessibility menu where supported; one-handed and tremor scenarios | Core tasks without drag-only controls or precision gestures |
| Hearing | Audio/video profile playback and system sound off | Text alternative; no audio-only critical status |
| Cognitive/language | Plain-language comprehension, error recovery, low literacy and translation expansion | Users can explain key states and recover from common errors |
| Age assurance | Large text, screen reader, timeout, denial, retry, vendor outage and appeal | No inaccessible dead end; no eligibility bypass |
| Billing and deletion | Accessible purchase/restore/cancel guidance and in-app deletion | Clear cost/state; deletion can be initiated and status understood |
| Assistive tech combinations | Screen reader + large text; switch + voice control where available | Combined settings do not make critical flows unusable |

## 5. Automated and manual quality gates
- Run static accessibility checks in CI where available, but do not treat passing automation as conformance.
- Add component-level tests for labels, roles, values, focus order, error announcement and state changes.
- Manual QA uses real devices and current OS versions; document model, OS, app build, assistive technology, tester and result.
- Any P0/P1 accessibility defect in age assurance, consent, blocking/reporting, messaging permission, account recovery or deletion blocks release.
- P2 defects need an owner, mitigation and dated fix; severity definitions must be documented and applied consistently.
- Conduct task-based tests with disabled participants before public launch and after major navigation changes.
- Maintain an accessibility statement/contact path and a simple way to report an accessibility problem.
- Include accessibility regressions in release notes, incident review and product metrics.

## 6. Requirements traceability IDs
- A11Y-01: Core journeys work with VoiceOver and TalkBack.
- A11Y-02: Large text and display accommodations do not hide critical content.
- A11Y-03: No essential action relies on colour, sound, animation or gesture alone.
- A11Y-04: Text profile content is sufficient; optional media never blocks use.
- A11Y-05: Consent, report, block, unmatch, pause and deletion are accessible.
- A11Y-06: Age-assurance failure/retry/appeal has an accessible path without weakening the gate.
- A11Y-07: Plain language and consistent states distinguish like/match/message permission.
- A11Y-08: AI coaching is optional, editable, user-reviewed and non-discriminatory in testing.
- A11Y-09: Automated checks plus manual device and participant testing are recorded.
- A11Y-10: Critical accessibility defects block release.

## 7. Open gates
1. Confirm legal accessibility obligations for the company, service, markets and native apps with qualified counsel; this document is not legal advice.
2. Decide the formal target standard and audit scope for the app and any marketing/support web pages.
3. Define supported OS versions and a practical device/assistive-technology matrix.
4. Recruit disabled testers and establish compensation, consent and privacy practices.
5. Verify any third-party age-assurance and billing SDKs with assistive technologies.
6. Test profile-photo descriptions and user-generated media policies with people with visual impairments and privacy reviewers.
7. Add these requirements to the eventual design and engineering acceptance process only after research sign-off.

## 8. Source notes
- W3C describes WCAG 2.2 around perceivable, operable, understandable and robust content, with testable A/AA/AAA success criteria: https://www.w3.org/WAI/standards-guidelines/wcag/
- WCAG 2.2 added criteria including focus not obscured, target size, dragging alternatives, consistent help, redundant entry and accessible authentication: https://www.w3.org/TR/WCAG22/
- Apple's accessibility guidance covers system text enlargement and use of Dynamic Type; VoiceOver guidance calls for meaningful labels and descriptions: https://developer.apple.com/design/human-interface-guidelines/accessibility/ and https://developer.apple.com/design/human-interface-guidelines/voiceover/
- Apple inclusion guidance emphasises supporting built-in accessibility features and avoiding assumptions about disability: https://developer.apple.com/design/human-interface-guidelines/inclusion/

No UI design is created by this document. It defines acceptance criteria to apply when the research phase has concluded.
