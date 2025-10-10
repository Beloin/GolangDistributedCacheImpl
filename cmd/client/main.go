package main

import (
	"log"
	"os"

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

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	log.Println("Ended client")
}
