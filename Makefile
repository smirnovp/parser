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
	protoc -I protos --grpc-gateway_out ./grpcgen --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative protos/parser.proto
	protoc -I protos --swagger_out=./docs protos/parser.proto

.PHONY: client
client:
	go run --race parserClient.go

.PHONY: docker
docker:
	docker build -t grpc-site-parser .

.PHONY: docker-run
docker-run:
	docker run --rm --name grpcSiteParser -p 8085:8085 -p 8084:8084 grpc-site-parser
