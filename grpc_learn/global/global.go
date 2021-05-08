package global

import (
	"learn-go/grpc_learn/helloworld"
	"log"

	"google.golang.org/grpc"
)

const (
	ADDRESS      = "localhost:50051"
	DEFAULT_NAME = "world"
)

var (
	ClientConn    *grpc.ClientConn
	GreeterClient helloworld.GreeterClient
)

func getClientConn(address string) *grpc.ClientConn {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func InitGrpcClient() {
	ClientConn = getClientConn(ADDRESS)
	GreeterClient = helloworld.NewGreeterClient(ClientConn)
}
