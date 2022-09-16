all: test lint

test:
	go test -cover ./...

coverage:
	go test -covermode count -coverprofile=coverage.out && go tool cover -func=coverage.out \
		| perl -ne 's{^'$$(awk '/^module/{print$$2}' go.mod)'/}{}; print unless m{^total:|100\.0%$$}' \
		| sort -rn -k3

lint:
	golangci-lint run
