.PHONY: all run

all: run

TARGET = ./Workspace/src/main/main.go

run:
	go run ${TARGET}
