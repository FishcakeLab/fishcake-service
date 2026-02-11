# Gas Estimate 功能说明

## 概述

在 `signInfo` 接口中新增了 gas 预估功能,解决了固定 gas limit 带来的问题:
- 复杂交易 gas limit 不够用
- 简单交易 gas limit 过高导致余额要求过高

## API 接口

### 请求参数

**必填参数:**
- `address`: 发送方地址

**可选参数 (用于 gas 预估):**
- `to`: 目标地址 (合约地址或转账接收地址)
- `data`: 交易的 calldata (hex 格式,如 `0x095ea7b3...`)
- `value`: 转账金额 (hex 格式,如 `0x0`,默认为 `0x0`)

### 返回值

```json
{
  "nonce": "123",
  "native_token_gas_limit": "21000",
  "erc20_token_gas_limit": "100000",
  "max_fee_per_gas": "1500000000000",
  "max_priority_fee_per_gas": "30000000000",
  "gas_price": "1530000000000",
  "gas_estimate": "65000"  // 新增字段,仅当提供 to 参数时返回
}
```

## 使用场景

### 场景 1: 普通查询 (不需要 gas 预估)

前端只需要获取 nonce 和 gas price 信息:

```
GET /v1/chain_info/sign_info?address=0x1234...
```

返回:
```json
{
  "nonce": "123",
  "native_token_gas_limit": "21000",
  "erc20_token_gas_limit": "100000",
  "max_fee_per_gas": "1500000000000",
  "max_priority_fee_per_gas": "30000000000",
  "gas_price": "1530000000000"
}
```

### 场景 2: ERC20 转账 (需要 gas 预估)

前端需要预估 ERC20 transfer 所需的 gas:

```
GET /v1/chain_info/sign_info?address=0x1234...&to=0xUSDT_CONTRACT&data=0xa9059cbb000000000000000000000000RECEIVER_ADDRESS0000000000000000000000000000000000000000000000000000000000000064
```

返回:
```json
{
  "nonce": "123",
  "native_token_gas_limit": "21000",
  "erc20_token_gas_limit": "100000",
  "max_fee_per_gas": "1500000000000",
  "max_priority_fee_per_gas": "30000000000",
  "gas_price": "1530000000000",
  "gas_estimate": "52000"  // 实际预估值 + 20% buffer
}
```

### 场景 3: ERC20 Approve (需要 gas 预估)

前端需要预估 approve 所需的 gas:

```
GET /v1/chain_info/sign_info?address=0x1234...&to=0xUSDT_CONTRACT&data=0x095ea7b3000000000000000000000000SPENDER_ADDRESS00000000000000000000000000000000000000000000000000000000000f4240
```

返回:
```json
{
  "gas_estimate": "48000"  // approve 通常需要 40k-50k gas
}
```

### 场景 4: 复杂合约调用 (需要 gas 预估)

前端需要预估复杂合约交互所需的 gas:

```
GET /v1/chain_info/sign_info?address=0x1234...&to=0xCOMPLEX_CONTRACT&data=0x...&value=0x0
```

返回:
```json
{
  "gas_estimate": "250000"  // 复杂交易可能需要更高的 gas
}
```

### 场景 5: 原生代币转账 (需要 gas 预估)

前端需要预估 POL 转账所需的 gas:

```
GET /v1/chain_info/sign_info?address=0x1234...&to=0xRECEIVER&value=0xde0b6b3a7640000
```

返回:
```json
{
  "gas_estimate": "21000"  // 原生代币转账固定 21000
}
```

## 前端使用建议

### 推荐流程

1. **用户准备发起交易时**,前端先调用 `signInfo` 并传入完整的交易参数 (`to`, `data`, `value`)
2. **获取 `gas_estimate`** 后,使用该值作为交易的 `gasLimit`
3. **如果 `gas_estimate` 为空** (预估失败或未提供 `to` 参数),则使用默认值:
   - 原生代币转账: 使用 `native_token_gas_limit` (21000)
   - ERC20 交易: 使用 `erc20_token_gas_limit` (100000)

### 示例代码

```javascript
// 1. 构造交易参数
const txParams = {
  from: userAddress,
  to: contractAddress,
  data: encodedData,  // 如 approve/transfer 的 calldata
  value: '0x0'
};

// 2. 调用 signInfo 获取 gas 预估
const params = new URLSearchParams({
  address: txParams.from,
  to: txParams.to,
  data: txParams.data,
  value: txParams.value
});

const response = await fetch(`/v1/chain_info/sign_info?${params}`);
const signInfo = await response.json();

// 3. 使用预估的 gas limit
const gasLimit = signInfo.gas_estimate || signInfo.erc20_token_gas_limit;

// 4. 构造最终交易
const finalTx = {
  ...txParams,
  nonce: signInfo.nonce,
  gasLimit: gasLimit,
  maxFeePerGas: signInfo.max_fee_per_gas,
  maxPriorityFeePerGas: signInfo.max_priority_fee_per_gas
};
```

## 技术细节

### Gas 预估逻辑

1. 后端调用 `eth_estimateGas` RPC 方法获取预估值
2. 在预估值基础上增加 **20% buffer** 以避免 out of gas
3. 如果预估失败,不影响其他数据的返回,只是 `gas_estimate` 字段为空

### 为什么需要 buffer?

- 区块链状态可能在预估和实际执行之间发生变化
- 某些合约的 gas 消耗可能因状态不同而变化
- 20% buffer 是行业标准做法,可以覆盖大部分情况

### 错误处理

- 如果 `eth_estimateGas` 调用失败 (如合约 revert),后端会记录错误日志但不会影响接口返回
- 前端应该检查 `gas_estimate` 是否存在,如果不存在则使用默认值

## 优势

1. **精确的 gas limit**: 根据实际交易内容预估,避免过高或过低
2. **降低余额要求**: 简单交易不再需要为高 gas limit 预留大量余额
3. **提高成功率**: 复杂交易可以获得足够的 gas limit
4. **向后兼容**: 不影响现有的调用方式,可选参数
