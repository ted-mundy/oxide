package packet

/**
 * Parse a packet from a byte array
 * @param data The byte array to parse
 * @return The parsed packet, or nil if the packet is invalid
 */
func Parse(data []byte) (packet Packet, valid bool) {
	// We'l start off by finding the start of header (0x01) and the start of the payload (0x02)
	startHeaderpos := -1
	startPayloadpos := -1
	for i := 0; i < len(data); i++ {
		if data[i] == 0x01 && startHeaderpos == -1 {
			startHeaderpos = i
			continue
		}
		if data[i] == 0x02 && startPayloadpos == -1 {
			startPayloadpos = i
			break
		}
	}

	// If we didn't find the start of the header or the start of the payload, the packet is invalid
	if startHeaderpos == -1 || startPayloadpos == -1 {
		return Packet{}, false
	}

	pkt := Packet{}

	pkt.EventType = string(data[startHeaderpos+1 : startPayloadpos])

	return pkt, true
}
