package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"parser/grpcgen"
	"parser/grpcservice"
	"parser/rusprofile"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

var port = ":8085"
var gwport = ":8084"

func runGrpcGw() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := grpcgen.RegisterParserServiceHandlerFromEndpoint(ctx, mux, port, opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Serving gRPC-gw requests on port %s...\n", gwport)
	err = http.ListenAndServe(gwport, mux)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	go runGrpcGw()

	server := grpc.NewServer()
	sh := rusprofile.NewSiteParser()

	gs := grpcservice.New(sh)
	grpcgen.RegisterParserServiceServer(server, gs)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("net.Listen: ", err)
	}

	fmt.Printf("Serving gRPC requests on port %s...\n", port)
	err = server.Serve(listen)
	if err != nil {
		log.Fatal("server.Serve:", err)
	}
}
