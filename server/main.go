package main

import (
	pb "github.com/girish/gRPC/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type helloServer struct {
	pb.GreetServiceServer
}

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to start server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server Started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("failed to start:", err)
	}
}
