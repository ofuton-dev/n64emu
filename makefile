NAME := n64emu
BINDIR := ./build
VERSION := $(shell git describe --tags 2>/dev/null)
LDFLAGS := -X 'main.version=$(VERSION)'
TAGS := debug

.PHONY: run
run:
	@go run ./cmd/

.PHONY: build
build:
	@go build --tags=$(TAGS) -o $(BINDIR)/darwin-amd64/$(NAME).app -ldflags "$(LDFLAGS)" ./cmd/

.PHONY: build-linux
build-linux:
	@GOOS=linux GOARCH=amd64 go build --tags=$(TAGS) -o $(BINDIR)/linux-amd64/$(NAME) -ldflags "$(LDFLAGS)" ./cmd/

.PHONY: build-windows
build-windows:
	@GOOS=windows GOARCH=amd64 go build --tags=$(TAGS) -o $(BINDIR)/windows-amd64/$(NAME).exe -ldflags "$(LDFLAGS)" ./cmd/

.PHONY: clean
clean:
	@-rm -rf $(BINDIR)

.PHONY: misspell
misspell:
	@hash misspell > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get -u github.com/client9/misspell/cmd/misspell; \
	fi
	@misspell -w $(shell find . -name "*.go")

.PHONY: test
test:
	@go test --tags=$(TAGS) -v ./...

.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)