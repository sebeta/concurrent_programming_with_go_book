package main

/*
	Consider listing 8.17 containing the player() function. This function creates a
	goroutine simulating a player in a game moving along a two-dimensional plane.
	The goroutine returns the movements at random times by writing UP, DOWN, LEFT,
	or RIGHT on an output channel. Create a main() function that creates four player
	goroutines and outputs on the console all movements from the four players. The
	main() function should terminate only when there is one player left in the game.
	Here is an example of what the output should look like:

	Player 1: DOWN
	Player 0: LEFT
	Player 3: DOWN
	Player 2 left the game. Remaining players: 3
	Player 1: UP
	. . .
	Player 0: LEFT
	Player 3 left the game. Remaining players: 2
	Player 1: RIGHT
	. . .
	Player 1: RIGHT
	Player 0 left the game. Remaining players: 1
	Game finished
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	player0 := player()
	player1 := player()
	player2 := player()
	player3 := player()

	for i := 4; i > 1; {
		select {
		case mov0, ok := <-player0:
			if !ok {
				i--
				player0 = nil
				fmt.Println("Player 0 left the game. Remaining players:", i)
			} else {
				fmt.Println("Player 0:", mov0)
			}
		case mov1, ok := <-player1:
			if !ok {
				i--
				player1 = nil
				fmt.Println("Player 1 left the game. Remaining players:", i)
			} else {
				fmt.Println("Player 1:", mov1)
			}
		case mov2, ok := <-player2:
			if !ok {
				i--
				player2 = nil
				fmt.Println("Player 2 left the game. Remaining players:", i)
			} else {
				fmt.Println("Player 2:", mov2)
			}
		case mov3, ok := <-player3:
			if !ok {
				i--
				player3 = nil
				fmt.Println("Player 3 left the game. Remaining players:", i)
			} else {
				fmt.Println("Player 3:", mov3)
			}
		}
	}
	fmt.Println("Game finished")
}

func player() chan string {
	output := make(chan string)
	count := rand.Intn(100)
	move := []string{"UP", "DOWN", "LEFT", "RIGHT"}
	go func() {
		defer close(output)
		for i := 0; i < count; i++ {
			output <- move[rand.Intn(4)]
			d := time.Duration(rand.Intn(200))
			time.Sleep(d * time.Millisecond)
		}
	}()

	return output
}
