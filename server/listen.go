package server

import (
	"fmt"
	"net"
	"os"
	"oxide/events"
	"oxide/packet"
	server_events "oxide/server/events"
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
	server_events.Initialize()
	// Finish initialising events here

	buffer := make([]byte, 1024)

	listen := true
	for listen {
		n, addr, err := udp_connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		pkt, valid := packet.Parse(buffer[:n])
		if !valid {
			if debugMode {
				fmt.Printf("Received invalid packet from %v - %v (%v)\n", addr, buffer[:n], string(buffer[:n]))
			}
			continue
		}

		if debugMode {
			fmt.Printf("Received %v bytes from %v - %v\n", n, addr, string(buffer[:n]))
			udp_connection.WriteToUDP([]byte("Hello World"), addr)
		}

		// Handle events here. For now we'll just print the event type
		events.Call(pkt.EventType)
	}
}
