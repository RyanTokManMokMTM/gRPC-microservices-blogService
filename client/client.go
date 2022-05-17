package main

import (
	pb "github.com/ryantokmanmokmtm/gRPC-microservices-blogService/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

const (
	ADDR = "localhost:50001"
	TLS  = false
)

func main() {
	client, err := grpc.Dial(ADDR, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()
	stub := pb.NewBlogServiceClient(client)

	//Create a Blog Article
	id := CreateBlog(stub)
	ReadBlog(stub, id) //valid
	//ReadBlog(stub, "TestingID") //invalid
	UpdateBlog(stub, id, "Tom-Sum", "Update Hello", "Hello world")

	ListBlog(stub)

	DeleteBlog(stub, id)
}
