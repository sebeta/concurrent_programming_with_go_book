package main

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

	fmt.Println(findFactors(4033836233))
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
