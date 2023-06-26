package main

import (
	"context"
	"io"
	"log"

	"github.com/deemanthak/go-grpc/greet/proto"
)

func doGreetEveryone(client proto.GreetServiceClient) {
	log.Print("doGreetEveryone was invoked")
	stream, err := client.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("error while creating stream %v", err)
	}

	reqs := []*proto.GreetRequest{
		{FirstName: "Dee"},
		{FirstName: "Dee1"},
		{FirstName: "Dee2"},
		{FirstName: "Dee3"},
	}

	wait := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send req %v\n", req)
			stream.Send(req)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("error receiving response %v", err)
				break
			}

			log.Printf("Response received : %v\n", res.Result)
		}
		close(wait)
	}()

	<-wait
}
