package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/amand143/grpc-demo/proto"
)

func callClientServerStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("bidirec client stream started")
	stream, err := client.BidirectionStreaming(context.Background())
	if err != nil {
		log.Fatalf("bidirec client error %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitc
	log.Printf("bidirec stream finished")
}
