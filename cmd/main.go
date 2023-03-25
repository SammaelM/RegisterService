package main

import (
	"log"
	"mainmod/internal/response"
	"mainmod/pkg/server"
	"net"
	"net/http"
)

func main() {
	//////
	//grpc

	pon := http.HandlerFunc(response.Handler)

	//run http website
	go func() {
		serv := new(server.Server)
		if err := serv.Run(":3002", pon); err != nil {
			log.Fatalf("error occured while running http server: %s", err)
		}
	}()

	//run connect grpc
	l, err := net.Listen("tcp", ":3001")
	if err != nil {
		log.Fatal(err)
	}

	//grpc
	if err := server.RegServ().Serve(l); err != nil {
		log.Fatal(err)
	}

}
