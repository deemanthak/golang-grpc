package main

import (
	"context"
	"fmt"
	"log"

	"github.com/deemanthak/go-grpc/calculator/proto"
)

func doAvg(client proto.CalculatorServiceClient) {
	fmt.Print("doAvg was invoked")

	reqs := []*proto.AvgRequest{
		{Number: 2},
		{Number: 2},
		{Number: 2},
		{Number: 2},
	}

	stream, err := client.Average(context.Background())

	if err != nil {
		log.Fatalf("error calling avg %v\n", err)
	}

	for _, req := range reqs {
		fmt.Print("send reqs %v\n", req.Number)
		stream.Send(req)
		//time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error recieving response %v\n", err)
	}

	log.Printf("AVG: %v\n", res.Avg)

}
