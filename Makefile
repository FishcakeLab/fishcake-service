# Fishcake Service Build Configuration
# =============================================================================

# 获取当前 Git 提交的完整哈希值，用于版本追踪
GITCOMMIT := $(shell git rev-parse HEAD)
# 获取当前 Git 提交的 Unix 时间戳
GITDATE := $(shell git show -s --format='%ct')

# 构建链接器标志，将 Git 信息注入到编译后的二进制文件中
LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
# 最终的 ldflags 参数，传递给 go build
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

# =============================================================================
# 合约 ABI 文件路径定义
# =============================================================================

# FishcakeEventManager 合约 ABI 文件路径
EVENT_ABI_ARTIFACT := ./abis/FishcakeEventManager.sol/FishcakeEventManager.json
# NftManager 合约 ABI 文件路径
NFT_ABI_ARTIFACT := ./abis/NftManager.sol/NftManager.json
# StakingManager 合约 ABI 文件路径
STAKING_ABI_ARTIFACT := ./abis/StakingManager.sol/StakingManager.json

# =============================================================================
# 构建目标
# =============================================================================

# 默认目标：编译 fishcake 服务二进制文件
fishcake:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/fishcake

# build 目标：与 fishcake 目标相同，提供更通用的命名
build:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/fishcake

# 清理编译产物，删除生成的二进制文件
clean:
	rm fishcake

# 运行所有包的单元测试，-v 输出详细信息
test:
	go test -v ./...

# 使用 golangci-lint 对所有包进行代码静态检查
lint:
	golangci-lint run ./...

# =============================================================================
# 合约 ABI Go Binding 生成
# =============================================================================

# 一键生成所有合约的 Go binding 代码
bindings: binding-event binding-nft binding-staking

# 从 FishcakeEventManager 合约 ABI 生成 Go binding
# 使用 jq 提取 abi 字段，通过 abigen 生成类型安全的 Go 代码
binding-event:
	$(eval temp := $(shell mktemp))
	cat $(EVENT_ABI_ARTIFACT) | jq .abi \
	| abigen --pkg abi \
	--abi - \
	--out event/polygon/abi/fish_cake_event_manager.go \
	--type FishcakeEventManager \
	rm $(temp)

# 从 NftManager 合约 ABI 生成 Go binding
binding-nft:
	$(eval temp := $(shell mktemp))
	cat $(NFT_ABI_ARTIFACT) | jq .abi \
	| abigen --pkg abi \
	--abi - \
	--out event/polygon/abi/nft_manager.go \
	--type NftManager \
	rm $(temp)

# 从 StakingManager 合约 ABI 生成 Go binding
binding-staking:
	$(eval temp := $(shell mktemp))
	cat $(STAKING_ABI_ARTIFACT) | jq .abi \
	| abigen --pkg abi \
	--abi - \
	--out event/polygon/abi/staking_manager.go \
	--type StakingManager \
	rm $(temp)

# =============================================================================
# 声明伪目标，确保这些目标不会与同名文件冲突
# =============================================================================
.PHONY: \
	fishcake \
	build \
	bindings \
	clean \
	test \
	lint
