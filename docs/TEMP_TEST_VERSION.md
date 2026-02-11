# 临时测试版本说明

## ⚠️ 重要提示
这是一个**临时测试版本**,用于在没有数据库的情况下测试 gas 预估功能。

**测试完成后必须恢复原代码!**

## 📝 修改内容

### 修改的文件
- `api/chain_info/chain_info.go` - signInfo 函数

### 备份文件
- `api/chain_info/chain_info.go.backup` - 原始代码备份

### 主要修改
1. **移除数据库依赖**: 将 `RpcService.GetAccount` 改为直接调用 `eth_getTransactionCount`
2. **添加 nil 检查**: 防止服务未初始化时 panic
3. **保留 gas 预估功能**: 完整的 gas estimate 逻辑保持不变

## 🧪 测试步骤

### 1. 停止并重启服务

```bash
# 停止当前服务 (Ctrl+C)
# 然后重新启动
./fishcake api
```

### 2. 运行测试

```bash
export TEST_ADDRESS="0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
export RECEIVER="0x1234567890123456789012345678901234567890"

# 测试 1: 基础查询 (不带 gas 预估)
echo "=== 测试 1: 基础查询 ==="
curl -s "http://localhost:8189/v1/chain_info/sign_info?address=${TEST_ADDRESS}" | jq

# 测试 2: 带 to 参数 (应该有 gas_estimate)
echo -e "\n=== 测试 2: Gas 预估 ==="
curl -s "http://localhost:8189/v1/chain_info/sign_info?address=${TEST_ADDRESS}&to=${RECEIVER}" | jq

# 测试 3: 带 value 参数
echo -e "\n=== 测试 3: 带 value 的 Gas 预估 ==="
curl -s "http://localhost:8189/v1/chain_info/sign_info?address=${TEST_ADDRESS}&to=${RECEIVER}&value=0xde0b6b3a7640000" | jq
```

### 3. 验证结果

**测试 1 - 应该返回**:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "nonce": "123",
    "native_token_gas_limit": "21000",
    "erc20_token_gas_limit": "100000",
    "max_fee_per_gas": "...",
    "max_priority_fee_per_gas": "...",
    "gas_price": "..."
    // 没有 gas_estimate
  }
}
```

**测试 2 & 3 - 应该返回**:
```json
{
  "code": "200",
  "msg": "success",
  "data": {
    "nonce": "123",
    "native_token_gas_limit": "21000",
    "erc20_token_gas_limit": "100000",
    "max_fee_per_gas": "...",
    "max_priority_fee_per_gas": "...",
    "gas_price": "...",
    "gas_estimate": "25200"  // ✅ 新增字段
  }
}
```

## 🔄 恢复原代码

测试完成后,恢复原代码:

```bash
# 方法 1: 使用备份文件
cp api/chain_info/chain_info.go.backup api/chain_info/chain_info.go

# 方法 2: 使用 git
git checkout api/chain_info/chain_info.go

# 重新编译
make build
```

## 📊 测试重点

关注以下几点:

1. **nonce 是否正确**: 应该是实际的账户 nonce
2. **gas_estimate 是否存在**: 
   - 不带 `to` 参数时不应该有
   - 带 `to` 参数时应该有
3. **gas_estimate 值是否合理**:
   - 简单转账: ~21,000 (+ 20% = ~25,200)
   - ERC20 操作: ~50,000-80,000
4. **max_fee_per_gas 和 max_priority_fee_per_gas**: 应该是合理的 Polygon 网络费用

## 🐛 可能的问题

### 问题 1: RPC 连接失败
**错误**: "RPC client not initialized"
**解决**: 检查 `config.yaml` 中的 `polygon_rpc` 配置

### 问题 2: gas_estimate 为空
**可能原因**:
- RPC 节点不支持 `eth_estimateGas`
- 提供的地址或参数无效
- 交易会 revert

**检查**: 查看服务日志中的 "eth_estimateGas failed" 错误

### 问题 3: nonce 获取失败
**错误**: "failed to get nonce"
**解决**: 确保 RPC 节点正常工作,可以访问 `eth_getTransactionCount`

## 📝 代码差异

主要修改在 `signInfo` 函数中:

**原代码** (依赖数据库):
```go
reqAccount := &account.AccountRequest{
    Chain:   "Polygon",
    Network: "mainnet",
    Address: address,
}
responseAccount, _ := service.BaseService.RpcService.GetAccount(context.Background(), reqAccount)
nonce := responseAccount.Sequence
```

**测试代码** (直接 RPC):
```go
var nonceHex string
err := service.BaseService.Client.CallContext(
    context.Background(),
    &nonceHex,
    "eth_getTransactionCount",
    address,
    "pending",
)
nonceBig, _ := parseHexBig(nonceHex)
nonce := nonceBig.String()
```

## ✅ 测试完成检查清单

- [ ] 基础查询正常返回 (无 gas_estimate)
- [ ] Gas 预估正常返回 (有 gas_estimate)
- [ ] gas_estimate 值合理
- [ ] nonce 正确
- [ ] gas price 合理
- [ ] 恢复原代码
- [ ] 重新编译验证
