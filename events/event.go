package events

type Event interface {
	getName() string
	handleEvent(map[string]interface{})
}
