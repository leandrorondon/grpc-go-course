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

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Prime function was invoked with %v\n", in)

	var k int32 = 2
	n := in.GetN()

	for n > 1 {
		if n%k == 0 {
			stream.SendMsg(&pb.PrimeResponse{
				Factor: k,
			})
			n = n / k
			continue
		}

		k = k + 1
	}

	return nil
}
