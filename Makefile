build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

install:
	go install ./cmd/hexlet-path-size/

uninstall:
	go clean -i ./cmd/hexlet-path-size/

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

fmt:
	golangci-lint fmt

test:
	go test

.PHONY: build install uninstall lint lint-fix fmt test