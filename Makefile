all: test

test:
	go test -race -coverprofile=coverage.txt -covermode=atomic

.PHONY: all test