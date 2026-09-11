// Package trywait provides a wait group implementation.
package trywait

import "sync"

/*
	In previous chapters, we saw the TryLock() operation on mutexes. This is a non-
	blocking call that returns immediately without waiting. If the lock is not avail-
	able, the function returns false; otherwise, it locks the mutex and returns true.
	Can you write a similar non-blocking function called TryWait() on our imple-
	mentation of a waitgroup from listing 6.8? This function will return immedi-
	ately with false if the waitgroup is not done; otherwise, it returns true.
*/

type BetaWaitGroup struct {
	groupSize int
	cond      *sync.Cond
}

func NewBetaWaitGroup() *BetaWaitGroup {
	return &BetaWaitGroup{
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (wg *BetaWaitGroup) Add(delta int) {
	wg.cond.L.Lock()
	wg.groupSize += delta
	wg.cond.L.Unlock()
}

func (wg *BetaWaitGroup) Wait() {
	wg.cond.L.Lock()
	for wg.groupSize > 0 {
		wg.cond.Wait()
	}
	wg.cond.L.Unlock()
}

// TryWait is a non-blocking version of Wait() that returns immediately
// with false if the waitgroup is not done; otherwise, it returns true.
func (wg *BetaWaitGroup) TryWait() bool {
	wg.cond.L.Lock()
	defer wg.cond.L.Unlock()
	return wg.groupSize == 0
}

// TryWait2 is the solution from the book.
func (wg *BetaWaitGroup) TryWait2() bool {
	wg.cond.L.Lock()
	result := wg.groupSize == 0
	wg.cond.L.Unlock()
	return result
}

func (wg *BetaWaitGroup) Done() {
	wg.cond.L.Lock()
	wg.groupSize--
	if wg.groupSize == 0 {
		wg.cond.Broadcast()
	}
	wg.cond.L.Unlock()
}
