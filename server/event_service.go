package server

import "fmt"

type EventService struct {
	mainService *MainService
}

func (e *EventService) getEvent(code int) string {
	return EventText(code)
}

func (e *EventService) getResponse(eventType int, code int, payload map[string]interface{}, errorCode int, message string) Response {
	res := Response{EventType: eventType, Code: code}
	if payload != nil {
		res.Data = payload
	}

	if code != SUCCESS {
		res.Error = ServerError{Code: errorCode, Message: message}
	}
	return res
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
		response := e.getResponse(payload.EventType, SERVER_ERROR, nil, APP_NAME_INVALID, ErrorMessages[APP_NAME_INVALID])
		e.pushMessage(payload, &response)
		return
	}

	app, ok := e.mainService.GetAppService(payload.AppName)
	if ok == false {
		response := e.getResponse(payload.EventType, SERVER_ERROR, nil, APP_NAME_INVALID, ErrorMessages[APP_NAME_INVALID])
		e.pushMessage(payload, &response)
		return
	}

	payload.RefApp = app
	if payload.RoomName != "" {
		payload.RefRoom, _ = payload.RefApp.GetRoomByName(payload.RoomName)
	}

	event := Event{payload: payload.Data, room: payload.RefRoom, app: payload.RefApp}

	user := getValidUser(payload)
	if (User{}.name) == user.name && payload.EventType != LOGIN {
		response := e.getResponse(payload.EventType, SERVER_ERROR, nil, USER_NOT_EXISTS, ErrorMessages[USER_NOT_EXISTS])
		e.pushMessage(payload, &response)
		return
	}
	payload.RefUser = user
	event.user = user

	response := Response{EventType: payload.EventType, Code: SUCCESS}

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
		response.Code = SERVER_ERROR
		response.Error = ServerError{Code: INVALID_EVENT, Message: ErrorMessages[INVALID_EVENT]}
	}

	fmt.Println("Finally sending response")
	if response.Code != SUCCESS {
		e.pushMessage(payload, &response)
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
		if error != nil {
			response.Code = SERVER_ERROR
			response.Error = ServerError{Code: DISCONNECTION_ERROR, Message: error.Error()}
			return
		}
	}
	for _, room := range user.roomMap {
		evtHandler, ok := room.eventHandler[DISCONNECTION]
		if ok == false {
			continue
		}
		evtHandler.handleEvent(event)
	}
	payload.RefApp.Logout(payload)
	e.mainService.connectionService.removeConnection(payload.RemoteAddress)
}

func (e *EventService) handleLogin(payload Payload, event Event, response *Response) {
	evtHandler, ok := payload.RefApp.eventHandler[LOGIN]
	if ok != false {
		res, error := evtHandler.handleEvent(event)
		if error != nil {
			response = &res
			if response.Code == SUCCESS {
				response.Code = SERVER_ERROR
			}
			response.Error = ServerError{Code: USER_LOGIN_ERROR, Message: error.Error()}
			return
		}
	}
	_, ok = payload.RefApp.userService.userMap[payload.RemoteAddress]
	if ok == true {
		response.Code = SERVER_ERROR
		response.Error = ServerError{Code: USER_ALREADY_LOGGED_IN, Message: ErrorMessages[USER_ALREADY_LOGGED_IN]}
		return
	}
	payload.RefApp.Login(payload)
}

func (e *EventService) handleLogout(payload Payload, event Event, response *Response) {
	evtHandler, ok := payload.RefApp.eventHandler[LOGOUT]
	if ok != false {
		res, error := evtHandler.handleEvent(event)
		if error != nil {
			response = &res
			if response.Code == SUCCESS {
				response.Code = SERVER_ERROR
			}
			response.Error = ServerError{Code: USER_LOGOUT_ERROR, Message: ErrorMessages[USER_LOGOUT_ERROR]}
			return
		}
	}

	e.handleDisconnection(payload, event, response)
}

func (e *EventService) handleJoinRoom(payload Payload, event Event, response *Response) {
	errorCode := payload.RefApp.JoinUserRoom(payload.RefUser, payload.Data["roomName"].(string), payload.Data)
	if errorCode != SUCCESS {
		response.Code = SERVER_ERROR
		response.Error = ServerError{Code: errorCode, Message: ErrorMessages[errorCode]}
		return
	}
}

func (e *EventService) handleLeaveRoom(payload Payload, event Event, response *Response) {
	errorCode := payload.RefApp.LeaveUserRoom(payload.RefUser, payload.Data["roomName"].(string), payload.Data)
	if errorCode != SUCCESS {
		response.Code = SERVER_ERROR
		response.Error = ServerError{Code: errorCode, Message: ErrorMessages[errorCode]}
		return
	}
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

func (e *EventService) pushMessage(payload Payload, response *Response) {
	(*response).Data = payload.Data
	(*response).Data["roomName"] = payload.RoomName
	(*response).Data["appName"] = payload.AppName
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
