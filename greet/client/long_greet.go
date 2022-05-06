package main

import (
	"context"
	pb "github.com/leandrorondon/grpc-go-course/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doLongGreet was invoked")
	reqs := []*pb.GreetRequest{
		{FirstName: "Leandro"},
		{FirstName: "Nathaly"},
		{FirstName: "Sofia"},
		{FirstName: "Joaquim"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		err := stream.Send(req)
		if err != nil {
			log.Fatalf("Error while sending: %v\n", err)
		}
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receive response from LongGreet: %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.GetResult())
}
