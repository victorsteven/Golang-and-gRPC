package main 

import (
	"fmt"
	"log"
	"github.com/victorsteven/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
) 	

func main(){
	fmt.Println("Hello client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatal("could not connect: %v", err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	fmt.Printf("Created client: %f", c)
}