CREATE TABLE IF NOT EXISTS stake_holder_staking (
    "id"                text COLLATE "pg_catalog"."default"
                        NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "user_address"      varchar COLLATE "pg_catalog"."default" NOT NULL,
    "amount"            "public"."uint256" NOT NULL,
    "staking_type"      int2,
    "start_time"        int8,
    "end_time"          int8,
    "token_id"          int8,
    "nft_apr"           int8,
    "is_auto_renew"     bool,
    "message_nonce"     int8 NOT NULL,
    "tx_message_hash"   varchar COLLATE "pg_catalog"."default",
    "staking_reward"    "public"."uint256",
    "staking_status"    int2 DEFAULT 0,  -- 0=质押中, 1=已结束
    "create_time"       timestamptz DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "stake_holder_staking_pkey" PRIMARY KEY ("id")
);

-- 为查询优化加索引
CREATE INDEX IF NOT EXISTS idx_stakeholder_nonce
ON stake_holder_staking ("user_address", "message_nonce");


-- ============================================
-- 🧩 Schema Fix: ensure activity_id is unique
-- ============================================

-- 1️⃣ 确保 activity_id 列非空（避免唯一索引报错）
ALTER TABLE activity_info
ALTER COLUMN activity_id SET NOT NULL;

-- 2️⃣ 若已存在相同 activity_id 的重复数据，先人工处理（这句仅提示，不执行删除）
-- SELECT activity_id, COUNT(*) FROM activity_info GROUP BY activity_id HAVING COUNT(*) > 1;

-- 3️⃣ 检查是否已有唯一约束（避免重复添加时报错）
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1
        FROM pg_constraint
        WHERE conname = 'activity_info_activity_id_key'
    ) THEN
        ALTER TABLE activity_info
        ADD CONSTRAINT activity_info_activity_id_key UNIQUE (activity_id);
    END IF;
END
$$;

-- ✅ 现在 activity_id 是全表唯一，可以支持 ON CONFLICT(activity_id)