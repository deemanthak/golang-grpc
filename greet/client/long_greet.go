package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deemanthak/go-grpc/greet/proto"
)

func doLongGreet(client proto.GreetServiceClient) {
	fmt.Print("goLongGreet was invoked")

	reqs := []*proto.GreetRequest{
		{FirstName: "tes1"},
		{FirstName: "tes2"},
		{FirstName: "tes3"},
		{FirstName: "tes4"},
		{FirstName: "tes4"},
		{FirstName: "tes4"},
		{FirstName: "tes4"},
		{FirstName: "tes4"},
		{FirstName: "tes4"},
		{FirstName: "tes4"},
	}

	stream, err := client.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("error callinf long greet %v\n", err)
	}

	for _, req := range reqs {

		stream.Send(req)
		//time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error recieving response %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}
