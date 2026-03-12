-- =====================================================================
-- 幂等性与物理追踪增强迁移 (Consolidated Migration)
-- =====================================================================

-- 1. [contract_events] 唯一约束：确保同步器拉取的原始事件不重复
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_constraint
        WHERE conname = 'uniq_event_log'
    ) THEN
        ALTER TABLE contract_events
        ADD CONSTRAINT uniq_event_log UNIQUE (block_hash, contract_address, transaction_hash, log_index);
    END IF;
END
$$;

-- 2. [processed_events] (废弃)：由于业务已实现真·幂等，不再需要此标记表
-- DROP TABLE IF EXISTS processed_events;

-- 3. [block_listener] 进度追踪增强：支持按配置名称分开记录高度
ALTER TABLE block_listener ADD COLUMN IF NOT EXISTS conf_name VARCHAR(255) DEFAULT 'event';
CREATE UNIQUE INDEX IF NOT EXISTS idx_block_listener_conf_name ON block_listener(conf_name);

-- 4. [drop_info] 增强
ALTER TABLE drop_info ADD COLUMN IF NOT EXISTS "block_number" int8 DEFAULT 0;
ALTER TABLE drop_info ADD COLUMN IF NOT EXISTS "log_index" int4 DEFAULT 0;

DROP INDEX IF EXISTS idx_drop_info_unique_action;
CREATE UNIQUE INDEX IF NOT EXISTS idx_drop_info_physical_unique 
ON drop_info (transaction_hash, address, log_index, drop_type);

-- 5. [stake_holder_staking] 增强
ALTER TABLE stake_holder_staking ADD COLUMN IF NOT EXISTS "block_number" int8 DEFAULT 0;
ALTER TABLE stake_holder_staking ADD COLUMN IF NOT EXISTS "log_index" int4 DEFAULT 0;

DROP INDEX IF EXISTS idx_stakeholder_nonce;
DROP INDEX IF EXISTS idx_stakeholder_nonce_unique;
CREATE UNIQUE INDEX IF NOT EXISTS idx_staking_physical_unique 
ON stake_holder_staking (tx_message_hash, log_index);

-- 6. [token_nft] 增强
ALTER TABLE token_nft ADD COLUMN IF NOT EXISTS "block_number" int8 DEFAULT 0;
ALTER TABLE token_nft ADD COLUMN IF NOT EXISTS "log_index" int4 DEFAULT 0;
ALTER TABLE token_nft ADD COLUMN IF NOT EXISTS "tx_hash" VARCHAR(80);
CREATE UNIQUE INDEX IF NOT EXISTS idx_token_nft_physical_unique 
ON token_nft (tx_hash, log_index);

-- 7. [token_transfer] 流水表增强 (针对 token_sent 和 token_received)
CREATE UNIQUE INDEX IF NOT EXISTS idx_token_sent_physical_unique 
ON token_sent (tx_hash, log_index);

CREATE UNIQUE INDEX IF NOT EXISTS idx_token_received_physical_unique 
ON token_received (tx_hash, log_index);

-- 8. [block_headers] 唯一索引：确保区块头不重复
CREATE UNIQUE INDEX IF NOT EXISTS idx_block_headers_hash ON block_headers (hash);

-- 9. [block_headers] 唯一索引：确保区块头不重复
CREATE UNIQUE INDEX IF NOT EXISTS idx_block_headers_number ON block_headers (number);
