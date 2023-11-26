ALL += main.json

JSONNET_ARGS  := -J ../../lib
JSONNET_FILES := $(shell find ./ -name '*.jsonnet')

main.json: ## Evaluate the Jsonnet solution.
main.json: main.jsonnet $(JSONNET_FILES) input.txt
	jsonnet $(JSONNET_ARGS) $< | tee $@

%.json: ## Evaluate any other Jsonnet.
%.json: %.jsonnet $(JSONNET_FILES) input.txt
	jsonnet $(JSONNET_ARGS) $< | tee $@
