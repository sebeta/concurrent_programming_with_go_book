package main

/*
	Connect the components developed in exercises 1 to 4 together in a main()
	function using the following pseudocode:

	Create quit channel
	Drain(quitChannel,
		Print(quitChannel,
			TakeUntil({ s <= 1000000 }, quitChannel,
					GenerateSquares(quitChannel))))
	Wait on quit channel
*/

import (
	ex1 "cpwg/chapter9/01_generatesquares"
	ex2 "cpwg/chapter9/02_takeuntil"
	ex3 "cpwg/chapter9/03_print"
	ex4 "cpwg/chapter9/04_drain"
)

func main() {
	quit := make(chan int)
	ex4.Drain(
		quit,
		ex3.Print(
			quit,
			ex2.TakeUntil(
				func(s int) bool { return s <= 1000000 },
				quit,
				ex1.GenerateSquares(quit),
			),
		),
	)

	<-quit
}
