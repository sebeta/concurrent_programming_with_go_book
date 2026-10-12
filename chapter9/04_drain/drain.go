// Package ex4 provides a generic pipeline sink, Drain, that reads and discards
// every message from its input channel until it closes or quit is signaled.
package ex4

/*
	Write a goroutine that drains the contents of its input channel without doing
	anything with them. The goroutine simply reads a message and throws it away:

	func Drain[T any](quit <-chan int, input <-chan T)
*/

func Drain[T any](quit <-chan int, input <-chan T) {
	go func() {
		moreData := true
		for moreData {
			select {
			case _, moreData = <-input:
			case <-quit:
				return
			}
		}
	}()
}
