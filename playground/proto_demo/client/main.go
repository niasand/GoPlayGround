package main

import (
	"log"
	"os"

	pb "proto_demo/demo"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect:%v", err)
	}

	defer conn.Close()

	c := pb.NewHelloWorldClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.SendPost(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatal("Could not say hello to :%v", err)
	}
	log.Printf("CLient POST:%s", r.Message)

	r, err = c.SendGet(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatal("Could not say hello to :%v", err)
	}
	log.Printf("CLient GET:%s", r.Message)
}
