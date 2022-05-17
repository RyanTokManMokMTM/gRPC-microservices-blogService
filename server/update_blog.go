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

func (s *Service) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Println("--- Updating blog info ---")

	OID, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("Cannot cast id to OID"))
	}

	updateData := Blog{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	//update mongo db
	//filter id and
	res, err := global.BlogCollection.UpdateOne(ctx, bson.M{"_id": OID}, bson.M{"$set": updateData})
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintln("Cannot update blog info"))
	}

	//check our result
	if res.MatchedCount == 0 {
		//nothing has been updated
		return nil, status.Errorf(codes.NotFound, fmt.Sprintln("Cannot find blog by OID"))
	}

	return &emptypb.Empty{}, nil
}
