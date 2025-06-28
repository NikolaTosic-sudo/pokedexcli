package pokecache

import (
	"time"
)

func (c *Cache) AddCache(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Entry[key] = cacheEntry{
		time.Now(),
		val,
	}
}

func (c *Cache) GetCache(key string) ([]byte, bool) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	cacheVal, exists := c.Entry[key]
	if !exists {
		return []byte{}, false
	}

	return cacheVal.Val, true
}

func (c *Cache) reapLoop(interval time.Duration) {

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for t := range ticker.C {
		for key, entry := range c.Entry {
			c.Mu.Lock()
			if t.Sub(entry.createdAt) > 5 {
				delete(c.Entry, key)
			}
			c.Mu.Unlock()
		}
	}
}
