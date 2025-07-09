package main

import (
	"context"
	"log"
	"time"

	pb "github.com/girish/gRPC/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, channel := context.WithTimeout(context.Background(), time.Second)
	defer channel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatal("could not greet:", err)
	}
	log.Printf("%s", res.Message)
}
