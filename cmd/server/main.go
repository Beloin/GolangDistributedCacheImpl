package main

import (
	"log"
	"net"
	"os"

	"beloin.com/distributed-cache/internal/cache"
	"beloin.com/distributed-cache/internal/network/proto"
	"beloin.com/distributed-cache/internal/network/server"
	"beloin.com/distributed-cache/pkg/cacher"
	"google.golang.org/grpc"
)

const addr = "0.0.0.0:50051"

type Server struct{}

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

	cacherServer := &server.CacherServer{
		Service: cache.CacheService{C: cacher.NewHashMapCacher()},
	}
	proto.RegisterRestoreServiceServer(s, cacherServer)
	proto.RegisterCacheServiceServer(s, cacherServer)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %s\n", err.Error())
	}

	log.Println("Server stopped suscessfully")
}
