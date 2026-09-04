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
	// It happens because the main goroutine reads the value of count before the countdown goroutine has finished.
	// For example, when the main goroutine reads 3, the countdown goroutine may still be decrementing it, but the
	// resource is not yet released, so the main goroutine may read the same value again.
}

func countdown(seconds *int) {
	for *seconds > 0 {
		time.Sleep(1 * time.Second)
		*seconds -= 1
	}
}
