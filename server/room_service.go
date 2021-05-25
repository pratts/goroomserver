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

func (rs *RoomService) createRoom(roomName string, maxUsers int, extension *Extension) Room {
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
