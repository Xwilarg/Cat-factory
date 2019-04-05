all: test build

test:
	go test

build:
	go tool pack r cat.a cat_impl.go cat.go factory.go