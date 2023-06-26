package main

import (
	"context"
	"io"
	"log"

	"github.com/deemanthak/go-grpc/calculator/proto"
)

func doGetMax(client proto.CalculatorServiceClient) {
	log.Print("doGetMax was invoked")

	stream, err := client.Max(context.Background())

	if err != nil {
		log.Fatalf("error calculating max %v\n", err)
	}

	reqs := []*proto.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
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

			log.Printf("Response received : %v\n", res.Max)
		}
		close(wait)
	}()

	<-wait

}
