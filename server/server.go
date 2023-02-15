package server

import (
	"sync"
)

var (
	// Wait group
	wg sync.WaitGroup
)

func Go(debugMode bool, tickrate int, memory int) {
	wg.Add(1)
	go loop(tickrate, memory)
	go listen(27089, debugMode)
	wg.Wait()
}
