package main

import (
	"fmt"
	"io"
	"log"

	"github.com/deemanthak/go-grpc/greet/proto"
)

func (s *Server) LongGreet(stream proto.GreetService_LongGreetServer) error {
	fmt.Print("LongGreet invoked")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&proto.GreetResponse{Result: res})
		}

		if err != nil {
			log.Fatalf("error reading client stream %v\n", err)
		}

		res += fmt.Sprintf("Hello %s!! \n", req.FirstName)
	}
}
