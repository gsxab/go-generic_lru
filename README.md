# Go - Generic LRU

This repo is based on the `github.com/golang/groupcache/lru` package.
See <https://github.com/golang/groupcache/tree/master/lru> for more information.

## Packages

### `github.com/gsxab/go-generic_lru` with package name `generic_lru`

`Cache[Key, Value]` as a generic interface for the concrete cache implementations.

### `github.com/gsxab/go-generic_lru/lru`

A naive LRU cache based on a linked list and a map.

### `github.com/gsxab/go-generic_lru/with_lock`

A decorator wrapping a `Cache[Key, Value]` with a `sync.Locker`. It locks the container for every r/w operations.

### `github.com/gsxab/go-generic_lru/with_rw_lock`

A decorator wrapping a `Cache[Key, Value]` with a `RWMutex`-like locker. It locks the container with a [readers-writer lock](https://en.wikipedia.org/wiki/Readers%E2%80%93writer_lock) for reading and writing operations accordingly.

### `github.com/gsxab/go-generic_lru/lru_with_lock`, `github.com/gsxab/go-generic_lru/lru_with_rw_lock`

Shortcuts to create lock-wrapped LRU caches.
