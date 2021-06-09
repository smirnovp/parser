package grpcgwservice

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"parser/grpcgen"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// Run ...
func Run(grpcPort, gwPort string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := grpcgen.RegisterParserServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts)
	if err != nil {
		log.Fatal(err)
	}

	err = mux.HandlePath("GET", "/docs/parser.swagger.json", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		http.FileServer(http.Dir("./")).ServeHTTP(w, r)
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Serving gRPC-gw requests on port %s...\n", gwPort)
	err = http.ListenAndServe(gwPort, cors(mux))
	if err != nil {
		log.Fatal(err)
	}
}

func cors(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		mux.ServeHTTP(w, r)
	})
}
