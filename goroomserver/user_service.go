package goroomserver

type UserService struct {
	userMap           map[string]User
	userCount         int
	connectionUserMap map[string]User
}

func (us *UserService) InitUserService() {
	us.userMap = make(map[string]User)
	us.connectionUserMap = make(map[string]User)
	us.userCount = 0
}

func (us *UserService) GetUserMap() map[string]User {
	return us.userMap
}

func (us *UserService) AddUser(user User) {
	us.userMap[user.GetName()] = user
}

func (us *UserService) AddUserConnection(ip string, user User) {
	us.connectionUserMap[ip] = user
}

func (us *UserService) GetUserForConnection(ip string) (User, bool) {
	user, ok := us.connectionUserMap[ip]
	return user, ok
}

func (us *UserService) CreateAndAddUser(name string, connection Connection) {
	us.userCount++
	user := User{name: name, id: us.userCount, connection: &connection}
	us.AddUser(user)
	us.AddUserConnection(connection.getRemoteAddress(), user)
}

func (us *UserService) RemoveUser(userName string) {
	user, ok := us.userMap[userName]
	if ok == true {
		user.DisconnectUser()
	}
	delete(us.userMap, userName)
}

func (us *UserService) GetUserByName(userName string) (User, bool) {
	user, ok := us.userMap[userName]
	return user, ok
}

func (us *UserService) NotifyUser(user User, payload Response) {
	user.SendMessageToUser(payload)
}

func (us *UserService) NotifyAll(payload Response) {
	for _, user := range us.userMap {
		us.NotifyUser(user, payload)
	}
}

func (us *UserService) NotifyUsers(userList []string, payload Response) {
	for i := 0; i < len(userList); i++ {
		user, _ := us.GetUserByName(userList[i])
		us.NotifyUser(user, payload)
	}
}
