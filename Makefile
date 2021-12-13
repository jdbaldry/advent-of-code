ALL = $(dir $(shell find . -wholename './*/*/Makefile' | sort))

define all
.PHONY: $1
$1: ## Run $1 solutions
$1:
	$(MAKE) -C $1 all
endef

$(foreach dir, $(ALL), $(eval $(call all, $(dir))))

include common.mk
