package goroomserver

type AppService struct {
	name         string
	roomService  RoomService
	userService  UserService
	eventHandler map[int]EventHandler
	extension    Extension
}

func (appService *AppService) InitData() {
	appService.roomService = RoomService{}
	appService.roomService.InitRoomService()

	appService.userService = UserService{}
	appService.userService.InitUserService()

	appService.CreateEventHandler()
}

func (appService *AppService) GetName() string {
	return appService.name
}

func (appService *AppService) initExtension() {
	payload := make(map[string]interface{})
	appService.extension.init(payload)
}

func (appService *AppService) setExtension(extension *Extension) {
	appService.extension = *extension
}

func (appService *AppService) GetExtension() Extension {
	return appService.extension
}

func (appService *AppService) GetRoomService() RoomService {
	return appService.roomService
}

func (appService *AppService) GetUserService() UserService {
	return appService.userService
}

func (appService *AppService) addExtension(extension Extension) {
	appService.extension = extension
}

func (appService *AppService) addEventHandler(code int, e EventHandler) {
	appService.eventHandler[code] = e
}

func (appService *AppService) removeEventHandler(code int) {
	delete(appService.eventHandler, code)
}

func (appService *AppService) CreateEventHandler() map[int]EventHandler {
	appService.eventHandler = make(map[int]EventHandler)
	return appService.eventHandler
}

func (appService *AppService) CreateRoom(roomName string, maxUsers int, extension Extension) Room {
	room := appService.roomService.createRoom(roomName, maxUsers, extension)
	return room
}

func (appService *AppService) RemoveRoom(roomName string) {
	appService.roomService.removeRoom(roomName)
}

func (appService *AppService) GetRoomByName(roomName string) (Room, bool) {
	return appService.roomService.getRoomByName(roomName)
}

func (appService *AppService) GetUserForRoom(roomName string) map[string]User {
	return appService.roomService.getUserForRoom(roomName)
}
