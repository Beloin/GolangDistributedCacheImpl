package main

import (
	"context"
	"errors"
	"time"

	pb "beloin.com/distributed-cache/internal/network/proto"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	results := make(chan string, 1) // Prevent starvation
	defer close(results)
	go internalWork(results)
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case res := <-results:
		return &pb.GreetResponse{Result: res + req.Name}, nil
	}

	return nil, errors.New("problem with timeoout")
}

func internalWork(results chan string) {
	time.Sleep(10 * time.Second)
	results <- "Hello "
}
