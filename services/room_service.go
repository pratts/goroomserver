package services

import (
	models "github.com/pratts/go-room-server/models"
)

type RoomService struct {
	roomMap map[string]models.Room
}

func (rs *RoomService) GetRoomMap() map[string]models.Room {
	return rs.roomMap
}

func (rs *RoomService) AddRoom(room models.Room) {
	rs.roomMap[room.GetRoomName()] = room
}

func (rs *RoomService) RemoveRoom(roomName string) {
	room, ok := rs.roomMap[roomName]
	if ok == true {
		room.RemoveAllUsers()
	}
	delete(rs.roomMap, roomName)
}

func (rs *RoomService) GetRoomByName(roomName string) models.Room {
	return rs.roomMap[roomName]
}
