.ONESHELL:
.DELETE_ON_ERROR:
export SHELL     := bash
export SHELLOPTS := pipefail:errexit
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rule

# Adapted from https://suva.sh/posts/well-documented-makefiles/
.PHONY: help
help: ## Display this help
help:
	@awk 'BEGIN {FS = ": ##"; printf "Usage:\n  make <target>\n\nTargets:\n"} /^[a-zA-Z0-9_\.\-\/%]+: ##/ { printf "  %-45s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: all
all: ## Run all solutions
all: main.json

main.json: ## Evaluate the Jsonnet solution
main.json: main.jsonnet input.txt test.txt
	jsonnet $< > $@

.INTERMEDIATE: i
i: ## Intermediate tiny file name for minified Jsonnet solution
i: input.txt
	cp input.txt i

minified.jsonnet: ## Produce a minified Jsonnet file
minified.jsonnet: main.jsonnet
	sed \
		-e '2d' \
		-e 's/std/p/g' \
		-e '1i local p = std;' \
		-e '/^\/\/.*$$/d' \
		-e '/^[[:space:]]*$$/d' \
		-e 's/input.txt/i/g' \
		-e 's/input/i/g' \
		-e 's/depths/d/g' \
		-e 's/line/l/g' \
		-e 's/zip/z/g' \
		-e 's/sum/s/g' \
		-e 's/window/w/g' \
		-e 's/countIncreases/c/g' \
		$< | tr -d '\n' | jsonnetfmt - > $@

minified.json: ## Evaulate minified Jsonnet solution
minified.json: minified.jsonnet i
	jsonnet $< > $@
