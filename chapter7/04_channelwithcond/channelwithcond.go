package main

/*
	Modify listings 7.14 through 7.17 to implement a channel using condition
	variables instead of semaphores. The implementation also needs to support
	channels with a zero-sized buffer.
*/

import (
	"container/list"
	"fmt"
	"sync"
	"time"
)

type Channel[M any] struct {
	capacity int
	cond     *sync.Cond
	buffer   *list.List
}

func NewChannel[M any](capacity int) *Channel[M] {
	return &Channel[M]{
		cond:     sync.NewCond(&sync.Mutex{}),
		buffer:   list.New(),
		capacity: capacity,
	}
}

// Send function needs to satisfy the following requirements:
// 1. Block the goroutine if the buffer is full.
// 2. Otherwise, safely add the message to the buffer.
// 3. If any receiver goroutines are blocked, waiting for messages, resume one of them.
func (c *Channel[M]) Send(message M) {
	c.cond.L.Lock()
	defer c.cond.L.Unlock()

	// Wait while the buffer is full; re-check after every wakeup.
	for c.buffer.Len() == c.capacity {
		c.cond.Wait()
	}

	// Buffer has room: enqueue the message.
	c.buffer.PushBack(message)

	// Wake all waiters so a blocked receiver can notice the new message.
	c.cond.Broadcast()
}

// Receive function needs to satisfy the following requirements:
// 1. Unblock a sender waiting for capacity space.
// 2. If the buffer is empty, block the receiver.
// 3. Otherwise, safely consume the next message from the buffer.
func (c *Channel[M]) Receive() M {
	c.cond.L.Lock()
	defer c.cond.L.Unlock()

	// Wait while the buffer is empty; re-check after every wakeup.
	for c.buffer.Len() == 0 {
		c.cond.Wait()
	}

	// Buffer has a message: dequeue it.
	v := c.buffer.Remove(c.buffer.Front()).(M)

	// Wake all waiters so a blocked sender can notice the freed slot.
	c.cond.Broadcast()

	return v
}

func receiver(messages *Channel[int], wGroup *sync.WaitGroup) {
	msg := 0
	for msg != -1 {
		time.Sleep(1 * time.Second)
		msg = messages.Receive()
		fmt.Println("Received:", msg)
	}
	wGroup.Done()
}

func main() {
	channel := NewChannel[int](2)
	wGroup := sync.WaitGroup{}
	wGroup.Add(1)
	go receiver(channel, &wGroup)
	for i := 1; i <= 6; i++ {
		fmt.Println("Sending: ", i)
		channel.Send(i)
	}
	channel.Send(-1)
	wGroup.Wait()
}
