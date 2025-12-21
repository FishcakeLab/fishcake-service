CREATE TABLE IF NOT EXISTS mining_info
(
    "id"                  text COLLATE "pg_catalog"."default" NOT NULL DEFAULT replace((uuid_generate_v4())::text, '-'::text, ''::text),
    "address"             varchar COLLATE "pg_catalog"."default",
    "mined_amount"        public.uint256,
    "mined_fishcakepower" public.uint256,
    "last_mint_time"      public.uint256 DEFAULT '0',
    CONSTRAINT "mining_info_pkey" PRIMARY KEY ("id")
)
;

CREATE INDEX IF NOT EXISTS idx_mining_info_address
ON mining_info("address");

INSERT INTO mining_info
(
    "id",
    "address",
    "mined_amount",
    "mined_fishcakepower",
    "last_mint_time"
)
SELECT
    replace((uuid_generate_v4())::text, '-'::text, ''::text) AS id,
    ai."business_account" AS address,
    SUM(ai."mined_amount") AS mined_amount,
    SUM(ai."mined_amount") AS mined_fishcakepower,
    '0'::public.uint256 AS last_mint_time
FROM activity_info ai
WHERE ai."business_account" IS NOT NULL
  AND length(ai."business_account") > 5
GROUP BY ai."business_account"
HAVING NOT EXISTS (
    SELECT 1 
    FROM mining_info mi
    WHERE mi."address" = ai."business_account"
);
