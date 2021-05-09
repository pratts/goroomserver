package interfaces

type Extension interface {
	init()
	addEventHandler(eventName string, e *Event)
	removeEventHandler(eventName string)
}
