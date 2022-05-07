package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/leandrorondon/grpc-go-course/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient, N []int32) {
	log.Println("doMax was invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error calling Avg: %v\n", err)
	}

	waitCh := make(chan struct{})

	go func() {
		for _, n := range N {
			log.Printf("Sending %v\n", n)
			err := stream.Send(&pb.MaxRequest{
				N: n,
			})
			if err != nil {
				log.Fatalf("Error sending Avg request: %v\n", err)
			}
			time.Sleep(time.Second)
		}
		err := stream.CloseSend()
		if err != nil {
			log.Fatalf("Error closing stream: %v\n", err)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error receiving: %v\n", err)
			}
			log.Printf("Received max: %d\n", res.GetMax())
		}

		close(waitCh)
	}()

	<-waitCh
}
