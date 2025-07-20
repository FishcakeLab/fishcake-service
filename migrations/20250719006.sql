-- table queue_tx
CREATE TABLE IF NOT EXISTS queue_tx
(
    "id"              text COLLATE "pg_catalog"."default" NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "raw_tx"          varchar COLLATE "pg_catalog"."default",
    "result"          varchar COLLATE "pg_catalog"."default",
    "transaction_hash" varchar COLLATE "pg_catalog"."default" NOT NULL,
    "status"           int2,
    "timestamp"        int8
);