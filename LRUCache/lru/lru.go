package lru

import "errors"

type LRU interface {
	Get(interface{}) (interface{}, error)
	Put(interface{}, interface{}) (error)
}


type lruCache struct {
	size int
	remaining int
	cache map[string]string
	queue []string
}

func (lru *lruCache) Get(key interface{}) (interface{}, error) {
	if val, ok := lru.cache[key.(string)]; ok {
		return val, nil
	}
	return  "", errors.New("Key not present")
}

func (lru *lruCache) Put(key, val interface{}) error {
	if _, ok := lru.cache[key.(string)]; ok { // Update
		lru.qDel(key.(string))
	} else if lru.remaining == 0 { // Insert with replacement
		victim := lru.queue[0]
		lru.qDel(victim)
		delete(lru.cache, key.(string))
	}
	lru.cache[key.(string)] = val.(string)
	lru.queue = append(lru.queue, key.(string))

	return nil
}

func NewLRU(size int) LRU {
	return &lruCache{size:size, remaining:size, cache:make(map[string]string), queue:make([]string, size)}
}

// Delete element from queue
func (lru *lruCache) qDel(ele string) {
	for i := 0; i < len(lru.queue); i++ {
		if lru.queue[i] == ele {
			oldlen := len(lru.queue)
			copy(lru.queue[i:],lru.queue[i+1:])
			lru.queue = lru.queue[:oldlen-1]
			break;
		}
	}
}


