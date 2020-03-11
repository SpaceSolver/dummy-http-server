.PHONY: default build exec

BIN := http-server-dump-request.exe
SRC := $(wildcard *.go)

default: build

build: $(BIN)

$(BIN): $(SRC)
	go build -o $@ $<

exec:
	./$(BIN)
