package main

import (
	"log"
	"net"
	"os"

	pb "beloin.com/distributed-cache/internal/network/proto"
	"google.golang.org/grpc"
)

const addr = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}



func main() {
	*log.Default() = *log.New(os.Stdout, "", log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v\n", addr, err)
	}

	log.Printf("Listening on %s\n", addr)

	s:= grpc.NewServer()

	if err = s.Serve(lis); err !=nil {
		log.Fatalf("Failed to serve %v\n", err)
	}

	log.Println("Server stopped suscessfully")
}
