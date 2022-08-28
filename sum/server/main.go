package main

import (
	"log"
	"net"

	pb "github.com/ben-c-bestow/grpc-course/sum/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.SumServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Listening failure on %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterSumServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Server failure on %v\n", err)
	}
}
