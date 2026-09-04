package main

/*
	Consider the following listing. Can you find the race condition in this program
	without running the race detector? Hint: Try running the program several
	times to see if it results in a race condition.
*/

import (
	"fmt"
	"time"
)

func addNextNumber(nextNum *[101]int) {
	i := 0
	for nextNum[i] != 0 {
		i++
	}
	nextNum[i] = nextNum[i-1] + 1
}

// Multiple goroutines read and write nextNum without synchronization.
// There's no atomicity between steps. Is very likely that two or more goroutines
// will read the same value of nextNum[i] and overwrite it before the other goroutine
// can increment it, so one position is lost. Since there are only 100 goroutines,
// the last position (100) will be always 0 and the loop will never exit.

func main() {
	nextNum := [101]int{1}
	for i := 0; i < 100; i++ {
		go addNextNumber(&nextNum)
	}
	for nextNum[100] == 0 {
		println("Waiting for goroutines to complete")
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println(nextNum)
}
