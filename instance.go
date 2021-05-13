package goroomserver

type Extension interface {
	init(map[string]interface{})
	addEventHandler(eventName string, e *Event)
	removeEventHandler(eventName string)
	sendResponseToUser(string, interface{})
	sendResponseToUserList([]string, interface{})
	sendResponseToAll(interface{})
}
