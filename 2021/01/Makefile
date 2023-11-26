include ../../jsonnet.mk
include ../../common.mk

minified.jsonnet: ## Produce a minified Jsonnet file
minified.jsonnet: main.jsonnet i
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
		-e 's/lshift/ls/g' \
		-e 's/zip/z/g' \
		-e 's/sum/s/g' \
		-e 's/windows/w/g' \
		-e 's/countIncreases/c/g' \
		$< | tr -d '\n' | jsonnetfmt - > $@

i: input.txt
	cp $< $@
