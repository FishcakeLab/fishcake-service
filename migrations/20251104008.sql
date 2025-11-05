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
