package main

import (
	"fmt"
	"time"
)

/*
	Run Go’s race detector on listing 3.1. Does the result contain a race condition?
	If it does, can you explain why it happens?

	go run -race countdown.go
*/

func main() {
	count := 5
	go countdown(&count)
	for count > 0 {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(count)
	}

	// Yes, this program has a race condition.
	// The main goroutine reads count (in the for condition and in Println) while the countdown
	// goroutine concurrently reads and writes it (*seconds -= 1), with no synchronization between
	// them at all — no mutex, no channel, nothing that establishes a happens-before relationship.
	// It's not that a "resource" is held or released; there's simply no mechanism preventing both
	// goroutines from touching the same memory location at the same instant. The race detector
	// flags exactly this: an unsynchronized read in one goroutine overlapping with an unsynchronized
	// write in another. In practice this means the read isn't guaranteed to be atomic or to see the
	// latest write, so main could observe a stale value, the same value twice, or (on other
	// architectures/compilers) even a torn value — the duplicated prints we see are just one
	// harmless-looking symptom of a bug that isn't guaranteed to stay harmless.
}

func countdown(seconds *int) {
	for *seconds > 0 {
		time.Sleep(1 * time.Second)
		*seconds -= 1
	}
}
