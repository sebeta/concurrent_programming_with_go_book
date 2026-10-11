package main

/*
	In listing 8.15, we have two goroutines. The generateTemp() function simulates
	reading and sending the temperature on a channel every 200 ms. The output-
	Temp() function simply outputs a message found on a channel every 2 seconds.
	Can you write a main() function, using a select statement, that reads messages
	coming from the generateTemp() goroutine and sends only the latest tempera-
	ture to the outputTemp() channel? Since the generateTemp() function outputs
	values faster than the outputTemp() function, you’ll need to discard some values
	so that only the most up-to-date temperature is displayed.
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Data flow: generateTemp --(temp)--> main (latest) --(out)--> outputTemp
	// The two channels are NOT connected to each other. The only bridge is the
	// `latest` variable inside main. Overwriting it is how old readings get
	// discarded.
	temp := generateTemp()
	out := make(chan int) // unbuffered: a send only completes if someone receives
	outputTemp(out)

	var latest int

	// sendCh is a "switch" for the send case below. A nil channel blocks
	// forever, and select never picks a case that can't proceed, so while
	// sendCh is nil the send case is disabled. If we used `out` directly,
	// we could send a bogus 0 before the first real reading arrives.
	var sendCh chan int
	timeout := time.After(10 * time.Second)

	// select runs once per iteration, so it must live inside a loop.
	// Each pass asks: "which of these operations can complete right now?"
	// If several can, one is picked at random; if none, it waits.
	for {
		select {
		// Receive half of the rendezvous with generateTemp's `output <- temp`.
		// Runs every ~200ms. Keep only the newest value and enable sending.
		case measure := <-temp:
			latest = measure
			sendCh = out

		// This is an OFFER to send, not an immediate send. It completes only
		// when outputTemp is waiting on `<-input`, i.e. about every 2s (after
		// its Sleep). In between, ~10 readings overwrite `latest`, so each
		// send carries the freshest value. Sends == receives == prints.
		// (Order of the two prints is not guaranteed: different goroutines.)
		case sendCh <- latest:
			fmt.Println("Sent: ", latest)

		// Without this, main would loop forever. If main returned early,
		// the other goroutines would be killed with it.
		case <-timeout:
			fmt.Println("Done")
			return
		}
	}
}

func generateTemp() chan int {
	output := make(chan int)
	go func() {
		temp := 50 // Fahrenheit
		for {
			output <- temp
			temp += rand.Intn(3) - 1
			time.Sleep(200 * time.Millisecond)
		}
	}()

	return output
}

func outputTemp(input chan int) {
	go func() {
		for {
			fmt.Println("Current temp:", <-input)
			time.Sleep(2 * time.Second)
		}
	}()
}
