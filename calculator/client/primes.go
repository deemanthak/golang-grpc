package main

import (
	"context"
	"io"
	"log"

	"github.com/deemanthak/go-grpc/calculator/proto"
)

func getPrimes(client proto.CalculatorServiceClient) error {
	log.Println("getPrimes was invoked")

	req := &proto.PrimeRequest{Number: 120}

	stream, err := client.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error calling greetmany %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error reading the stream %v\n", err)
		}

		log.Printf("GreetManyTimes: %v\n", msg)
	}

	return nil
}
