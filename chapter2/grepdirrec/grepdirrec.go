package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/*
	Adapt the program in the third exercise to continue searching recursively in any subdirectories.
	If you give your search goroutine a file, it should search for a string match in that file, just
	like in the previous exercises. Otherwise, if you give it a directory, it should recursively
	spawn a new goroutine for each file or directory found inside. Call the program grepdirrec.go,
	and execute it by running this command:

	go run grepdirrec.go bubbles ../../commonfiles
*/

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Use: go run grepdirrec.go <search_term> <dir>")
		return
	}

	searchTerm := os.Args[1]
	dirPath := os.Args[2]

	grepFilesFromDir(dirPath, searchTerm)

	// We need this to wait for the goroutines to finish before the main function exits. In a real
	// application, you would use synchronization primitives like WaitGroup instead of sleep.
	time.Sleep(time.Second * 2)
}

func grepFilesFromDir(dirPath, searchTerm string) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, entry := range entries {
		var fullPath string
		if dirPath == "./" {
			fullPath = dirPath + entry.Name()
		} else {
			fullPath = dirPath + "/" + entry.Name()
		}

		if entry.IsDir() {
			go grepFilesFromDir(fullPath, searchTerm)
		} else if strings.HasSuffix(fullPath, ".txt") {
			go grepFile(fullPath, searchTerm)
		}
	}
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
