// Package ex2 provides generic pipeline stages, Take and TakeUntil, that
// forward messages from an input channel and signal upstream to stop via quit.
package ex2

/*
	In listing 9.18, we developed a take(n) goroutine. Extend the functionality of
	this goroutine to implement TakeUntil(f), where f is a function returning a
	Boolean. The goroutine needs to continue consuming and forwarding the mes-
	sages on its input channel while the return value of f is true. Using generics
	ensures that we can reuse the TakeUntil(f) function and plug it into many other
	pipelines. Here’s the function signature:

	func TakeUntil[K any](f func(K) bool, quit chan int, input <-chan K) <-chan K
*/

func TakeUntil[K any](f func(K) bool, quit chan int, input <-chan K) <-chan K {
	output := make(chan K)
	go func() {
		defer close(output)
		// keepGoing becomes false only when f rejects a REAL message.
		// moreData becomes false when input is closed (no more real messages).
		keepGoing, moreData := true, true
		var msg K
		for keepGoing && moreData {
			select {
			case msg, moreData = <-input:
				// Call f only if msg is real. On a closed channel msg is just
				// the zero value of K, and f's answer about it is meaningless.
				if moreData {
					if f(msg) {
						output <- msg
					} else {
						// Stop, without forwarding the message that failed f.
						keepGoing = false
					}
				}
			case <-quit:
				return
			}
		}
		// Tell the upstream generator to stop, but only if WE decided to stop
		// (f returned false). If input closed by itself, nobody needs a signal.
		if !keepGoing {
			close(quit)
		}
	}()

	return output

}

func Take[K any](quit chan int, n int, input <-chan K) <-chan K {
	output := make(chan K)
	go func() {
		defer close(output)
		moreData := true
		var msg K
		for n > 0 && moreData {
			select {
			case msg, moreData = <-input:
				if moreData {
					output <- msg
					n--
				}
			case <-quit:
				return
			}
		}
		if n == 0 {
			close(quit)
		}
	}()
	return output
}
