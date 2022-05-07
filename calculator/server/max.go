package main

import (
	"io"
	"log"
	"math"

	pb "github.com/leandrorondon/grpc-go-course/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("Max function was invoked")

	max := int32(math.MinInt32)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error reading client stream: %v\n", err)
		}

		log.Printf("Received %d\n (max = %d)", req.GetN(), max)
		if req.GetN() > max {
			max = req.GetN()

			err := stream.Send(&pb.MaxResponse{
				Max: max,
			})
			if err != nil {
				log.Fatalf("Error while sending response: %v\n", err)
			}
		}
	}

	return nil
}
