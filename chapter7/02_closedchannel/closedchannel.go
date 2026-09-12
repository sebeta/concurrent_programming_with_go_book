package main

/*
	In listing 7.8, the receiver reads a 0 when the channel is closed. Can you try it
	with different data types? What happens if the channel is of type string? What if
	it is of type slice?
*/

import (
	"fmt"
	"time"
)

func main() {
	intChannel := make(chan int)
	stringChannel := make(chan string)
	sliceChannel := make(chan []string)
	go intReceiver(intChannel)
	go stringReceiver(stringChannel)
	go sliceReceiver(sliceChannel)
	for i := 1; i <= 3; i++ {
		fmt.Println(time.Now().Format("15:04:05"), "Sending to int channel:", i)
		intChannel <- i
		stringChannel <- fmt.Sprintf("string %d", i)
		sliceChannel <- []string{fmt.Sprintf("slice %d", i)}
		time.Sleep(1 * time.Second)
	}
	close(intChannel)
	close(stringChannel)
	close(sliceChannel)
	time.Sleep(3 * time.Second)
}

// When the string channel is closed, the receiver reads a "" when the channel is closed.
// When the slice of strings channel is closed, the receiver reads a []string{} when the channel is closed.

func intReceiver(messages <-chan int) {
	for {
		msg := <-messages
		fmt.Println(time.Now().Format("15:04:05"), "Received:", msg)
		time.Sleep(1 * time.Second)
	}
}

func stringReceiver(messages <-chan string) {
	for {
		msg := <-messages
		fmt.Println(time.Now().Format("15:04:05"), "Received:", msg)
		time.Sleep(1 * time.Second)
	}
}

func sliceReceiver(messages <-chan []string) {
	for {
		msg := <-messages
		fmt.Println(time.Now().Format("15:04:05"), "Received:", msg)
		time.Sleep(1 * time.Second)
	}
}
