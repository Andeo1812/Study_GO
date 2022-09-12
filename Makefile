.PHONY: all launch download run_tests

all: launch

TARGET = ./project/main.go

launch:
	go run ${TARGET}

build:
	go build ${TARGET}

launch_race:
	go run -race ${TARGET}

download:
	go mod download

run_tests:
	go test -race ./... -cover -coverpkg ./...
