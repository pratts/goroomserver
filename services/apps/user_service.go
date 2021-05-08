package services

import (
	models "github.com/pratts/go-room-server/models"
)

type UserService struct {
	userMap map[string]models.User
}

func (us *UserService) GetUserMap() map[string]models.User {
	return us.userMap
}

func (us *UserService) AddUser(user models.User) {
	us.userMap[user.GetName()] = user
}

func (us *UserService) RemoveUser(userName string) {
	user, ok := us.userMap[userName]
	if ok == true {
		user.DisconnectUser()
	}
	delete(us.userMap, userName)
}

func (us *UserService) GetUserByName(userName string) models.User {
	return us.userMap[userName]
}
