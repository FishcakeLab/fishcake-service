#!/bin/bash

# Gas Estimate 功能测试脚本
# 使用方法: ./test_gas_estimate.sh

# 配置
BASE_URL="http://localhost:8080"
API_ENDPOINT="/v1/chain_info/sign_info"

# 颜色输出
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# 测试用地址 (请替换为实际地址)
TEST_ADDRESS="0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb"
USDT_CONTRACT="0xc2132D05D31c914a87C6611C10748AEb04B58e8F"  # Polygon USDT
RECEIVER_ADDRESS="0x1234567890123456789012345678901234567890"

echo -e "${YELLOW}========================================${NC}"
echo -e "${YELLOW}Gas Estimate 功能测试${NC}"
echo -e "${YELLOW}========================================${NC}\n"

# 测试 1: 基础查询 (不带 gas 预估)
echo -e "${GREEN}测试 1: 基础查询 (不带 gas 预估)${NC}"
echo "请求: GET ${API_ENDPOINT}?address=${TEST_ADDRESS}"
echo ""
curl -s "${BASE_URL}${API_ENDPOINT}?address=${TEST_ADDRESS}" | jq '.'
echo -e "\n"

# 测试 2: ERC20 Transfer (需要 gas 预估)
echo -e "${GREEN}测试 2: ERC20 Transfer Gas 预估${NC}"
# transfer(address to, uint256 amount) 的函数签名
# 0xa9059cbb + 接收地址(32字节) + 金额(32字节)
TRANSFER_DATA="0xa9059cbb000000000000000000000000${RECEIVER_ADDRESS:2}0000000000000000000000000000000000000000000000000000000000000064"
echo "请求: GET ${API_ENDPOINT}?address=${TEST_ADDRESS}&to=${USDT_CONTRACT}&data=${TRANSFER_DATA}&value=0x0"
echo ""
curl -s "${BASE_URL}${API_ENDPOINT}?address=${TEST_ADDRESS}&to=${USDT_CONTRACT}&data=${TRANSFER_DATA}&value=0x0" | jq '.'
echo -e "\n"

# 测试 3: ERC20 Approve (需要 gas 预估)
echo -e "${GREEN}测试 3: ERC20 Approve Gas 预估${NC}"
# approve(address spender, uint256 amount) 的函数签名
# 0x095ea7b3 + spender地址(32字节) + 金额(32字节)
SPENDER_ADDRESS="0x1111111254EEB25477B68fb85Ed929f73A960582"  # 1inch Router
APPROVE_DATA="0x095ea7b3000000000000000000000000${SPENDER_ADDRESS:2}ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff"
echo "请求: GET ${API_ENDPOINT}?address=${TEST_ADDRESS}&to=${USDT_CONTRACT}&data=${APPROVE_DATA}&value=0x0"
echo ""
curl -s "${BASE_URL}${API_ENDPOINT}?address=${TEST_ADDRESS}&to=${USDT_CONTRACT}&data=${APPROVE_DATA}&value=0x0" | jq '.'
echo -e "\n"

# 测试 4: 原生代币转账 (需要 gas 预估)
echo -e "${GREEN}测试 4: 原生代币 (POL) 转账 Gas 预估${NC}"
VALUE_HEX="0xde0b6b3a7640000"  # 1 POL = 1e18 wei
echo "请求: GET ${API_ENDPOINT}?address=${TEST_ADDRESS}&to=${RECEIVER_ADDRESS}&value=${VALUE_HEX}"
echo ""
curl -s "${BASE_URL}${API_ENDPOINT}?address=${TEST_ADDRESS}&to=${RECEIVER_ADDRESS}&value=${VALUE_HEX}" | jq '.'
echo -e "\n"

# 测试 5: 只提供 to 地址 (简单转账预估)
echo -e "${GREEN}测试 5: 简单转账 Gas 预估 (只提供 to)${NC}"
echo "请求: GET ${API_ENDPOINT}?address=${TEST_ADDRESS}&to=${RECEIVER_ADDRESS}"
echo ""
curl -s "${BASE_URL}${API_ENDPOINT}?address=${TEST_ADDRESS}&to=${RECEIVER_ADDRESS}" | jq '.'
echo -e "\n"

echo -e "${YELLOW}========================================${NC}"
echo -e "${YELLOW}测试完成${NC}"
echo -e "${YELLOW}========================================${NC}"
echo ""
echo -e "${GREEN}检查要点:${NC}"
echo "1. 测试 1 应该返回基础信息,不包含 gas_estimate 字段"
echo "2. 测试 2-5 应该返回 gas_estimate 字段,值应该合理:"
echo "   - ERC20 Transfer: 约 50,000 - 80,000"
echo "   - ERC20 Approve: 约 45,000 - 65,000"
echo "   - 原生代币转账: 约 21,000"
echo "3. 所有测试都应该返回 nonce, max_fee_per_gas, max_priority_fee_per_gas 等字段"
