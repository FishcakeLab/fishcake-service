# Fishcake Service API 接口文档

Base URL: `{host}/v1`

## 通用响应格式

所有接口统一返回格式：

```json
{
  "success": true,
  "msg": "success",
  "code": "0000",
  "obj": {}
}
```

错误时：

```json
{
  "success": false,
  "msg": "error message",
  "code": "1001",
  "obj": null
}
```

分页接口返回格式（`obj` 内容）：

```json
{
  "result": [],
  "total": 100,
  "currentPage": 1,
  "pageSize": 10
}
```

---

## 一、历史 Received / Dropped 数据

与 FishcakeEventManager 合约相关的交互记录。

### 1.1 活动列表（Activity List）

`GET /v1/activity/list`

查询活动列表，对应合约 `ActivityAdd` 事件创建的活动。

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| pageNum | int | 是 | 页码，从 1 开始 |
| pageSize | int | 是 | 每页条数 |
| activityId | int | 否 | 活动 ID |
| activityStatus | string | 否 | 状态：`1`=进行中，`2`=已结束，`3`=已过期 |
| businessName | string | 否 | 商家名称（模糊搜索） |
| businessAccount | string | 否 | 商家地址（模糊搜索） |
| tokenContractAddr | string | 否 | Token 合约地址（模糊搜索） |
| latitude | string | 否 | 纬度（需配合 longitude、scope 使用） |
| longitude | string | 否 | 经度 |
| scope | string | 否 | 搜索范围（米） |
| activityFilter | string | 否 | NFT 过滤：`1`=Pro NFT 有效，`2`=Basic NFT 有效 |

响应 `obj` 字段（分页格式）：

```json
{
  "result": [
    {
      "id": "abc123",
      "activityId": 1,
      "businessAccount": "0x...",
      "businessName": "Shop A",
      "activityContent": "...",
      "latitudeLongitude": "31.23,121.47",
      "activityCreateTime": 1700000000,
      "activityDeadline": 1700086400,
      "dropType": 1,
      "dropNumber": 100,
      "minDropAmt": "1000000000000000000",
      "maxDropAmt": "5000000000000000000",
      "tokenContractAddr": "0x...",
      "activityStatus": 1,
      "alreadyDropNumber": 50,
      "basicDeadline": 1700086400,
      "proDeadline": 1700086400,
      "returnAmount": "0",
      "minedAmount": "0",
      "isExpired": false
    }
  ],
  "total": 100,
  "currentPage": 1,
  "pageSize": 10
}
```

### 1.2 活动详情（Activity Info）

`GET /v1/activity/info`

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| activityId | int | 是 | 活动 ID |

响应 `obj` 字段：同上单条 activity 对象。

---

### 1.3 Drop 记录列表

`GET /v1/drop/list`

查询 drop/receive 记录，对应合约 `Drop` 事件。

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| pageNum | int | 是 | 页码，从 1 开始 |
| pageSize | int | 是 | 每页条数 |
| address | string | 否 | 钱包地址（不区分大小写） |
| dropType | string | 否 | `1`=用户领取（received），`2`=商家发放（dropped） |

说明：
- `dropType=1` 时会自动合并 system_drop_info（系统空投）记录
- 返回数据会 JOIN activity_info 补充 `tokenContractAddr`、`businessName`、`returnAmount`、`minedAmount`

响应 `obj` 字段（分页格式）：

```json
{
  "result": [
    {
      "id": "abc123",
      "activityId": 1,
      "address": "0x...",
      "dropAmount": "1000000000000000000",
      "dropType": 1,
      "timestamp": 1700000000,
      "transactionHash": "0x...",
      "eventSignature": "0x...",
      "tokenContractAddr": "0x...",
      "businessName": "Shop A",
      "returnAmount": "0",
      "minedAmount": "0",
      "blockNumber": 12345678,
      "logIndex": 0
    }
  ],
  "total": 50,
  "currentPage": 1,
  "pageSize": 10
}
```

---

### 1.4 Token Sent 历史

`GET /v1/token/sent/list`

查询 token 发送记录。包含 `ActivityAdd` 事件产生的商家发送记录，以及 FCC Token 的 ERC20 Transfer 记录。

类型区分：

- `ActivityAdd: ...`
  表示活动创建时，商家为活动预存 token
