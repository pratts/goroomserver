package goroomserver

type EventHandler1 interface {
	handleEvent(map[string]interface{})
}
