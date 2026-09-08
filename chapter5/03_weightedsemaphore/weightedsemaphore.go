// Package weightedsemaphore implements a weighted semaphore.
package weightedsemaphore

import "sync"

/*
	A weighted semaphore is a variation on a semaphore that allows you to acquire and
	release more than one permit at the same time. The function signatures for a
	weighted semaphore are as follows:

	func (rw *WeightedSemaphore) Acquire(permits int)
	func (rw *WeightedSemaphore) Release(permits int)

	Use these function signatures to implement a weighted semaphore with a simi-
	lar functionality to a counting semaphore. It should allow you to acquire or
	release more than one permit.
*/

type WeightedSemaphore struct {
	permits int
	cond    *sync.Cond
}

func NewWeightedSemaphore(permits int) *WeightedSemaphore {
	return &WeightedSemaphore{
		permits: permits,
		cond:    sync.NewCond(&sync.Mutex{}),
	}
}

func (ws *WeightedSemaphore) Acquire(permits int) {
	ws.cond.L.Lock()
	defer ws.cond.L.Unlock()
	for ws.permits < permits {
		ws.cond.Wait()
	}
	ws.permits -= permits
}

func (ws *WeightedSemaphore) Release(permits int) {
	ws.cond.L.Lock()
	defer ws.cond.L.Unlock()
	ws.permits += permits
	ws.cond.Broadcast()
}
