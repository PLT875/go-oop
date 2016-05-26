clean:
	go clean

deps:
	go get

test: deps
	go test ./...

ci_test: test

build: deps
	go build -o $(GOPATH)/src/oop

run:
	./oop