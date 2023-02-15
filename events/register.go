package events

func Register(identifier string, callback func()) {
	pool.Events = append(pool.Events, Event{identifier, callback})
}
