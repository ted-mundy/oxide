package packet_test

import (
	"oxide/packet"
	"testing"
)

func TestParse(t *testing.T) {
	// Create a new packet.Packet

	// This packet contains the following information:
	// 1. The header (0x01)
	// 2. The event (handshake - ASCII)
	// 3. The payload (0x02)
	// 4. 'lorem=ipsum;foo=bar'
	pkt, valid := packet.Parse([]byte{
		0x01, 0x68, 0x61, 0x6E, 0x64, 0x73,
		0x68, 0x61, 0x6B, 0x65, 0x02, 0x6C,
		0x6F, 0x72, 0x65, 0x6D, 0x3D, 0x69,
		0x70, 0x73, 0x75, 0x6D, 0x3B, 0x66,
		0x6F, 0x6F, 0x3D, 0x62, 0x61, 0x72,
	})
	if !valid {
		t.Error("Packet is invalid")
	}
	if pkt.EventType != "handshake" {
		t.Error("Packet event type is not 'handshake'")
	}
}
