.PHONY: test test-verbose test-race test-cover cover-html lint build

test:
	go test ./...

test-verbose:
	go test -v ./...

test-race:
	go test -race ./...

test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

cover-html: test-cover
	go tool cover -html=coverage.out

lint:
	golangci-lint run ./...

build:
	go build ./...
