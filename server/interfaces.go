package server

type Extension interface {
	init(Event)
}

type EventHandler interface {
	handleEvent(Event) (Response, error)
}

type Response struct {
	Data  map[string]interface{} `json:"data"`
	Code  int                    `json:"code"`
	Error ServerError            `json:"error"`
}

type ServerError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
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
