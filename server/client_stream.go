package main

import (
	"io"
	"log"

	pb "github.com/amand143/grpc-demo/proto"
)

func (s *helloServer) ClientStreaming(stream pb.GreetService_ClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			log.Fatalf("erro while receiving stream %v", err)
		}
		messages = append(messages, "Hello "+req.Name)
	}
}
