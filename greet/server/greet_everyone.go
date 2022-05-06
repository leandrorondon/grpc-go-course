package main

import (
	"fmt"
	pb "github.com/leandrorondon/grpc-go-course/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("GreetEveryone function was invoked")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		err = stream.Send(&pb.GreetResponse{
			Result: fmt.Sprintf("Hello %s\n", req.GetFirstName()),
		})
		if err != nil {
			log.Fatalf("Error while sending response: %v\n", err)
		}
	}

	return nil
}
