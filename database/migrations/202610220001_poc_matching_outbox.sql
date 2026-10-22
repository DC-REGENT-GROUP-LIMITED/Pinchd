-- Disposable synthetic PoC schema only. Not a production migration.
-- Apply only to the isolated pinchd_poc database using the local PoC role.
BEGIN;

CREATE TABLE IF NOT EXISTS poc_accounts (
    account_id text PRIMARY KEY CHECK (account_id ~ '^test-[a-z0-9-]+$'),
    eligible boolean NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS poc_likes (
    actor_id text NOT NULL REFERENCES poc_accounts(account_id),
    recipient_id text NOT NULL REFERENCES poc_accounts(account_id),
    state text NOT NULL CHECK (state IN ('active', 'withdrawn')),
    created_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (actor_id, recipient_id),
    CHECK (actor_id <> recipient_id)
);

CREATE TABLE IF NOT EXISTS poc_matches_core (
    match_id bigserial PRIMARY KEY,
    pair_low_id text NOT NULL REFERENCES poc_accounts(account_id),
    pair_high_id text NOT NULL REFERENCES poc_accounts(account_id),
    state text NOT NULL CHECK (state IN ('active', 'ended')),
    created_at timestamptz NOT NULL DEFAULT now(),
    ended_at timestamptz,
    end_reason text,
    CHECK (pair_low_id < pair_high_id),
    UNIQUE (pair_low_id, pair_high_id)
);

CREATE TABLE IF NOT EXISTS poc_conversations (
    conversation_id bigserial PRIMARY KEY,
    match_id bigint NOT NULL UNIQUE REFERENCES poc_matches_core(match_id),
    state text NOT NULL CHECK (state IN ('active', 'closed')),
    created_at timestamptz NOT NULL DEFAULT now(),
    closed_at timestamptz
);

CREATE TABLE IF NOT EXISTS poc_blocks (
    blocker_id text NOT NULL REFERENCES poc_accounts(account_id),
    blocked_id text NOT NULL REFERENCES poc_accounts(account_id),
    created_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (blocker_id, blocked_id),
    CHECK (blocker_id <> blocked_id)
);

CREATE TABLE IF NOT EXISTS poc_pair_exclusions (
    pair_low_id text NOT NULL REFERENCES poc_accounts(account_id),
    pair_high_id text NOT NULL REFERENCES poc_accounts(account_id),
    reason text NOT NULL CHECK (reason IN ('unmatch')),
    created_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (pair_low_id, pair_high_id),
    CHECK (pair_low_id < pair_high_id)
);

CREATE TABLE IF NOT EXISTS poc_messages (
    message_id bigserial PRIMARY KEY,
    match_id bigint NOT NULL REFERENCES poc_matches_core(match_id),
    sender_id text NOT NULL REFERENCES poc_accounts(account_id),
    client_message_id text NOT NULL,
    body text NOT NULL CHECK (length(body) BETWEEN 1 AND 2000),
    created_at timestamptz NOT NULL DEFAULT now(),
    UNIQUE (sender_id, client_message_id)
);

CREATE TABLE IF NOT EXISTS poc_idempotency_requests (
    actor_id text NOT NULL REFERENCES poc_accounts(account_id),
    idempotency_key text NOT NULL,
    request_hash text NOT NULL,
    response_json jsonb NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (actor_id, idempotency_key)
);

CREATE TABLE IF NOT EXISTS poc_outbox (
    event_id bigserial PRIMARY KEY,
    event_key text NOT NULL UNIQUE,
    event_type text NOT NULL,
    aggregate_key text NOT NULL,
    payload jsonb NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    published_at timestamptz,
    attempts integer NOT NULL DEFAULT 0 CHECK (attempts >= 0)
);

CREATE TABLE IF NOT EXISTS poc_inbox (
    consumer_name text NOT NULL,
    event_key text NOT NULL,
    processed_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (consumer_name, event_key)
);

CREATE INDEX IF NOT EXISTS poc_outbox_pending_idx
    ON poc_outbox (event_id) WHERE published_at IS NULL;
CREATE INDEX IF NOT EXISTS poc_messages_match_idx
    ON poc_messages (match_id, message_id);

COMMIT;
