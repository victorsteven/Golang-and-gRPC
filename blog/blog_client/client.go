package main

import (
	"context"
	"fmt"
	"log"

	"github.com/victorsteven/go-grpc/blog/blogpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	fmt.Println("Blog client")

	tls := false
	opts := grpc.WithInsecure()
	if tls {
		certFile := "ssl/ca.crt" //Certificate Authority Trust Certificate
		creds, sslErr := credentials.NewClientTLSFromFile(certFile, "")
		if sslErr != nil {
			log.Fatal("Error while loading CA trust certificates: %v", sslErr)
			return
		}
		opts = grpc.WithTransportCredentials(creds)
	}

	cc, err := grpc.Dial("localhost:50051", opts)
	// cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %v", err)
	}

	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	blog := &blogpb.Blog{
		AuthorId: "steven",
		Title:    "My first blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})

	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has been created: %v\n", res.GetBlog())
	blogID := res.GetBlog().GetId()

	// Read Blog
	fmt.Println("Reading the blog")

	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "234wefdsgdssdf"})

	if err2 != nil {
		fmt.Printf("Error happened while reading: %v\n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	readRes, readErr := c.ReadBlog(context.Background(), readBlogReq)
	if readErr != nil {
		fmt.Printf("Error happened while reading: %v\n", readErr)
	}
	fmt.Printf("Blog was read: %v\n", readRes)
}
