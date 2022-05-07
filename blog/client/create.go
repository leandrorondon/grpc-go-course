package main

import (
	"context"
	pb "github.com/leandrorondon/grpc-go-course/blog/proto"
	"log"
)

func createBlog(client pb.BlogServiceClient) string {
	log.Println("createBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "Leandro",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	res, err := client.Create(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.GetId())

	return res.GetId()
}
