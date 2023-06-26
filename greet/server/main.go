package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/deemanthak/go-grpc/greet/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen %v \n", err)
	}

	log.Printf("listning to %v\n", addr)

	svr := grpc.NewServer()

	pb.RegisterGreetServiceServer(svr, &Server{})

	if err = svr.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v\n", err)
	}
}
