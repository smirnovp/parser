package grpcservice

import (
	"context"
	"fmt"
	"log"
	"net"
	"parser/grpcgen"

	"google.golang.org/grpc"
)

// ISiteParser ...
type ISiteParser interface {
	GetDataFromSite(*grpcgen.ParserRequest) (*grpcgen.ParserResponse, error)
}

// GrpcService ...
type GrpcService struct {
	sp ISiteParser
	grpcgen.UnimplementedParserServiceServer
}

// New ...
func New(sp ISiteParser) *GrpcService {
	return &GrpcService{sp: sp}
}

// GetData ...
func (g *GrpcService) GetData(ctx context.Context, req *grpcgen.ParserRequest) (*grpcgen.ParserResponse, error) {
	fmt.Println("Запрос:", req.INN)
	resp, err := g.sp.GetDataFromSite(req)
	return resp, err
}

// Run ...
func (g *GrpcService) Run(grpcPort string) {
	server := grpc.NewServer()
	grpcgen.RegisterParserServiceServer(server, g)

	listen, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatal("net.Listen: ", err)
	}

	fmt.Printf("Serving gRPC requests on port %s...\n", grpcPort)
	err = server.Serve(listen)
	if err != nil {
		log.Fatal("server.Serve:", err)
	}
}
