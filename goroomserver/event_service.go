package goroomserver

import "fmt"

type EventService struct {
	mainService *MainService
}

type Payload struct {
	AppName       string                 `json:"appName"`
	RoomName      string                 `json:"roomName"`
	EventType     int                    `json:"eventType"`
	Payload       map[string]interface{} `json:"payload"`
	RemoteAddress string
}

func (e *EventService) getEvent(code int) string {
	return EventText(code)
}

func (e *EventService) handleEvent(payload Payload) {
	fmt.Println("Received:", payload)
	if payload.AppName == "" {
		return
	}
	app := e.mainService.getAppService(payload.AppName)
	eventId := payload.EventType
	if payload.RoomName == "" {
		app.eventHandler[eventId].handleEvent(payload.Payload)
		return
	}

	room := app.roomService.GetRoomByName(payload.RoomName)
	room.eventHandler[eventId].handleEvent(payload.Payload)
}
