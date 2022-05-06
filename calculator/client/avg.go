package main

import (
	"context"
	"log"

	pb "github.com/leandrorondon/grpc-go-course/calculator/proto"
)

func doAvg(c pb.CalculatorServiceClient, N []int32) {
	log.Println("doAvg was invoked")

	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("Error calling Avg: %v\n", err)
	}

	for _, n := range N {
		err := stream.Send(&pb.AvgRequest{
			N: n,
		})
		if err != nil {
			log.Fatalf("Error sending Avg request: %v\n", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error reading Avg response: %v\n", err)
	}

	log.Printf("Avg: %f\n", res.Result)
}
