package math

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"sync"
	"time"
)

type CacheManager interface {
	AddToCache(hashKey string, data []tgbotapi.Chattable)
	GetFromCache(hashKey string) ([]tgbotapi.Chattable, bool)
}

type cacheManager struct {
	cache    sync.Map
	Duration time.Duration
}

func NewCacheManager() CacheManager {
	return &cacheManager{Duration: DefaultCacheDuration}
}

type CacheItem struct {
	Data       []tgbotapi.Chattable
	LastAccess time.Time
}

func (s *cacheManager) AddToCache(hashKey string, data []tgbotapi.Chattable) {
	s.cache.Store(hashKey, CacheItem{
		Data:       data,
		LastAccess: time.Now(),
	})
}

func (s *cacheManager) GetFromCache(hashKey string) ([]tgbotapi.Chattable, bool) {
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
