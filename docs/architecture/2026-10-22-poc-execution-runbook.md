# Pinchd — Local PoC Execution Runbook
**Date:** 22 October 2026  
**Status:** Draft execution procedure. No command in this document is evidence of a test having run.  
**Scope:** Local, disposable, synthetic-only engineering experiments; no mobile UI work and no production deployment.

## 1. Preconditions
- [ ] Engineering, security and privacy review the fixture specification and tests/fixtures/poc/core-scenarios.json.
- [ ] Product decisions LD-01–LD-10 are confirmed or each dependent behavior is explicitly kept out of scope.
- [ ] Use a clean working tree or a dedicated branch; record repository commit.
- [ ] Confirm installed Go, Docker/Compose and PostgreSQL tooling versions before selecting commands.
- [ ] Confirm no production/vendor secrets are present in the environment; verify .gitignore, secret scan and dependency inventory.
- [ ] Ensure outbound network is disabled or allowlisted to local-only dependencies. No live Didit, store billing, email, SMS or push calls.
- [ ] Approve a disposable local database and define teardown; no cloud spend is authorised by this runbook.

## 2. Execution order
1. **Fixture integrity:** parse JSON, validate unique IDs/references, confirm synthetic-only fields and deterministic seed.
2. **POC-01 foundation:** create a disposable Go module/API shell only after stack/versions are recorded; verify configuration validation, safe errors and database connection lifecycle.
3. **Database transaction behavior:** apply a throwaway schema; test rollback, constraint violations, duplicate idempotency keys and restart behavior.
4. **POC-02 reciprocal-like concurrency:** run concurrent reciprocal requests; assert one match and one conversation, with retries and conflicting idempotency-key payloads.
5. **POC-03 block/unmatch races:** race block/unmatch against reads, likes, message writes and queued notifications; assert no prohibited action succeeds after the authoritative state transition.
6. **POC-04 outbox reliability:** crash at commit/publish/consumer boundaries; assert durable event delivery, idempotent consumers and visible dead-letter/backlog states.
7. **POC-10 eligibility:** property-test age, gender, radius, blocks, account status, relationship goals and unknown values; assert hard exclusions never appear in output.
8. **POC-06 media:** only after selecting the local storage/test adapter; assert private-by-default access, MIME/size checks, metadata removal and deletion behavior.
9. **POC-07 notification sink:** local sink only; assert state recheck before delivery and no real recipients.
10. **POC-09 deletion:** simulate failures, retries, stale callbacks and restore; assert no resurrection and truthful completion status.
11. **POC-11 accessibility:** when the client exists, run automated semantic checks and manual assistive-technology testing. Fixture assertions alone do not close this PoC.
12. **POC-12 reliability/cost:** document backup/restore, measured runtime/resource use, alerts and teardown. Local-only results cannot prove AWS production cost or managed-service recovery.

POC-05 (Didit) and POC-08 (store billing) remain blocked until vendor terms, test workflow/account and counsel/privacy review are in hand. Do not replace these with a fabricated “passed” local mock; mocks prove our handling logic only.

## 3. Required evidence per run
Record POC ID, requirement IDs, date/time UTC, commit SHA, tool/dependency versions, environment ID, fixture schema/version/seed, exact command, expected/observed result, exit code, redacted output reference, defects, reviewer, teardown result and cost where applicable. Preserve only the minimum output needed; never attach credentials, raw personal data, message bodies or sensitive vendor payloads.

## 4. Status rules
- **NOT RUN:** no execution evidence.
- **BLOCKED:** a required approval, dependency or environment is absent.
- **FAIL:** assertions failed or safety boundary violated.
- **CONDITIONAL:** limited result with explicit non-critical limitation, owner and due date; cannot waive a critical gate.
- **PASS:** reproducible output and independent review confirm every required assertion.
A fixture parser passing does not mean the API, database or product requirement passed.

## 5. Stop immediately if
Production credentials/data are detected; a provider endpoint is contacted unintentionally; a notification reaches a real person; a hard eligibility gate is bypassed; private media becomes public; a deleted account is resurrected; sensitive content appears in logs; spend exceeds a written cap; or teardown cannot account for every created resource. Isolate, preserve minimal redacted evidence, revoke affected credentials, and follow the incident process.

## 6. Current state
- Fixture specification: drafted; approval pending.
- Fixture JSON: initial deterministic synthetic set; JSON parse, unique account/profile IDs, account-profile links, relationship endpoints and provider-event account references validated on 22 October 2026. Engineering/privacy review and behavioral tests remain pending.
- Host tooling check on 22 October 2026: Go 1.25.6 and Git 2.52.0 were available. PostgreSQL 17.10 binaries exist under `C:\Program Files\PostgreSQL\17\bin`; the `postgresql-x64-17` service is Running/Automatic and `pg_isready` reports port 5432 accepting connections. `psql` is not on PATH and Docker Desktop was not found at its common install path. No database login, database creation, service change or data modification was performed. Do not reuse the existing service for PoCs until Engineering/Security/Privacy approve isolation, dedicated credentials/database, network boundaries and teardown. These observed versions are not yet dependency pins.
- Local PoC environment: dedicated PostgreSQL 17.10 cluster provisioned under `%LOCALAPPDATA%\Pinchd\poc-postgres-20261022`, loopback-only on `127.0.0.1:55433`, SCRAM authentication, separate `pinchd_poc` database and least-privilege `pinchd_poc_app` role. Credential files are ACL-restricted and outside the repository. The existing PostgreSQL service/data was not used or modified. SQL transaction baseline and restart check have a partial-pass record at `../research/2026-10-22-poc01-postgres-baseline-record.md`; full POC-01 remains incomplete.
- POC-01: PARTIAL PASS for isolated PostgreSQL transaction/constraint/rollback/restart checks plus a minimal Go API readiness baseline. The API unit tests, vet, opt-in database integration test and loopback HTTP smoke checks passed; full POC-01 remains incomplete because migrations/CI, security hardening, broader API behavior and independent review are outstanding. See `../research/2026-10-22-poc01-postgres-baseline-record.md` and `../research/2026-10-22-poc01-go-api-baseline-record.md`. POC-02–POC-12 remain NOT RUN/BLOCKED unless separate dated evidence records say otherwise.
- Mobile UI design: remains paused.
