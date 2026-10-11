package main

/*
	In listing 8.16, we have a goroutine in the generateNumbers() function that out-
	puts random numbers. Can you write a main() function using a select state-
	ment that continuously consumes from the output channel, printing the output
	on the console until 5 seconds have elapsed from the start of the program?
	After 5 seconds, the function should stop consuming from the output channel,
	and the program should terminate.
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	timeout := time.After(3 * time.Second)
	numbers := generateNumbers()

	for {
		select {
		case n := <-numbers:
			fmt.Println("Number: ", n)
		case <-timeout:
			fmt.Println("Finished after 3 seconds")
			return
		}
	}
}

func generateNumbers() chan int {
	output := make(chan int)
	go func() {
		for {
			output <- rand.Intn(10)
			time.Sleep(200 * time.Millisecond)
		}
	}()

	return output
}
