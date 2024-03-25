GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD)fmt
BINARY_NAME=$(GOPATH)/bin/jaws
BINARY_UNIX=$(BINARY_NAME)_unix

help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  build: Build the jaws binary"
	@echo "  test: Run the jaws tests"
	@echo "  run: Run the jaws binary with the specified parameters"
	@echo "  help: Display this help message"

build: 
	@echo "Building jaws"
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	@echo "Running jaws tests"
	$(GOTEST) -v ./...

run:
	@echo "Running jaws with params: $(filter-out $@,$(MAKECMDGOALS))"
	$(GOBUILD) -o $(BINARY_NAME)
	$(BINARY_NAME) $(filter-out $@,$(MAKECMDGOALS))