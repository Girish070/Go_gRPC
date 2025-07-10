package main

import (
	"context"
	"log"
	"time"

	pb "github.com/girish/gRPC/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Print("Client streaming Started")
	stream, err := client.SayHelloClientStream(context.Background())
	if err != nil {
		log.Fatal("Could not greet:", err)
	}
	for _, name := range names.Name {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending request: %v", err)
		}
		log.Printf("Sent request with name: %v", name)
		time.Sleep(2 * time.Second)
	}
	resp , err := stream.CloseAndRecv()
	log.Printf("Client Streaming Finished")
	if err != nil {
		log.Fatalf("Error while receiving %v", err)
	}
	log.Printf("%v", resp.Massages)
}
