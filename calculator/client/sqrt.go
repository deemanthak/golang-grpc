package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/deemanthak/go-grpc/calculator/proto"
)

func doSqrt(client proto.CalculatorServiceClient, n int32) {
	log.Println("doSqrt was invoked")

	res, err := client.Sqrt(context.Background(), &proto.SqrtRequest{Number: n})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server %v\n", e.Message())
			log.Printf("error code from server %v\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("a negative number")
				return
			}
		} else {
			log.Fatalf("NON GPC ERROR %v", err)
		}
	}

	fmt.Printf("response %v", res.Result)
}
