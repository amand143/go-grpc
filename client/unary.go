package main

import (
	"context"
	"log"
	"time"

	pb "github.com/amand143/grpc-demo/proto"
)

func callSayHellp(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.SayHellp(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("error while rpc sayHellp %v", err)
	}
	log.Printf("%s", res.Message)
}
