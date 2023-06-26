package main

import (
	"context"
	"log"

	"github.com/deemanthak/go-grpc/calculator/proto"
)

func doCalculate(c proto.CalculatorServiceClient) {

	log.Print("doCalculate was invoked")

	res, err := c.Calculate(context.Background(), &proto.CalculatorRequest{
		Num1: 10,
		Num2: 30,
	})

	if err != nil {
		log.Fatalf("error occurred %v", err)
	}
	log.Printf("Greeting: %v\n", res.Total)
}
