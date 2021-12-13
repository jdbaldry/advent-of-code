scheme.tests: ## Run MIT scheme tests
scheme.tests: lib_test.scm lib.scm
	scheme --quiet < $< > $@

bags.dot: ## Produce a graphviz dot file representing the bag rules
bags.dot: dot.scm lib.scm
	scheme --quiet < $< | sed -e '1d' > $@

include ../../scheme.mk
include ../../common.mk
