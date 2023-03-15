package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCacheEntry(createdAt time.Time, val []byte) CacheEntry {
	return CacheEntry{
		createdAt,
		val,
	}
}

type Cache struct {
	mux     *sync.Mutex
	entries map[string]CacheEntry
}

func NewCache(interval time.Duration) Cache {
	c := Cache{&sync.Mutex{}, map[string]CacheEntry{}}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(pageUrl string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	cEntry := NewCacheEntry(time.Now(), val)
	c.entries[pageUrl] = cEntry

}

func (c Cache) Get(pageUrl string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	entry, exists := c.entries[pageUrl]

	if !exists {
		return []byte{}, false
	}

	return entry.val, true
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}

}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()

	for pageUrl, entry := range c.entries {
		if entry.createdAt.Before(now.Add(last)) {
			delete(c.entries, pageUrl)
		}
	}
}
