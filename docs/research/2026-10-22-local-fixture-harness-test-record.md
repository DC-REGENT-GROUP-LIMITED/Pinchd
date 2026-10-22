# Pinchd — Local Synthetic Fixture Harness Test Record

- **Run date:** 2026-10-22
- **Environment:** DISHAUN-ROG, Windows amd64
- **Scope:** Local validation of synthetic fixture structure and safety labels only
- **Fixture:** `tests/fixtures/poc/core-scenarios.json`
- **Harness:** `tests/poc-harness/core_test.go`
- **Data:** Synthetic-only; no live vendor, database, cloud, notification or production calls

## Commands and results

| Command | Result |
|---|---|
| `go version` | `go1.25.6 windows/amd64` (previously observed on PATH) |
| `gofmt -w core_test.go` | Completed |
| `go vet ./...` | PASS (exit code 0; run before test command) |
| `go test -count=1 -v ./...` | PASS — 6 tests |
| `git diff --check` | PASS; Git emitted line-ending conversion warnings for existing working-copy files, not whitespace errors |
| `go test -race ./...` | NOT RUN — toolchain reports race detector requires cgo; cgo is not enabled in the current shell |

## Assertions exercised

1. Fixture is labelled synthetic-only and prohibits live-provider use.
2. Offline/no-real-data/no-production-credential constraints are present and true.
3. Account/profile/relationship/provider-event references and identifiers are consistent.
4. Pending/expired age state and suspended/deletion-pending account states do not satisfy the harness's basic access predicate.
5. Unknown relationship goal is not converted into explicit open-to intent.
6. Block and post-unmatch pair suppression are represented as separate relationship types.
7. Scenario identifiers are unique and scenarios include names/assertions. (Six Go test functions; several test functions cover multiple assertions.)

## Limitations and next gate

This result validates fixture shape and a few static policy guardrails only. It is **not** a production matching test, PostgreSQL transaction/concurrency test, Didit callback test, outbox/push test, deletion test, security assessment, privacy approval or legal compliance result. PostgreSQL 17.10 is installed and its Windows service is running; `pg_isready` reports port 5432 accepting connections, but `psql` is not on PATH. No login or database modification was attempted. POC-01 remains NOT STARTED until a dedicated disposable database/user, isolation boundaries, teardown and environment approval are recorded. Do not reuse the existing service for PoCs without that review, provision cloud resources, or send fixture payloads to live providers as a workaround.
