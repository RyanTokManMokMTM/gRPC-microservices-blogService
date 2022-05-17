package main

import (
	"context"
	"github.com/ryantokmanmokmtm/gRPC-microservices-blogService/global"
	pb "github.com/ryantokmanmokmtm/gRPC-microservices-blogService/proto"
	db "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	ADDR       = "localhost:50001"
	TLS        = false
	MongoDBURI = "mongodb://root:admin@localhost:27017/"
)

func init() {
	log.Println("MongoDB initializing...")
	err := mongoDBInit()
	if err != nil {
		log.Fatalln(err)
	}
}

type Service struct {
	pb.UnimplementedBlogServiceServer
}

func main() {
	tcp, err := net.Listen("tcp", ADDR)
	if err != nil {
		log.Fatalln(err)
	}

	var opt []grpc.ServerOption
	if TLS {
		//load cred file
		//load sign file
		//append to option
	}
	grpcServer := grpc.NewServer(opt...)
	pb.RegisterBlogServiceServer(grpcServer, &Service{})
	log.Printf("RPC Server is listening on %v", ADDR)
	log.Fatalln(grpcServer.Serve(tcp))
}

func mongoDBInit() error {
	client, err := db.NewClient(options.Client().ApplyURI(MongoDBURI))
	if err != nil {
		return err
	}
	global.DBConnection = client

	//connection to db
	err = client.Connect(context.Background())
	if err != nil {
		return nil
	}

	global.BlogCollection = client.Database("blogdb").Collection("blog")

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		panic(err)
	}
	return nil
}
