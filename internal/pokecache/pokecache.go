package pokecache

import (
	"fmt"
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

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}
	go cache.reapLoop(interval)
	return &cache
} 

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Printf("Adding %v to cache...\n", key)
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if cacheEntry, ok := c.cache[key]; ok {
		return cacheEntry.val, true
	}
	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	fmt.Println("starting reap cycle...")
	
	func() {
		for {
			t := <-ticker.C
			fmt.Println("reaping cache...")

			func() {
				c.mu.Lock()
				defer c.mu.Unlock()

				fmt.Printf("Current time: %v\n", t)
				fmt.Printf("Interval: %v\n", interval)
				fmt.Println("Cache contents:")

				for key, cacheEntry := range c.cache {
					fmt.Printf("%v: created at %v\n", key, cacheEntry.createdAt)
					if t.Sub(cacheEntry.createdAt) > interval {
						fmt.Printf("deleting %v from cache...\n", key)
						delete(c.cache, key)
					}
				}
			}()
		}
	}()
}