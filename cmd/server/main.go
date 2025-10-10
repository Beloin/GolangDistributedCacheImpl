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
	pb.UnimplementedGreetServiceServer
}

func main() {
	*log.Default() = *log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// TODO: See this
	// tls.Listen(network string, laddr string, config *tls.Config)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on %s: %v\n", addr, err)
	}
	defer lis.Close()

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	service := Server{}
	pb.RegisterGreetServiceServer(s, &service)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %s\n", err.Error())
	}

	log.Println("Server stopped suscessfully")
}
