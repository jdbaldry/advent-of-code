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

DAYS := $(wildcard 0[1-9])
DAYS += $(wildcard 1[0-9])
DAYS += $(wildcard 2[0-5])

.PHONY: all
all: ## Make all 2022 targets
	for day in $(DAYS); do
		$(MAKE) -C "$${day}" go.test.out go.bench.out go.out
	done
