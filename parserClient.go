package main

import (
	"context"
	"fmt"
	"log"
	"parser/grpcgen"

	"google.golang.org/grpc"
)

var port = ":8085"

func main() {
	fmt.Println("grpc Client started ...")
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Dial error:", err)
	}
	client := grpcgen.NewParserServiceClient(conn)

	req := &grpcgen.ParserRequest{
		// INN: "7721679536",
		INN: "7703735562",
	}
	ctx := context.Background()
	res, err := client.GetData(ctx, req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("ответ от сервера:\nИНН:%s\nКПП:%s\nОрганизация:%s\nРуководитель:%s\n", res.INN, res.KPP, res.Company, res.Manager)
}
