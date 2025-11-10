CREATE TABLE IF NOT EXISTS token_sent
(
    "id"              text COLLATE "pg_catalog"."default" NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "address"         text COLLATE "pg_catalog"."default",
    "token_address"           varchar COLLATE "pg_catalog"."default",
    "description"    text COLLATE "pg_catalog"."default",
    "amount"          "public"."uint256",
    "timestamp"        int8
);


CREATE TABLE IF NOT EXISTS token_received
(
    "id"              text COLLATE "pg_catalog"."default" NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "address"         text COLLATE "pg_catalog"."default",
    "token_address"           varchar COLLATE "pg_catalog"."default",
    "description"    text COLLATE "pg_catalog"."default",
    "amount"          "public"."uint256",
    "timestamp"        int8
);
