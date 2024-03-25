package main

import (
	"io"
	"log"

	pb "github.com/amand143/grpc-demo/proto"
)

func (s *helloServer) BidirectionStreaming(stream pb.GreetService_BidirectionStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("error bidirec stream receiving %v", err)
			return err
		}
		log.Printf("got req with name %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
