#!/bin/bash

# 快速测试 Gas Estimate 功能
# 使用前请先启动服务: ./fishcake api

# 配置
TEST_ADDRESS="0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
RECEIVER="0x1234567890123456789012345678901234567890"
BASE_URL="http://localhost:8080"

echo "🧪 快速测试 Gas Estimate 功能"
echo "================================"
echo ""

# 测试 1: 基础查询
echo "📋 测试 1: 基础查询 (应该没有 gas_estimate)"
echo "---"
curl -s "${BASE_URL}/v1/chain_info/sign_info?address=${TEST_ADDRESS}" | jq '{
  nonce: .data.nonce,
  gas_estimate: .data.gas_estimate // "❌ 无 (正常)",
  max_fee_per_gas: .data.max_fee_per_gas,
  has_all_fields: (.data.nonce != null and .data.max_fee_per_gas != null)
}'
echo ""
echo ""

# 测试 2: 带 to 参数
echo "📋 测试 2: 带 to 参数 (应该有 gas_estimate)"
echo "---"
curl -s "${BASE_URL}/v1/chain_info/sign_info?address=${TEST_ADDRESS}&to=${RECEIVER}" | jq '{
  nonce: .data.nonce,
  gas_estimate: .data.gas_estimate // "❌ 预估失败",
  max_fee_per_gas: .data.max_fee_per_gas,
  has_gas_estimate: (.data.gas_estimate != null)
}'
echo ""
echo ""

# 测试 3: 完整参数
echo "📋 测试 3: 完整参数 (to + value)"
echo "---"
curl -s "${BASE_URL}/v1/chain_info/sign_info?address=${TEST_ADDRESS}&to=${RECEIVER}&value=0xde0b6b3a7640000" | jq '{
  nonce: .data.nonce,
  gas_estimate: .data.gas_estimate // "❌ 预估失败",
  gas_estimate_number: (.data.gas_estimate | tonumber? // 0),
  is_reasonable: ((.data.gas_estimate | tonumber? // 0) > 20000 and (.data.gas_estimate | tonumber? // 0) < 30000)
}'
echo ""
echo ""

echo "✅ 测试完成!"
echo ""
echo "💡 验证要点:"
echo "  - 测试 1: has_all_fields 应该为 true, gas_estimate 应该为 null"
echo "  - 测试 2: has_gas_estimate 应该为 true"
echo "  - 测试 3: is_reasonable 应该为 true (gas 在 20k-30k 之间)"
