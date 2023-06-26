package main

import (
	"fmt"
	"io"
	"log"

	"github.com/deemanthak/go-grpc/calculator/proto"
)

func (s *Server) Max(stream proto.CalculatorService_MaxServer) error {
	log.Print("Max was invoked")

	var numberList []int32

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while reading client stream %v", err)
		}

		numberList = append(numberList, req.Number)
		max := int32(0)
		fmt.Println(numberList, max)
		// Iterate through the slice and update the max if a higher number is found
		for _, num := range numberList {
			if num > max {
				max = num
			}
		}

		err = stream.Send(&proto.MaxResponse{Max: max})

		if err != nil {
			log.Fatalf("error while sending response %v", err)
		}
	}
}
