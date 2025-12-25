-- ALTER TABLE contract_events
-- ADD CONSTRAINT uniq_event_log UNIQUE
-- (
--   block_hash,
--   contract_address,
--   transaction_hash,
--   log_index
-- );

CREATE TABLE IF NOT EXISTS processed_events (
                                                tx_hash        VARCHAR(80) NOT NULL,
    log_index      INTEGER NOT NULL,
    contract       VARCHAR(80) NOT NULL,
    block_number   NUMERIC NOT NULL,
    created_at     TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (tx_hash, log_index)
    );