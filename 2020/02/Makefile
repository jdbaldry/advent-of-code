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

main.json: ## Evaluate the Jsonnet solution
main.json: main.jsonnet input.txt parser-combinators.libsonnet
	jsonnet $< > $@

scheme.tests: ## Run MIT scheme tests
scheme.tests: lib_test.scm lib.scm
	scheme --quiet < $< > $@

scheme.out: ## Evaluate the MIT scheme solution
scheme.out: main.scm lib.scm input.txt
	scheme --quiet < $< > $@

.PHONY: repl
repl: ## Load the MIT scheme library in a REPL
repl:
	rlwrap scheme --load lib.scm
