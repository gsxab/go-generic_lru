package generic_lru

type Cache[Key comparable, Value any] interface {
	Add(key Key, value Value)
	Get(key Key) (value Value, ok bool)
	Remove(key Key) (value Value, ok bool)
	RemoveOldest() (key Key, value Value, ok bool)
	GetOldest() (key Key, value Value, ok bool)
	ApplyRO(f func(Cache[Key, Value]))
	ApplyRW(f func(Cache[Key, Value]))
	Len() int
	Clear()
}
