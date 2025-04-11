-- table system_drop_info drop_type 1 Fishcake Wallet 🍥 2 Congrats First Claim 🍥
CREATE TABLE IF NOT EXISTS system_drop_info
(
    "id"               text COLLATE "pg_catalog"."default"    NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "address"          varchar COLLATE "pg_catalog"."default",
    "drop_amount"      "public"."uint256",
    "drop_type"        int2,
    "timestamp"        int8,
    "transaction_hash" varchar COLLATE "pg_catalog"."default" NOT NULL,
    "status"           int2,
    CONSTRAINT "system_drop_info_pkey" PRIMARY KEY ("id")
);


