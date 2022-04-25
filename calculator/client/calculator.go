package main

import (
	"context"
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
