2015 := $(dir $(shell find . -wholename './2015/*/Makefile' | sort))
2020 := $(dir $(shell find . -wholename './2020/*/Makefile' | sort))
2021 := $(dir $(shell find . -wholename './2021/*/Makefile' | sort))
2022 := $(dir $(shell find . -wholename './2022/*/Makefile' | sort))

ALL    := $(2015) $(2020) $(2021) $(2022)
BROKEN := $(dir $(shell find . -wholename './*/*/Makefile.broken' | sort))

include common.mk

.PHONY: info
info: ## Display informational values.
	@printf "2015:\t$(2015)\n"
	@printf "2020:\t$(2020)\n"
	@printf "2021:\t$(2021)\n"
	@printf "2022:\t$(2022)\n"
	@printf "ALL:\t$(ALL)\n"
	@printf "BROKEN:\t$(BROKEN)\n"

define all
.PHONY: $(1)
$(1): ## Run $(1) solutions.
$(1):
	$(MAKE) -C $(1) all
endef

$(foreach dir,$(ALL),$(eval $(call all,$(dir))))
$(foreach dir,$(ALL),$(eval $(call all,$(dir:./%/=%))))

.PHONY: 2015
2015: ## Run all 2015 solutions.
2015: $(2015)

.PHONY: 2020
2020: ## Run all 2020 solutions.
2020: $(2020)

.PHONY: 2021
2021: ## Run all 2021 solutions.
2021: $(2021)

.PHONY: 2022
2022: ## Run all 2022 solutions.
2022: $(2022)
