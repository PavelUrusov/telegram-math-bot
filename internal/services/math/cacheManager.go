package math

import (
	"sync"
	"time"
)

type CacheManager interface {
	AddToCache(hashKey string, data interface{})
	GetFromCache(hashKey string) (interface{}, bool)
}

type cacheManager struct {
	cache    sync.Map
	Duration time.Duration
}

func NewCacheManager() CacheManager {
	return &cacheManager{Duration: DefaultCacheDuration}
}

type CacheItem struct {
	Data       interface{}
	LastAccess time.Time
}

func (s *cacheManager) AddToCache(hashKey string, data interface{}) {
	s.cache.Store(hashKey, CacheItem{
		Data:       data,
		LastAccess: time.Now(),
	})
}

func (s *cacheManager) GetFromCache(hashKey string) (interface{}, bool) {
	item, found := s.cache.Load(hashKey)
	if !found {
		return nil, false
	}

	cacheItem := item.(CacheItem)
	if time.Since(cacheItem.LastAccess) > s.Duration*time.Minute {
		s.cache.Delete(hashKey)
		return nil, false
	}

	cacheItem.LastAccess = time.Now()
	s.cache.Store(hashKey, cacheItem)
	return cacheItem.Data, true
}
