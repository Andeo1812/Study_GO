.PHONY: all run download run_tests

all: run

TARGET = ./Project/main.go

run:
	go run ${TARGET}

download:
	go mod download

run_tests:
	go test ./... -cover -coverpkg ./...
