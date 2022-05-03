package main

import (
	"context"
	"log"

	pb "github.com/leandrorondon/grpc-go-course/calculator/proto"
)

func (s *Server) Sum(_ context.Context, req *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", req)
	return &pb.SumResponse{
		Result: req.GetA() + req.GetB(),
	}, nil
}
