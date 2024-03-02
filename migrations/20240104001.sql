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
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" cascade;

CREATE TABLE IF NOT EXISTS block_headers
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


CREATE TABLE IF NOT EXISTS block_listener
(
    "guid"         text COLLATE "pg_catalog"."default" NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "block_number" "public"."uint256"                           DEFAULT 0,
    "created"      int4,
    "updated"      int4,
    CONSTRAINT "block_listener_pkey" PRIMARY KEY ("guid"),
    CONSTRAINT "block_listener_created_check" CHECK (created > 0),
    CONSTRAINT "block_listener_updated_check" CHECK (updated > 0)
);


CREATE TABLE IF NOT EXISTS contract_events
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
);

CREATE TABLE IF NOT EXISTS token_nft
(
    "id"               text COLLATE "pg_catalog"."default" NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "token_id"         int8,
    "address"          varchar COLLATE "pg_catalog"."default",
    "contract_address" varchar COLLATE "pg_catalog"."default",
    "token_url"        varchar COLLATE "pg_catalog"."default",
    "token_amount"     int8,
    "timestamp"        int8,
    CONSTRAINT "token_nft_pkey" PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS activity_info
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
    "min_drop_amt"         "public"."uint256",
    "max_drop_amt"         "public"."uint256",
    "token_contract_addr"  text COLLATE "pg_catalog"."default",
    "activity_status"      int2,
    "already_drop_number"  int8,
    CONSTRAINT "activity_info_pkey" PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS activity_info_ext
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
CREATE TABLE IF NOT EXISTS drop_info
(
    "id"          text COLLATE "pg_catalog"."default" NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "activity_id" int8,
    "address"     varchar COLLATE "pg_catalog"."default",
    "drop_amount" "public"."uint256",
    "timestamp"   int8,
    CONSTRAINT "drop_info_pkey" PRIMARY KEY ("id")
)
;
