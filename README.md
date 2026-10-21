# Pinchd

**Product:** Pinchd  
**Intended domain:** `pinchd.co`  
**Product direction:** Inclusive adult dating app; open to casual dating while helping users find meaningful relationships.  
**Status:** Project scaffold and architecture proposal — not a running application.

## What this repository is for

Pinchd is organised as a monorepo so the mobile app, admin console, API, shared code, tests, infrastructure, and product research can evolve together without mixing responsibilities.

This scaffold does not lock the team into a final programming language, database, hosting provider, or billing vendor. The `apps/mobile` layout is compatible with a recommended Expo / React Native approach for shared iOS and Android development, but confirm the stack decision before generating application code.

## Start here

1. `docs/product/` — product principles and links to the existing requirements blueprint.
2. `docs/architecture/folder-structure.md` — explains every major directory and dependency boundaries.
3. `docs/architecture/decisions/ADR-0001-stack-and-deployment.md` — decisions still to be made.
4. `docs/security/` — security and privacy foundations.
5. `docs/trust-and-safety/` — moderation, reporting, blocking, appeals, and incident-response requirements.
6. `research/2026-folder-architecture-research.md` — current guidance informing this structure.

## Suggested initial architecture

- **Mobile:** one shared cross-platform app for iOS and Android, with platform-specific code only when needed.
- **Admin:** a separate web app for authorised support and trust-and-safety staff.
- **API:** start with a modular monolith organised by domain boundaries. Split services only when scaling, reliability, team ownership, or deployment needs justify it.
- **Shared packages:** small packages for API contracts, design tokens/components, analytics event schemas, configuration, and test utilities.
- **Data and infrastructure:** versioned migrations, environment-specific configuration, infrastructure-as-code, observability, backup/restore and runbooks.
- **Safety and privacy:** first-class domains, least-privilege staff access, audit trails, data minimisation, and defined retention/deletion.

## Product-critical domains

`identity`, `profiles`, `preferences`, `discovery`, `matching`, `messaging`, `trust-safety`, `media`, `notifications`, `billing`, `analytics`, and `admin`.

Safety is not a paid feature. Reporting, blocking, account controls, and the ability to message a mutual match must remain core product capabilities.

## Important

This is a **folder scaffold**, not a completed or production-ready codebase. Empty directories contain `.gitkeep` files. No credentials, user data, live services, or production configuration are included.
