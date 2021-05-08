package interfaces

type instance interface {
	init()
	addEventHandler(eventName string, e *Event)
	removeEventHandler(eventName string)
}
