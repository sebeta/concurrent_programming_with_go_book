package main

/*
	Change the game-sync listings 5.8 and 5.9 so that, still using condition variables,
	the players wait for a fixed number of seconds. If the players haven’t all joined
	within this time, the goroutines should stop waiting and let the game start with-
	out all the players. Hint: try using another goroutine with an expiry timer.
*/

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	var cancel bool
	go timeout(cond, &cancel)
	playersInGame := 4
	for playerID := range 4 {
		go playerHandler(cond, &playersInGame, playerID, &cancel)
		time.Sleep(1 * time.Second)
	}
}

func playerHandler(
	cond *sync.Cond,
	playersRemaining *int,
	playerID int,
	cancel *bool,
) {
	cond.L.Lock()
	fmt.Println(playerID, ": Connected")
	*playersRemaining--
	if *playersRemaining == 0 {
		cond.Broadcast()
	}
	for *playersRemaining > 0 && !*cancel {
		fmt.Println(playerID, ": Waiting for more players")
		cond.Wait()
	}
	cond.L.Unlock()
	if *cancel {
		fmt.Println(playerID, ": Time out, game cancelled")
		return
	}
	fmt.Println("All players connected. Ready player", playerID)
	// Game started
}

func timeout(cond *sync.Cond, cancel *bool) {
	time.Sleep(10 * time.Second)
	cond.L.Lock()
	*cancel = true
	cond.Broadcast()
	cond.L.Unlock()
}
