package main

import (
	"context"
	"fmt"
	"io"
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

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("ComputeAverage function was invoked with a streaming request %v\n", stream)

	sum := float64(0)
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// We have finished reading the client stream
			return stream.SendAndClose(&calculatorpb.NumberResponse{
				Average: sum / float64(count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		number := req.GetNumber()

		sum += float64(number)

		count++
	}
}

func (*server) FindMaximum(stream calculatorpb.CalculatorService_FindMaximumServer) error {
	fmt.Printf("FindMaximum function was invoked with a streaming request %v\n", stream)

	maximum := int32(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while receiving client stream to find max: %v", err)
			return err
		}
		number := req.GetNumber()
		if number > maximum {
			maximum = number
			err = stream.Send(&calculatorpb.FindMaximumResponse{
				Maximum: maximum,
			})
			if err != nil {
				log.Fatalf("Error while sending data to  client: %v", err)
				return err
			}
		}
	}
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
