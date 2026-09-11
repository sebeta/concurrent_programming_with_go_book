package main

/*
	In listings 6.14 and 6.15 and again in 6.16 and 6.17, we implemented single- and
	multi-threaded matrix multiplication programs. Can you measure the time it
	takes to compute the multiplication for large matrices of size 1000 × 1000 or
	larger? For the time measurement to be accurate, you should remove the
	Println() calls because large matrices will take a long time to be printed on the
	console. You might notice a difference only if your system has multiple cores.
*/

import (
	"fmt"
	"math/rand"
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
	fmt.Println("Sequential:", sequentialElapsed)
}

// This is the kind of response I received after running the program:
// Concurrent: 67.783667ms
// Sequential: 426.7165ms
// The concurrent version is much faster than the sequential version.

func generateRandMatrix(matrix *[matrixSize][matrixSize]int) {
	for row := range matrixSize {
		for col := range matrixSize {
			matrix[row][col] = rand.Intn(10) - 5
		}
	}
}

func concurrentMatrixMultiply(
	matrixA, matrixB, result *[matrixSize][matrixSize]int,
) {
	barrier := NewBarrier(matrixSize + 1)
	for row := range matrixSize {
		go rowMultiply(matrixA, matrixB, result, row, barrier)
	}

	// Releases the barrier so the goroutines can start their computations
	barrier.Wait()
	// Waits for the goroutines to finish their computations
	barrier.Wait()
}

func rowMultiply(matrixA, matrixB, result *[matrixSize][matrixSize]int,
	row int, barrier *Barrier) {
	for {
		barrier.Wait()
		for col := range matrixSize {
			sum := 0
			for i := range matrixSize {
				sum += matrixA[row][i] * matrixB[i][col]
			}
			result[row][col] = sum
		}
		barrier.Wait()
	}
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
