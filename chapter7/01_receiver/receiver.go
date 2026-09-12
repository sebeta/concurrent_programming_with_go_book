package main

/*
	In listings 7.1 and 7.2, the receiver doesn’t output the last message STOP. This is
	because the main() goroutine terminates before the receiver() goroutine gets
	the chance to print out the last message. Can you change the logic, without
	using extra concurrency tools and without using the sleep function, so that the
	last message is printed?
*/

import "fmt"

// Basically, the way to deal with this exercise is to use the
// bidirectional default property of the channels.

func main() {
	msgChannel := make(chan string)
	go receiver(msgChannel)
	fmt.Println("Sending HELLO...")
	msgChannel <- "HELLO"
	fmt.Println("Sending THERE...")
	msgChannel <- "THERE"
	fmt.Println("Sending STOP...")
	msgChannel <- "STOP"

	// ← blocks main until receiver makes: messages <- ""
	<-msgChannel
}

func receiver(messages chan string) {
	msg := ""
	for msg != "STOP" {
		msg = <-messages
		fmt.Println("Received:", msg)
	}
	// ← it only executes AFTER printing "Received: STOP"
	messages <- ""
}
