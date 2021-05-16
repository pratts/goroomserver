package goroomserver

import (
	"sync"
)

type MainService struct {
	connectionService ConnectionService
	eventService      EventService
	appServices       map[string]AppService
	webSocketService  WebSocketService
}

func (mainService *MainService) Init(wg *sync.WaitGroup) {
	defer wg.Done()
	mainService.eventService = EventService{mainService: mainService}
	mainService.connectionService = ConnectionService{}
	mainService.connectionService.Init(&mainService.eventService)

	mainService.webSocketService = WebSocketService{
		eventService: &mainService.eventService,
	}
	mainService.webSocketService.StartConnectionService()
}

func (mainService *MainService) createAppService(appName string, extension Extension) {
	appService := AppService{Name: appName, Extension: extension}
	appService.InitData()
	mainService.appServices[appName] = appService
}

func (mainService *MainService) getAppService(appName string) AppService {
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
		}
	}
	return mainServiceInstance
}
