# 本地测试 Gas Estimate 功能指南

## 步骤 1: 编译项目

```bash
cd /home/octavius/theWeb3Work/fishcake-service
make build
```

## 步骤 2: 启动服务

### 方式 A: 只启动 API 服务 (推荐用于测试)
```bash
./fishcake api
```

### 方式 B: 启动完整服务 (包括索引器)
```bash
./fishcake index
```

服务默认会在 `http://localhost:8080` 启动 (端口在 `config.yaml` 中配置)

## 步骤 3: 验证服务启动

```bash
curl http://localhost:8080/ping
```

**期望输出**:
```json
{"message":"pong"}
```

## 步骤 4: 运行测试

### 方式 A: 使用测试脚本 (推荐)

```bash
# 编辑脚本,替换测试地址
nano test_gas_estimate.sh

# 运行测试
./test_gas_estimate.sh
```

### 方式 B: 手动测试单个接口

```bash
# 设置测试地址
export TEST_ADDRESS="0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"

# 测试基础查询
curl "http://localhost:8080/v1/chain_info/sign_info?address=${TEST_ADDRESS}" | jq

# 测试 gas 预估
export RECEIVER="0x1234567890123456789012345678901234567890"
curl "http://localhost:8080/v1/chain_info/sign_info?address=${TEST_ADDRESS}&to=${RECEIVER}" | jq
```

## 步骤 5: 验证结果

### 基础查询 (不带 to 参数)
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "nonce": "123",
    "native_token_gas_limit": "21000",
    "erc20_token_gas_limit": "100000",
    "max_fee_per_gas": "1500000000000",
    "max_priority_fee_per_gas": "30000000000",
    "gas_price": "1530000000000"
    // 注意: 没有 gas_estimate 字段
  }
}
```

### Gas 预估查询 (带 to 参数)
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "nonce": "123",
    "native_token_gas_limit": "21000",
    "erc20_token_gas_limit": "100000",
    "max_fee_per_gas": "1500000000000",
    "max_priority_fee_per_gas": "30000000000",
    "gas_price": "1530000000000",
    "gas_estimate": "25200"  // ✅ 新增字段
  }
}
```

## 常见问题排查

### 问题 1: 服务启动失败

**检查配置文件**:
```bash
cat config.yaml
```

确保以下配置正确:
- `polygon_rpc`: Polygon RPC 节点地址
- `http_port`: HTTP 服务端口 (默认 8080)
- 数据库配置

### 问题 2: gas_estimate 字段为空

**可能原因**:
1. RPC 节点连接失败
2. 提供的地址或参数无效
3. 合约调用会 revert

**检查日志**:
```bash
# 查看服务日志,搜索 "eth_estimateGas failed"
```

**测试 RPC 连接**:
```bash
# 从 config.yaml 获取 RPC URL
RPC_URL=$(grep polygon_rpc config.yaml | awk '{print $2}')

# 测试连接
curl -X POST $RPC_URL \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'
```

### 问题 3: 返回错误 "missing address"

**原因**: 没有提供 address 参数

**解决**: 确保 URL 中包含 `?address=0x...`

### 问题 4: nonce 或 gas price 异常

**可能原因**: RPC 节点数据不同步

**检查区块高度**:
```bash
curl -X POST $RPC_URL \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'
```

## 高级测试

### 测试真实的 ERC20 交易

```bash
# 1. 获取 USDT 合约地址
USDT="0xc2132D05D31c914a87C6611C10748AEb04B58e8F"

# 2. 构造 transfer calldata
# transfer(address to, uint256 amount)
# 函数选择器: 0xa9059cbb
RECEIVER="0x1234567890123456789012345678901234567890"
AMOUNT="0000000000000000000000000000000000000000000000000000000005f5e100"  # 100 USDT (6 decimals)
DATA="0xa9059cbb000000000000000000000000${RECEIVER:2}${AMOUNT}"

# 3. 测试
curl "http://localhost:8080/v1/chain_info/sign_info?address=${TEST_ADDRESS}&to=${USDT}&data=${DATA}&value=0x0" | jq
```

### 使用 Postman 测试

1. 导入 `docs/test_gas_estimate.md` 中的 Postman 集合
2. 设置环境变量:
   - `TEST_ADDRESS`: 你的测试地址
   - `RECEIVER`: 接收方地址
   - `USDT_CONTRACT`: USDT 合约地址
3. 运行测试集合

## 性能测试

```bash
# 测试响应时间
time curl -s "http://localhost:8080/v1/chain_info/sign_info?address=${TEST_ADDRESS}&to=${RECEIVER}" > /dev/null

# 并发测试 (需要安装 apache2-utils)
ab -n 100 -c 10 "http://localhost:8080/v1/chain_info/sign_info?address=${TEST_ADDRESS}"
```

## 下一步

测试通过后:
1. 更新前端代码,使用新的 `gas_estimate` 字段
2. 在测试环境部署
3. 进行端到端测试
4. 部署到生产环境
