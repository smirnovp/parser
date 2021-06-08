package main

import (
	"fmt"
	"log"
	"net"
	"parser/grpcgen"
	"parser/grpcservice"
	"parser/rusprofile"

	"google.golang.org/grpc"
)

var port = ":8085"

func main() {
	server := grpc.NewServer()
	sh := rusprofile.NewSiteParser()

	gs := grpcservice.New(sh)
	grpcgen.RegisterParserServiceServer(server, gs)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("net.listen: ", err)
	}

	fmt.Println("Serving requests...")
	err = server.Serve(listen)
	if err != nil {
		log.Fatal("server.Serve:", err)
	}
}
