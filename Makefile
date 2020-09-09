GOCMD=go
GOTEST=$(GOCMD) test
GOLINT=golangci-lint

all: lint test

test:
	$(GOTEST) -v ./...

test.race:
	$(GOTEST) -v -race ./...

lint:
	$(GOLINT) run -v ./... -c .golangci.yml
