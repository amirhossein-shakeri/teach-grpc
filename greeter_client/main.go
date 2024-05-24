package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/amirhossein-shakeri/teach-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "World"
)

var (
	addr = flag.String("addr", ":50051", "the greeter server address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("greeter client failed to connect to %v: %v", *addr, err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	response, err := client.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("greeter client could not greet: %v", err)
	}
	log.Printf("Greeting: %s", response.GetMessage())
}
