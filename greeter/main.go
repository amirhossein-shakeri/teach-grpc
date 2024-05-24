package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/amirhossein-shakeri/teach-grpc/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The greeter server port")
)

type greeterServer struct {
	pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello, " + in.GetName() + "!"}, nil
}

func (s *greeterServer) Health(ctx context.Context, in *pb.HealthRequest) (*pb.HealthReply, error) {
	log.Printf("Health Check")
	return &pb.HealthReply{Ok: true, Message: "I'm Alive!"}, nil
}

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("greeter server failed to listen to %d: %v", *port, err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &greeterServer{})
	log.Printf("greeter server listening at %v", listener.Addr())
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("greeter server failed to serve: %v", err)
	}
	log.Print("DONE!")
}
