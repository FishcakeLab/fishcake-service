# Gas Estimate API 测试命令

## 前置条件
确保服务已启动: `make build && ./fishcake api`

## 测试地址配置
```bash
# 替换为你的实际测试地址
export TEST_ADDRESS="0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
export USDT_CONTRACT="0xc2132D05D31c914a87C6611C10748AEb04B58e8F"
export RECEIVER="0x1234567890123456789012345678901234567890"
```

## 测试命令

### 1. 基础查询 (不带 gas 预估)
```bash
curl "http://localhost:8080/v1/chain_info/sign_info?address=${TEST_ADDRESS}" | jq
```

**期望结果**: 返回 nonce 和 gas price 信息,不包含 `gas_estimate` 字段

---

### 2. 原生代币转账 Gas 预估
```bash
curl "http://localhost:8080/v1/chain_info/sign_info?address=${TEST_ADDRESS}&to=${RECEIVER}&value=0xde0b6b3a7640000" | jq
```

**期望结果**: 
- 包含 `gas_estimate` 字段
- 值应该约为 21000

---

### 3. ERC20 Transfer Gas 预估
```bash
# transfer(address to, uint256 amount)
# 转账 100 USDT (6 decimals)
DATA="0xa9059cbb000000000000000000000000${RECEIVER:2}0000000000000000000000000000000000000000000000000000000005f5e100"

curl "http://localhost:8080/v1/chain_info/sign_info?address=${TEST_ADDRESS}&to=${USDT_CONTRACT}&data=${DATA}&value=0x0" | jq
```

**期望结果**: 
- 包含 `gas_estimate` 字段
- 值应该约为 50,000 - 80,000

---

### 4. ERC20 Approve Gas 预估
```bash
# approve(address spender, uint256 amount)
# 授权无限额度
SPENDER="0x1111111254EEB25477B68fb85Ed929f73A960582"
DATA="0x095ea7b3000000000000000000000000${SPENDER:2}ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"

curl "http://localhost:8080/v1/chain_info/sign_info?address=${TEST_ADDRESS}&to=${USDT_CONTRACT}&data=${DATA}&value=0x0" | jq
```

**期望结果**: 
- 包含 `gas_estimate` 字段
- 值应该约为 45,000 - 65,000

---

### 5. 简单查询 (只提供 to 地址)
```bash
curl "http://localhost:8080/v1/chain_info/sign_info?address=${TEST_ADDRESS}&to=${RECEIVER}" | jq
```

**期望结果**: 
- 包含 `gas_estimate` 字段
- 值应该约为 21000 (简单转账)

---

## 快速测试 (一键执行所有测试)

```bash
# 设置测试地址
TEST_ADDRESS="0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
USDT_CONTRACT="0xc2132D05D31c914a87C6611C10748AEb04B58e8F"
RECEIVER="0x1234567890123456789012345678901234567890"

echo "=== 测试 1: 基础查询 ==="
curl -s "http://localhost:8080/v1/chain_info/sign_info?address=${TEST_ADDRESS}" | jq '.data.gas_estimate // "无 gas_estimate"'

echo -e "\n=== 测试 2: 原生代币转账 ==="
curl -s "http://localhost:8080/v1/chain_info/sign_info?address=${TEST_ADDRESS}&to=${RECEIVER}&value=0xde0b6b3a7640000" | jq '.data.gas_estimate'

echo -e "\n=== 测试 3: ERC20 Transfer ==="
DATA="0xa9059cbb000000000000000000000000${RECEIVER:2}0000000000000000000000000000000000000000000000000000000005f5e100"
curl -s "http://localhost:8080/v1/chain_info/sign_info?address=${TEST_ADDRESS}&to=${USDT_CONTRACT}&data=${DATA}&value=0x0" | jq '.data.gas_estimate'

echo -e "\n=== 测试 4: ERC20 Approve ==="
SPENDER="0x1111111254EEB25477B68fb85Ed929f73A960582"
DATA="0x095ea7b3000000000000000000000000${SPENDER:2}ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
curl -s "http://localhost:8080/v1/chain_info/sign_info?address=${TEST_ADDRESS}&to=${USDT_CONTRACT}&data=${DATA}&value=0x0" | jq '.data.gas_estimate'
```

---

## 验证要点

✅ **测试 1**: 不应该有 `gas_estimate` 字段  
✅ **测试 2**: `gas_estimate` ≈ 21000  
✅ **测试 3**: `gas_estimate` ≈ 50000-80000  
✅ **测试 4**: `gas_estimate` ≈ 45000-65000  
✅ **所有测试**: 都应该返回 `nonce`, `max_fee_per_gas`, `max_priority_fee_per_gas` 等字段

---

## 错误处理测试

### 测试无效地址
```bash
curl "http://localhost:8080/v1/chain_info/sign_info?address=invalid&to=${RECEIVER}" | jq
```

### 测试无效 data
```bash
curl "http://localhost:8080/v1/chain_info/sign_info?address=${TEST_ADDRESS}&to=${USDT_CONTRACT}&data=0xinvalid" | jq
```

**期望结果**: 即使 gas 预估失败,也应该返回其他字段,只是 `gas_estimate` 为空或不存在

---

## 使用 Postman/Insomnia 测试

### 请求配置
- **Method**: GET
- **URL**: `http://localhost:8080/v1/chain_info/sign_info`
- **Query Parameters**:
  - `address`: (必填) 发送方地址
  - `to`: (可选) 目标地址
  - `data`: (可选) 交易 calldata
  - `value`: (可选) 转账金额

### 示例集合
导入以下 JSON 到 Postman:

```json
{
  "info": {
    "name": "Gas Estimate API Tests",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "1. Basic Query",
      "request": {
        "method": "GET",
        "url": {
          "raw": "http://localhost:8080/v1/chain_info/sign_info?address={{TEST_ADDRESS}}",
          "query": [
            {"key": "address", "value": "{{TEST_ADDRESS}}"}
          ]
        }
      }
    },
    {
      "name": "2. Native Transfer",
      "request": {
        "method": "GET",
        "url": {
          "raw": "http://localhost:8080/v1/chain_info/sign_info?address={{TEST_ADDRESS}}&to={{RECEIVER}}&value=0xde0b6b3a7640000",
          "query": [
            {"key": "address", "value": "{{TEST_ADDRESS}}"},
            {"key": "to", "value": "{{RECEIVER}}"},
            {"key": "value", "value": "0xde0b6b3a7640000"}
          ]
        }
      }
    }
  ],
  "variable": [
    {"key": "TEST_ADDRESS", "value": "0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"},
    {"key": "RECEIVER", "value": "0x1234567890123456789012345678901234567890"},
    {"key": "USDT_CONTRACT", "value": "0xc2132D05D31c914a87C6611C10748AEb04B58e8F"}
  ]
}
```
