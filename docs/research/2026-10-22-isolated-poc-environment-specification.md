# Pinchd — Isolated Proof-of-Concept Environment Specification
**Date:** 22 October 2026  
**Status:** Proposed specification for Platform/Security/Privacy review; no environment is provisioned or approved by this document.  
**Purpose:** Define bounded, disposable conditions for POC-01–POC-12 without production access or real-user data.

## 1. Separation and permitted scope
Maintain distinct local development, isolated PoC, any controlled beta and production environments. The PoC must use synthetic fixtures only. It must have separate accounts, credentials, encryption keys, databases, object storage, queues, webhook endpoints, analytics and billing configuration. No trust relationship, shared secret, production IAM role, production network route, or production data export is permitted.

Default first stage: local disposable Go/PostgreSQL services using pinned versions and synthetic data. Cloud resources are optional and require the cost owner’s written budget, region, spend cap, alert recipient and teardown approval before provisioning. This specification does not authorize spend or select a cloud region.

## 2. Required controls before first run
- Repository-reviewed infrastructure/configuration; no committed secrets; secret scanner and dependency inventory recorded.
- Least-privilege, short-lived identities per service; MFA for human access; no shared administrator credential.
- Network egress allowlist; deny production endpoints and unrestricted inbound access; webhook callbacks use local test signing keys unless a vendor sandbox is formally approved.
- Separate encrypted storage and synthetic-only database; private object bucket/container with short-lived access and no public listing.
- Logging redaction by default; prohibit message bodies, auth tokens, exact location, age documents, biometric data and vendor fraud signals in logs/traces.
- Test queues and notification sinks cannot reach real users, email/SMS numbers, push tokens or public channels.
- CI uses disposable databases and fake provider endpoints; no live credentials or vendor billable actions.
- Spend limits, budget alerts and resource expiry labels where cloud is approved; automated teardown plus an owner-confirmed inventory check.
- Document backup behavior; do not rely on backups to clean test data. Verify tombstone replay before any restore exercise.

## 3. Environment inventory template
Record before execution: environment ID; owner role; creation/expiry; region (if any); service versions; database/object/queue resources; identity and permissions; ingress/egress rules; secret references (never values); data classification; estimated and maximum cost; monitoring/alert route; backup settings; teardown method; residual-resource check; approval/evidence links.

Unknown fields block provisioning if they affect privacy, security, cost, or production isolation. Do not fill them with assumed values.

## 4. Staged execution boundaries
**Stage A — local:** schema/migration, transaction, idempotency, state-machine, outbox, block/unmatch and deletion logic against deterministic fixtures. No network access to live providers.  
**Stage B — isolated cloud (optional):** only when a POC requires cloud-managed behavior; use a separate account/project, region and synthetic-only data, non-production endpoints, explicit spend guardrails and time-bounded resources.  
**Stage C — vendor sandbox (optional):** only after vendor terms, sandbox rules, data flow and test costs are reviewed. Synthetic payloads and approved test identities only; no real identity or biometric material.  
**Stage D — production or beta:** outside this POC specification and requires separately approved gates.

## 5. Stop conditions
Immediately stop and isolate the environment if production data/credentials are detected, a notification reaches a real recipient, a test can grant public/production eligibility, a public object URL is exposed, spend exceeds the approved cap, logs contain prohibited content, or teardown cannot establish resource inventory. Preserve minimal redacted evidence; rotate/revoke affected credentials and record incident owner and review.

## 6. POC evidence record
For each run record: POC ID and requirement IDs; commit/ref and dependency versions; environment ID/config hash; fixture version/seed; exact command/procedure; start/end time; result (PASS/FAIL/BLOCKED/NOT RUN); raw output location after secret/content review; expected vs observed behavior; defect links; reviewer; cleanup result; cost if applicable; limitations. A PASS requires reproducible output plus independent review. A plan, CI configuration or successful dashboard setup is not a test result.

## 7. Authorization gates and relation to tracker
ACT-12 remains OPEN until Platform, Security and Privacy review the concrete environment plan, cost owner approves any spend, a specific environment inventory exists, and isolation/teardown are demonstrated. EV-09 remains review-open until controls are evidenced. POC results remain NOT RUN until actually executed and recorded. See ACT-12/13–21, EV-09/10/12/13 and the technical PoC plan. Mobile UI design remains paused.
