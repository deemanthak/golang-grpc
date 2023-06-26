package main

import (
	"context"
	"log"

	"github.com/deemanthak/go-grpc/calculator/proto"
)

func (s Server) Calculate(ctx context.Context, cr *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	log.Printf("Greet function was invoked with %v\n", cr)

	return &proto.CalculatorResponse{Total: cr.Num1 + cr.Num2}, nil
}
