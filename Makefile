all: test

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic  -timeout 2s

.PHONY: all test