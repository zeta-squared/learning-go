.DEFAULT_GOAL := build

.PHONY: mod fmt vet build clean

mod:
	go mod tidy

fmt: mod
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build

clean:
	go clean
