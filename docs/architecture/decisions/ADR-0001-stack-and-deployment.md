# ADR-0001: Approved stack and deployment direction

- **Status:** Approved direction; implementation gated by proof-of-concept and legal/security review
- **Decision date:** 2026-10-22

## Decision

Pinchd will be engineered to major-platform standards from the beginning while keeping initial capacity and operating spend proportional to demand.

1. **Mobile:** React Native + Expo + TypeScript for shared iOS and Android delivery.
2. **Backend:** Go, organised around explicit domain boundaries and typed interfaces. Services may be independently deployable where operationally justified; do not create a large fleet of microservices without concrete ownership, scaling or reliability reasons.
3. **Cloud:** AWS as the intended cloud platform.
4. **Transactional data:** PostgreSQL as source of truth. Evaluate PostGIS if location query requirements justify it.
5. **Media:** Private S3 buckets, controlled short-lived access, upload validation, image processing, metadata stripping and moderation states.
6. **Async processing:** Managed queues/workers for notifications, media processing, deletion propagation and other workflows when needed; define idempotency, retries, dead-letter handling and observability.
7. **Admin:** Separately deployed, role-restricted console with MFA, least privilege and audit trails.
8. **Operations:** CI/CD, automated tests, infrastructure as code, secret management, dependency/secret scanning, observability with redaction, budget alerts, backups, restoration exercises, incident response and release gates.
9. **Security/privacy:** Server-side authorisation and entitlement checks; private-by-default media; explicit retention/deletion; data-flow and vendor reviews; DPIA and legal review before high-risk profiling, biometric processing or novel AI.
10. **Billing:** Server-authoritative entitlements, platform-compliant purchase flows, restore support and reconciliation of purchase notifications.
11. **Initial geography:** One primary AWS region. Use Multi-AZ managed database resilience in production where the selected service and budget support it. Automate backups and test recovery. Do not deploy a global database cluster at launch by default.
12. **Scale path:** Design for future independent services, cross-region recovery and, if justified, multi-region data services. Add them based on measured load, recovery objectives, user geography, team ownership and legal/data-residency review.

## Cost discipline

Working estimates only, not vendor quotes:
- Development/testing: approximately USD $25–$100/month.
- Lean early production: approximately USD $150–$500+/month.
- Growth planning: approximately USD $500–$2,000+/month.

These ranges exclude engineering labour, legal/compliance work, age-assurance charges, paid third-party vendors, app-store fees and exceptional usage. Validate current regional pricing, data transfer, logs, backup, networking and database-resilience costs before provisioning. Treat startup credits as conditional, not guaranteed. Set budgets and alerts before production deployment.

## Alternatives and rationale

- TypeScript backend: credible and can share language with the client; Go is selected for the performance-oriented backend direction, while TypeScript remains the mobile language.
- Python: not required for the selected stack; add a separate ML/data component only if a concrete need emerges.
- Full microservices from day one: not a requirement. Preserve service boundaries but avoid operational fragmentation without a measurable benefit.
- Multi-region database from day one: deferred due to cost and consistency/data-residency complexity.
- Self-managed servers/database: not preferred for the initial phase because they transfer operational burden to the team.

## Gates before implementation/production

- Validate Expo/native SDK compatibility, including Didit age-assurance flows, push, secure storage, deep links, media upload and accessibility.
- Benchmark matching and pagination with synthetic data.
- Test block propagation, account deletion, media access revocation and subscription reconciliation end to end.
- Complete threat modelling, DPIA/privacy review, data-flow and processor/transfer reviews, and legal review.
- Define SLOs, recovery objectives, incident ownership, alerting, backup restoration and migration rollback/forward recovery.
- Confirm team capacity, service ownership, CI/CD and infrastructure-as-code approach.
- Verify actual AWS costs for development, beta and expected growth scenarios.

## Review triggers

Revisit the architecture when measured traffic, team size, availability requirements, international expansion, regulatory scope, vendor capabilities or unit economics materially change.

This ADR records the approved direction, not proof that integrations, legal compliance, or production readiness have already been completed.
