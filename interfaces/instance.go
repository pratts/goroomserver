package interfaces

type Instance interface {
	init()
	addEventHandler(eventName string, e *Event)
	removeEventHandler(eventName string)
}
