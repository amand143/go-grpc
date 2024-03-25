package main

import (
	"context"
	"log"
	"time"

	pb "github.com/amand143/grpc-demo/proto"
)

func callClientStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("client stream starts")
	stream, err := client.ClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("error while client stream")
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("error while streaming client %v", err)
		}
		log.Printf("sent req with name %v", name)
		time.Sleep(2 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("client stream ended")
	if err != nil {
		log.Fatalf("error while sendin client %v", err)
	}
	log.Printf("%v", res.Messages)
}
