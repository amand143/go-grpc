package main

import (
	"log"
	"time"

	pb "github.com/amand143/grpc-demo/proto"
)

func (s *helloServer) ServerStreaming(req *pb.NamesList, stream pb.GreetService_ServerStreamingServer) error {
	log.Printf("got request with name %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(2 * time.Second)
	}
	return nil
}
