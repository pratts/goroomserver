package server

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

	appService.initExtension()
}

func (appService *AppService) GetName() string {
	return appService.name
}

func (appService *AppService) initExtension() {
	event := Event{app: *appService}
	appService.extension.InitExtension(event)
}

func (appService *AppService) setExtension(extension Extension) {
	appService.extension = extension
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
	response := Response{EventType: ROOM_ADD, Code: SUCCESS}
	data := map[string]interface{}{
		"roomName": roomName,
		"appName":  appService.name,
	}
	response.Data = data
	appService.userService.NotifyAll(response)
	return room
}

func (appService *AppService) RemoveRoom(roomName string) {
	appService.roomService.removeRoom(roomName)
	response := Response{EventType: ROOM_REMOVE, Code: SUCCESS}
	data := map[string]interface{}{
		"roomName": roomName,
		"appName":  appService.name,
	}
	response.Data = data
	appService.userService.NotifyAll(response)
}

func (appService *AppService) GetRoomByName(roomName string) (Room, bool) {
	return appService.roomService.getRoomByName(roomName)
}

func (appService *AppService) GetUserForRoom(roomName string) map[string]User {
	return appService.roomService.getUserForRoom(roomName)
}

func (appService *AppService) Login(payload Payload) {
	user := appService.userService.CreateAndAddUser(payload.Data["username"].(string), payload.Connection)
	response := Response{EventType: LOGIN, Code: SUCCESS}
	data := map[string]interface{}{
		"userName": user.name,
		"appName":  appService.name,
	}
	response.Data = data
	appService.userService.NotifyUser(user, response)
}

func (appService *AppService) Logout(payload Payload) {
	response := Response{EventType: LOGIN, Code: SUCCESS}
	data := map[string]interface{}{
		"userName": payload.RefUser.name,
		"appName":  appService.name,
	}
	response.Data = data
	appService.userService.NotifyUser(payload.RefUser, response)
	appService.userService.RemoveUser(payload.RefUser.name)
}

func (appService *AppService) JoinRoom(userName string, roomName string, payload map[string]interface{}) int {
	user, ok := appService.userService.GetUserByName(userName)
	if ok == false {
		return USER_NOT_EXISTS
	}
	return appService.JoinUserRoom(user, roomName, payload)
}

func (appService *AppService) JoinUserRoom(user User, roomName string, payload map[string]interface{}) int {
	return appService.roomService.joinRoom(appService, user, roomName, payload)
}

func (appService *AppService) LeaveRoom(userName string, roomName string, payload map[string]interface{}) int {
	user, ok := appService.userService.GetUserByName(userName)
	if ok == false {
		return USER_NOT_EXISTS
	}
	return appService.LeaveUserRoom(user, roomName, payload)
}

func (appService *AppService) LeaveUserRoom(user User, roomName string, payload map[string]interface{}) int {
	return appService.roomService.leaveRoom(appService, user, roomName, payload)
}
