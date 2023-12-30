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

package lru_with_rw_lock

import (
	"sync"

	"github.com/gsxab/go-generic_lru"
	"github.com/gsxab/go-generic_lru/lru"
	"github.com/gsxab/go-generic_lru/with_rw_lock"
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
