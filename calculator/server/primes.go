package main

import (
	"log"

	"github.com/deemanthak/go-grpc/calculator/proto"
)

func (s Server) Primes(pr *proto.PrimeRequest, cp proto.CalculatorService_PrimesServer) error {
	log.Print("Primes function invoked %v\n", pr)

	var k int32 = 2
	var N int32 = pr.Number

	for N > 1 {
		if N%k == 0 {
			N = N / k
			cp.Send(&proto.PrimeResponse{Prime: k})
		} else {
			k++
		}
	}

	return nil
}
