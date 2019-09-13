package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/victorsteven/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// doAverage(c)
	// doMaximum(c)
	doSquareRoot(c)
	// fmt.Printf("Created client: %f", c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Sum Unary RPC...")
	req := &calculatorpb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 30,
	}

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

func doMaximum(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do a Maximum Numbers RPC...")

	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while calling FindMaximum RPC: %v", err)
	}

	waitc := make(chan struct{})

	// send go routine
	go func() {
		// numbers := []int32{1, 5, 3, 9, 8, 2}
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}

		for _, number := range numbers {
			fmt.Printf("Sending req: %v\n", number)
			stream.Send(&calculatorpb.FindMaximumRequest{
				Number: number,
			})
			time.Sleep(1000 * time.Millisecond)
		}
		err = stream.CloseSend()
		if err != nil {
			log.Fatalf("Error while sending response %v", err)
		}
	}()

	// Receive go routine
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("problem while reading server stream: %v", err)
				break
			}
			maximum := res.GetMaximum()
			fmt.Printf("Received a maximum of: %v\n", maximum)
		}
		close(waitc)
	}()
	<-waitc
}

func doSquareRoot(c calculatorpb.CalculatorServiceClient) {

	doCalculation(c, 20)

	doCalculation(c, -20)

}

func doCalculation(c calculatorpb.CalculatorServiceClient, n int32) {

	req := &calculatorpb.SquareRootRequest{
		Number: n,
	}
	res, err := c.SquareRoot(context.Background(), req)
	if err != nil {
		respErr, ok := status.FromError(err)
		if ok {
			// Actual error from gRPC (user error)
			fmt.Printf("Error message from server: %v\n", respErr.Message())
			fmt.Printf("Error Code: %v\n", respErr.Code())
			if respErr.Code() == codes.InvalidArgument {
				fmt.Println("We probably sent a negative number\n")
				return
			}
		} else {
			log.Fatalf("error while calling SquareRoot RPC: %v\n", err)
			return
		}
	}
	fmt.Printf("This is the root: %v\n", res.GetNumberRoot())
}
