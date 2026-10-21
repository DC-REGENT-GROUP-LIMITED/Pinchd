# Contributing to Pinchd

## Working agreements

- Read the product blueprint and relevant architecture decision records before changing a domain.
- Keep changes small, reviewable, and covered by tests.
- Never commit secrets, real user data, raw identity documents, moderation evidence, or production exports.
- Use synthetic fixtures for development and automated tests.
- Add a decision record when changing a major architectural choice.
- Treat safety, privacy, accessibility, and abuse resistance as acceptance criteria.

## Branch and pull-request conventions

Suggested branches: `feat/<short-name>`, `fix/<short-name>`, `docs/<short-name>`, `security/<short-name>`.

Pull requests should describe the user impact, data impact, safety implications, tests run, migrations, observability changes, and rollback plan.

## Domain boundaries

Avoid importing another domain's database implementation directly. Prefer documented service interfaces, typed contracts, or domain events. Cross-domain operations that require atomicity must be designed explicitly.
