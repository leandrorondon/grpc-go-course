package main

import (
	"context"
	"log"

	pb "github.com/leandrorondon/grpc-go-course/blog/proto"
)

func updateBlog(client pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Leandro Rondon",
		Title:    "A new title",
		Content:  "Content of the blog",
	}
	_, err := client.Update(context.Background(), newBlog)
	if err != nil {
		log.Printf("Error updating blog: %v\n", err)
		return
	}

	log.Printf("Blog was updated")
}
