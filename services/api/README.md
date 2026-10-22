# Pinchd API — local PoC shell

Approved direction: Go backend, initially a modular monolith with explicit domain boundaries; PostgreSQL is the relational database direction. This scaffold is a local proof-of-concept only, not production API functionality.

## Endpoints

- `GET /healthz`: process-level liveness; does not query the database.
- `GET /readyz`: bounded database ping; returns only `ready` or `not_ready`, never driver or credential details.
- Default bind: `127.0.0.1:8080`. Override with `PINCHD_LISTEN_ADDR` only for an explicitly controlled test.

The server requires `PGHOST`, `PGDATABASE`, `PGUSER`, and `PGPASSWORD`; `PGPORT` defaults to 5432. Never put secrets in source control, command-line arguments, logs, screenshots, or evidence files. For the isolated Windows PoC, load the app password from the ACL-restricted `%LOCALAPPDATA%\Pinchd\poc-postgres-20261022\app.secret` into the process environment in a private local shell; do not print it.

## Run and test (PowerShell)

From this directory:

```powershell
go mod tidy
go test ./...
$env:PGHOST = '127.0.0.1'
$env:PGPORT = '55433'
$env:PGDATABASE = 'pinchd_poc'
$env:PGUSER = 'pinchd_poc_app'
$env:PGPASSWORD = (Get-Content (Join-Path $env:LOCALAPPDATA 'Pinchd\poc-postgres-20261022\app.secret') -Raw).Trim()
$env:PINCHD_POC_INTEGRATION = '1'
go test -count=1 ./...
go run .
```

Do not echo environment variables. Unset `PGPASSWORD` after the test. The integration test uses only the isolated local cluster, not live providers or cloud services.

## Security and limitations

- Configuration is validated before serving; DB connection is tested with a bounded timeout.
- HTTP server has request timeouts and a 1 MiB header cap.
- Health/readiness responses are non-cacheable and include basic defensive headers.
- Database errors are intentionally not returned to HTTP callers or logs.
- The PoC has no authentication, domain routes, migrations, authorization, rate limiting, telemetry, production TLS, deployment configuration or production-grade secret manager. Do not expose it beyond loopback or use it with real user data.
- The PostgreSQL DSN disables TLS because this PoC connects only to the isolated local loopback cluster. Production must use a reviewed transport/security configuration.
- This is only part of POC-01; it does not prove transaction concurrency, matching, outbox delivery, block/unmatch races, deletion, vendor integration, or production readiness.
