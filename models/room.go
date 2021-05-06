package models

type Room struct {
	id       int
	name     string
	usersMap map[string]User
}

func (r *Room) GetId() int {
	return r.id
}

func (r *Room) GetRoomName() string {
	return r.name
}

func (r *Room) GetUserMap() map[string]User {
	return r.usersMap
}

func (r *Room) GetUserByName(userName string) User {
	return r.usersMap[userName]
}

func (r *Room) AddUser(u User) map[string]User {
	r.usersMap[u.name] = u
	u.AddRoom(*r)
	return r.usersMap
}

func (r *Room) RemoveUser(u User) map[string]User {
	u.RemoveRoom(*r)
	delete(r.usersMap, u.GetName())
	return r.usersMap
}

func (r *Room) ClearUsers() map[string]User {
	r.usersMap = make(map[string]User)
	return r.usersMap
}

func (r *Room) RemoveAllUsers() {
	for _, user := range r.usersMap {
		r.RemoveUser(user)
	}
}
