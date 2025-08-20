package pokecache

import (
	_ "fmt"
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	mu       sync.Mutex
	entry    map[string]cacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := Cache{sync.Mutex{}, make(map[string]cacheEntry), interval}
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			c.reapLoop()
		}
	}()
	return &c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.entry[key] = cacheEntry{time.Now(), val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if item, ok := c.entry[key]; ok {
		item.createdAt = time.Now()
		c.entry[key] = item
		return item.val, ok
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, item := range c.entry {
		tempTime := time.Time(item.createdAt).Add(c.interval)
		timeNow := time.Now()
		if item.createdAt.Nanosecond()+timeNow.Nanosecond() > tempTime.Nanosecond() {
			delete(c.entry, key)
		}
	}
}
