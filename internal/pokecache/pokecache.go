package pokecache

import (
	"sync"
	"time"
)

type Cache struct{
	cache map[string]cacheEntry
	mu *sync.Mutex
}

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return cache
} 

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheEntry, ok := c.cache[key]
	return cacheEntry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	
	for t := range ticker.C {
		c.reap(t.UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, cacheEntry := range c.cache {
		if now.Sub(cacheEntry.createdAt) > interval {
			delete(c.cache, key)
		}
	}
}