package main

type MainService struct {
	connectionService ConnectionService
	eventService      EventService
}

func (mainService *MainService) Init() {
	mainService.eventService = EventService{}

	mainService.connectionService = ConnectionService{}
	mainService.connectionService.Init(&mainService.eventService)
	mainService.connectionService.StartConnectionService()
}
