.PHONY: tidy build run

tidy:
	go mod tidy

build: tidy
	go build -o bin/haste ./cmd/main.go

run: tidy
	go run ./cmd/main.go