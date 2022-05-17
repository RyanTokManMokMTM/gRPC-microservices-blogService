package main

import (
	"context"
	"fmt"
	"github.com/ryantokmanmokmtm/gRPC-microservices-blogService/global"
	pb "github.com/ryantokmanmokmtm/gRPC-microservices-blogService/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Service) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogID, error) {
	log.Printf("Creating Blog information...")

	//create a Blog instance that information from a incoming request
	data := Blog{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	//insert to our collection
	res, err := global.BlogCollection.InsertOne(ctx, data)

	//return RPC error
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("INternal Error %v", err))
	}

	//casting db InsertedID to Objection ID
	OID, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil,
			status.Errorf(codes.Internal,
				fmt.Sprintf("Cannot convert to OID,%v", err))
	}

	return &pb.BlogID{
		Id: OID.Hex(), //encode to string
	}, nil
}
