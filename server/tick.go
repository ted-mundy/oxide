package server

import (
	"fmt"
	"time"
)

func loop(tickrate int, memory int) {
	interval := 1000 / tickrate // Interval in milliseconds. How many milliseconds between each tick
	run := true
	for run {
		// Do tick stuff here
		fmt.Println("I'm sending out updates to all clients!")

		// Sleep for the interval
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}

	wg.Done()
}
