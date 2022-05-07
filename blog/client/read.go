package main

import (
	"context"
	"log"

	pb "github.com/leandrorondon/grpc-go-course/blog/proto"
)

func readBlog(client pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("readBlog was invoked")

	req := &pb.BlogId{Id: id}
	res, err := client.Read(context.Background(), req)
	if err != nil {
		log.Printf("Error reading blog: %v\n", err)
		return nil
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}
