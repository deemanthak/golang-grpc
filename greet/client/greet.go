package main

import (
	"context"
	"log"

	pb "github.com/deemanthak/go-grpc/greet/proto"
)

func doGreet(client pb.GreetServiceClient) {
	log.Print("do greet was invoked")

	res, err := client.Greet(context.Background(), &pb.GreetRequest{FirstName: "Dee"})

	if err != nil {
		log.Fatalf("error occurred %v", err)
	}
	log.Printf("Greeting: %v\n", res.Result)
}
