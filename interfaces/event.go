package interfaces

type Event interface {
	handleEvent(map[string]interface{})
}
