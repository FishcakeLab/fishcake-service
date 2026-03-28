CREATE TABLE IF NOT EXISTS mining_record (
    id varchar PRIMARY KEY DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    address varchar NOT NULL,
    record_type varchar NOT NULL,
    mined_amount_delta numeric(78, 0) NOT NULL DEFAULT 0,
    power_increase numeric(78, 0) NOT NULL DEFAULT 0,
    power_decrease numeric(78, 0) NOT NULL DEFAULT 0,
    activity_id int8 NULL,
    token_id int8 NULL,
    description text NOT NULL DEFAULT '',
    timestamp int8 NOT NULL DEFAULT 0,
    tx_hash varchar NOT NULL,
    log_index int4 NOT NULL,
    block_number int8 NOT NULL DEFAULT 0
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_mining_record_physical_unique
ON mining_record (tx_hash, log_index, record_type);

CREATE INDEX IF NOT EXISTS idx_mining_record_address_timestamp
ON mining_record (address, timestamp DESC);

CREATE INDEX IF NOT EXISTS idx_mining_record_activity_id
ON mining_record (activity_id);

CREATE INDEX IF NOT EXISTS idx_mining_record_token_id
ON mining_record (token_id);
