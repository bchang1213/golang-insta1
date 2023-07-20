GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=bin/app
BINARY_UNIX=$(BINARY_NAME)_unix

all: test
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
	rm -rf dist
	rm -rf .code-coverage
clean-all:
	make clean
	rm -rf node_modules
	rm -rf vendor
dev:
	bash run_dev.sh
