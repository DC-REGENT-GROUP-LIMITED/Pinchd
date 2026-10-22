# ADR-0002: Technical feasibility and provisional architecture
**Status:** Superseded for stack selection by ADR-0001; retained as supporting feasibility checklist  
**Date:** 21 October 2026  
**Related:** `docs/product/consolidated-requirements-v1.md`, `ADR-0001-stack-and-deployment.md`

## Context
Pinchd needs iOS and Android clients, a restricted admin console, a backend for matching and messaging, sensitive-data controls, reliable safety workflows, and subscription handling. The repository is a scaffold, not a running production application. The requirements are domain-heavy but do not currently justify independently deployed microservices.

## Recommendation
1. **Mobile:** evaluate Expo / React Native with TypeScript as the leading option for shared iOS and Android delivery, subject to proof-of-concept tests for push notifications, secure storage, media upload, deep links, accessibility, background behavior and any age-assurance SDK.
2. **Admin console:** separate web application with strong authentication, MFA, role-based access, audit trails and no broad access to intimate data by default.
3. **Backend (historical proposal, superseded):** TypeScript modular monolith with explicit domain modules and server-authoritative policy enforcement. Use Go per ADR-0001; retain the domain-boundary and policy-enforcement requirements.
4. **Primary database:** managed relational PostgreSQL for transactional entities (accounts, preferences, likes, matches, conversations, reports, subscriptions, audit references). Evaluate PostGIS only if spatial queries justify it; coarse location may be enough initially.
5. **Media:** private object storage with short-lived signed access, validated uploads, image processing, metadata stripping and moderation state. Do not use public permanent media URLs for private profiles.
6. **Async work:** managed queue for notifications, media processing, report workflows and deletion propagation only where asynchronous processing is needed. Define retries, idempotency, dead-letter handling and observability.
7. **Caching/search:** add Redis or a dedicated search system only when measured latency or query requirements justify it. Cache keys and invalidation must respect blocks, account deletion and privacy.
8. **Hosting:** choose a managed UK/EU-compatible deployment model after a data-flow, processor and transfer review; do not assume hosting region alone settles UK GDPR compliance.
9. **Billing:** use platform-compliant purchase flows; validate receipts/notifications server-side and reconcile entitlements. Keep billing status distinct from safety access.
10. **Observability:** managed logs, metrics, traces and alerting with strict redaction. Separate audit events from product analytics.
11. **Architecture:** avoid microservices until scale, team ownership, reliability isolation or independent deployment provide a concrete benefit.

## Domain boundaries
- Identity and account lifecycle
- Profiles and media
- Preferences and eligibility
- Discovery and matching
- Likes and matches
- Messaging and notifications
- Trust, safety and moderation
- Age assurance and verification
- Billing and entitlements
- Privacy rights and data lifecycle
- Analytics and experimentation
- Admin and support

Domains own their business rules. Transport/API controllers should not contain the authoritative matching, billing or moderation policy. Cross-domain actions use explicit interfaces and typed contracts.

## Feasibility risks and proof-of-concept gates
| Risk / unknown | Proof required before stack freeze |
|---|---|
| Age-assurance vendor SDK or flow | Confirm platform support, privacy model, failure/retry paths, accessibility and data returned |
| Mobile push and notification controls | Verify token lifecycle, per-conversation preferences and privacy-safe notification payloads |
| Media and intimate content safety | Test private storage, upload validation, metadata stripping, access revocation and moderation states |
| Block propagation | Prove block suppresses discovery, likes, messages, notifications, cached data and queued jobs |
| Account deletion | Prove deletion/retention jobs across primary DB, object storage, search/cache, analytics and providers |
| Subscription entitlements | Test purchase, restore, renewal, expiry, refund, server notification and cross-device consistency |
| Matching query performance | Build synthetic-data benchmark for eligible candidate generation, reciprocal filters and pagination |
| Moderator console | Test role separation, case evidence permissions, audit logs, dual review and appeal actions |
| Data residency and vendor terms | Complete vendor/subprocessor inventory, data-flow map, contract and transfer review |
| Launch-scale economics | Compare managed service costs for conservative beta and growth scenarios |

