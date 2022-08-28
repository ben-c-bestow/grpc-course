package main

import (
	"context"
	"log"

	pb "github.com/ben-c-bestow/grpc-course/sum/proto"
)

func doSum(c pb.SumServiceClient) {
	log.Println("doSum invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		Addend_1: 3,
		Addend_2: 10,
	})

	if err != nil {
		log.Fatalf("Sum failure: %v\n", err)
	}

	log.Printf("Sum: %d\n", res.Sum)
}
