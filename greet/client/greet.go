package main

import (
	"context"
	"log"

	pb "github.com/ben-c-bestow/grpc-course/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Dave",
	})

	if err != nil {
		log.Fatalf("Greet failure: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
