GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOFMT=$(GOCMD)fmt
BINARY_NAME=$(GOPATH)/bin/jaws
BINARY_UNIX=$(BINARY_NAME)_unix

build: 
	@echo "Building jaws"
	$(GOBUILD) -o $(BINARY_NAME) -v

test:
	@echo "Running jaws tests"
	$(GOTEST) -v ./...

clean: 
	@echo "Cleaning jaws"
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run:
	@echo "Running jaws with params: $(filter-out $@,$(MAKECMDGOALS))"
	$(GOBUILD) -o $(BINARY_NAME)
	./$(BINARY_NAME) $(filter-out $@,$(MAKECMDGOALS))

fmt:
	@echo "Running gofmt for all project files"
	$(GOFMT) -w *.go