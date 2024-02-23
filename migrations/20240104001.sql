DO
$$
BEGIN
        IF
NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'uint256') THEN
CREATE DOMAIN UINT256 AS NUMERIC
    CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
ELSE
ALTER DOMAIN UINT256 DROP CONSTRAINT uint256_check;
ALTER DOMAIN UINT256 ADD
    CHECK (VALUE >= 0 AND VALUE < POWER(CAST(2 AS NUMERIC), CAST(256 AS NUMERIC)) AND SCALE(VALUE) = 0);
END IF;
END
$$;
DROP
EXTENSION IF EXISTS "uuid-ossp" cascade;
CREATE
EXTENSION "uuid-ossp";

DROP TABLE IF EXISTS block_headers;
CREATE TABLE block_headers
(
    "guid"        text COLLATE "pg_catalog"."default"    NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "hash"        varchar COLLATE "pg_catalog"."default" NOT NULL,
    "parent_hash" varchar COLLATE "pg_catalog"."default" NOT NULL,
    "number"      "public"."uint256"                     NOT NULL,
    "timestamp"   int4                                   NOT NULL,
    "rlp_bytes"   varchar COLLATE "pg_catalog"."default" NOT NULL,
    CONSTRAINT "block_headers_pkey" PRIMARY KEY ("guid"),
    CONSTRAINT "block_headers_timestamp_check" CHECK ("timestamp" > 0)
);

DROP TABLE IF EXISTS block_listener;
CREATE TABLE block_listener
(
    "guid"         text COLLATE "pg_catalog"."default" NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "block_number" "public"."uint256"                           DEFAULT 0,
    "created"      int4,
    "updated"      int4,
    CONSTRAINT "block_listener_pkey" PRIMARY KEY ("guid"),
    CONSTRAINT "block_listener_created_check" CHECK (created > 0),
    CONSTRAINT "block_listener_updated_check" CHECK (updated > 0)
);

DROP TABLE IF EXISTS contract_events;
CREATE TABLE contract_events
(
    "guid"             text COLLATE "pg_catalog"."default"    NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "block_hash"       varchar COLLATE "pg_catalog"."default" NOT NULL,
    "contract_address" varchar COLLATE "pg_catalog"."default" NOT NULL,
    "transaction_hash" varchar COLLATE "pg_catalog"."default" NOT NULL,
    "log_index"        int4                                   NOT NULL,
    "block_number"     "public"."uint256"                     NOT NULL,
    "event_signature"  varchar COLLATE "pg_catalog"."default" NOT NULL,
    "timestamp"        int4                                   NOT NULL,
    "rlp_bytes"        varchar COLLATE "pg_catalog"."default" NOT NULL,
    CONSTRAINT "contract_events_pkey" PRIMARY KEY ("guid")
)
;

DROP TABLE IF EXISTS token_nft;
CREATE TABLE token_nft
(
    "id"                   text COLLATE "pg_catalog"."default" NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "activity_id"          int8,
    "business_account"     text COLLATE "pg_catalog"."default",
    "business_name"        text COLLATE "pg_catalog"."default",
    "activity_content"     text COLLATE "pg_catalog"."default",
    "latitude_longitude"   text COLLATE "pg_catalog"."default",
    "activity_create_time" int8,
    "activity_deadline"    int8,
    "drop_type"            int2,
    "drop_number"          int8,
    "min_drop_amt"         int8,
    "max_drop_amt"         int8,
    "token_contract_addr"  text COLLATE "pg_catalog"."default",
    "activity_status"      int2,
    CONSTRAINT "activity_info_pkey" PRIMARY KEY ("id")
)
;

DROP TABLE IF EXISTS activity_info;
CREATE TABLE activity_info
(
    "id"                   text COLLATE "pg_catalog"."default" NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "activity_id"          int8,
    "business_account"     text COLLATE "pg_catalog"."default",
    "business_name"        text COLLATE "pg_catalog"."default",
    "activity_content"     text COLLATE "pg_catalog"."default",
    "latitude_longitude"   text COLLATE "pg_catalog"."default",
    "activity_create_time" int8,
    "activity_deadline"    int8,
    "drop_type"            int2,
    "drop_number"          int8,
    "min_drop_amt"         int8,
    "max_drop_amt"         int8,
    "token_contract_addr"  text COLLATE "pg_catalog"."default",
    "activity_status"      int2,
    CONSTRAINT "activity_info_pkey" PRIMARY KEY ("id")
)
;

DROP TABLE IF EXISTS activity_info;
CREATE TABLE activity_info
(
    "id"                            text COLLATE "pg_catalog"."default" NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "activity_id"                   int8,
    "already_drop_amts"             int8,
    "already_drop_number"           int8,
    "business_mined_amt"            int8,
    "business_mined_withdrawed_amt" int8,
    "activity_status"               int2,
    CONSTRAINT "activity_info_ext_pkey" PRIMARY KEY ("id")
)
;
