package synx

import (
	"sync"
)

type locker struct {
	lock *sync.Mutex
}

func newLock() locker {
	return locker{
		lock: &sync.Mutex{},
	}
}

func (l *locker) Lock() {
	l.lock.Lock()
}

func (l *locker) Unlock() {
	l.lock.Unlock()
}
