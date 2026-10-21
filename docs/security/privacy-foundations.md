# Pinchd security and privacy foundations

This is a starting checklist, not a completed security review or legal assessment.

## Before collecting real user data

- Create a data inventory and data-flow diagram.
- Identify purpose, lawful basis, retention, access roles, processors and deletion behavior for every data category.
- Assess whether a DPIA is required; give extra scrutiny to location, relationship preferences, profiling, verification and moderation evidence.
- Minimise collection and avoid inferring protected/sensitive traits from photos or message content for compatibility.
- Define export, correction, account deletion and downstream deletion workflows.
- Keep secrets in a managed secret store; use MFA and least privilege for staff.
- Use TLS in transit and appropriate encryption at rest; manage keys separately.
- Redact message bodies, exact location, identity documents, tokens and sensitive attributes from logs and analytics.
- Threat-model account takeover, scraping, stalking, location triangulation, impersonation, scams, harassment, credential stuffing and insider access.
- Test authorisation server-side for every profile, photo, conversation, report, admin action and paid entitlement.
- Establish incident response, backup/restore tests, vulnerability handling and security review gates.

## Mobile security baseline

Use OWASP MASVS / MASTG as the mobile security baseline. Check secure storage, authentication/session handling, network communication, platform permissions, code/dependency integrity, privacy controls, and resilience appropriate to the threat model.

## Location

Store only the precision required for the feature. Prefer distance bands or approximate location for discovery; never expose exact coordinates to other users. Prevent repeated-query triangulation and define location retention.

## Verification and biometrics

Do not implement face matching or biometric verification until the purpose, legal basis, applicable special-category condition, vendor processing, retention, deletion, accessibility, false-positive handling and DPIA decision have been reviewed and documented.

## Release gate

No production launch until a security review, data-protection review, moderation readiness review, deletion/export test, incident plan, and backup recovery test have named owners and recorded outcomes.
