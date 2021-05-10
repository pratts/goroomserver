package main

import (
	interfaces "github.com/pratts/go-room-server/interfaces"
	services "github.com/pratts/go-room-server/services/apps"
)

type MainService struct {
	connectionService ConnectionService
	eventService      EventService
	appServices       map[string]services.AppService
}

func (mainService *MainService) Init() {
	mainService.eventService = EventService{}

	mainService.connectionService = ConnectionService{}
	mainService.connectionService.Init(&mainService.eventService)
	mainService.connectionService.StartConnectionService()
}

func (mainService *MainService) createAppService(appName string, extension interfaces.Extension) {
	appService := services.AppService{Name: appName, Extension: extension}
	appService.InitData()
}
