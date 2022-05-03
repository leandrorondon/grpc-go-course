package main

import (
	"context"
	"io"
	"log"

	pb "github.com/leandrorondon/grpc-go-course/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		A: 3, B: 10,
	})
	if err != nil {
		log.Fatalf("Could not Sum: %v\n", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}

func doPrime(c pb.CalculatorServiceClient, n int32) {
	log.Println("doPrime was invoked")
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
