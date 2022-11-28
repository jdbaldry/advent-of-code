DAYS = $(shell find . -wholename './*/*/Makefile' | sort)
2015 = $(dir $(shell find . -wholename './2015/*/Makefile' | sort))
2020 = $(dir $(shell find . -wholename './2020/*/Makefile' | sort))
ALL  = $(2020)
2021 = $(dir $(shell find . -wholename './2021/*/Makefile' | sort))
ALL += $(2021)

include common.mk

.PHONY: info
info: ## Display informational values.
	@printf "2015:\t$(2015)\n"
	@printf "2020:\t$(2020)\n"
	@printf "2021:\t$(2021)\n"
	@printf "ALL:\t$(ALL)\n"
	@printf "DAYS:\t$(DAYS)\n"

define all
.PHONY: $(1)
$(1): ## Run $(1) solutions.
$(1):
	$(MAKE) -C $(1) all
endef

$(foreach dir,$(ALL),$(eval $(call all,$(dir))))

.PHONY: 2015
2015: ## Run all 2015 solutions.
2015: $(2015)
	:

.PHONY: 2020
2020: ## Run all 2020 solutions.
2020: $(2020)
	:

.PHONY: 2021
2021: ## Run all 2021 solutions.
: $()
	:
