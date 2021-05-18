package goroomserver

type Room struct {
	Id           int
	Name         string
	UsersMap     map[string]User
	MaxUserCount int
	eventHandler map[int]EventHandler
	Extension    Extension
}

func (r *Room) InitRoomData() {
	r.CreateUserMap()
	r.CreateEventHander()
	r.initExtension()
}

func (r *Room) initExtension() {
	payload := make(map[string]interface{})
	payload["room"] = r
	r.Extension.init(payload)
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

func (r *Room) GetUserByName(userName string) (User, bool) {
	user, ok := r.UsersMap[userName]
	return user, ok
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

func (r *Room) addEventHandler(code int, e EventHandler) {
	r.eventHandler[code] = e
}

func (r *Room) removeEventHandler(code int) {
	delete(r.eventHandler, code)
}

func (r *Room) CreateEventHander() map[int]EventHandler {
	r.eventHandler = make(map[int]EventHandler)
	return r.eventHandler
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
	for _, user := range r.UsersMap {
		user.SendMessageToUser(payload)
	}
}
