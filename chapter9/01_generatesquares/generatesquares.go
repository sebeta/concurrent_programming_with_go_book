// Package ex1 provides a generator goroutine that streams the
// square numbers (1, 4, 9, 16, ...) on a channel until the caller
// signals it to stop through a quit channel.
package ex1

/*
	Write a generator goroutine similar to listing 9.2 that, instead of generating
	URL strings, generates an infinite stream of square numbers (1, 4, 9, 16, 25 . . .)
	on an output channel. Here is the signature:

	func GenerateSquares(quit <-chan int) <-chan int
*/

func GenerateSquares(quit <-chan int) <-chan int {
	squares := make(chan int)
	go func() {
		defer close(squares)
		for i := 1; ; i++ {
			select {
			case squares <- i * i:
			case <-quit:
				return
			}
		}
	}()

	return squares
}
