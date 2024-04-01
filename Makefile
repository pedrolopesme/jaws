GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD)fmt
BINARY_NAME=$(GOPATH)/bin/jaws
BINARY_UNIX=$(BINARY_NAME)_unix

help:
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo ""
	@echo "  build:      Build the jaws binary"
	@echo "  test:       Run the jaws tests"
	@echo "  run:        Run the jaws binary with the specified parameters"
	@echo "  coverage:   Check the jaws tests coverage"
	@echo "  help:       Display this help message"
	@echo ""


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

test:
	@echo "Running jaws tests"
	$(GOTEST) -v ./...

coverage:
	@echo "Checking jaws tests coverage"
	$(GOTEST) -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
