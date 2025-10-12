// Package server
package server

import (
	"beloin.com/distributed-cache/internal/network/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func (cs *CacherServer) Restore(req *proto.RestoreRequest, stream grpc.ServerStreamingServer[proto.RestoreResponse]) error {
	// TODO: proper implement pagination

	res := cs.Service.Paginate(0, 0)
	for len(res) > 0 {
		var caches []*proto.Cache
		for key, value := range res {
			protoString := wrapperspb.String(value)
			newVar, err := anypb.New(protoString)
			if err != nil {
				return err
			}
			caches = append(caches, &proto.Cache{
				Key:   key,
				Value: newVar,
				Type:  proto.CacheType_string,
			})
		}

		response := &proto.RestoreResponse{Caches: caches}
		stream.Send(response)
		// res = cs.service.Paginate(0, 0)
		res = map[string]string{}
	}

	return nil
}
