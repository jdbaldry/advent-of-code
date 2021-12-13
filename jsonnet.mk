ALL += main.json
JSONNET_ARGS ?=

main.json: ## Evaluate the Jsonnet solution
main.json: main.jsonnet input.txt
	jsonnet $(JSONNET_ARGS) $< > $@
