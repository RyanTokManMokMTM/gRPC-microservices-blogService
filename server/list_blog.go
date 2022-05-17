package main

import (
	"context"
	"fmt"
	"github.com/ryantokmanmokmtm/gRPC-microservices-blogService/global"
	pb "github.com/ryantokmanmokmtm/gRPC-microservices-blogService/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Service) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("--- List All Blogs ---")

	cur, err := global.BlogCollection.Find(context.Background(), primitive.D{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintln("Unknown Internal Error"))
	}
	defer cur.Close(context.Background()) // close database connection

	for cur.Next(context.Background()) {
		//decoding
		res := &Blog{}
		if err := cur.Decode(res); err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Cannot decode MongoDB's data %v\n", err))
		}

		//send to client
		stream.Send(DocumentToBlog(res))
		//time.Sleep(time.Second * 1)
	}

	//check cur error
	if cur.Err() != nil {
		return status.Errorf(codes.Internal, fmt.Sprintln("Unknown Internal Error"))
	}

	return nil
}
