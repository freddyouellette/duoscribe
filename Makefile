.PHONY: *

COVERAGE_REQUIREMENT := 90

build:
	mkdir -p ./bin && go build -v -o ./bin ./...

utest:
	go test `go list ./... | grep -v /integration` -coverprofile=coverage.out

itest:
	go test `go list ./... | grep /integration`

test: utest itest

format:
	go fmt ./...

lint:
	golangci-lint run

ci-lint:
	actionlint

coverage:
	@go tool cover -html=coverage.out -o coverage.html
	@go tool cover -func=coverage.out
	@COVERAGE=`go tool cover -func=coverage.out | grep "^total:" | grep -Eom 1 '[0-9]+' | head -1`;\
	echo Coverage report available at ./coverage.html;\
	if [ "$$COVERAGE" -lt "${COVERAGE_REQUIREMENT}" ]; then\
		echo "Test coverage $$COVERAGE% does not meet minimum ${COVERAGE_REQUIREMENT}% requirement";\
		exit 1;\
	else\
		echo "Test Coverage $$COVERAGE% (OK)";\
	fi

check-pipeline: build test lint ci-lint coverage