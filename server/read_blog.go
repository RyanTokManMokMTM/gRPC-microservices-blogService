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
	"log"
)

func (s *Service) ReadBlog(ctx context.Context, in *pb.BlogID) (*pb.Blog, error) {
	//get blog info by id

	//casting to Mongodb object id

	OID, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("InvaildArgument err: %v", err),
		)
	}

	//get data by id
	res := &Blog{}
	log.Println(OID)
	filter := bson.M{"_id": OID} // searching condition
	data := global.BlogCollection.FindOne(ctx, filter)
	//convert data to our blog type
	if err := data.Decode(res); err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find blog by ID : err %v", err))
	}

	return DocumentToBlog(res), nil
}
