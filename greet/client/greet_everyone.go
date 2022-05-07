package main

import (
	"context"
	pb "github.com/leandrorondon/grpc-go-course/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while calling GreetEveryone: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Leandro"},
		{FirstName: "Nathaly"},
		{FirstName: "Sofia"},
		{FirstName: "Joaquim"},
	}

	waitCh := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request %v\n", req)
			err := stream.SendMsg(req)
			if err != nil {
				log.Fatalf("Error sending request: %v\n", err)
			}
			time.Sleep(time.Second)
		}
		stream.CloseSend()
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
			log.Printf("Received: %s\n", res.GetResult())
		}
		close(waitCh)
	}()

	<-waitCh
}
