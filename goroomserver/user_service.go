package goroomserver

type UserService struct {
	userMap map[string]User
}

func (us *UserService) InitUserService() {
	us.userMap = make(map[string]User)
}

func (us *UserService) GetUserMap() map[string]User {
	return us.userMap
}

func (us *UserService) AddUser(user User) {
	us.userMap[user.GetName()] = user
}

func (us *UserService) RemoveUser(userName string) {
	user, ok := us.userMap[userName]
	if ok == true {
		user.DisconnectUser()
	}
	delete(us.userMap, userName)
}

func (us *UserService) GetUserByName(userName string) User {
	return us.userMap[userName]
}
