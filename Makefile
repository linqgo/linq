all: test lint

test:
	go test -cover ./...

lint:
	golangci-lint run
