.INTERMEDIATE: i
i: ## Intermediate tiny file name for minified Jsonnet solution
i: input.txt
	cp input.txt i

minified.jsonnet: ## Produce a minified Jsonnet file
minified.jsonnet: main.jsonnet
	$(SED) \
		-e '2d' \
		-e 's/std/p/g' \
		-e '1i local p = std;' \
		-e '/^\/\/.*$$/d' \
		-e '/^[[:space:]]*$$/d' \
		-e 's/input.txt/i/g' \
		-e 's/input/i/g' \
		-e 's/depths/d/g' \
		-e 's/line/l/g' \
		-e 's/zip/z/g' \
		-e 's/sum/s/g' \
		-e 's/window/w/g' \
		-e 's/countIncreases/c/g' \
		$< | tr -d '\n' | jsonnetfmt - > $@

minified.json: ## Evaulate minified Jsonnet solution
minified.json: minified.jsonnet i
	jsonnet $< > $@

DEPFILES := main.jsonnet.d
$(DEPFILES):
include $(wildcard $(DEPFILES))

include ../../jsonnet.mk
include ../../common.mk
