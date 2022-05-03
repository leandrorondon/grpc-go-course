package main

import (
	"fmt"
	pb "github.com/leandrorondon/grpc-go-course/greet/proto"
	"io"
	"log"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")

	var res string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res += fmt.Sprintf("Hello %s\n", req.GetFirstName())
	}

	return nil
}
