GO_BIN=go

.PHONY: build clean fmt test

build:
	$(GO_BIN) build

clean:
	rm -rf *.png && rm fen2image

fmt:
	$(GO_BIN) fmt

test:
	$(GO_BIN) test
