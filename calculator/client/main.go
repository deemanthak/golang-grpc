package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/deemanthak/go-grpc/calculator/proto"
)

var addr string = "localhost:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("faild to dial %v\n", err)
	}

	c := proto.NewCalculatorServiceClient(conn)

	//doCalculate(c)
	//getPrimes(c)
	//doAvg(c)
	//doGetMax(c)
	doSqrt(c, -10)
}
