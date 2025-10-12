// Package cache
// Contains cache services
package cache

import "beloin.com/distributed-cache/pkg/cacher"

type CacheService struct {
	C cacher.Cacher
}

func (service *CacheService) GetString(s string) (string, bool) {
	return service.C.GetString(s)
}

func (service *CacheService) SetString(s string, value string) bool {
	return service.C.SetString(s, value)
}

func (service *CacheService) Paginate(from, limit int) map[string]string {
	return service.C.Paginate(from, limit)
}
