package services

import (
	interfaces "github.com/pratts/go-room-server/interfaces"
)

type AppService struct {
	Name         string
	roomService  RoomService
	userService  UserService
	eventHandler map[int]interfaces.Event
	Extension    interfaces.Extension
}

func (appService *AppService) InitData() {
	appService.roomService = RoomService{}
	appService.roomService.InitRoomService()

	appService.userService = UserService{}
	appService.userService.InitUserService()

	appService.CreateEventHander()
}

func (appService *AppService) setExtension(extension *interfaces.Extension) {
	appService.Extension = *extension
}

func (appService *AppService) getRoomService() RoomService {
	return appService.roomService
}

func (appService *AppService) getUserService() UserService {
	return appService.userService
}

func (appService *AppService) addExtension(extension interfaces.Extension) {
	appService.Extension = extension
}

func (appService *AppService) addEventHandler(code int, e interfaces.Event) {
	appService.eventHandler[code] = e
}

func (appService *AppService) removeEventHandler(code int) {
	delete(appService.eventHandler, code)
}

func (appService *AppService) CreateEventHander() map[int]interfaces.Event {
	appService.eventHandler = make(map[int]interfaces.Event)
	return appService.eventHandler
}
