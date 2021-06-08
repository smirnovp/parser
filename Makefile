.PHONY: all
all: run

.PHONY: build
build:
	go build .

.PHONY: run
run:
	go run --race main.go

.PHONY: protos
protos:
	protoc -I protos --go_out=. --go-grpc_out=. protos/parser.proto

.PHONY: client
client:
	go run --race parserClient.go
