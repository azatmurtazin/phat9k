.PHONY: test test-coverage lint build clean install

BIN := phat9k
COVER_PROFILE := coverage.out
LINT_CONFIG := .golangci.yml

default: test

test:
	go test -v -race ./...

test-coverage:
	go test -v -race -coverprofile=$(COVER_PROFILE) ./...
	go tool cover -html=$(COVER_PROFILE) -o coverage.html

lint:
	golangci-lint run ./...

lint-fix:
	golangci-lint run --fix ./...

govuln:
	govulncheck ./...

build:
	go build -o $(BIN) .

clean:
	rm -f $(BIN) $(COVER_PROFILE) coverage.html

install:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest

fmt:
	go fmt ./...
	go mod tidy

race:
	go test -race -msan ./...

all: fmt lint test build
