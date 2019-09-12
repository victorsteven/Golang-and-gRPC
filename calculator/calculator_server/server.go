package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/victorsteven/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Received Sum RPC: %v", req)
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumber
	sum := firstNumber + secondNumber
	res := &calculatorpb.SumResponse{
		SumResult: sum,
	}
	return res, nil
}

func (*server) PrimeNumber(req *calculatorpb.PrimeRequest, stream calculatorpb.CalculatorService_PrimeNumberServer) error {
	fmt.Printf("Received Prime RPC: %v", req)
	theNumber := req.GetTheNumber()
	var n int32 = 2
	for theNumber > 1 {
		if theNumber%n == 0 {
			stream.Send(&calculatorpb.PrimeResponse{
				PrimeFactor: n,
			})
			theNumber = theNumber / n
		} else {
			n++
		}
	}
	return nil
}

func main() {
	fmt.Println("Calculator server")

	// lis, err := net.Listen(network string, address string)
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	// greetpb.RegisterGreetServiceServer(s, &server{})
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
