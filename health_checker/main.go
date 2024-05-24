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

var (
	addr     = flag.String("addr", ":50051", "the greeter server address to connect to")
	interval = flag.Int("interval", 3, "health check interval")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("health checker failed to connect to %v: %v", *addr, err)
	}
	defer conn.Close()
	client := pb.NewGreeterClient(conn)
	var checkInterval time.Duration = time.Second * time.Duration(*interval)
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()
			response, err := client.Health(ctx, &pb.HealthRequest{})
			if err != nil {
				log.Fatalf("health checker could not greet: %v", err)
			}
			log.Printf("Health Check: %v %v", response.GetOk(), response.GetMessage())
		}
	}
}
