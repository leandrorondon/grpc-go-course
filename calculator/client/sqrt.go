package main

import (
	"context"
	pb "github.com/leandrorondon/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{N: n})
	e, ok := status.FromError(err)
	if e != nil && ok {
		// GRPC error
		log.Printf("Error message from server: %s\n", e.Message())
		log.Printf("Error code from server: %v\n", e.Code())
		if e.Code() == codes.InvalidArgument {
			log.Println("Negative number was sent.")
		}
		return
	}
	if err != nil && !ok {
		log.Fatalf("Could not Sqrt: %v\n", err)
	}

	log.Printf("Sqrt: %v\n", res.Result)
}
