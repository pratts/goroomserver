package goroomserver

type AppService struct {
	Name         string
	roomService  RoomService
	userService  UserService
	eventHandler map[int]Event
	Extension    Extension
}

func (appService *AppService) InitData() {
	appService.roomService = RoomService{}
	appService.roomService.InitRoomService()

	appService.userService = UserService{}
	appService.userService.InitUserService()

	appService.CreateEventHander()
}

func (appService *AppService) setExtension(extension *Extension) {
	appService.Extension = *extension
}

func (appService *AppService) getRoomService() RoomService {
	return appService.roomService
}

func (appService *AppService) getUserService() UserService {
	return appService.userService
}

func (appService *AppService) addExtension(extension Extension) {
	appService.Extension = extension
}

func (appService *AppService) addEventHandler(code int, e Event) {
	appService.eventHandler[code] = e
}

func (appService *AppService) removeEventHandler(code int) {
	delete(appService.eventHandler, code)
}

func (appService *AppService) CreateEventHander() map[int]Event {
	appService.eventHandler = make(map[int]Event)
	return appService.eventHandler
}
