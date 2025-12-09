package pokecache

import (
	"fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	data     map[string]cacheEntry
	mu       sync.RWMutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		data:     map[string]cacheEntry{},
		mu:       sync.RWMutex{},
		interval: interval,
	}
	go cache.reapLoop(interval)
	return &cache
}

func (c *Cache) Add(key string, val []byte) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	return nil
}

func (c *Cache) Get(key string) ([]byte, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	cached, ok := c.data[key]
	if !ok {
		return nil, fmt.Errorf("couldn't get key %s from cache", key)
	}
	return cached.val, nil
}

func (c *Cache) reapLoop(interval time.Duration) {
	for range time.Tick(interval) {
		for key, value := range c.data {
			isOlder := time.Since(value.createdAt) >= interval
			if isOlder {
				c.mu.Lock()
				delete(c.data, key)
				c.mu.Unlock()
			}
		}
	}
}
