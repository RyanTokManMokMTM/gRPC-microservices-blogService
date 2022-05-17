package main

import (
	pb "github.com/ryantokmanmokmtm/gRPC-microservices-blogService/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Blog - Mongo DB Field
type Blog struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"` //ignore empty
	AuthorId string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func DocumentToBlog(blogInfo *Blog) *pb.Blog {
	return &pb.Blog{
		Id:       blogInfo.ID.Hex(), //encode to string
		AuthorId: blogInfo.AuthorId,
		Title:    blogInfo.Title,
		Content:  blogInfo.Content,
	}
}
