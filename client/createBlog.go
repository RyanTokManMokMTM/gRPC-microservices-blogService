package main

import (
	"context"
	pb "github.com/ryantokmanmokmtm/gRPC-microservices-blogService/proto"
	"log"
	"time"
)

func CreateBlog(client pb.BlogServiceClient) string {
	log.Println("-- Creating a Blog --")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	data := pb.Blog{
		AuthorId: "Jackson",
		Title:    "Hello,world",
		Content:  "My First blog",
	}

	blogID, err := client.CreateBlog(ctx, &data)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Your blog ID is %v", blogID.Id)
	return blogID.Id
}