- `FCC Transfer Sent`
  表示 FCC Token 的普通转出记录（不包含与 FishcakeEventManager 合约直接交互而被过滤掉的记录）

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| address | string | 是 | 钱包地址 |
| tokenType | string | 否 | Token 合约地址过滤 |
| lastTimestamp | uint64 | 否 | 游标：上一页最后一条的 timestamp，默认 `0` |
| lastId | string | 否 | 游标：上一页最后一条的 id |
| limit | int | 否 | 每页条数，默认 `10` |

说明：使用游标分页（cursor-based pagination），按 `(timestamp DESC, id DESC)` 排序。

响应格式（非通用格式，直接返回）：

```json
{
  "data": [
    {
      "id": "abc123",
      "address": "0x...",
      "token_address": "0x...",
      "description": "activity content or ERC20 Token Transfer Sent",
      "amount": "5000000000000000000",
      "timestamp": 1700000000,
      "tx_hash": "0x...",
      "log_index": 0
    }
  ],
  "nextTimestamp": 1699999999,
  "nextId": "abc122"
}
```

> 注意：前端下次请求时将 `nextTimestamp` 和 `nextId` 作为参数传入即可翻页。

---

### 1.5 Token Received 历史

`GET /v1/token/received/list`

查询 token 接收记录。包含 `Drop` 事件产生的用户领取记录、`ActivityFinish` 事件产生的商家退款记录，以及 FCC Token 的 ERC20 Transfer 记录。

类型区分：

- `Drop Receive: ...`
  表示用户收到活动投放的 token
- `ActivityFinish Return: ...`
  表示活动结束时，系统返还给活动创建者的剩余 token
- `FCC Transfer Received`
  表示 FCC Token 的普通转入记录（不包含与 FishcakeEventManager 合约直接交互而被过滤掉的记录）

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| address | string | 是 | 钱包地址 |
| tokenType | string | 否 | Token 合约地址过滤 |
| lastTimestamp | uint64 | 否 | 游标：上一页最后一条的 timestamp，默认 `0` |
| lastId | string | 否 | 游标：上一页最后一条的 id |
| limit | int | 否 | 每页条数，默认 `10` |

响应格式：同 1.4 Token Sent，字段一致。

---

### 1.6 Mining Record 历史

`GET /v1/mining/record/list`

查询挖矿流水记录。包含两类：

- `activity_finish`：活动结束时给商家增加的挖矿收益
- `mint_booster_nft`：Mint Booster NFT 时消耗的 Fishcake Power

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| address | string | 是 | 钱包地址 |
| recordType | string | 否 | 记录类型过滤：`activity_finish` 或 `mint_booster_nft` |
| lastTimestamp | uint64 | 否 | 游标：上一页最后一条的 timestamp，默认 `0` |
| lastId | string | 否 | 游标：上一页最后一条的 id |
| limit | int | 否 | 每页条数，默认 `10` |

说明：使用游标分页（cursor-based pagination），按 `(timestamp DESC, id DESC)` 排序。

响应格式（非通用格式，直接返回）：

```json
{
  "data": [
    {
      "id": "abc123",
      "address": "0x...",
      "record_type": "activity_finish",
      "mined_amount_delta": "1000000000000000000",
      "power_increase": "1000000000000000000",
      "power_decrease": "0",
      "activity_id": 123,
      "token_id": null,
      "description": "ActivityFinish reward",
      "timestamp": 1700000000,
      "tx_hash": "0x...",
      "log_index": 0,
      "block_number": 12345678
    }
  ],
  "nextTimestamp": 1699999999,
  "nextId": "abc122"
}
```

> `mint_booster_nft` 类型记录中，`activity_id` 为空，`token_id` 有值；`activity_finish` 类型则相反。

---

## 二、Earnings（用户个人数据）

### 2.1 用户 Staking 记录

`GET /v1/staking/stakingInfo`

查询用户的质押记录列表，包含 staking / renew / unlock 状态。

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| address | string | 是 | 钱包地址 |
| status | int | 否 | `0`=质押中，`1`=已结束（已 unlock）。不传返回全部 |
| page | int | 否 | 页码，默认 `1` |
| size | int | 否 | 每页条数，默认 `10`，最大 `100` |

响应 `obj` 字段：

```json
{
  "list": [
    {
      "tokenId": 5,
      "amount": "1000000000000000000000",
      "stakingType": 1,
      "startTime": 1700000000,
      "endTime": 1702592000,
      "nftApr": 3,
      "isAutoRenew": true,
      "messageNonce": 1,
      "txMessageHash": "0x...",
      "stakingReward": "50000000000000000000",
      "stakingStatus": 0,
      "createTime": 1700000000
    }
  ],
  "total": 5,
  "page": 1,
  "size": 10
}
```

