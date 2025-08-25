.DEFAULT_GOAL := build

.PHONY: build vet fmt

build: vet
	go build

vet: fmt
	go vet ./...

fmt:
	go fmt ./...
