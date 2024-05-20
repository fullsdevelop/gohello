.PHONY: clean tidy vendor test build
.ONESHELL:
SHELL:=/bin/bash
TARGET:=build/gohello

clean:
	rm -rf build
	go clean -r

tidy:
	go mod tidy

vendor:
	go mod vendor

test:
	go test -v -cover -timeout 10s ./...

build:
	go build -o $(TARGET) cmd/hello/main.go
	strip $(TARGET)
	upx $(TARGET)