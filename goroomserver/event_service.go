package goroomserver

import "fmt"

type EventService struct {
	mainService *MainService
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

	payload.Connection, _ = e.mainService.connectionService.GetConnectionByIp(payload.RemoteAddress)
	response := Response{}

	if payload.AppName == "" {
		response.error = ServerError{code: APP_NAME_INVALID, message: ErrorMessages[APP_NAME_INVALID]}
		e.pushMessage(payload, response)
		return
	}

	payload.RefApp, _ = e.mainService.GetAppService(payload.AppName)
	if payload.RoomName != "" {
		payload.RefRoom, _ = payload.RefApp.GetRoomByName(payload.RoomName)
	}

	event := Event{payload: payload.Payload, refRoom: payload.RefRoom, refApp: payload.RefApp}

	user := getValidUser(payload)
	if (User{}.name) == user.name {
		response.error = ServerError{code: USER_NOT_EXISTS, message: ErrorMessages[USER_NOT_EXISTS]}
		e.pushMessage(payload, response)
		return
	}
	payload.RefUser = user
	event.user = user

	switch payload.EventType {
	case DISCONNECTION:
		e.handleDisconnection(payload, event)
		break
	case LOGIN:
		e.handleLogin(payload, event)
		break
	case LOGOUT:
		e.handleLogout(payload, event)
		break
	case JOIN_ROOM:
		e.handleJoinRoom(payload, event)
		break
	case LEAVE_ROOM:
		e.handleLeaveRoom(payload, event)
		break
	case MESSAGE:
		e.handleMessage(payload, event)
		break
	default:
		response.error = ServerError{code: INVALID_EVENT, message: ErrorMessages[INVALID_EVENT]}
		e.pushMessage(payload, response)
	}
}

func (e *EventService) handleConnection(payload Payload) {
	e.mainService.connectionService.addConnectionInstance(payload.Connection)
}

func (e *EventService) handleDisconnection(payload Payload, event Event) {
	user := payload.RefUser
	disconnectionHandler, ok := payload.RefApp.eventHandler[DISCONNECTION]
	if ok == true {
		disconnectionHandler.handleEvent(event)
	}
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
		_, error := evtHandler.handleEvent(event)
		if error != nil {
			return
		}
	}
	_, ok = payload.RefApp.userService.userMap[payload.RemoteAddress]
	if ok == true {
		//handle login duplicate and send event
		return
	}
	payload.RefApp.userService.CreateAndAddUser(payload.Payload["username"].(string), payload.Connection)
}

func (e *EventService) handleLogout(payload Payload, event Event) {
	evtHandler, ok := payload.RefApp.eventHandler[LOGOUT]
	if ok != false {
		response, error := evtHandler.handleEvent(event)
		if error != nil {
			//handle logout error and send event
			return
		}
	}

	e.handleDisconnection(payload, event)
}

func (e *EventService) handleJoinRoom(payload Payload, event Event) {
	roomName := payload.Payload["roomName"].(string)
	if roomName == "" {
		//handle case when roomname is blank
		return
	}
	room, ok := payload.RefApp.GetRoomByName(roomName)
	if ok == false {
		//handle case when room does not exist
		return
	}
	user := payload.RefUser
	_, userExists := room.GetUserByName(user.name)
	if userExists == true {
		//handle case when user already is in room
		return
	}

	handler, evtExists := room.eventHandler[JOIN_ROOM]
	if evtExists == true {
		handler.handleEvent(event)
	}

	room.addUser(user)
}

func (e *EventService) handleLeaveRoom(payload Payload, event Event) {
	roomName := payload.Payload["roomName"].(string)
	if roomName == "" {
		//handle case when roomname is blank
		return
	}
	room, ok := payload.RefApp.GetRoomByName(roomName)
	if ok == false {
		//handle case when room does not exist
		return
	}
	user := payload.RefUser
	handler, evtExists := room.eventHandler[LEAVE_ROOM]

	_, userExists := room.GetUserByName(user.name)
	if userExists == false {
		//handle case when user already is not in room
		return
	}

	if evtExists == true {
		handler.handleEvent(event)
	}
	room.removeUser(user)
}

func (e *EventService) handleMessage(payload Payload, event Event) {
	if payload.RoomName != "" {
		evtHandler, ok := payload.RefRoom.eventHandler[MESSAGE]
		if ok != false {
			evtHandler.handleEvent(event)
			return
		}
	}
	evtHandler, ok := payload.RefApp.eventHandler[MESSAGE]
	if ok != false {
		evtHandler.handleEvent(event)
	}
}

func (e *EventService) pushMessage(payload Payload, response Response) {
	if payload.RemoteAddress != "" {
		connection, ok := e.mainService.connectionService.GetConnectionByIp(payload.RemoteAddress)
		if ok == false {
			return
		}
		connection.WriteToSocket(response)
	} else if payload.Connection != (Connection{}) {
		payload.Connection.WriteToSocket(response)
	}
}

func getValidUser(payload Payload) User {
	user, isValidUser := payload.RefApp.userService.GetUserForConnection(payload.RemoteAddress)
	if isValidUser == true {
		return user
	}
	return User{}
}
