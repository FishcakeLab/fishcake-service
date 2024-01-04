CREATE EXTENSION "uuid-ossp";

CREATE TABLE IF NOT EXISTS block_headers (
    hash        VARCHAR PRIMARY KEY,
    parent_hash VARCHAR NOT NULL UNIQUE,
    number      BIGINT NOT NULL UNIQUE,
    timestamp   BIGINT NOT NULL UNIQUE CHECK (timestamp > 0),
    rlp_bytes   VARCHAR NOT NULL
);
CREATE INDEX IF NOT EXISTS block_headers_timestamp ON block_headers(timestamp);
CREATE INDEX IF NOT EXISTS block_headers_number ON block_headers(number);


CREATE TABLE IF NOT EXISTS activity_info
(
    id                   UUID PRIMARY KEY DEFAULT uuid_generate_v4(), -- 唯一标识符，主键
    activity_id          BIGINT,                                      -- 活动ID
    business_account     TEXT,                                        -- 发起人账户（商家0x...）
    business_name        TEXT,                                        -- 商家名称
    activity_content     TEXT,                                        -- 活动内容
    latitude_longitude   TEXT,                                        -- 经纬度，以_分割经纬度
    activity_create_time BIGINT,                                      -- 活动创建时间
    activity_deadline    BIGINT,                                      -- 活动结束时间
    drop_type            SMALLINT,                                    -- 奖励规则：1表示平均获得  2表示随机
    drop_number          BIGINT,                                      -- 奖励份数
    min_drop_amt         BIGINT,                                      -- 当dropType为1时，填0，为2时，填每份最少领取数量
    max_drop_amt         BIGINT,                                      -- 当dropType为1时，填每份奖励数量，为2时，填每份最多领取数量
    token_contract_addr  TEXT                                         -- Token合约地址，例如USDT合约地址
);

CREATE TABLE IF NOT EXISTS activity_info_ext
(
    id                            UUID PRIMARY KEY  DEFAULT uuid_generate_v4(), -- 唯一标识符，主键
    activity_id                   BIGINT,           -- 活动ID
    already_drop_amts             BIGINT,           -- 总共已奖励数量
    already_drop_number           BIGINT,           -- 总共已奖励份数
    business_mined_amt            BIGINT,           -- 商家获得的挖矿奖励
    business_mined_withdrawed_amt BIGINT,           -- 商家已提取的挖矿奖励
    activity_status               SMALLINT          -- 活动状态：1表示进行中  2表示已结束
);

CREATE TABLE IF NOT EXISTS token_nft
(
    id                            UUID PRIMARY KEY  DEFAULT uuid_generate_v4(), -- 唯一标识符，主键
    token_id                      BIGINT,           -- TokenID
    address                       VARCHAR,          -- Token 地址
    contract_address              VARCHAR,          -- 合约地址
    token_url                     VARCHAR,          -- NFT URL
    token_amount                  BIGINT,           -- token 数量
    timestamp                     BIGINT            -- nft 链上时间戳
);
