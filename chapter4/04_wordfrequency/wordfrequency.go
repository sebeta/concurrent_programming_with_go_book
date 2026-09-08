package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
	// "time"
)

/*
	In the previous chapter, in exercise 3.1, we developed a program to output the
	frequencies of words from downloaded web pages. If you used a shared mem-
	ory map to store the word frequencies, access to the shared map would need to
	be protected. Can you use a mutex to guarantee exclusive access to the map?
*/

var wordPattern = regexp.MustCompile(`[a-zA-Z]+`)

func main() {
	var frequency = make(map[string]int)
	mutex := sync.Mutex{}
	for i := 1000; i <= 1003; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countWords(url, frequency, &mutex)
	}

	time.Sleep(5 * time.Second) // Wait for all goroutines to finish
	for word, count := range frequency {
		fmt.Printf("%s -> %d\n", word, count)
	}
}

func countWords(url string, frequency map[string]int, mutex *sync.Mutex) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server returning error status code: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	words := wordPattern.FindAllString(strings.ToLower(string(body)), -1)
	mutex.Lock()
	for _, word := range words {
		frequency[word] += 1
	}
	mutex.Unlock()
	fmt.Println("Completed:", url)
}
