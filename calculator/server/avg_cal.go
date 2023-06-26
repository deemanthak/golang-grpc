package main

import (
	"fmt"
	"io"
	"log"

	"github.com/deemanthak/go-grpc/calculator/proto"
)

func (s *Server) Average(stream proto.CalculatorService_AverageServer) error {
	fmt.Print("Avegage func was invokded \n")

	var avg float32 = 0
	var icount float32 = 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&proto.AvgResponse{Avg: avg / icount})
		}

		if err != nil {
			log.Fatalf("error reading client stream %v\n", err)
		}

		icount++
		avg += req.Number

		fmt.Printf("AVG : %v\n", avg)
	}

	return nil
}
