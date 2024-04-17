all: test run build

test:
	go test ./...

run:
	go run ./cmd

build:
	go build -o bin/main ./cmd/main.go
