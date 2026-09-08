package main

/*
	Listing 4.15 (originally from chapter 3) does not use any mutexes to protect
	access to its shared variable. This is bad practice. Change this program so that
	access to the shared seconds variable is protected by a mutex. Hint: you might
	need to copy a variable.
*/

import (
	"fmt"
	"sync"
	"time"
)

// Note: the duplicate numbers in the output (e.g. 4,4,3,3...) are NOT a race
// condition — run this with -race and it's clean. It's just that main polls
// every 500ms while countdown only decrements every 1s, so main sees the same
// value twice before it changes. The mutex's job here is only to make every
// read/write of the shared int (including the loop conditions) safe, not to
// make the printed cadence look like a clean 5,4,3,2,1 countdown.
func main() {
	count := 5
	mutex := sync.Mutex{}
	go countdown(&count, &mutex)
	c := count

	for c > 0 {
		time.Sleep(500 * time.Millisecond)
		mutex.Lock()
		fmt.Println(count)
		c = count
		mutex.Unlock()
	}
}

func countdown(seconds *int, mutex *sync.Mutex) {
	mutex.Lock()
	s := *seconds
	mutex.Unlock()
	for s > 0 {
		time.Sleep(1 * time.Second)
		mutex.Lock()
		*seconds -= 1
		s = *seconds
		mutex.Unlock()
	}
}
