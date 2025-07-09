package main

import (
	"context"
	"io"
	"log"

	pb "github.com/girish/gRPC/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Striming Statrted")
	stream, err := client.SayHelloServerStream(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}

	for{
		message, err := stream.Recv()
		if err == io.EOF {
			break // End of the strean
		} 

		if err != nil {
			log.Fatalf("Error while receving %v", err)
		}
		log.Println(message)

	}
	log.Printf("Streaming Finishe")
}