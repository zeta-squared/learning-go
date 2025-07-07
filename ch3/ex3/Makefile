.DEFAULT_GOAL := build

.PHONY: mod vet fmt build clean

mod:
	go mod tidy

fmt: mod
	go fmt ./...

vet: fmt
	go vet ./...

build: fmt
	go build

clean:
	go clean ./...
