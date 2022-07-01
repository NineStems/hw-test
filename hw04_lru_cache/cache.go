package hw04lrucache

import (
	"sync"
)

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

func (l *lruCache) Set(key Key, value interface{}) bool {
	l.mu.Lock()
	defer l.mu.Unlock()
	item, has := l.items[key]
	if !has {
		item = l.queue.PushFront(cacheItem{key: key, value: value})
	} else {
		item.Value = cacheItem{key: key, value: value}
		l.queue.MoveToFront(item)
	}
	l.items[key] = item
	if l.queue.Len() > l.capacity {
		item = l.queue.Back()
		l.queue.Remove(item)
		ci := item.Value.(cacheItem)
		delete(l.items, ci.key)
	}
	return has
}

func (l *lruCache) Get(key Key) (interface{}, bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.queue.Len() == 0 {
		return nil, false
	}
	if item, has := l.items[key]; has {
		l.queue.MoveToFront(item)
		ci := item.Value.(cacheItem)
		return ci.value, true
	}
	return nil, false
}

func (l *lruCache) Clear() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.queue = NewList()
	l.items = make(map[Key]*ListItem, l.capacity)
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
