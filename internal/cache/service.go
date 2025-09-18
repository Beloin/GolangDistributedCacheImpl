// Package cache
// Contains cache services
package cache

import "beloin.com/distributed-cache/pkg/cacher"

type CacheService struct {
	c cacher.Cacher
}

func (service *CacheService) GetString(s string) (string, bool) {
	return service.c.GetString(s)
}

func (service *CacheService) SetString(s string, value string) bool {
	return service.c.SetString(s, value)
}
