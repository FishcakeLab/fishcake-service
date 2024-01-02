GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

fishcake:
	env GO111MODULE=on go build -v $(LDFLAGS) ./cmd/fishcake

clean:
	rm fishcake

test:
	go test -v ./...

lint:
	golangci-lint run ./...

.PHONY: \
	fishcake \
	bindings \
	bindings-scc \
	clean \
	test \
	lint
