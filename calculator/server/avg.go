package main

import (
	"io"
	"log"

	pb "github.com/leandrorondon/grpc-go-course/calculator/proto"
)

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function was invoked")

	sum := 0.0
	n := 0.0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: sum / n,
			})
		}

		if err != nil {
			log.Fatalf("Error reading client stream: %v\n", err)
		}

		n++
		sum += float64(req.GetN())
	}
}
