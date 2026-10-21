# Pinchd folder structure

## Repository tree

```text
pinchd/
├── apps/
│   ├── mobile/                 # iOS + Android client
│   └── admin/                  # Restricted internal web console
├── services/
│   └── api/                    # Backend API; modular monolith initially
├── packages/
│   ├── api-contracts/          # Request/response schemas and generated types
│   ├── design-system/          # Tokens and reusable UI primitives
│   ├── analytics-contracts/    # Typed event names and event schemas
│   ├── config/                 # Shared lint, formatting and TypeScript config
│   └── test-utils/             # Shared fakes, fixtures and test helpers
├── docs/
│   ├── product/                # Product principles, requirements and journeys
│   ├── architecture/
│   │   └── decisions/          # Architecture Decision Records (ADRs)
│   ├── api/                    # API conventions and endpoint contracts
│   ├── data/                   # Data dictionary, retention and lifecycle maps
│   ├── security/               # Threat models, security and privacy reviews
│   ├── trust-and-safety/       # Moderation policy, escalation and appeals
│   ├── operations/             # Incident, backup, recovery and release runbooks
│   └── research/               # Research notes and source register
├── infra/
│   ├── local/                  # Local development dependencies
│   ├── environments/
│   │   ├── staging/
│   │   └── production/
│   ├── modules/                # Reusable infrastructure modules
│   └── observability/          # Dashboards, alerts and service-level objectives
├── database/
│   ├── migrations/             # Version-controlled schema changes
│   ├── seeds/                  # Synthetic, non-production seed data
│   └── diagrams/               # ERDs and data-flow diagrams
├── tests/
│   ├── e2e/                    # End-to-end journeys
│   ├── contract/               # API/event compatibility tests
│   ├── performance/            # Load and latency tests
│   ├── security/               # Security and abuse-case test plans
│   └── accessibility/          # Accessibility checks and test plans
├── scripts/                    # Repeatable developer and release scripts
├── research/                   # Dated research and architecture source notes
├── .github/
│   └── workflows/              # CI workflows (to be added with the chosen stack)
├── .gitignore
├── CONTRIBUTING.md
└── README.md
```

## Mobile app organisation

Once the framework is confirmed, organise mobile code by **feature/domain first**, with explicit UI, application/state, domain, and data boundaries inside complex features. Avoid one giant `screens/`, `components/`, or `utils/` folder.

Suggested feature areas:

- `auth/` — sign-in, account recovery, sessions and age eligibility.
- `onboarding/` — intent, preferences, profile setup and privacy preview.
- `profile/` — profile editing, prompts, photos and visibility.
- `discovery/` — candidate feed, filters and recommendation explanations.
- `likes/` and `matches/` — reciprocal interest and match lifecycle.
- `messaging/` — conversation UI, delivery state and client-side safety affordances.
- `trust-safety/` — report, block, unmatch, appeal and safety centre.
- `settings/` — privacy, notification preferences, export and deletion.
- `subscription/` — entitlements and purchase restoration; never authoritative for server access control.

Keep API clients and storage behind repositories/adapters. UI components should not call the database or remote service directly. Keep platform-specific implementations behind narrow interfaces.

## API organisation

Start as a modular monolith, with domain modules such as:

- `identity`
- `profiles`
- `preferences`
- `discovery`
- `matching`
- `messaging`
- `media`
- `trust-safety`
- `notifications`
- `billing`
- `analytics`
- `admin`

Each domain should own its business rules and service interfaces. Keep transport controllers, business/application logic, persistence adapters, schemas, and tests separate. Use explicit contracts for cross-domain calls and events. Avoid premature microservices; split a domain out only when there is a concrete operational or team reason.

## Trust, safety and privacy

Trust and safety is a core product domain, not just a moderation queue. Include:

- Report and block flows from profiles, messages and matches.
- Abuse prevention, spam/scam signals and rate limits.
- Moderation queues, evidence access controls and auditable staff actions.
- Appeal and reversal workflows for significant enforcement decisions.
- Age eligibility and any additional age assurance decisioning.
- User privacy controls, data export, account deletion and retention/deletion jobs.
- Secure media upload, access controls and moderation state.
- Location minimisation; never expose precise user coordinates to other users.
- Staff MFA, least privilege, audit logs and separation of support/admin roles.

Keep intimate data, verification data and moderation evidence out of ordinary application logs and analytics payloads. Define purpose, lawful basis, retention, access policy and deletion behavior before collecting sensitive data.

## Billing and analytics boundaries

Billing owns products, entitlements, subscription state reconciliation, refunds and store receipt validation. The server—not the client—must decide access to paid entitlements.

Analytics uses a documented event taxonomy and data minimisation. Do not treat swipe volume or time spent as the sole success metric. Measure reciprocal matches, two-sided conversations, user-reported usefulness, fairness, safety outcomes, reliability and business sustainability together.

## Testing and release quality

- Unit tests for domain logic and state transitions.
- Contract tests for API and event compatibility.
- Integration tests for persistence, queues and external providers.
- End-to-end tests for onboarding, discovery, mutual match, messaging, report/block, subscription restore and account deletion.
- Abuse-case and permission tests for every sensitive action.
- Accessibility and device coverage.
- Dependency, secret, static-analysis and mobile security checks.
- Staging verification, database migration checks, backup/restore drills, monitoring and rollback plans.

## Naming conventions

- Lowercase kebab-case for folders and documentation filenames.
- Use feature/domain names rather than generic buckets.
- Keep ADRs numbered: `ADR-0001-short-title.md`.
- Document environment variables in `.env.example`, never commit secrets.
- Every new event, database table, sensitive field and retention rule should be documented.
