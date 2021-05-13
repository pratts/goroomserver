package goroomserver

type Event interface {
	handleEvent(map[string]interface{})
}
