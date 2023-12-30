/*
Copyright 2023 gsxab.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package with_lock

import (
	"errors"
	"reflect"
	"sync"

	"github.com/gsxab/go-generic_lru"
)

type WithLock[Key comparable, Value any, C generic_lru.Cache[Key, Value]] struct {
	lock  sync.Locker
	Cache C
}

func New[Key comparable, Value any, C generic_lru.Cache[Key, Value]](c C, lock sync.Locker) *WithLock[Key, Value, C] {
	if lock == nil || reflect.ValueOf(lock).IsNil() {
		lock = &sync.Mutex{}
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
	w.lock.Lock()
	defer w.lock.Unlock()
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
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.Cache.GetOldest()
}

func (w *WithLock[Key, Value, C]) ApplyRO(f func(generic_lru.Cache[Key, Value])) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.Cache.ApplyRO(f)
}

func (w *WithLock[Key, Value, C]) ApplyRW(f func(generic_lru.Cache[Key, Value])) {
	w.lock.Lock()
	defer w.lock.Unlock()
	w.Cache.ApplyRW(f)
}

func (w *WithLock[Key, Value, C]) Len() int {
	w.lock.Lock()
	defer w.lock.Unlock()
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
