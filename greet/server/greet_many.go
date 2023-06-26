package main

import (
	"fmt"
	"log"

	"github.com/deemanthak/go-grpc/greet/proto"
)

func (s *Server) GreetManyTimes(gr *proto.GreetRequest, stream proto.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes was invoked %v\n", gr)

	for i := 0; i < 10000; i++ {
		res := fmt.Sprintf("Hello %s,Number %d", gr.FirstName, i)
		stream.Send(&proto.GreetResponse{Result: res})
	}

	return nil
}
