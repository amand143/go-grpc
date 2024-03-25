package main

import (
	"context"

	pb "github.com/amand143/grpc-demo/proto"
)

func (s *helloServer) SayHellp(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
