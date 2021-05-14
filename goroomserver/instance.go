package goroomserver

type Extension1 interface {
	init(map[string]interface{})
	addEventHandler(eventName string, e *EventHandler)
	removeEventHandler(eventName string)
	sendResponseToUser(string, interface{})
	sendResponseToUserList([]string, interface{})
	sendResponseToAll(interface{})
}
