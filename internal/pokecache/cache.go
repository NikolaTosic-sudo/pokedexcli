package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entry map[string]cacheEntry
	Mu    *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	Val       []byte
}

func NewCache(interval time.Duration) Cache {

	cache := Cache{
		make(map[string]cacheEntry),
		&sync.Mutex{},
	}

	go cache.reapLoop(interval)

	return cache
}
