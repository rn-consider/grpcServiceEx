package helloworld_server

import (
	"context"
	"github.com/rn-consider/grpcservice/protos/helloworld"
	"log"
)

// Server 用来实现pb中的helloworld.GreeterServer.
type Server struct {
	helloworld.UnimplementedGreeterServer
}

// SayHello 实现hello-world.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloworld.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *Server) SayHelloAgain(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	return &helloworld.HelloReply{Message: "Hello again " + in.GetName()}, nil
}
