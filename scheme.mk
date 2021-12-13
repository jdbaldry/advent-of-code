ALL += scheme.out

scheme.out: ## Evaluate the MIT scheme solution
scheme.out: main.scm lib.scm input.txt
	scheme --quiet < $< > $@

.PHONY: repl
repl: ## Load the MIT scheme library in a REPL
repl:
	rlwrap scheme --load lib.scm
