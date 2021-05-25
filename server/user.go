package server

type User struct {
	id         int
	name       string
	roomMap    map[string]Room
	connection *Connection
}

func (u *User) GetId() int {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetRoomMap() map[string]Room {
	return u.roomMap
}

func (u *User) GetConnection() *Connection {
	return u.connection
}

func (u *User) AddRoom(r Room) map[string]Room {
	u.roomMap[r.GetRoomName()] = r
	return u.roomMap
}

func (u *User) RemoveRoom(r Room) map[string]Room {
	delete(u.GetRoomMap(), r.GetRoomName())
	return u.roomMap
}

func (u *User) GetRoomByName(roomName string) (Room, bool) {
	room, ok := u.GetRoomMap()[roomName]
	return room, ok
}

func (u *User) DisconnectUser() {
	for _, room := range u.roomMap {
		room.removeUser(*u)
	}
}

func (u *User) SendMessageToUser(payload Response) {
	u.connection.WriteToSocket(&payload)
}
