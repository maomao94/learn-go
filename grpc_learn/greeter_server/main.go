/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"learn-go/grpc_learn/helloworld"
	"learn-go/grpc_learn/helloworld/errcodepb"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	//return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
	return nil, Error(errcodepb.ErrCode_LoginWechatCreateUser, errors.New("must supply Type"))
}

func main() {
	ctx := context.Background()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloworld.RegisterGreeterServer(s, &server{})
	c := make(chan os.Signal, 1)
	go func() {
		for range c {
			log.Printf("shutting down gRPC server...")
			s.GracefulStop()
			<-ctx.Done()
		}
	}()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// 记录grpc error信息
func Error(code errcodepb.ErrCode, err error) error {
	log.Println("grpc error", zap.Int32("code", int32(code)), zap.Error(err))
	var cause string
	if err != nil {
		cause = err.Error()
	}
	return status.Error(codes.Code(code), fmt.Sprintf("error name: %s, cause: %s", errcodepb.ErrCode_name[int32(code)], cause))
}
