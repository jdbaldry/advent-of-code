.ONESHELL:
.DELETE_ON_ERROR:
export SHELL     := bash
export SHELLOPTS := pipefail:errexit
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rule

# Adapted from https://www.thapaliya.com/en/writings/well-documented-makefiles/
.PHONY: help
help: ## Display this help.
help:
	@awk 'BEGIN {FS = ": ##"; printf "Usage:\n  make <target>\n\nTargets:\n"} /^[a-zA-Z0-9_\.\-\/%]+: ##/ { printf "  %-45s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

GO_FILES      = one.go one_test.go two.go two_test.go
GO_TEST_BENCH = Benchmark.*
GO_TEST_RUN   = Test.*

all: go.test.out go.bench.out go.out

go.bench.out go.bench.mem.prof: ## Run benchmarks.
go.bench.out go.bench.mem.prof: go.test
	./$< -test.bench $(GO_TEST_BENCH) -test.benchmem -test.memprofile go.bench.mem.prof -test.run $(GO_TEST_RUN) ./ | tee go.bench.out

go.fuzz.out: ## Run fuzz tests.
go.fuzz.out: go.test
	./$< -test.fuzz 'Fuzz.*' -test.fuzzcachedir $$(mktemp -d) ./ | tee $@

go.test.out: ## Run tests.
go.test.out: go.test
	./$< -test.count=1 -test.v -test.run $(GO_TEST_RUN) | tee $@

go.test: ## Build the test binary.
go.test: $(GO_FILES)
	go test -c ./ -o $@

go.build: ## Build the main binary.
	go build ./ -o $@

go.out: ## Run solution.
go.out: one.go one_test.go two.go two_test.go
	go run ./ | tee $@
