package main

import (
	"fmt"
	"log"
	"net"

	"github.com/IfuryI/WEB_BACK/internal/proto"
	fileServerGrpc "github.com/IfuryI/WEB_BACK/internal/services/fileserver/delivery/grpc"
	constants "github.com/IfuryI/WEB_BACK/pkg/const"
	"google.golang.org/grpc"
)

func main() {
	handler := fileServerGrpc.NewFileServerHandlerServer()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", constants.FileServerPort))

	if err != nil {
		log.Fatalln("Can't listen session microservice port", err)
	}
	defer lis.Close()

	server := grpc.NewServer()
	proto.RegisterFileServerHandlerServer(server, handler)

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
