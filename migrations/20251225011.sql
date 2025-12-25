ALTER TABLE contract_events
ADD CONSTRAINT uniq_event_log UNIQUE
(
  block_hash,
  contract_address,
  transaction_hash,
  log_index
);
