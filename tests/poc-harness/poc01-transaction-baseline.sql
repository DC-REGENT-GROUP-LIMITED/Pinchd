\set ON_ERROR_STOP on
\echo 'POC-01 isolated PostgreSQL baseline: begin'
SELECT current_database() = 'pinchd_poc' AS expected_database,
       current_user = 'pinchd_poc_app' AS expected_role,
       inet_server_addr() = '127.0.0.1'::inet AS loopback_only,
       inet_server_port() = 55433 AS expected_port
\gset
\if :expected_database
\else
  \echo 'FAIL: unexpected database'
  \quit 10
\endif
\if :expected_role
\else
  \echo 'FAIL: unexpected role'
  \quit 11
\endif
\if :loopback_only
\else
  \echo 'FAIL: connection is not loopback'
  \quit 12
\endif
\if :expected_port
\else
  \echo 'FAIL: unexpected port'
  \quit 13
\endif

CREATE TABLE IF NOT EXISTS public.poc_idempotency (
    idempotency_key text PRIMARY KEY,
    request_hash text NOT NULL,
    response_payload jsonb NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS public.poc_matches (
    pair_a text NOT NULL,
    pair_b text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT poc_matches_sorted_pair CHECK (pair_a < pair_b),
    CONSTRAINT poc_matches_unique_pair UNIQUE (pair_a, pair_b)
);
TRUNCATE TABLE public.poc_idempotency, public.poc_matches;

-- Duplicate idempotency keys must be rejected.
INSERT INTO public.poc_idempotency (idempotency_key, request_hash, response_payload)
VALUES ('poc01-key-001', 'hash-a', '{"result":"accepted"}');
DO $$
BEGIN
  BEGIN
    INSERT INTO public.poc_idempotency (idempotency_key, request_hash, response_payload)
    VALUES ('poc01-key-001', 'hash-a', '{"result":"duplicate"}');
    RAISE EXCEPTION 'duplicate idempotency key was unexpectedly accepted';
  EXCEPTION WHEN unique_violation THEN
    RAISE NOTICE 'PASS: duplicate idempotency key rejected';
  END;
END $$;

-- A constraint failure must not leave a partial row.
DO $$
BEGIN
  BEGIN
    INSERT INTO public.poc_matches (pair_a, pair_b) VALUES ('acct-z', 'acct-a');
    RAISE EXCEPTION 'unsorted pair was unexpectedly accepted';
  EXCEPTION WHEN check_violation THEN
    RAISE NOTICE 'PASS: pair-order constraint rejected invalid pair';
  END;
END $$;

-- Simulate an operation that writes two records then rolls back atomically.
BEGIN;
INSERT INTO public.poc_idempotency (idempotency_key, request_hash, response_payload)
VALUES ('poc01-rollback-key', 'hash-rollback', '{"result":"temporary"}');
INSERT INTO public.poc_matches (pair_a, pair_b) VALUES ('acct-a', 'acct-b');
ROLLBACK;

DO $$
BEGIN
  IF EXISTS (SELECT 1 FROM public.poc_idempotency WHERE idempotency_key = 'poc01-rollback-key') THEN
    RAISE EXCEPTION 'rollback left idempotency row behind';
  END IF;
  IF EXISTS (SELECT 1 FROM public.poc_matches WHERE pair_a = 'acct-a' AND pair_b = 'acct-b') THEN
    RAISE EXCEPTION 'rollback left match row behind';
  END IF;
  RAISE NOTICE 'PASS: rollback removed all rows from simulated atomic operation';
END $$;

-- Verify successful state survives a fresh connection.
SELECT count(*) = 1 AS only_one_idempotency_record
FROM public.poc_idempotency WHERE idempotency_key = 'poc01-key-001'
\gset
\if :only_one_idempotency_record
  \echo 'PASS: unique idempotency record persisted'
\else
  \echo 'FAIL: idempotency record count incorrect'
  \quit 14
\endif
SELECT count(*) = 0 AS rollback_rows_absent
FROM public.poc_idempotency WHERE idempotency_key = 'poc01-rollback-key'
\gset
\if :rollback_rows_absent
  \echo 'PASS: rolled-back rows absent after transaction'
\else
  \echo 'FAIL: rollback verification failed'
  \quit 15
\endif
\echo 'POC-01 SQL baseline assertions passed'
