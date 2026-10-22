# Pinchd local synthetic fixture checks

This is a small, standard-library-only Go test harness for validating the structure and safety labels of `tests/fixtures/poc/core-scenarios.json`.

Run from this directory:

```powershell
go test ./...
```

Scope is deliberately narrow: fixture integrity, synthetic/offline guardrails, reference consistency, basic account/age-gate invariants, unknown-goal handling, and presence of separate block/unmatch-suppression examples. It does **not** implement the production matching engine, database transactions, provider callback verification, push delivery, billing, deletion orchestration, or end-to-end behavior. Passing tests are not proof of those systems or of legal compliance. The fixture contains no real users and must never be sent to live vendors.
