package events

func Call(identifier string) {
	for _, event := range pool.Events {
		if event.Identifier == identifier {
			event.Callback()
			break
		}
	}
}
