package main

import (
	"context"
	"io"
	"log"
	"os"
	"time"

	"beloin.com/distributed-cache/internal/network/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	c := proto.NewRestoreServiceClient(conn)
	c2 := proto.NewCacheServiceClient(conn)
	storeKv(ctx, c2)
	doRead(ctx, c2)

	doRestore(ctx, c)

	log.Println("Ended client")
}

func storeKv(ctx context.Context, c proto.CacheServiceClient) {
	protoString := wrapperspb.String("john smith")
	newVar, err := anypb.New(protoString)
	if err != nil {
		return
	}

	request := &proto.StoreCacheRequest{
		Key:   "key",
		Value: newVar,
	}
	_, err = c.StoreCache(ctx, request)
	if err != nil {
		log.Fatalf("error here: %v", err)
	}
}

func doRestore(ctx context.Context, c proto.RestoreServiceClient) {
	req := &proto.RestoreRequest{
		Batch: 10,
	}
	stream, err := c.Restore(ctx, req)
	if err != nil {
		log.Fatalln("Could not recieve restore")
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			log.Println("Done while recieveing")
			break
		}

		if err != nil {
			log.Fatalf("error %v", err)
		}

		log.Printf("Recieved all data:\n%v", msg.Caches)
	}
}

func doRead(ctx context.Context, c proto.CacheServiceClient) {
	req := &proto.ReadCacheRequest{Key: "key", MaxAge: &proto.Duration{Value: 1, Unit: proto.Unit_Day}}
	res, err := c.ReadCache(ctx, req)
	if err != nil {
		log.Fatalln("Could not read cache")
	}

	log.Printf("Value read: %v\n", res.Cache)
}
