.ONESHELL:
.DELETE_ON_ERROR:
SHELL       := bash
SHELLOPTS   := -euf -o pipefail
MAKEFLAGS   += --warn-undefined-variables
MAKEFLAGS   += --no-builtin-rule

# Adapted from https://suva.sh/posts/well-documented-makefiles/
.PHONY: help
help: ## Display this help
help:
	@awk 'BEGIN {FS = ": ##"; printf "Usage:\n  make <target>\n\nTargets:\n"} /^[a-zA-Z0-9_\.\-\/%]+: ##/ { printf "  %-45s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: test
test: ## Run tests
test:
	scheme --quiet < lib_test.scm

.PHONY: scheme
scheme: ## Run the MIT scheme solution
scheme:
	scheme --quiet < main.scm

bags.dot: ## Produce a graphviz dot file representing the bag rules
bags.dot: dot.scm lib.scm
	scheme --quiet < $< | sed -e '1d' > $@

.PHONY: repl
repl: ## Load the MIT scheme solution in a REPL
repl:
	rlwrap scheme --load lib.scm
