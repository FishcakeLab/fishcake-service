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

-- 3. [block_listener] 进度追踪增强：支持按配置名称(如 'event', 'native') 分开记录同步高度
ALTER TABLE block_listener ADD COLUMN IF NOT EXISTS conf_name VARCHAR(255) DEFAULT 'event';
CREATE UNIQUE INDEX IF NOT EXISTS idx_block_listener_conf_name ON block_listener(conf_name);

-- 4. [drop_info] 空投信息增强：引入物理坐标(block_number, log_index) 并建立最强物理唯一性约束
ALTER TABLE drop_info ADD COLUMN IF NOT EXISTS "block_number" int8 DEFAULT 0;
ALTER TABLE drop_info ADD COLUMN IF NOT EXISTS "log_index" int4 DEFAULT 0;

-- 清除旧的/弱的索引，建立基于物理坐标的唯一索引
DROP INDEX IF EXISTS idx_drop_info_unique_action;
CREATE UNIQUE INDEX IF NOT EXISTS idx_drop_info_physical_unique 
ON drop_info (transaction_hash, log_index, drop_type);

-- 5. [stake_holder_staking] 质押信息增强
ALTER TABLE stake_holder_staking ADD COLUMN IF NOT EXISTS "block_number" int8 DEFAULT 0;
ALTER TABLE stake_holder_staking ADD COLUMN IF NOT EXISTS "log_index" int4 DEFAULT 0;

DROP INDEX IF EXISTS idx_stakeholder_nonce;
DROP INDEX IF EXISTS idx_stakeholder_nonce_unique;

CREATE UNIQUE INDEX IF NOT EXISTS idx_staking_physical_unique 
ON stake_holder_staking (tx_message_hash, log_index);

-- 6. [token_nft] NFT 信息增强：引入物理坐标并建立唯一索引
ALTER TABLE token_nft ADD COLUMN IF NOT EXISTS "block_number" int8 DEFAULT 0;
ALTER TABLE token_nft ADD COLUMN IF NOT EXISTS "log_index" int4 DEFAULT 0;
ALTER TABLE token_nft ADD COLUMN IF NOT EXISTS "tx_hash" VARCHAR(80);

CREATE UNIQUE INDEX IF NOT EXISTS idx_token_nft_physical_unique 
ON token_nft (tx_hash, log_index);
