.PHONY: pre-commit

generate-all:
	go generate ./...

build-all:
	go build ./...

test:
	go test -v ./...

goreport:
	goreportcard-cli -v -t 100.0

pre-commit: generate-all build-all goreport test
