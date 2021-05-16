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
	RefRoom       Room
	RefApp        AppService
	connection    Connection
}

func (e *EventService) getEvent(code int) string {
	return EventText(code)
}

func (e *EventService) handleEvent(payload Payload) {
	fmt.Println("Received:", payload.EventType, " Payload:", payload)
	eventType := payload.EventType
	switch eventType {
	case CONNECTION:
		e.mainService.connectionService.addConnectionInstance(payload.connection)
	case DISCONNECTION:
		e.mainService.connectionService.removeConnection(payload.RemoteAddress)
	}
	if payload.AppName == "" {
		return
	}

	payload.RefApp = e.mainService.getAppService(payload.AppName)
	if payload.RoomName != "" {
		payload.RefRoom = payload.RefApp.roomService.GetRoomByName(payload.RoomName)
		return
	}
}
