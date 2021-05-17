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
	Connection    Connection
	RefUser       User
}

func (e *EventService) getEvent(code int) string {
	return EventText(code)
}

func (e *EventService) handleEvent(payload Payload) {
	fmt.Println("Received:", payload.EventType, " Payload:", payload)

	if payload.EventType == CONNECTION {
		e.handleConnection(payload)
		return
	}

	payload.Connection = e.mainService.connectionService.connectionMap[payload.RemoteAddress]

	if payload.AppName == "" {
		return
	}

	payload.RefApp = e.mainService.getAppService(payload.AppName)
	if payload.RoomName != "" {
		payload.RefRoom = payload.RefApp.roomService.GetRoomByName(payload.RoomName)
	}

	event := Event{payload: payload.Payload, refRoom: payload.RefRoom, refApp: payload.RefApp}

	switch payload.EventType {
	case DISCONNECTION:
		e.handleDisconnection(payload, event)
	case LOGIN:
		e.handleLogin(payload, event)
	case LOGOUT:
		e.handleLogout(payload, event)
	case JOIN_ROOM:
		e.handleJoinRoom(payload, event)
	case LEAVE_ROOM:
		e.handleLeaveRoom(payload, event)
	}
}

func (e *EventService) handleConnection(payload Payload) {
	e.mainService.connectionService.addConnectionInstance(payload.Connection)
}

func (e *EventService) handleDisconnection(payload Payload, event Event) {
	user := payload.RefApp.userService.connectionUserMap[payload.RemoteAddress]
	for _, room := range user.roomMap {
		evtHandler, ok := room.eventHandler[DISCONNECTION]
		if ok == false {
			continue
		}
		evtHandler.handleEvent(event)
	}
	payload.RefApp.userService.RemoveUser(user.name)
	e.mainService.connectionService.removeConnection(payload.RemoteAddress)
}

func (e *EventService) handleLogin(payload Payload, event Event) {
	evtHandler, ok := payload.RefApp.eventHandler[LOGIN]
	if ok != false {
		response := evtHandler.handleEvent(event)
		if response.err != nil {

		}
	}
	_, ok = payload.RefApp.userService.userMap[payload.RemoteAddress]
	if ok == true {
		//handle login duplicate and send event
		return
	}
	payload.RefApp.userService.CreateAndAddUser(payload.Payload["username"].(string), e.mainService.connectionService.connectionMap[payload.RemoteAddress])
}

func (e *EventService) handleLogout(payload Payload, event Event) {
	evtHandler, ok := payload.RefApp.eventHandler[LOGOUT]
	if ok != false {
		response := evtHandler.handleEvent(event)
		if response.err != nil {
			//handle logout error and send event
		}
	}

	e.handleDisconnection(payload, event)
}

func (e *EventService) handleJoinRoom(payload Payload, event Event) {

}

func (e *EventService) handleLeaveRoom(payload Payload, event Event) {

}
