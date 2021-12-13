ALL += go.out

GO_DEPS ?= main input.txt

main: main.go
main: ## Build the Go binary
	go build -o $@ ./

go.out: ## Run the Go solution
go.out: $(GO_DEPS)
	./main > $@
