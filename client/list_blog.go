package main

import (
	"context"
	pb "github.com/ryantokmanmokmtm/gRPC-microservices-blogService/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func ListBlog(client pb.BlogServiceClient) {
	log.Println("List all Blog data...")

	stream, err := client.ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalln(err)
	}

	for {
		blogData, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Revevied Blog Info : %v\n", blogData)
	}

	//close the sending streaming channel
	if err := stream.CloseSend(); err != nil {
		log.Fatalln(err)
	}
}
