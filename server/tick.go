package server

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func loop(tickrate int, memory int) {
	interval := 1000 / tickrate // Interval in milliseconds. How many milliseconds between each tick
	run := true
	for run {
		// Do tick stuff here
		fmt.Println("I'm sending out updates to all clients!")

		// Get memory usage
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		// Convert to MB
		usedMemory := m.Alloc / 1024 / 1024
		// If we have gone over the memory limit, we need to stop the server using os.exit
		if usedMemory > uint64(memory) {
			fmt.Printf("Memory limit exceeded! Used %vMB, limit %vMB\n", usedMemory, memory)
			os.Exit(1)
		}
		// Sleep for the interval
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
