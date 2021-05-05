package models

type Room struct {
	id       int
	name     string
	usersMap map[string]User
}

func (r *Room) getId() int {
	return r.id
}

func (r *Room) getRoomName() string {
	return r.name
}

func (r *Room) getUserMap() map[string]User {
	return r.usersMap
}

func (r *Room) getUserByName(userName string) User {
	return r.usersMap[userName]
}

func (r *Room) addUser(u User) {
	r.usersMap[u.name] = u
	u.addRoom(*r)
}

func (r *Room) removeUser(u User) {
	delete(r.usersMap, u.getName())
	u.removeRoom(*r)
}
