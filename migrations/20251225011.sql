ALTER TABLE contract_events
ADD CONSTRAINT uniq_event_log UNIQUE
(
  block_hash,
  contract_address,
  transaction_hash,
  log_index
);

CREATE TABLE processed_events (
    tx_hash        VARCHAR NOT NULL,
    log_index      INTEGER NOT NULL,
    contract       VARCHAR NOT NULL,
    block_number   NUMERIC NOT NULL,
    PRIMARY KEY (tx_hash, log_index)
);
