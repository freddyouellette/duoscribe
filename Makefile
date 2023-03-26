build:
	go build -v -o ./bin ./...

test: build
	go test ./...

build-binary:
	go build -o bin/duolingo_extractor cmd/duolingo_extractor/main.go 