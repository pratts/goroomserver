package main

import (
	"fmt"
	"sync"

	interfaces "github.com/pratts/go-room-server/interfaces"
	services "github.com/pratts/go-room-server/services"
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
	mainService.appServices[appName] = appService
}

func (mainService *MainService) getAppService(appName string) services.AppService {
	return mainService.appServices[appName]
}

var mainServiceInstance *MainService
var lock = &sync.Mutex{}

func GetInstance() *MainService {
	if mainServiceInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if mainServiceInstance == nil {
			mainServiceInstance = &MainService{}
			mainServiceInstance.Init()
		}
	}
	return mainServiceInstance
}

func main() {
	GetInstance()
}

// var mainService = GetInstance()
