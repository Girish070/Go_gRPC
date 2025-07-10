package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/girish/gRPC/proto"
)

func callHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Bidirectional streaming Started")
	stream, err := client.SayHelloBidirectionalStream(context.Background())
	if err != nil {
		log.Fatalf("Could not send Names : %v", err)
	}

	wait := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal("Error While Striming:", err)
			}
			log.Println(message)
		}
		close(wait)
	}()

	for _, name := range names.Name {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while sending %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-wait
	log.Printf("Bidirectional Striming Finished")
}
