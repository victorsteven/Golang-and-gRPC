package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/victorsteven/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	// doUnary(c)
	// doPrimeNumbers(c)
	doAverage(c)

	// fmt.Printf("Created client: %f", c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 30,
	}
	// c.Greet(context.Background(), in *greetpb.GreetRequest, opts ...grpc.CallOption)
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.SumResult)
}

func doPrimeNumbers(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Prime Numbers RPC...")
	req := &calculatorpb.PrimeRequest{
		TheNumber: 25,
	}
	stream, err := c.PrimeNumber(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeNumber RPC: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Something happened: %v", err)
		}
		fmt.Println(res.GetPrimeFactor())
	}
}

func doAverage(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Aveage Numbers RPC...")
	// numbers := []*calculatorpb.NumberRequest{
	// &calculatorpb.NumberRequest{
	// 	Number: 5,
	// },
	// 	&calculatorpb.NumberRequest{
	// 		Number: 10,
	// 	},
	// 	&calculatorpb.NumberRequest{
	// 		Number: 15,
	// 	},
	// 	&calculatorpb.NumberRequest{
	// 		Number: 14,
	// 	},
	// }

	numbers := []int64{1, 23, 4, 5, 6}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Average RPC: %v", err)
	}

	for _, number := range numbers {
		fmt.Printf("Sending req: %v\n", number)
		stream.Send(&calculatorpb.NumberRequest{
			Number: number,
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response %v", err)
	}

	fmt.Printf("The Average is: %v\n", res.GetAverage())
}
