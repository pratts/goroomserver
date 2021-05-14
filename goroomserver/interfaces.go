package goroomserver

type Extension interface {
	init(map[string]interface{})
}

type EventHandler interface {
	handleEvent(map[string]interface{})
}
