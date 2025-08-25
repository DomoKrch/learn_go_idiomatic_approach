.DEFAULT_GOAL := build

.PHONY: clean build vet fmt

clean:
	go clean

build: vet
	go build

vet: fmt
	go vet ./...

fmt:
	go fmt ./...


