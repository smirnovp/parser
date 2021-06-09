package main

import (
	"fmt"
	"os"
	"os/signal"
	"parser/grpcgwservice"
	"parser/grpcservice"
	"parser/rusprofile"
)

var grpcPort = ":8085"
var gwPort = ":8084"

func main() {

	go grpcgwservice.Run(grpcPort, gwPort)

	parser := rusprofile.NewSiteParser()
	gs := grpcservice.New(parser)
	go gs.Run(grpcPort)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill)
	s := <-sigs
	fmt.Println("\nСигнал на выход: ", s)
}
