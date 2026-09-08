// Package readerswritermutex implements a readers-writer mutex with non-blocking TryLock() and TryReadLock() functions.
package readerswritermutex

/*
	2. Add a non-blocking TryLock() function to the implementation of the readers-
	writer mutex. The function should try to lock the writer's side of the lock. If the
	lock is acquired, it should return a true value; otherwise, the function should
	return immediately, without blocking, with a false return value.

	3. Add a non-blocking TryReadLock() function to the implementation of the readers-
	writer lock. The function should try to lock the readers’ side of the lock. Just like
	in exercise 2, the function should return immediately with true if it managed to
	obtain the lock and return false otherwise.
*/

import "sync"

type ReadersWriterMutex struct {
	readersCount int
	// Internal lock: protects readersCount from concurrent reads/writes. Never guards user data.
	readersLock sync.Mutex
	// Business lock: grants mutual exclusion between readers and writers.
	// Held by the first reader in and released by the last reader out; held directly by a writer.
	globalLock sync.Mutex
}

func (rw *ReadersWriterMutex) ReadLock() {
	rw.readersLock.Lock()
	defer rw.readersLock.Unlock()

	rw.readersCount++
	if rw.readersCount == 1 {
		rw.globalLock.Lock()
	}
}

func (rw *ReadersWriterMutex) ReadUnlock() {
	rw.readersLock.Lock()
	defer rw.readersLock.Unlock()

	rw.readersCount--
	if rw.readersCount == 0 {
		rw.globalLock.Unlock()
	}
}

func (rw *ReadersWriterMutex) TryReadLock() bool {
	if !rw.readersLock.TryLock() { // If TryLock() is true, it locks the readers lock.
		return false
	}
	defer rw.readersLock.Unlock()

	if rw.readersCount == 0 {
		if !rw.globalLock.TryLock() { // If TryLock() is true, it locks the writers lock.
			return false
		}
	}
	rw.readersCount++
	return true
}

func (rw *ReadersWriterMutex) WriteLock() {
	rw.globalLock.Lock()
}

func (rw *ReadersWriterMutex) WriteUnlock() {
	rw.globalLock.Unlock()
}

func (rw *ReadersWriterMutex) TryWriteLock() bool {
	return rw.globalLock.TryLock()
}
