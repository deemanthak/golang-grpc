package main

import (
	"context"
	"io"
	"log"

	"github.com/deemanthak/go-grpc/greet/proto"
)

func doGreetManyTimes(c proto.GreetServiceClient) error {
	log.Println("doGreetManyTimes was invoked")

	req := &proto.GreetRequest{FirstName: "tester"}

	stream, err := c.GreetManyTimes(context.Background(), req)

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
