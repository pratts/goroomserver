// User

//     - id
//     - username
//     - roomlist
//     - connection
package models

type User struct {
	id         int
	name       string
	roomMap    map[string]Room
	connection Connection
}

func (u *User) getId() int {
	return u.id
}

func (u *User) getName() string {
	return u.name
}

func (u *User) getRoomMap() map[string]Room {
	return u.roomMap
}

func (u *User) getConnection() Connection {
	return u.connection
}

func (u *User) addRoom(r Room) map[string]Room {
	u.roomMap[r.getRoomName()] = r
	return u.roomMap
}

func (u *User) removeRoom(r Room) map[string]Room {
	delete(u.getRoomMap(), r.getRoomName())
	return u.roomMap
}

func (u *User) getRoomByName(roomName string) Room {
	return u.getRoomMap()[roomName]
}
