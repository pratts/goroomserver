package goroomserver

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

type MainService struct {
	connectionService ConnectionService
	eventService      EventService
	appServices       map[string]AppService
	webSocketService  WebSocketService
	serverConfig      ServerConfig
}

func (mainService *MainService) Init() {
	mainService.loadConfigFile()
	mainService.eventService = EventService{mainService: mainService}
	mainService.connectionService = ConnectionService{}
	mainService.connectionService.Init()

	mainService.webSocketService = WebSocketService{
		eventService: &mainService.eventService,
	}
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

func (mainService *MainService) StartServer(wg *sync.WaitGroup) {
	defer wg.Done()
	mainService.webSocketService.startWebSocketServer(&mainService.serverConfig)
}

func (mainService *MainService) loadConfigFile() {
	data, err := ioutil.ReadFile("config/server.json")
	if err != nil {
		fmt.Println("Error parsing:", err.Error())
	}

	err = json.Unmarshal(data, &(mainService.serverConfig))
	if err != nil {
		fmt.Println("Error setting data:", err.Error())
	}
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
