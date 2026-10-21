# API conventions

To decide after the backend stack is selected:

- Version public API contracts deliberately.
- Validate inputs at the boundary; enforce authorisation in the server/domain layer.
- Use idempotency for actions that may be retried, such as likes, match creation, billing webhooks and deletion requests.
- Define pagination, rate limits, error codes, correlation IDs and retry semantics.
- Avoid logging message content, exact location, auth tokens or sensitive profile fields.
- Document event schemas and compatibility rules.
- Keep provider-specific integrations behind adapters.
