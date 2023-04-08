package lru_with_rw_lock

import (
	"github.com/gsxab/go-generic_lru"
	"github.com/gsxab/go-generic_lru/lru"
	"github.com/gsxab/go-generic_lru/with_rw_lock"
	"sync"
)

func NewWithLock[Key comparable, Value any](maxEntries int, lock *sync.RWMutex) generic_lru.Cache[Key, Value] {
	return with_rw_lock.New[Key, Value, *lru.LRU[Key, Value]](lru.New[Key, Value](maxEntries), lock)
}

func New[Key comparable, Value any](maxEntries int) generic_lru.Cache[Key, Value] {
	return with_rw_lock.New[Key, Value, *lru.LRU[Key, Value]](lru.New[Key, Value](maxEntries), nil)
}

func NewWithLockWithEvicted[Key comparable, Value any](maxEntries int, lock *sync.RWMutex, onEvicted func(Key, Value)) generic_lru.Cache[Key, Value] {
	return with_rw_lock.New[Key, Value, *lru.LRU[Key, Value]](lru.NewWithOnEvicted[Key, Value](maxEntries, onEvicted), lock)
}

func NewWithEvicted[Key comparable, Value any](maxEntries int, onEvicted func(Key, Value)) generic_lru.Cache[Key, Value] {
	return with_rw_lock.New[Key, Value, *lru.LRU[Key, Value]](lru.NewWithOnEvicted[Key, Value](maxEntries, onEvicted), nil)
}
