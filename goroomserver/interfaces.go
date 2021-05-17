package goroomserver

type Extension interface {
	init(map[string]interface{})
}

type EventHandler interface {
	handleEvent(Event) Response
}

type Response struct {
	data map[string]interface{}
	err  error
}

type Event struct {
	payload map[string]interface{}
	refRoom Room
	refApp  AppService
	user    User
}
