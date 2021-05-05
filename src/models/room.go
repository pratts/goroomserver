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

func (r *Room) addUser(u User) map[string]User {
	r.usersMap[u.name] = u
	u.addRoom(*r)
	return r.usersMap
}

func (r *Room) removeUser(u User) map[string]User {
	u.removeRoom(*r)
	delete(r.usersMap, u.getName())
	return r.usersMap
}

func (r *Room) clearUsers() map[string]User {
	r.usersMap = make(map[string]User)
	return r.usersMap
}

func (r *Room) removeAllUsers() {
	for _, user := range r.usersMap {
		r.removeUser(user)
	}
}
