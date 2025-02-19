GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"


EVENT_ABI_ARTIFACT := ./abis/FishcakeEventManager.sol/FishcakeEventManager.json
NFT_ABI_ARTIFACT := ./abis/NftManager.sol/NftManager.json

fishcake:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/fishcake

build:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/fishcake

clean:
	rm fishcake

test:
	go test -v ./...

lint:
	golangci-lint run ./...

bindings: binding-event binding-nft

binding-event:
	$(eval temp := $(shell mktemp))
	cat $(EVENT_ABI_ARTIFACT) | jq .abi \
	| abigen --pkg abi \
	--abi - \
	--out event/polygon/abi/fish_cake_event_manager.go \
	--type FishcakeEventManager \
	rm $(temp)

binding-nft:
	$(eval temp := $(shell mktemp))
	cat $(NFT_ABI_ARTIFACT) | jq .abi \
	| abigen --pkg abi \
	--abi - \
	--out event/polygon/abi/nft_manager.go \
	--type NftManager \
	rm $(temp)


.PHONY: \
	fishcake \
	build \
	bindings \
	clean \
	test \
	lint
