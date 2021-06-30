package main

import (
	"fmt"
	sonic "github.com/ProjectAthenaa/sonic/protos"
	"google.golang.org/grpc"
	"log"
	"main/module"
	"net"
	"os"
)

var port = "5000"

func main() {
	if a := os.Getenv("PORT"); a != "" {
		port = a
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	sonic.RegisterModuleServer(server, &module.Server{})

	if err = server.Serve(listener); err != nil {
		panic(err)
	}

}
