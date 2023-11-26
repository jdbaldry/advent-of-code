ALL += go.out

GO_FILES := $(shell find ./ -name '*.go')

go.bench.out go.bench.mem.prof: ## Run benchmarks.
go.bench.out go.bench.mem.prof: go.test $(GO_FILES) input.txt
	./$< -test.bench 'Benchmark.*' -test.benchmem -test.memprofile go.bench.mem.prof ./ | tee go.bench.out

go.fuzz.out: ## Run fuzz tests.
go.fuzz.out: go.test $(GO_FILES) input.txt
	./$< -test.fuzz 'Fuzz.*' -test.fuzzcachedir $$(mktemp -d) ./ | tee $@

go.test.out: ## Run tests.
go.test.out: go.test $(GO_FILES) input.txt
	./$< -test.count=1 -test.v | tee $@

go.test: ## Build the test binary.
go.test: $(GO_FILES) input.txt
	go test -c ./ -o $@

go.build: ## Build the main binary.
go.build: $(shell find ./ -name '*.go') input.txt
	go build -o $@ ./

go.out: ## Run the Go solution.
go.out: go.build $(shell find ./ -name '*.go') input.txt
	./$< | tee $@
