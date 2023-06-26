package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/deemanthak/go-grpc/calculator/proto"
)

type Server struct {
	proto.CalculatorServiceServer
}

var addr string = "0.0.0.0:50052"

func main() {

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	log.Printf("listening %v", addr)
	s := grpc.NewServer()

	proto.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v\n", err)
	}
}
