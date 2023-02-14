package main

import (
	"flag"
	"fmt"
)

func main() {
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	if *debug {
		fmt.Println("\033[1;44m[INFO]\033[0;34m Debug Mode Enabled\033[0m")
	}
}
