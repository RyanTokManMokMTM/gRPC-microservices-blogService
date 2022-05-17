package main

import (
	"context"
	"fmt"
	"github.com/ryantokmanmokmtm/gRPC-microservices-blogService/global"
	pb "github.com/ryantokmanmokmtm/gRPC-microservices-blogService/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Service) DeleteBlog(ctx context.Context, in *pb.BlogID) (*emptypb.Empty, error) {
	log.Println("--- Deleting Blog data ---")

	//casting to OID
	OID, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("Cannot cast thd id to OID"))
	}

	//Delete data from mongoDB by OID
	_, err = global.BlogCollection.DeleteOne(ctx, bson.M{"_id": OID})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Deleted Failed.Cannot delete from mongoDB : %v", err))
	}

	return &emptypb.Empty{}, nil
}
