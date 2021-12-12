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
	@awk 'BEGIN {FS = ": ##"; printf "Usage:\n  make <target>\n\nTargets:\n"} /^[a-zA-Z0-9_\.\-\/%]+: ##/ { printf "  %-15s %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

ALL = $(dir $(shell find . -wholename './*/*/Makefile' | sort))

.PHONY: all
all: ## Run all solutions
all: $(ALL)

define all
.PHONY: $1
$1: ## Run $1 solutions
$1:
	$(MAKE) -C $1 all
endef

$(foreach dir, $(ALL), $(eval $(call all, $(dir))))
