#!/bin/bash

# 重启服务并测试的快速脚本

echo "🔄 重启服务..."
echo "请在运行 ./fishcake api 的终端按 Ctrl+C 停止服务"
echo "然后运行: ./fishcake api"
echo ""
echo "等待服务启动后,运行以下测试:"
echo ""
echo "export TEST_ADDRESS=\"0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb\""
echo "export RECEIVER=\"0x1234567890123456789012345678901234567890\""
echo ""
echo "# 测试 1: 基础查询"
echo "curl -s \"http://localhost:8189/v1/chain_info/sign_info?address=\${TEST_ADDRESS}\" | jq"
echo ""
echo "# 测试 2: Gas 预估"
echo "curl -s \"http://localhost:8189/v1/chain_info/sign_info?address=\${TEST_ADDRESS}&to=\${RECEIVER}\" | jq"
