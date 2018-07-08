package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"github.com/yuuyamad/grpc-sample/service"
	pb "github.com/yuuyamad/grpc-sample/grpcsample"
)

func main(){
	port := ":8080"
	lintenPort, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterFileServer(server, &service.MyFileService{})


		log.Printf("start grpc server port: %s",port)
		server.Serve(lintenPort)

	server.Serve(lintenPort)
}