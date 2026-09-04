package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/*
	Expand the program you wrote in the first exercise so that instead of printing the contents of
	the text files, it searches for a string match. The string to search for is the first argument
	on the command line. When you spawn a new goroutine, instead of printing the file’s contents, it
	should read the file and search for a match. If the goroutine finds a match, it should output a
	message saying that	the filename contains a match. Call the program grepfiles.go. Here’s how you
	can execute this Go program (“bubbles” is the search string in this example):

	go run grepfiles.go bubbles txtfile1 txtfile2 txtfile3
*/

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Use: go run grepfiles.go <search_tern> <file1> <file2> ...")
		return
	}

	searchTerm := os.Args[1]
	files := os.Args[2:]
	for i := 0; i < len(files); i++ {
		go grepFile(files[i], searchTerm) // Starts a new goroutine that calls grepFiles function.
	}

	// We need this to wait for the goroutines to finish before the main function exits. In a real
	// application, you would use synchronization primitives like WaitGroup instead of sleep.
	time.Sleep(time.Second * 2)
}

func grepFile(filename, searchTerm string) {
	content, err := os.ReadFile(filename + ".txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	if strings.Contains(string(content), searchTerm) {
		fmt.Printf("%s: ✓ '%s' found\n", filename, searchTerm)
	} else {
		fmt.Printf("%s: ✗ '%s' not found\n", filename, searchTerm)
	}
}
