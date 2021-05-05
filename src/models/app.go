package models

type App struct {
	name    string
	userMap map[string]User
	roomMap map[string]Room
}

func (app *App) getName() string {
	return app.name
}

func (app *App) getAllUsers() map[string]User {
	return app.userMap
}

func (app *App) getAllRooms() map[string]Room {
	return app.roomMap
}

func (app *App) getRoomByName(roomName string) Room {
	return app.roomMap[roomName]
}

func (app *App) getUserByName(userName string) User {
	return app.userMap[userName]
}

func (app *App) addRoom(r Room) {
	app.roomMap[r.getRoomName()] = r
}

func (app *App) addUser(u User) {
	app.userMap[u.getName()] = u
}

func (app *App) removeRoom(r Room) {
	r.removeAllUsers()
	delete(app.roomMap, r.getRoomName())
}

func (app *App) removeUser(u User) {
	u.disconnectUser()
	delete(app.userMap, u.getName())
}
