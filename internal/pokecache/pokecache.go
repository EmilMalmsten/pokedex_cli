package pokecache

import (
	"fmt"
	"time"
)

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	fmt.Printf("adding cache with key %s\n", key)
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	fmt.Printf("trying to find data with key %s\n", key)
	entry, ok := c.cache[key]
	if !ok {
		fmt.Println("data not found")
		return []byte{}, false
	}
	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) reap(interval time.Duration) {
	cutoffTime := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(cutoffTime) {
			delete(c.cache, k)
		}
	}
}
