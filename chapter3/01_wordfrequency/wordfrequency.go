package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	// "time"
)

/*
	Modify our sequential-letter frequency program to produce a list of word fre-
	quencies rather than letter frequencies. You can use the same URLs for the RFC
	web pages as were used in listing 3.3. Once it’s finished, the program should
	output a list of words with the frequency with which each word appears in the
	web page. Here’s some sample output:

	$ go run wordfrequency.go
	the -> 5
	a -> 8
	car -> 1
	program -> 3

	What happens when you try to convert the sequential program into a concur-
	rent one, creating a goroutine for each page? We will fix these errors in the
	next chapter.
*/

var wordPattern = regexp.MustCompile(`[a-zA-Z]+`)

func main() {
	var frequency = make(map[string]int)
	for i := 1000; i <= 1001; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		countWords(url, frequency)
		// go countWords(url, frequency)
	}

	// time.Sleep(10 * time.Second) // Wait for all goroutines to finish
	for word, count := range frequency {
		fmt.Printf("%s -> %d\n", word, count)
	}

	// When I make this program concurrent, the output is not deterministic.
	// The goroutines are changing the same resource (the frequency map) concurrently,
	// so a correct output is not guaranteed. The number per word is different than
	// the sequential version.
}

func countWords(url string, frequency map[string]int) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server returning error status code: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	words := wordPattern.FindAllString(strings.ToLower(string(body)), -1)
	for _, word := range words {
		frequency[word] += 1
	}
	fmt.Println("Completed:", url)
}
