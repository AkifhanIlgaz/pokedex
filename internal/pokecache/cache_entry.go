package pokecache

import "time"

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
