package main

import (
	"fmt"
	"time"
)

const (
	passwordToGuess = "go far"
	alphabet        = " abcdefghijklmnopqrstuvwxyz"
)

func main() {
	// Shared across all goroutines as `stop`. Never sent on directly — it's only
	// ever closed, which is what lets one goroutine broadcast "stop searching"
	// to every other goroutine at once (a single send would only wake one).
	finished := make(chan int)
	passwordFound := make(chan string)
	for i := 1; i <= 387_420_488; i += 10_000_000 {
		go guessPassword(i, i+10_000_000, finished, passwordFound)
	}
	fmt.Println("password found:", <-passwordFound)
	close(passwordFound)
	time.Sleep(5 * time.Second)
}

func guessPassword(from int, upto int, stop chan int, result chan string) {
	for guessN := from; guessN < upto; guessN += 1 {
		select {
		// Receiving from a closed channel never blocks and returns immediately
		// (the zero value), so once stop is closed, every goroutine's select
		// hits this case right away instead of the default branch.
		case <-stop:
			fmt.Printf("Stopped at %d [%d,%d)\n", guessN, from, upto)
			return
		default:
			if toBase27(guessN) == passwordToGuess {
				result <- toBase27(guessN)
				// Closing (not sending) lets ALL other goroutines notice at once.
				close(stop)
				return
			}
		}
	}
	fmt.Printf("Not found between [%d,%d)\n", from, upto)
}

func toBase27(n int) string {
	result := ""
	for n > 0 {
		result = string(alphabet[n%27]) + result
		n /= 27
	}
	return result
}
