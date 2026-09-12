package main

import (
	"sync"
)

type Semaphore struct {
	mu       sync.Mutex
	cond     *sync.Cond
	count    int // Available permits
	capacity int // Maximum number of permits
}

func NewSemaphore(capacity int) *Semaphore {
	s := &Semaphore{
		count:    capacity,
		capacity: capacity,
	}
	s.cond = sync.NewCond(&s.mu)
	return s
}

// Acquire takes a permit, blocking if none are available
func (s *Semaphore) Acquire() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for s.count == 0 {
		s.cond.Wait() // Sleeps until a permit is released
	}
	s.count--
}

// Release frees a permit and wakes up a waiter
func (s *Semaphore) Release() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.count++
	s.cond.Signal() // Wakes up ONE waiting goroutine
}
