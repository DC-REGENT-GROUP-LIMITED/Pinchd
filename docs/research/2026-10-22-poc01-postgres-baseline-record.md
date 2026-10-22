# Pinchd — POC-01 Local PostgreSQL Transaction Baseline

- **Run date:** 22 October 2026
- **Environment:** Dedicated local PostgreSQL cluster on DISHAUN-ROG
- **Status:** PARTIAL PASS — database transaction baseline only; POC-01 is not complete
- **PostgreSQL:** 17.10, x86_64 Windows
- **Connection:** TCP loopback only, 127.0.0.1:55433
- **Database:** `pinchd_poc`
- **Application role:** `pinchd_poc_app` (verified not superuser, cannot create databases or roles)
- **Cluster directory:** `%LOCALAPPDATA%\Pinchd\poc-postgres-20261022\data`
- **Credentials:** Random secrets stored outside the repository in ACL-restricted files under the PoC AppData directory. Secret values are intentionally not recorded here.

## Isolation controls observed

- Created a new PostgreSQL cluster in the current Windows user's AppData, separate from the existing `postgresql-x64-17` service and its data directory.
- Cluster listens on `127.0.0.1` only and uses port 55433. Port 55432 was already occupied, so it was not used.
- SCRAM-SHA-256 authentication configured for local and host connections.
- Dedicated application role has no superuser, database-creation or role-creation privileges.
- Test credentials were not printed or placed in the repository.
- No cloud resources, vendor endpoints, real notifications, production credentials or real user data were used.
- The cluster is intentionally left running for subsequent approved local PoCs. It has not been added as a Windows service or configured to start automatically.

## Executed checks

| Check | Result |
|---|---|
| Connect as `pinchd_poc_app` to `pinchd_poc` | PASS |
| Verify current database, role, loopback address and port | PASS |
| Create minimal idempotency and match-pair tables | PASS |
| Duplicate idempotency key rejected by primary-key constraint | PASS |
| Unsorted pair rejected by check constraint | PASS |
| Simulated two-row operation rolled back atomically | PASS |
| Rolled-back rows absent on subsequent query | PASS |
| Restart isolated cluster with `pg_ctl` | PASS |
| Reconnect after restart and verify committed idempotency row persisted | PASS |
| Verify application role has `rolsuper=false`, `rolcreatedb=false`, `rolcreaterole=false` | PASS |

Command execution was performed using the installed PostgreSQL 17.10 binaries and the application role. SQL source is retained in `tests/poc-harness/poc01-transaction-baseline.sql`.

## Limitations

This is a narrow SQL baseline, not the full POC-01 exit criteria. It does not yet test the Go API/server lifecycle, API configuration validation, safe error mapping, concurrent reciprocal likes, conflicting idempotency-key payloads, outbox semantics, block/unmatch races, migrations under CI, security hardening, backup restoration or production reliability. The tables are throwaway PoC tables, not the proposed production schema or migration.

No independent reviewer has signed this result. Engineering/Security/Privacy review of the environment and fixture remains open; the current work is local-only and does not authorize production or cloud use.

## Teardown

Stop only this dedicated cluster (not the existing PostgreSQL service):

```powershell
$base = Join-Path $env:LOCALAPPDATA 'Pinchd\poc-postgres-20261022'
& 'C:\Program Files\PostgreSQL\17\bin\pg_ctl.exe' -D (Join-Path $base 'data') -m fast -w stop
```

Before deleting the cluster, confirm no subsequent PoC depends on it, verify the data directory path is exactly the dedicated AppData path above, stop the cluster, then remove only that dedicated PoC directory. Never delete or modify `C:\Program Files\PostgreSQL\17\data` or stop the existing `postgresql-x64-17` service as part of this PoC.
