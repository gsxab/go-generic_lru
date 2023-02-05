package with_rw_lock

import "sync"

type RWLocker interface {
	RLock()
	TryRLock() bool
	RUnlock()
	sync.Locker
	RLocker() sync.Locker
}
