package with_rw_lock

import (
	"errors"
	"reflect"
	"sync"

	"github.com/gsxab/go-generic_lru"
)

type WithLock[Key comparable, Value any, C generic_lru.Cache[Key, Value]] struct {
	lock  RWLocker
	Cache C
}

func New[Key comparable, Value any, C generic_lru.Cache[Key, Value]](c C, lock RWLocker) *WithLock[Key, Value, C] {
	if lock == nil || reflect.ValueOf(lock).IsNil() {
		lock = &sync.RWMutex{}
	}
	return &WithLock[Key, Value, C]{
		lock:  lock,
		Cache: c,
	}
}

func (w *WithLock[Key, Value, C]) Add(key Key, value Value) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.Cache.Add(key, value)
}

func (w *WithLock[Key, Value, C]) Get(key Key) (value Value, ok bool) {
	w.lock.RLock()
	defer w.lock.RUnlock()
	return w.Cache.Get(key)
}

func (w *WithLock[Key, Value, C]) Remove(key Key) (value Value, ok bool) {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.Cache.Remove(key)
}

func (w *WithLock[Key, Value, C]) RemoveOldest() (key Key, value Value, ok bool) {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.Cache.RemoveOldest()
}

func (w *WithLock[Key, Value, C]) GetOldest() (key Key, value Value, ok bool) {
	w.lock.RLock()
	defer w.lock.RUnlock()
	return w.Cache.GetOldest()
}

func (w *WithLock[Key, Value, C]) ApplyRO(f func(generic_lru.Cache[Key, Value])) {
	w.lock.RLock()
	defer w.lock.RUnlock()
	w.Cache.ApplyRO(f)
}

func (w *WithLock[Key, Value, C]) ApplyRW(f func(generic_lru.Cache[Key, Value])) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.Cache.ApplyRW(f)
}

func (w *WithLock[Key, Value, C]) Len() int {
	w.lock.RLock()
	defer w.lock.RUnlock()
	return w.Cache.Len()
}

func (w *WithLock[Key, Value, C]) Container() (interface{}, error) {
	return nil, errors.New("cannot get underlying container without breaking lock protocol")
}

func (w *WithLock[Key, Value, C]) Clear() {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.Cache.Clear()
}
