package server

type Extension interface {
	init(Event)
}

type EventHandler interface {
	handleEvent(Event) (Response, error)
}

type Response struct {
	data  map[string]interface{}
	code  int
	error ServerError
}

type ServerError struct {
	code    int
	message string
}

type Event struct {
	payload map[string]interface{}
	room    Room
	app     AppService
	user    User
}

type Payload struct {
	AppName       string                 `json:"appName"`
	RoomName      string                 `json:"roomName"`
	EventType     int                    `json:"eventType"`
	Data          map[string]interface{} `json:"data"`
	RemoteAddress string
	RefRoom       Room
	RefApp        AppService
	Connection    Connection
	RefUser       User
}
