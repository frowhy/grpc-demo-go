package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-demo-go/protobuf/message"
	"grpc-demo-go/protobuf/service/testA"
	"grpc-demo-go/protobuf/service/testB"
	"log"
	"net"
)

const (
	port = ":50051"
)

type testAServer struct{}
type testBServer struct{}

func (s testAServer) Echo(ctx context.Context, req *testA.EchoRequest) (*testA.EchoResponse, error) {
	return &testA.EchoResponse{Meta: &message.Meta{Message: "Req " + req.Message}, Data: &message.Echo{Message: "Test Data Message"}}, nil
}

func (s testBServer) Echo(ctx context.Context, req *testB.EchoRequest) (*testB.EchoResponse, error) {
	return &testB.EchoResponse{Meta: &message.Meta{Message: "Req " + req.Message}, Data: &message.Echo{Message: "Demo Data Message"}}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	testA.RegisterGreeterServer(s, &testAServer{})
	testB.RegisterGreeterServer(s, &testBServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
