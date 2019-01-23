package basestruct

import (
	"sync"
)

type Cache struct {
	sync.Mutex
	mapping map[string]string
}

func NewCache() *Cache {
	cache := &Cache{
		mapping: make(map[string]string),
	}
	return cache
}

func (cache *Cache) Set(key string, value string) {
	cache.Lock()
	cache.mapping[key] = value
	cache.Unlock()
}

func (cache *Cache) Get(key string) string {
	cache.Lock()
	res := cache.mapping[key]
	cache.Unlock()
	return res
}
