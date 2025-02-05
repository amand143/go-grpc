package main

import (
	"log"

	pb "github.com/amand143/grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("client could not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	names := &pb.NamesList{
		Names: []string{"A", "B", "C"},
	}

	//callSayHellp(client)
	//callServerStreaming(client, names)
	//callClientStreaming(client, names)
	callClientServerStreaming(client, names)

}
