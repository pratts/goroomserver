package goroomserver

type App struct {
	name    string
	userMap map[string]User
	roomMap map[string]Room
}

func (app *App) GetName() string {
	return app.name
}

func (app *App) GetAllUsers() map[string]User {
	return app.userMap
}

func (app *App) GetAllRooms() map[string]Room {
	return app.roomMap
}

func (app *App) GetRoomByName(roomName string) Room {
	return app.roomMap[roomName]
}

func (app *App) GetUserByName(userName string) User {
	return app.userMap[userName]
}

func (app *App) AddRoom(r Room) {
	app.roomMap[r.GetRoomName()] = r
}

func (app *App) AddUser(u User) {
	app.userMap[u.GetName()] = u
}

func (app *App) RemoveRoom(r Room) {
	r.RemoveAllUsers()
	delete(app.roomMap, r.GetRoomName())
}

func (app *App) RemoveUser(u User) {
	u.DisconnectUser()
	delete(app.userMap, u.GetName())
}
