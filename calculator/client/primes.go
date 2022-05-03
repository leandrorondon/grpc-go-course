package main

import (
	"context"
	"io"
	"log"

	pb "github.com/leandrorondon/grpc-go-course/calculator/proto"
)

func doPrimes(c pb.CalculatorServiceClient, n int32) {
	log.Println("doPrimes was invoked")
	req := &pb.PrimeRequest{
		N: n,
	}

	stream, err := c.Primes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Prime: %v\n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the stream: %v\n", err)
		}

		log.Printf("Prime: %d\n", msg.GetFactor())
	}
}
