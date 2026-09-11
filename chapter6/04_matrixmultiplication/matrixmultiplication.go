package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const matrixSize = 1000

func main() {
	var matrixA, matrixB, conResult, seqResult [matrixSize][matrixSize]int

	generateRandMatrix(&matrixA)
	generateRandMatrix(&matrixB)

	concurrentStart := time.Now()
	concurrentMatrixMultiply(&matrixA, &matrixB, &conResult)
	concurrentElapsed := time.Since(concurrentStart)

	sequentialStart := time.Now()
	sequentialMatrixMultiply(&matrixA, &matrixB, &seqResult)
	sequentialElapsed := time.Since(sequentialStart)

	fmt.Println("Concurrent:", concurrentElapsed)
	// fmt.Println("Concurrent result:")
	// printMatrix(&conResult)
	fmt.Println("Sequential:", sequentialElapsed)
	// fmt.Println("Sequential result:")
	// printMatrix(&seqResult)
}

// I found something interesting: for minor matrix sizes, the sequential
// implementation is faster than the concurrent one.
//
// For 3x3 matrixes:
// Concurrent: 62.458µs
// Sequential: 166ns
//
// For 1000x1000 matrixes:
// Concurrent: 66.403666ms
// Sequential: 427.6725ms

func generateRandMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := range matrixSize {
		for col := range matrixSize {
			matrix[row][col] = rand.Intn(10) - 5
		}
	}
}

func printMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := range matrixSize {
		for col := range matrixSize {
			fmt.Printf("%d ", matrix[row][col])
		}
		fmt.Println()
	}
}

func concurrentMatrixMultiply(
	matrixA, matrixB, result *[matrixSize][matrixSize]int,
) {
	wg := sync.WaitGroup{}
	for row := range matrixSize {
		wg.Add(1)
		go rowMultiply(matrixA, matrixB, result, row, &wg)
	}
	wg.Wait()
}

func rowMultiply(
	matrixA, matrixB, result *[matrixSize][matrixSize]int,
	row int,
	wg *sync.WaitGroup,
) {
	for col := range matrixSize {
		sum := 0
		for i := range matrixSize {
			sum += matrixA[row][i] * matrixB[i][col]
		}
		result[row][col] = sum
	}
	wg.Done()
}

func sequentialMatrixMultiply(matrixA, matrixB, result *[matrixSize][matrixSize]int) {
	for row := range matrixSize {
		for col := range matrixSize {
			sum := 0
			for i := range matrixSize {
				sum += matrixA[row][i] * matrixB[i][col]
			}
			result[row][col] = sum
		}
	}
}
