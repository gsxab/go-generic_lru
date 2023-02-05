package with_rw_lock

import "sync"

type RWLocker interface {
	RLock()
	TryRLock() bool
	RUnlock()
	Lock()
	TryLock() bool
	Unlock()
	RLocker() sync.Locker
}
