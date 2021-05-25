package server

type Room struct {
	id           int
	name         string
	usersMap     map[string]User
	maxUserCount int
	eventHandler map[int]EventHandler
	extension    *Extension
}

func (r *Room) InitRoomData() {
	r.createUserMap()
	r.createEventHandler()
	r.initExtension()
}

func (r *Room) initExtension() {
	event := Event{room: *r}
	(*r.extension).init(event)
}

func (r *Room) GetId() int {
	return r.id
}

func (r *Room) GetRoomName() string {
	return r.name
}

func (r *Room) createUserMap() map[string]User {
	r.usersMap = make(map[string]User)
	return r.usersMap
}

func (r *Room) GetUserMap() map[string]User {
	return r.usersMap
}

func (r *Room) GetMaxUserCount() int {
	return r.maxUserCount
}

func (r *Room) GetUserByName(userName string) (User, bool) {
	user, ok := r.usersMap[userName]
	return user, ok
}

func (r *Room) addUser(u User) map[string]User {
	r.usersMap[u.name] = u
	u.AddRoom(*r)
	return r.usersMap
}

func (r *Room) removeUser(u User) map[string]User {
	u.RemoveRoom(*r)
	delete(r.usersMap, u.GetName())
	return r.usersMap
}

func (r *Room) ClearUsers() map[string]User {
	return r.createUserMap()
}

func (r *Room) RemoveAllUsers() {
	for _, user := range r.usersMap {
		r.removeUser(user)
	}
}

func (r *Room) addEventHandler(code int, e EventHandler) {
	r.eventHandler[code] = e
}

func (r *Room) removeEventHandler(code int) {
	delete(r.eventHandler, code)
}

func (r *Room) createEventHandler() map[int]EventHandler {
	r.eventHandler = make(map[int]EventHandler)
	return r.eventHandler
}

func (r *Room) getEventHandler() map[int]EventHandler {
	return r.eventHandler
}

func (r *Room) getExtension() Extension {
	return *r.extension
}

func (r *Room) sendResponseToUser(userName string, payload Response) {
	u, ok := r.GetUserByName(userName)
	if ok == true {
		u.SendMessageToUser(payload)
	}
}

func (r *Room) sendResponseToUserList(userList []string, payload Response) {
	for i := 0; i < len(userList); i++ {
		u, ok := r.GetUserByName(userList[i])
		if ok == true {
			u.SendMessageToUser(payload)
		}
	}
}

func (r *Room) sendResponseToAll(payload Response) {
	for _, user := range r.usersMap {
		user.SendMessageToUser(payload)
	}
}
