package main

/*
	In listing 7.13, we use a child goroutine to calculate the factors of one number
	and the main() goroutine to work out the factors of the other. Modify this listing
	so that, using multiple goroutines, we collect the factors of 10 random numbers.
*/

import (
	"fmt"
	"math/rand"
)

func main() {
	resultCh := make(chan []int, 10)
	for range 10 {
		go func() {
			resultCh <- findFactors(rand.Intn(1e6))
		}()
	}

	// We need to read 10 times because 10 goroutines each send one result;
	// the buffer just lets all 10 sends complete without blocking on a reader.
	for range 10 {
		fmt.Println(<-resultCh)
	}
}

func findFactors(number int) []int {
	result := make([]int, 0)
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			result = append(result, i)
		}
	}
	return result
}
