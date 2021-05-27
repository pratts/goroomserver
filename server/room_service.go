package server

type RoomService struct {
	roomMap         map[string]Room
	numRoomsCreated int
}

func (rs *RoomService) InitRoomService() {
	rs.roomMap = make(map[string]Room)
	rs.numRoomsCreated = 0
}

func (rs *RoomService) GetRoomMap() map[string]Room {
	return rs.roomMap
}

func (rs *RoomService) createRoom(roomName string, maxUsers int, extension Extension) Room {
	room := Room{id: rs.numRoomsCreated + 1, name: roomName, maxUserCount: maxUsers, extension: extension}
	room.InitRoomData()
	rs.numRoomsCreated += 1
	rs.addRoom(room)
	return room
}

func (rs *RoomService) addRoom(room Room) {
	rs.roomMap[room.GetRoomName()] = room
}

func (rs *RoomService) removeRoom(roomName string) {
	room, ok := rs.roomMap[roomName]
	if ok == true {
		room.RemoveAllUsers()
	}
	delete(rs.roomMap, roomName)
}

func (rs *RoomService) getRoomByName(roomName string) (Room, bool) {
	room, ok := rs.roomMap[roomName]
	return room, ok
}

func (rs *RoomService) getUserForRoom(roomName string) map[string]User {
	room, ok := rs.roomMap[roomName]
	if ok == true {
		return room.GetUserMap()
	}
	return nil
}

func (rs *RoomService) joinRoom(appService *AppService, user User, roomName string, payload map[string]interface{}) int {
	// user, ok := appService.userService.GetUserByName(userName)
	if roomName == "" {
		return ROOM_NAME_INVALID
	}
	room, ok := rs.getRoomByName(roomName)
	if ok == false {
		//handle case when room does not exist
		return ROOM_NOT_EXISTS
	}
	_, userExists := room.GetUserByName(user.name)
	if userExists == true {
		//handle case when user already is in room
		return USER_ALREADY_IN_ROOM
	}

	handler, evtExists := room.eventHandler[JOIN_ROOM]
	if evtExists == true {
		event := Event{
			payload: payload,
			room:    room,
			app:     *appService,
			user:    user,
		}
		_, err := handler.handleEvent(event)
		if err != nil {
			return ROOM_JOIN_ERROR
		}
	}

	room.addUser(user)
	return SUCCESS
}

func (rs *RoomService) leaveRoom(appService *AppService, user User, roomName string, payload map[string]interface{}) int {
	if roomName == "" {
		//handle case when roomname is blank
		return ROOM_NAME_INVALID
	}
	room, ok := rs.getRoomByName(roomName)
	if ok == false {
		//handle case when room does not exist
		return ROOM_NOT_EXISTS
	}
	handler, evtExists := room.eventHandler[LEAVE_ROOM]

	_, userExists := room.GetUserByName(user.name)
	if userExists == false {
		//handle case when user already is not in room
		return USER_NOT_IN_ROOM
	}

	if evtExists == true {
		event := Event{
			payload: payload,
			room:    room,
			app:     *appService,
			user:    user,
		}
		_, err := handler.handleEvent(event)
		if err != nil {
			return LEAVE_ROOM_ERROR
		}
	}
	room.removeUser(user)
	return SUCCESS
}
