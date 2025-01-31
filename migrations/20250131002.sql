CREATE TABLE IF NOT EXISTS wallet_addresses(
    "id"         text COLLATE "pg_catalog"."default" NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "address"    varchar COLLATE "pg_catalog"."default" NOT NULL,
    "created_at" int8 NOT NULL,                                -- 创建时间戳
    "status"     int2 DEFAULT 1,                               -- 钱包状态：1-正常 0-禁用
    "remark"     text COLLATE "pg_catalog"."default",          -- 备注信息
    CONSTRAINT "wallet_addresses_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "wallet_addresses_address_unique" UNIQUE ("address")
);

CREATE TABLE IF NOT EXISTS activity_participants_addresses(
    "activity_id" int8 NOT NULL,                               -- 活动ID
    "address"     varchar COLLATE "pg_catalog"."default" NOT NULL, -- 参与者钱包地址
    "join_time"   int8 NOT NULL,                               -- 参与时间戳
    CONSTRAINT "activity_participants_unique" UNIQUE ("activity_id", "address")
);

-- 创建索引提升查询性能
CREATE INDEX IF NOT EXISTS idx_wallet_addresses_address ON wallet_addresses("address");
CREATE INDEX IF NOT EXISTS idx_activity_participants_address ON activity_participants_addresses("address");

