package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/*
	Change the program you wrote in the second exercise so that instead of passing a list of text
	filenames, you pass a directory path. The program will look inside this directory and list the
	files. For each file, you can spawn a goroutine that will search for a string match (the same as
	before). Call the program grepdir.go. Here’s how you can execute this Go program:

	go run grepdir.go bubbles ../../commonfiles
*/

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Use: go run grepdirrec.go <search_term> <dir>")
		return
	}

	searchTerm := os.Args[1]
	dirPath := os.Args[2]

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, entry := range entries {
		filename := entry.Name()
		if strings.HasSuffix(filename, ".txt") {
			go grepFile(filename, searchTerm) // Starts a new goroutine that calls grepFiles function.
		}
	}

	// We need this to wait for the goroutines to finish before the main function exits. In a real
	// application, you would use synchronization primitives like WaitGroup instead of sleep.
	time.Sleep(time.Second * 2)
}

func grepFile(filename, searchTerm string) {
	content, err := os.ReadFile(filename)
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
