package lru

import (
	"container/list"
	"time"
)

// LRU is the interface for the LRU cache
type LRU interface {
	// Adds a value to the cache, or updates an item in the cache.
	// It returns true if an item needed to be removed for storing the new item.
	Add(key interface{}, value interface{}, ttl int64) bool

	// Returns the value of the provided key, and updates status of the item
	// in the cache.
	Get(key interface{}) (value interface{}, ok bool)

	// Check if a key exsists in the cache.
	Contains(key interface{}) (ok bool)

	// Expires returns the time of expiration.
	Expires(key interface{}) (expires time.Time, ok bool)

	// Fetches a value which has expired, or does not exits and fills the cache.
	Fetch(key interface{}, ttl int64, call func() (interface{}, error)) (value interface{}, ok bool, err error)

	// Removes a key from the cache.
	Remove(key interface{}) bool

	// Removes the oldest entry from cache.
	RemoveOldest() (interface{}, interface{}, bool)

	// Returns the oldest entry from the cache.
	GetOldest() (interface{}, interface{}, bool)

	// Returns a slice of the keys in the cache, from oldest to newest.
	Keys() []interface{}

	// Returns the number of items in the cache.
	Len() int

	// Purge is purging the full cache.
	Purge()
}

type Sized interface {
	Size() int64
}

type LRUCache struct {
	size  int
	items map[interface{}]*list.Element
	list  *list.List
}

// Item represents the internal presentation of a cache entry
type Item struct {
	key        interface{}
	group      string // this is not yet used
	promotions int32
	refs       int32
	expires    int64
	ttl        int64
	timestamp  int64
	size       int64
	value      interface{}
}
