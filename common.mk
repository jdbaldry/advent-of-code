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

GIT_ROOT := $(shell git rev-parse --show-toplevel)

input.txt: ## Fetch the input for the puzzle.
ifndef AOC_SESSION_COOKIE
	$(error the AOC_SESSION_COOKIE environment variable is required to fetch the input file)
endif
	curl -sLo $(@) \
		-H 'Cookie: session=$(AOC_SESSION_COOKIE)' \
		https://adventofcode.com/$(eval YEAR := $(notdir $(realpath ..)))$(YEAR)/day/$(eval DAY := $(notdir $(CURDIR)))$(DAY:0%=%)/input

.PHONY: all
all: ## Run all solutions
all: $(ALL)
