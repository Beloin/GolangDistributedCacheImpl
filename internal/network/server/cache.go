package server

import (
	"context"

	"beloin.com/distributed-cache/internal/cache"
	"beloin.com/distributed-cache/internal/network/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type CacherServer struct {
	proto.UnimplementedCacheServiceServer
	proto.UnimplementedRestoreServiceServer

	Service cache.CacheService
}

func (cs *CacherServer) ReadCache(ctx context.Context, request *proto.ReadCacheRequest) (*proto.ReadCacheResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	key := request.Key
	v, found := cs.Service.GetString(key)
	protoString := wrapperspb.String(v)

	if found {
		newVar, err := anypb.New(protoString)
		if err != nil {
			return nil, err
		}
		cache := proto.Cache{Key: key, Value: newVar, Ttl: nil, Type: proto.CacheType_string}
		return &proto.ReadCacheResponse{Cache: &cache, Hit: true}, nil
	}

	cache := proto.Cache{Key: key, Value: nil, Ttl: nil, Type: proto.CacheType_string}
	return &proto.ReadCacheResponse{Cache: &cache, Hit: false}, nil
}

func (cs *CacherServer) StoreCache(ctx context.Context, request *proto.StoreCacheRequest) (*proto.StoreCacheResponse, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	if request.Type != proto.CacheType_string {
		return nil, status.Errorf(codes.Unimplemented, "Cannot store anything other than string")
	}

	str := wrapperspb.StringValue{}
	err := request.Value.UnmarshalTo(&str)
	if err != nil {
		return nil, err
	}

	key := request.Key
	value := str.String()

	_, hasOv := cs.Service.GetString(key)
	stored := cs.Service.SetString(key, value)
	if stored {
		return &proto.StoreCacheResponse{Key: key, HasOverriden: hasOv}, nil
	}

	return nil, status.Errorf(codes.Internal, "could not store desired data")
}
