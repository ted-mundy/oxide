package server

import (
	"fmt"
	"net"
	"os"
)

func listen(port int, debugMode bool) {
	udp_address, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%v", port))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	udp_connection, err := net.ListenUDP("udp", udp_address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer udp_connection.Close()

	// Initialise events here
	// ...
	// Finish initialising events here

	buffer := make([]byte, 1024)

	listen := true
	for listen {
		n, addr, err := udp_connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if debugMode {
			fmt.Printf("Received %v bytes from %v - %v\n", n, addr, string(buffer[:n]))
			udp_connection.WriteToUDP([]byte("Hello World"), addr)
		}

		// Simple event. To be changed into our event system later on.
		if string(buffer[:n]) == "handshake" && !debugMode {
			fmt.Printf("Received Handshake request from %v\n", addr)
		}
	}
}
