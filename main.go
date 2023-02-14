package main

import (
	"flag"
	"fmt"
	"oxide/server"
)

func main() {
	debug := flag.Bool("debug", false, "Enable debug mode")
	tickrate := flag.Int("tickrate", 4, "Game ticks per second")
	memory := flag.Int("memory", 2, "Memory in MB") // On every tick, we need to check if we have enough memory to continue
	flag.Parse()

	fmt.Printf("Debug mode: %v\nTickrate: %v\nMemory Allocated: %v\n", *debug, *tickrate, *memory)
	server.Go(*debug, *tickrate, *memory)
}
