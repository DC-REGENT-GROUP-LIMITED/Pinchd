# Pinchd — 2026 architecture and folder-organisation research

- **Research date:** 21 October 2026
- **Purpose:** Inform the initial repository structure, not to claim a final stack decision.
- **Scope:** Cross-platform mobile structure, modular backend organisation, mobile security, platform safety requirements, and UK dating/social-discovery regulation.

## Findings and implications

### 1. Cross-platform monorepo

Expo's official monorepo guide describes workspaces with `apps/` for applications and `packages/` for shared code, and warns that monorepos add tooling complexity and are most useful when apps share code or native modules.

**Implication:** Use a small monorepo for mobile, admin, API and a few genuinely shared packages. Avoid creating a package for every tiny helper. Keep root scripts and workspace configuration to be added only after the stack is confirmed.

Source: https://docs.expo.dev/guides/monorepos/

### 2. Clear app layers and domain boundaries

Android's official architecture recommendations favour clear UI and data layers, repositories, a single source of truth, unidirectional data flow and tests. Its modularisation guide recommends loosely coupled, self-contained modules and high cohesion.

**Implication:** Organise client code by feature, with clear UI/application/domain/data boundaries where complexity warrants them. Start the backend as a modular monolith rather than prematurely splitting every domain into a separate service.

Sources:
- https://developer.android.com/topic/architecture/recommendations
- https://developer.android.com/topic/modularization
- https://developer.android.com/topic/modularization/patterns

### 3. Mobile security and privacy

OWASP's Mobile Application Security project provides MASVS security/privacy requirements, MASWE weakness classifications and MASTG testing guidance.

**Implication:** Include security test plans and privacy reviews in the repository from day one. Treat authentication, secure storage, network communication, permissions, dependency integrity and sensitive-data handling as explicit engineering concerns.

Sources:
- https://owasp.org/projects/mobile-application-security
- https://github.com/OWASP/masvs

### 4. User-generated content and safety

Apple's App Review Guidelines for user-generated content require filtering, reporting, timely response, user blocking and published contact information for apps in scope.

**Implication:** Trust and safety must be a product domain with its own code, admin workflows, policy documentation and tests. Report/block are not optional launch extras.

Source: https://developer.apple.com/app-store/review/guidelines/

### 5. UK dating and social-discovery obligations

Ofcom published dating and social-discovery guidance on 18 May 2026. It explains that the Online Safety Act may apply to user-to-user services with links to the UK and advises services to assess scope, risks and applicable duties.

**Implication:** Include legal/compliance, risk assessment, moderation operations, data mapping and incident procedures in project planning. Applicability must be assessed for the actual service; this note is not legal advice.

Source: https://www.ofcom.org.uk/online-safety/illegal-and-harmful-content/dating-and-social-discovery-know-the-online-safety-risks-rules-and-how-to-comply

## Recommended folder principles

1. **Domain-first, not file-type-only:** group code by capabilities such as matching, messaging, billing and trust/safety.
2. **Explicit boundaries:** UI should not directly query data sources; domains should not reach into another domain's database internals.
3. **Privacy by design:** document data purpose, access, retention and deletion before collection.
4. **Safety by design:** moderation, reporting, blocking, appeals and auditability are core capabilities.
5. **Modular monolith first:** split services only for demonstrated scaling, reliability, ownership or deployment needs.
6. **Tests reflect risk:** cover critical user journeys, abuse cases, authorisation, billing and deletion—not just happy-path UI.
7. **Document decisions:** use ADRs for stack, database, hosting, authentication, matching strategy, verification and monetisation choices.
8. **Do not overbuild:** placeholders represent planned boundaries, not a commitment to implement every subsystem in the first release.

## Decisions still required

- Mobile framework: Expo/React Native vs Flutter vs native Swift/Kotlin.
- Backend language/framework and deployment model.
- Database, object storage, queues/search and realtime messaging approach.
- Identity, age-assurance and optional verification providers.
- Subscription and app-store purchase validation strategy.
- Initial launch geography, data protection scope and applicable online-safety duties.
- Moderation staffing, escalation SLAs, retention and appeals policy.
