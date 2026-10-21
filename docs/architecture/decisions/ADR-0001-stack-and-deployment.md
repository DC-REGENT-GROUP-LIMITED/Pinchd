# ADR-0001: Stack and deployment are not final

- **Status:** Proposed / decision required
- **Date:** 2026-10-21

## Context

Pinchd needs iOS and Android clients, a restricted admin console, a backend for matching and messaging, and first-class trust-and-safety capabilities. The existing product blueprint does not settle every implementation technology.

## Provisional recommendation

1. Evaluate Expo / React Native with TypeScript for a shared iOS/Android app.
2. Use a TypeScript backend if it fits the team's skills and operational needs; begin with a modular monolith and clear domain boundaries.
3. Choose a relational database with transactional integrity for accounts, likes, matches, billing state, and moderation actions; select storage/queues/search only when requirements justify them.
4. Keep the admin console separately deployed and role-restricted.
5. Use managed infrastructure where it reduces operational burden, but document data residency, processor terms, export/deletion, backup, incident response, and vendor exit paths.
6. Select a billing/subscription approach that complies with the relevant app-store policies and supports purchase restoration and server-side entitlement validation.

## Alternatives

- Fully native Swift + Kotlin: stronger platform-specific control, but higher implementation and maintenance cost.
- Flutter: another credible cross-platform option; evaluate team skills and required native integrations.
- Separate backend microservices from day one: rejected provisionally due to added deployment and operational complexity before traffic and team boundaries are known.

## Decision criteria

Evaluate accessibility, performance, native SDK requirements, hiring/team familiarity, testability, dependency health, security support, deployment complexity, observability, cost, data portability, and lock-in.

## Consequences

This repository is intentionally a scaffold. Do not generate framework-specific application files until the team records the stack decision and pins supported versions.
