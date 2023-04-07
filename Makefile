.PHONY: *

build:
	go build -v -o ./bin ./...

utest:
	go test `go list ./... | grep -v /integration` -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo Coverage report available at ./coverage.html

itest:
	go test `go list ./... | grep /integration`

test: utest itest

format:
	go fmt ./...

lint:
	golangci-lint run

ci-lint:
	actionlint

check-pipeline: build test lint ci-lint