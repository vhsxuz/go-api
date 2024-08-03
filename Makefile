run: build
	@./bin/go-api

build:
	@go build -o bin/go-api

test:
	@go test -v ./...
