# Pinchd — POC-01 Go API and PostgreSQL Readiness Baseline

- **Run date:** 22 October 2026
- **Environment:** DISHAUN-ROG, dedicated local PostgreSQL cluster only
- **Status:** CONDITIONAL PASS — minimal API smoke checks passed; POC-01 remains incomplete
- **Go:** 1.25.6
- **PostgreSQL:** 17.10, loopback-only at 127.0.0.1:55433
- **Database/role:** `pinchd_poc` / `pinchd_poc_app`
- **API module:** `services/api` (`pinchd.local/api`)
- **Review:** No independent engineering/security/privacy review recorded

## Implemented for this experiment

- Go HTTP server with a loopback default bind at `127.0.0.1:8080`.
- `GET /healthz` for process liveness without a database query.
- `GET /readyz` for a bounded database ping, returning a generic `ready` or `not_ready` response.
- Required DB environment validation and port/address validation.
- Bounded startup connection attempt, connection pool limits, HTTP timeouts, graceful-shutdown path, non-cacheable responses and basic defensive headers.
- Generic database-unavailable errors; database driver/credential details are not returned by readiness responses or startup logs.
- URL construction uses structured URL fields to encode credentials rather than string concatenation.
- Unit tests for missing configuration, invalid ports, liveness, readiness failure and non-leaking errors.
- Opt-in integration test controlled by `PINCHD_POC_INTEGRATION=1`.

## Executed checks

| Check | Observed result |
|---|---|
| `gofmt` | PASS |
| `go test -count=1 ./...` in `services/api` | PASS |
| `go vet ./...` in `services/api` | PASS |
| Opt-in DB integration test against isolated cluster | PASS |
| Actual server HTTP smoke check `GET /healthz` on 127.0.0.1:18080 | HTTP 200, `{"status":"ok"}` |
| Actual server HTTP smoke check `GET /readyz` on 127.0.0.1:18080 | HTTP 200, `{"status":"ready"}` |
| Existing PostgreSQL service/database | Not used; port 5432 service/data left untouched |
| Live providers, cloud services, real user data, notifications | Not used |

The integration command loaded the app password from the ACL-restricted file under the user's local AppData directory into a temporary PowerShell process environment. The secret was not printed or stored in the repository. The environment variables were removed after the test command.

## Limitations and open gates

This is not production readiness and does not complete POC-01. The API currently has no authentication, domain endpoints, authorization, schema migrations, CI workflow, rate limiting, production TLS configuration, production secret manager, structured observability, deployment/infra, load testing, concurrency testing, outbox semantics, or block/unmatch behavior. The local-only TLS-disabled DSN is explicitly unsuitable for remote production connections. The server was smoke-tested over loopback; independent review and an explicit teardown demonstration remain outstanding.

POC-02 through POC-12 remain NOT RUN/BLOCKED according to their dependency gates. No legal, privacy, vendor, age-assurance or safety approval is implied.
