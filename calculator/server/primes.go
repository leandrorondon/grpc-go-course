package main

import (
	"log"

	pb "github.com/leandrorondon/grpc-go-course/calculator/proto"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("Primes function was invoked with %v\n", in)

	var k int32 = 2
	n := in.GetN()

	for n > 1 {
		if n%k == 0 {
			err := stream.SendMsg(&pb.PrimeResponse{
				Factor: k,
			})
			if err != nil {
				log.Fatalf("Error sending Prime response: %v\n", err)
			}
			n = n / k
			continue
		}

		k = k + 1
	}

	return nil
}
