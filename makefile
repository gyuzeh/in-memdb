.PHONY: build unit-test benchmark-test

build:
	go build -v ./...

unit-test:
	go test -v ./...

benchmark-test:
	go test -bench=. ./...