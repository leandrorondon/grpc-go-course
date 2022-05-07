package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	pb "github.com/leandrorondon/grpc-go-course/blog/proto"
)

func (s *Server) Create(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("Create was invoked with %v\n", in)

	newItem := BlogItem{
		AuthorID: in.GetAuthorId(),
		Title:    in.GetTitle(),
		Content:  in.GetContent(),
	}

	res, err := collection.InsertOne(ctx, newItem)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal Error: %v", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to oid: %v", err),
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}
