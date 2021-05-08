package global

import (
	"context"
	"learn-go/grpc_learn/helloworld"
	"log"
	"time"

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

func initClientConn(address string) {
	// Set up a connection to the server.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	conn, err := grpc.DialContext(ctx, address, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ClientConn = conn
}

func InitGrpcClient() {
	initClientConn(ADDRESS)
	GreeterClient = helloworld.NewGreeterClient(ClientConn)
}
