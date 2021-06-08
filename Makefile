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

.PHONY: docker
docker:
	docker build -t grpc-site-parser .

.PHONY: docker-run
docker-run:
	docker run --rm --name grpcSiteParser -p 8085:8085 grpc-site-parser
