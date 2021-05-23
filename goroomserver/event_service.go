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
	if payload.Data == nil {
		payload.Data = make(map[string]interface{})
	}

	if payload.EventType == CONNECTION {
		e.handleConnection(payload)
		return
	}

	payload.Connection, _ = e.mainService.connectionService.GetConnectionByIp(payload.RemoteAddress)

	if payload.AppName == "" {
		response := Response{code: SERVER_ERROR, error: ServerError{code: APP_NAME_INVALID, message: ErrorMessages[APP_NAME_INVALID]}}
		e.pushMessage(payload, response)
		return
	}

	payload.RefApp, _ = e.mainService.GetAppService(payload.AppName)
	if payload.RoomName != "" {
		payload.RefRoom, _ = payload.RefApp.GetRoomByName(payload.RoomName)
	}

	event := Event{payload: payload.Data, room: payload.RefRoom, app: payload.RefApp}

	user := getValidUser(payload)
	if (User{}.name) == user.name {
		response := Response{code: SERVER_ERROR, error: ServerError{code: USER_NOT_EXISTS, message: ErrorMessages[USER_NOT_EXISTS]}}
		e.pushMessage(payload, response)
		return
	}
	payload.RefUser = user
	event.user = user

	response := Response{code: SUCCESS}

	switch payload.EventType {
	case DISCONNECTION:
		e.handleDisconnection(payload, event, &response)
		break
	case LOGIN:
		e.handleLogin(payload, event, &response)
		break
	case LOGOUT:
		e.handleLogout(payload, event, &response)
		break
	case JOIN_ROOM:
		e.handleJoinRoom(payload, event, &response)
		break
	case LEAVE_ROOM:
		e.handleLeaveRoom(payload, event, &response)
		break
	case MESSAGE:
		e.handleMessage(payload, event, &response)
		break
	default:
		response.code = SERVER_ERROR
		response.error = ServerError{code: INVALID_EVENT, message: ErrorMessages[INVALID_EVENT]}
		e.pushMessage(payload, response)
	}
}

func (e *EventService) handleConnection(payload Payload) {
	e.mainService.connectionService.addConnectionInstance(payload.Connection)
}

func (e *EventService) handleDisconnection(payload Payload, event Event, response *Response) {
	user := payload.RefUser
	disconnectionHandler, ok := payload.RefApp.eventHandler[DISCONNECTION]
	if ok == true {
		_, error := disconnectionHandler.handleEvent(event)
		response.error = ServerError{code: SERVER_ERROR, message: error.Error()}
		return
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

func (e *EventService) handleLogin(payload Payload, event Event, response *Response) {
	evtHandler, ok := payload.RefApp.eventHandler[LOGIN]
	if ok != false {
		res, error := evtHandler.handleEvent(event)
		if error != nil {
			response = &res
			return
		}
	}
	_, ok = payload.RefApp.userService.userMap[payload.RemoteAddress]
	if ok == true {
		response.code = SERVER_ERROR
		response.error = ServerError{code: USER_ALREADY_IN_ROOM, message: ErrorMessages[USER_ALREADY_LOGGED_IN]}
		return
	}
	payload.RefApp.userService.CreateAndAddUser(payload.Data["username"].(string), payload.Connection)
}

func (e *EventService) handleLogout(payload Payload, event Event, response *Response) {
	evtHandler, ok := payload.RefApp.eventHandler[LOGOUT]
	if ok != false {
		res, error := evtHandler.handleEvent(event)
		if error != nil {
			response = &res
			//handle logout error and send event
			return
		}
	}

	e.handleDisconnection(payload, event, response)
}

func (e *EventService) handleJoinRoom(payload Payload, event Event, response *Response) {
	roomName := payload.Data["roomName"].(string)
	if roomName == "" {
		response.code = SERVER_ERROR
		response.error = ServerError{code: ROOM_NAME_INVALID, message: ErrorMessages[ROOM_NAME_INVALID]}
		return
	}
	room, ok := payload.RefApp.GetRoomByName(roomName)
	if ok == false {
		//handle case when room does not exist
		response.code = SERVER_ERROR
		response.error = ServerError{code: ROOM_NOT_EXISTS, message: ErrorMessages[ROOM_NOT_EXISTS]}
		return
	}
	user := payload.RefUser
	_, userExists := room.GetUserByName(user.name)
	if userExists == true {
		//handle case when user already is in room
		response.code = SERVER_ERROR
		response.error = ServerError{code: USER_ALREADY_IN_ROOM, message: ErrorMessages[USER_ALREADY_IN_ROOM]}
		return
	}

	handler, evtExists := room.eventHandler[JOIN_ROOM]
	if evtExists == true {
		handler.handleEvent(event)
	}

	room.addUser(user)
}

func (e *EventService) handleLeaveRoom(payload Payload, event Event, response *Response) {
	roomName := payload.Data["roomName"].(string)
	if roomName == "" {
		//handle case when roomname is blank
		response.code = SERVER_ERROR
		response.error = ServerError{code: ROOM_NAME_INVALID, message: ErrorMessages[ROOM_NAME_INVALID]}
		return
	}
	room, ok := payload.RefApp.GetRoomByName(roomName)
	if ok == false {
		//handle case when room does not exist
		response.code = SERVER_ERROR
		response.error = ServerError{code: ROOM_NOT_EXISTS, message: ErrorMessages[ROOM_NOT_EXISTS]}
		return
	}
	user := payload.RefUser
	handler, evtExists := room.eventHandler[LEAVE_ROOM]

	_, userExists := room.GetUserByName(user.name)
	if userExists == false {
		//handle case when user already is not in room
		response.code = SERVER_ERROR
		response.error = ServerError{code: USER_NOT_IN_ROOM, message: ErrorMessages[USER_NOT_IN_ROOM]}
		return
	}

	if evtExists == true {
		handler.handleEvent(event)
	}
	room.removeUser(user)
}

func (e *EventService) handleMessage(payload Payload, event Event, response *Response) {
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
	response.data = payload.Data
	response.data["roomName"] = payload.RoomName
	response.data["appName"] = payload.AppName
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
