-- 给 token_sent 补链上唯一键
ALTER TABLE token_sent
ADD COLUMN IF NOT EXISTS tx_hash varchar,
ADD COLUMN IF NOT EXISTS log_index int;

-- 给 token_received 补链上唯一键
ALTER TABLE token_received
ADD COLUMN IF NOT EXISTS tx_hash varchar,
ADD COLUMN IF NOT EXISTS log_index int;

-- 建唯一索引（幂等核心）
CREATE UNIQUE INDEX IF NOT EXISTS uniq_token_sent_tx
ON token_sent (tx_hash, log_index);

CREATE UNIQUE INDEX IF NOT EXISTS uniq_token_received_tx
ON token_received (tx_hash, log_index);
