package main

import (
	"fmt"
	"io"
	"log"

	"github.com/deemanthak/go-grpc/greet/proto"
)

func (s *Server) GreetEveryone(stream proto.GreetService_GreetEveryoneServer) error {
	fmt.Print("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while reading client stream %v", err)
		}

		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&proto.GreetResponse{Result: res})

		if err != nil {
			log.Fatalf("error sending response %v", err)
		}
	}
}
