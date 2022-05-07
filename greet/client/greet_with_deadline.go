package main

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"

	pb "github.com/leandrorondon/grpc-go-course/greet/proto"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{
		FirstName: "Leandro",
	})
	e, ok := status.FromError(err)
	if ok && e != nil {
		// GRPC error
		if e.Code() != codes.DeadlineExceeded {
			log.Fatalf("Unexpected gRPC error: %v\n", e)
		}
		log.Println("Deadline exceeded")
		return
	}
	if !ok && err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
