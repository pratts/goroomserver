package goroomserver

type AppService struct {
	Name         string
	roomService  RoomService
	userService  UserService
	eventHandler map[int]EventHandler
	Extension    Extension
}

func (appService *AppService) InitData() {
	appService.roomService = RoomService{}
	appService.roomService.InitRoomService()

	appService.userService = UserService{}
	appService.userService.InitUserService()

	appService.CreateEventHander()
}

func (appService *AppService) initExtension() {
	payload := make(map[string]interface{})
	appService.Extension.init(payload)
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

func (appService *AppService) addEventHandler(code int, e EventHandler) {
	appService.eventHandler[code] = e
}

func (appService *AppService) removeEventHandler(code int) {
	delete(appService.eventHandler, code)
}

func (appService *AppService) CreateEventHander() map[int]EventHandler {
	appService.eventHandler = make(map[int]EventHandler)
	return appService.eventHandler
}

func (appService *AppService) HandleEvent(payload map[string]interface{}) {

}
