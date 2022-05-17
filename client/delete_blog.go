package main

import (
	"context"
	pb "github.com/ryantokmanmokmtm/gRPC-microservices-blogService/proto"
	"log"
	"time"
)

func DeleteBlog(client pb.BlogServiceClient, id string) {
	log.Println("Delete Blog By Id...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if _, err := client.DeleteBlog(ctx, &pb.BlogID{
		Id: id,
	}); err != nil {
		log.Fatalln(err)
	}

	log.Printf("Blog data with %v was deleted!\n", id)
}
