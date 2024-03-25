package main

import (
	"context"
	"io"
	"log"

	pb "github.com/amand143/grpc-demo/proto"
)

func callServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("streaming started")
	stream, err := client.ServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("error while stream: %v", err)
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("erro while serverstraem: %v", err)
		}
		log.Println(message)
	}
	log.Printf("stream ended")
}