## Security and privacy requirements
- Server-side authorization for every user-owned resource and admin action.
- MFA and least privilege for staff; audit access to sensitive evidence.
- TLS, appropriate encryption at rest, managed secrets and key rotation.
- No raw message text, exact coordinates, ID documents or secrets in logs/analytics.
- Private media by default; short-lived access and revocation.
- Rate limits for registration, recovery, likes, messaging, uploads and report abuse.
- Explicit retention and deletion for data, logs, backups and vendors.
- DPIA/legal review before high-risk profiling, biometric processing or novel AI.
- No internet-wide face recognition or sensitive-trait inference.
- Keep safety/risk scores separate from romantic ranking.

## Operational requirements
- Staging and production separated; production data not copied to development by default.
- Versioned migrations and tested rollback/forward-recovery strategy.
- Backups with periodic restoration exercises.
- CI checks for formatting, types, unit tests, dependency and secret scanning, contract tests and build verification.
- Release checklist includes privacy, safety, accessibility, security, data migration and support readiness.
- Alert ownership and on-call process for critical messaging, blocking, reports, authentication and billing events.
- Vendor outage behavior documented; safety-critical actions should fail safely.

## Alternatives considered
- **Fully native Swift/Kotlin:** strongest platform-specific control but duplicates implementation and maintenance.
- **Flutter:** credible cross-platform option; compare team skills, accessibility, vendor SDKs and long-term maintenance.
- **Microservices from day one:** rejected provisionally due to operational overhead and unknown traffic/team structure.
- **NoSQL-first data model:** not preferred for the transactional core because preferences, matches, billing and moderation actions need relational constraints and auditability; specialised stores may be added for measured needs.
- **Self-hosted infrastructure:** not preferred for the initial phase unless control requirements or economics outweigh the operational burden.

## Decision criteria
Evaluate team familiarity, SDK compatibility, accessibility, testability, security support, performance, deployment complexity, data portability, vendor lock-in, total cost, observability, backup/restore, operational support and ability to meet launch legal/privacy obligations.

## Acceptance gate before implementation
Record the final decision only after:
- Team and maintenance capability are confirmed.
- A small technical spike validates cross-platform needs.
- Hosting/data-flow and vendor privacy reviews are complete.
- Database and API ownership boundaries are agreed.
- Build/test/CI/deployment strategy is documented.
- Security, privacy and operational risks have owners and mitigations.
- The decision record includes rejected alternatives and review triggers.

## Consequences if approved
The architecture will be simple enough for a small team while preserving domain boundaries. It still requires explicit implementation decisions for framework versions, hosting, database provider, media storage, queueing, observability, authentication, age-assurance vendor and billing. This ADR does not authorise production code or settle the open product decisions.

## Review triggers
Revisit when team size, traffic, reliability needs, regulatory scope, vendor capabilities, native SDK requirements or unit economics materially change.


## Product decision dependency

This ADR remains **Proposed; decision required before implementation**. It describes a candidate architecture, not a selected stack or completed proof-of-concept. Close the MVP and requirements reconciliation first; then use the proof gates above before approving implementation. See `../../research/2026-10-21-workstream-1-reconciled-decision-baseline.md`.


## Approved direction update — 2026-10-22

The user has approved the stack direction in ADR-0001. This ADR's previous TypeScript-backend and “avoid microservices until scale” recommendation is superseded as the default language decision, but its domain-boundary, proof-of-concept, safety, privacy, cost and operational gates remain in force. Go is the backend language direction; design clear domain boundaries and allow independently deployable services when justified, without multiplying services solely for appearance. React Native + Expo + TypeScript, AWS, PostgreSQL, private S3 media and managed async processing when useful are the approved direction. Multi-region database is deferred; start with one primary AWS region and Multi-AZ managed database resilience in production where service choice and budget support it. Refer to ADR-0001 for cost assumptions and implementation gates.