字段说明：

| 字段 | 说明 |
|------|------|
| stakingType | 质押类型：`1`=30天/3%，`2`=60天/6%，`3`=90天/9%，`4`=180天/15% |
| stakingStatus | `0`=质押中（staking），`1`=已结束（unlocked） |
| isAutoRenew | 是否自动续期（renew） |
| stakingReward | 已领取的奖励金额（unlock 后写入），质押中为 `"0"` |
| nftApr | Booster NFT 额外 APR |
| tokenId | 绑定的 Booster NFT Token ID，`0` 表示无绑定 |
| txMessageHash | 交易哈希 |

---

### 2.2 用户 Mining 累计数量

`GET /v1/activity/minedAmount`

查询用户作为商家的累计挖矿产出（基于已结束活动的 mined_amount 汇总）。

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| address | string | 是 | 钱包地址 |
| month | int | 否 | `1`=仅统计近 30 天，不传为全部 |

响应 `obj` 字段：

```json
{
  "userAddress": "0x...",
  "totalMined": "500000000000000000000",
  "month": false
}
```

---

### 2.3 用户 Mining 详情 & Fishcake Power

`GET /v1/activity/miningInfo`

查询用户的实时挖矿信息，包含 Fishcake Power。

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| address | string | 是 | 钱包地址 |

响应 `obj` 字段：

```json
{
  "id": "abc123",
  "address": "0x...",
  "minedAmount": "500000000000000000000",
  "minedFishCakePower": "300000000000000000000",
  "lastMintTime": "1700000000"
}
```

字段说明：

| 字段 | 说明 |
|------|------|
| minedAmount | 累计挖矿总量（实时） |
| minedFishCakePower | 当前可用 Fishcake Power（挖矿累计 - 已消耗铸造 Booster NFT） |
| lastMintTime | 最近一次铸造 Booster NFT 的时间戳 |

---

### 2.4 用户 Booster NFT 列表

`GET /v1/booster/list`

查询用户持有的 Booster NFT 列表。

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| address | string | 是 | 钱包地址 |
| pageNum | int | 否 | 页码，默认 `1` |
| pageSize | int | 否 | 每页条数，默认 `50`，最大 `100` |

响应 `obj` 字段（分页格式）：

```json
{
  "result": [
    {
      "tokenId": 5,
      "nftType": 2
    }
  ],
  "total": 3,
  "currentPage": 1,
  "pageSize": 50
}
```

字段说明：

| 字段 | 说明 |
|------|------|
| nftType | NFT 类型。`nftType >= 10` 表示已被 staking 绑定使用（原始类型 + 10） |

---

## 三、Leaderboard（排行榜）

### 3.1 Staking 排行榜（质押总量）

`GET /v1/staking/stakingRank`

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| month | int | 否 | `1`=本月排行，不传为全部（All-time） |

响应 `obj` 字段：

```json
[
  {
    "rank": 1,
    "userAddress": "0x...",
    "totalStake": "10000000000000000000000",
    "month": false
  }
]
```

---

### 3.2 Staking 已领取奖励排行

`GET /v1/staking/claimedRank`

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| month | int | 否 | `1`=本月排行，不传为全部 |

响应 `obj` 字段：

```json
[
  {
    "rank": 1,
    "userAddress": "0x...",
    "claimed": "500000000000000000000",
    "month": false
  }
]
```

---

### 3.3 Staking 总奖励排行（已领取 + 预估未领取）

`GET /v1/staking/totalRewardRank`

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| month | int | 否 | `1`=本月排行，不传为全部 |

响应 `obj` 字段：

```json
[
  {
    "rank": 1,
    "userAddress": "0x...",
    "totalReward": "800000000000000000000",
    "claimed": "500000000000000000000",
    "unclaimed": "300000000000000000000",
    "month": false
  }
]
```

说明：`unclaimed` 为动态计算的预估值，基于当前时间和质押 APR 实时计算。

---

### 3.4 Mining 排行榜

`GET /v1/activity/miningRank`

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| month | int | 否 | `1`=近 30 天排行，不传为全部（All-time） |

响应 `obj` 字段：

```json
[
  {
    "rank": 1,
    "businessAccount": "0x...",
    "totalMined": "1000000000000000000000",
    "month": false
  }
]
```

---

## 附录：金额说明

所有金额字段（`amount`、`totalStake`、`minedAmount`、`dropAmount` 等）均为 Wei 单位的字符串，前端需除以 `10^18` 转换为可读数值。
