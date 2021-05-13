package goroomserver

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

func (rs *RoomService) CreateRoom(roomName string, maxUsers int, extension interface{}) {
	room := Room{Id: rs.numRoomsCreated + 1, Name: roomName, MaxUserCount: maxUsers, Extension: extension}
	room.InitRoomData()
	rs.numRoomsCreated += 1
	rs.AddRoom(room)
}

func (rs *RoomService) AddRoom(room Room) {
	rs.roomMap[room.GetRoomName()] = room
}

func (rs *RoomService) RemoveRoom(roomName string) {
	room, ok := rs.roomMap[roomName]
	if ok == true {
		room.RemoveAllUsers()
	}
	delete(rs.roomMap, roomName)
}

func (rs *RoomService) GetRoomByName(roomName string) Room {
	return rs.roomMap[roomName]
}

func (rs *RoomService) GetUserForRoom(roomName string) map[string]User {
	room, ok := rs.roomMap[roomName]
	if ok == true {
		return room.GetUserMap()
	}
	return nil
}
