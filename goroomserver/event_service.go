package goroomserver

import "fmt"

type EventService struct {
	mainService *MainService
}

type Payload struct {
	AppName       string      `json:"appName"`
	RoomName      string      `json:"roomName"`
	EventType     int         `json:"eventType"`
	Payload       interface{} `json:"payload"`
	RemoteAddress string
}

func (e *EventService) getEvent(code int) string {
	return EventText(code)
}

func (e *EventService) handleEvent(payload Payload) {
	fmt.Println("Received:", payload)
	// if payload.appName == "" {

	// }
	// appName, ok := payload.appName
	// if ok == false {

	// }
	// appService := e.mainService.getAppService(appName)

	// appName := payload["appName"]
	// roomName := payload["roomName"]
	// eventType := payload["eventType"]
	// remoteAddr := payload["remoteAddr"]
	// data := payload["data"]
}
