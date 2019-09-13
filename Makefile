greet_server:
	go run greet/greet_server/server.go

greet_client:
	go run greet/greet_client/client.go

calculator_server:
	go run calculator/calculator_server/server.go

calculator_client:
	go run calculator/calculator_client/client.go

calculator:
	protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.

greet:
	protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.
