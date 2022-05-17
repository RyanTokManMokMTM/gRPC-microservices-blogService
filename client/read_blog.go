package main

import (
	"context"
	pb "github.com/ryantokmanmokmtm/gRPC-microservices-blogService/proto"
	"log"
	"time"
)

func ReadBlog(client pb.BlogServiceClient, blogID string) {
	log.Println("--Retrieving Blog Information ---")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	data, err := client.ReadBlog(ctx, &pb.BlogID{Id: blogID})
	if err != nil {
		//e, ok := status.FromError(err)
		//if ok {
		//	if e.Code() == codes.InvalidArgument {
		//		log.Fatalf("InvalidArgument err %v", e)
		//	} else if e.Code() == codes.Internal {
		//		log.Fatalf("Internal err %v", e)
		//	}
		//} else {
		//	log.Fatalln(err)
		//}
		log.Fatalln(err)
	}

	log.Printf("Your blog info : %v", data)

}
