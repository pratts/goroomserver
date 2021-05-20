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
	mainService.connectionService.Init()

	mainService.webSocketService = WebSocketService{
		eventService: &mainService.eventService,
	}
	mainService.webSocketService.StartWebSocketServer()
}

func (mainService *MainService) CreateAppService(appName string, extension Extension) {
	appService := AppService{name: appName, extension: extension}
	appService.InitData()
	mainService.appServices[appName] = appService
}

func (mainService *MainService) GetAppService(appName string) (AppService, bool) {
	app, ok := mainService.appServices[appName]
	return app, ok
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
