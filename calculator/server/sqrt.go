package main

import (
	"context"
	"fmt"
	"math"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/deemanthak/go-grpc/calculator/proto"
)

func (s *Server) Sqrt(ctx context.Context, request *proto.SqrtRequest) (*proto.SqrtResponse, error) {
	fmt.Println("Sqrt invoked")

	number := request.Number

	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("received negative number %d", number))
	}

	return &proto.SqrtResponse{Result: math.Sqrt(float64(number))}, nil
}
