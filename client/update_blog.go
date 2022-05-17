package main

import (
	"context"
	pb "github.com/ryantokmanmokmtm/gRPC-microservices-blogService/proto"
	"log"
	"time"
)

func UpdateBlog(client pb.BlogServiceClient, id, authorID, title, content string) {
	log.Println("Updating Blog info...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	_, err := client.UpdateBlog(ctx, &pb.Blog{
		Id:       id,
		AuthorId: authorID,
		Title:    title,
		Content:  content,
	})

	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Blog info updated successfully!")
}
