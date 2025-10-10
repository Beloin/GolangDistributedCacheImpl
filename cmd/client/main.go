package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "beloin.com/distributed-cache/internal/network/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const addr = "localhost:50051"

func main() {
	*log.Default() = *log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	// creds, err := credentials.NewClientTLSFromFile("", "localhost")
	// dialopt := grpc.WithTransportCredentials(creds)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}
	defer conn.Close()

	service := pb.NewGreetServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := pb.GreetRequest{Name: "Juan"}
	resp, err :=service.Greet(ctx, &req)
	if err != nil {
		panic(err)
	}

	log.Printf("Response: %s", resp.GetResult())

	log.Println("Ended client")
}
