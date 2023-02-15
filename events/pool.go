package events

var (
	// Pool is the global event pool.
	pool = EventPool{}
)

type EventPool struct {
	Events []Event
}

type Event struct {
	Identifier string
	Callback   func()
}
