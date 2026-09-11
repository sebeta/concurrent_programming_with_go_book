package main

/*
	In listings 6.5 and 6.6, we developed a recursive concurrent file search. When a
	goroutine finds a file match, it outputs it on the console. Can you change the
	implementation of this file search so that it prints all the file matches, sorted
	into alphabetical order, after the search completes? Hint: try collecting the
	results in a shared data structure instead of printing them on the console from
	the goroutine.

	Example of use:
	$ go run ./chapter6/01_sortedfilesearch/sortedfilesearch.go ./ countdown
*/

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	paths := make([]string, 0)
	wg.Add(1)
	go fileSearch(os.Args[1], os.Args[2], &paths, &wg, &mu)
	wg.Wait()

	if len(paths) == 0 {
		fmt.Println("no matches found for ", os.Args[2])
		return
	}

	mu.Lock()
	slices.Sort(paths)
	mu.Unlock()

	for _, path := range paths {
		fmt.Println(path)
	}
}

func fileSearch(
	dir string,
	filename string,
	paths *[]string,
	wg *sync.WaitGroup,
	mu *sync.Mutex,
) {
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		fPath := filepath.Join(dir, file.Name())
		if strings.Contains(file.Name(), filename) {
			mu.Lock()
			*paths = append(*paths, fPath)
			mu.Unlock()
		}
		if file.IsDir() {
			wg.Add(1)
			go fileSearch(fPath, filename, paths, wg, mu)
		}
	}
	wg.Done()
}
