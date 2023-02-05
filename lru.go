/*
Copyright 2013 Google Inc.

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
/*
 * Modification copyright 2023 gsxab.
 *
 * For changes against the original version, see git history.
 */

// Package lru implements an LRU cache.
package generic_lru

import "container/list"

// LRU is an LRU cache. It is not safe for concurrent access.
type LRU[Key comparable, Value any] struct {
	// MaxEntries is the maximum number of cache entries before
	// an item is evicted. Zero means no limit.
	MaxEntries int

	// OnEvicted optionally specifies a callback function to be
	// executed when an entry is purged from the cache.
	OnEvicted func(key Key, value Value)

	ll    *list.List
	cache map[interface{}]*list.Element
}

type entry[Key comparable, Value any] struct {
	key   Key
	value Value
}

// New creates a new LRU.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func New[Key comparable, Value any](maxEntries int) *LRU[Key, Value] {
	return &LRU[Key, Value]{
		MaxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

// Add adds a value to the cache.
func (c *LRU[Key, Value]) Add(key Key, value Value) {
	if c.cache == nil {
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}
	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		ee.Value.(*entry[Key, Value]).value = value
		return
	}
	ele := c.ll.PushFront(&entry[Key, Value]{key, value})
	c.cache[key] = ele
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.RemoveOldest()
	}
}

// Get looks up a key's value from the cache.
func (c *LRU[Key, Value]) Get(key Key) (value Value, ok bool) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry[Key, Value]).value, true
	}
	return
}

// Remove removes the provided key from the cache.
func (c *LRU[Key, Value]) Remove(key Key) (value Value, ok bool) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		_, value = c.removeElement(ele)
		return value, true
	}
	return
}

// RemoveOldest removes the oldest item from the cache.
func (c *LRU[Key, Value]) RemoveOldest() (key Key, value Value, ok bool) {
	if c.cache == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		key, value = c.removeElement(ele)
		return key, value, true
	}
	return
}

// RemoveOldest gets the oldest item from the cache.
func (c *LRU[Key, Value]) GetOldest() (key Key, value Value, ok bool) {
	if c.cache == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		kv := ele.Value.(*entry[Key, Value])
		return kv.key, kv.value, true
	}
	return
}

func (c *LRU[Key, Value]) removeElement(e *list.Element) (key Key, value Value) {
	c.ll.Remove(e)
	kv := e.Value.(*entry[Key, Value])
	delete(c.cache, kv.key)
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
	return kv.key, kv.value
}

func (c *LRU[Key, Value]) ApplyRO(f func(Cache[Key, Value])) {
	f(c)
}

func (c *LRU[Key, Value]) ApplyRW(f func(Cache[Key, Value])) {
	f(c)
}

// Len returns the number of items in the cache.
func (c *LRU[Key, Value]) Len() int {
	if c.cache == nil {
		return 0
	}
	return c.ll.Len()
}

// Clear purges all stored items from the cache.
func (c *LRU[Key, Value]) Clear() {
	if c.OnEvicted != nil {
		for _, e := range c.cache {
			kv := e.Value.(*entry[Key, Value])
			c.OnEvicted(kv.key, kv.value)
		}
	}
	c.ll = nil
	c.cache = nil
}
