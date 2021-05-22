package goroomserver

type Extension interface {
	init(map[string]interface{})
}

type EventHandler interface {
	handleEvent(Event) (Response, error)
}

type Response struct {
	data  map[string]interface{}
	error ServerError
}

type ServerError struct {
	code    int
	message string
	err     error
}

type Event struct {
	payload map[string]interface{}
	refRoom Room
	refApp  AppService
	user    User
}

type Payload struct {
	AppName       string                 `json:"appName"`
	RoomName      string                 `json:"roomName"`
	EventType     int                    `json:"eventType"`
	Payload       map[string]interface{} `json:"payload"`
	RemoteAddress string
	RefRoom       Room
	RefApp        AppService
	Connection    Connection
	RefUser       User
}
