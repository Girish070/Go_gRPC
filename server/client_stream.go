package main

import (
	"io"
	"log"

	pb "github.com/girish/gRPC/proto"
)

func (s *helloServer) SayHelloClientStream(stream pb.GreetService_SayHelloClientStreamServer) error {
	var messages []string
	for {
		req , err :=  stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MassageList{Massages: messages})
		}
		log.Printf("Got request with name: %v", req.Name)
		messages = append(messages, "Hello "+req.Name)
	}
}