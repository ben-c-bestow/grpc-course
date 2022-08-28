package main

import (
	"context"
	"log"

	pb "github.com/ben-c-bestow/grpc-course/sum/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function invoked w/ %v\n", in)
	return &pb.SumResponse{
		Sum: in.Addend_1 + in.Addend_2,
	}, nil
}
