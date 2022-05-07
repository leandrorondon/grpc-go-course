package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"

	pb "github.com/leandrorondon/grpc-go-course/calculator/proto"
)

func (s *Server) Sqrt(ctx context.Context, req *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("ctx function was invoked with %v\n", req)

	n := req.GetN()

	if n < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", n),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(n)),
	}, nil
}
