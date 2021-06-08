package grpcservice

import (
	"context"
	"fmt"
	"parser/grpcgen"
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
