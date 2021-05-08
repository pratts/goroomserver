package models

type Room struct {
	Id           int
	Name         string
	UsersMap     map[string]User
	MaxUserCount int
}

func (r *Room) GetId() int {
	return r.Id
}

func (r *Room) GetRoomName() string {
	return r.Name
}

func (r *Room) CreateUserMap() map[string]User {
	r.UsersMap = make(map[string]User)
	return r.UsersMap
}

func (r *Room) GetUserMap() map[string]User {
	return r.UsersMap
}

func (r *Room) GetMaxUserCount() int {
	return r.MaxUserCount
}

func (r *Room) GetUserByName(userName string) User {
	return r.UsersMap[userName]
}

func (r *Room) AddUser(u User) map[string]User {
	r.UsersMap[u.name] = u
	u.AddRoom(*r)
	return r.UsersMap
}

func (r *Room) RemoveUser(u User) map[string]User {
	u.RemoveRoom(*r)
	delete(r.UsersMap, u.GetName())
	return r.UsersMap
}

func (r *Room) ClearUsers() map[string]User {
	return r.CreateUserMap()
}

func (r *Room) RemoveAllUsers() {
	for _, user := range r.UsersMap {
		r.RemoveUser(user)
	}
}
