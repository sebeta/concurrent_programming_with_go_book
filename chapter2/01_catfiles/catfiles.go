package main

import (
	"fmt"
	"os"
	"time"
)

/*
	Write a program similar to the one in listing 2.3 that accepts a list of text filenames as
	arguments. For each filename, the program should spawn a new goroutine that will output the
	contents of that file to the console. You can use the time.Sleep() function to wait for the
	child goroutines to complete (until you know how to do this better). Call the program
	catfiles.go. Here’s how you can execute this Go program:

	go run catfiles.go txtfile1 txtfile2 txtfile3
*/

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Use: go run catfiles.go <file1> <file2>...")
		return
	}

	files := os.Args[1:]
	for i := 0; i < len(files); i++ {
		go catFiles(files[i]) // Starts a new goroutine that calls catFiles function.
	}

	// We need this to wait for the goroutines to finish before the main function exits. In a real
	// application, you would use synchronization primitives like WaitGroup instead of sleep.
	time.Sleep(time.Second * 2)
}

func catFiles(filename string) {
	content, err := os.ReadFile(filename + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s: %s\n", filename, content)
}
