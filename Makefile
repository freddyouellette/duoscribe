build:
	go build -v -o ./bin ./...

test: build
	go test ./...

format:
	go fmt ./...

lint:
	golangci-lint run