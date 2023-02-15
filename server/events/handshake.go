package server_events

import (
	"fmt"
	"oxide/events"
)

func InitializeHandshake() {
	events.Register("handshake", func() {
		fmt.Println("Handshake event called.")
	})
}
