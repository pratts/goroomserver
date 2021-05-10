package services

import (
	interfaces "github.com/pratts/go-room-server/interfaces"
	models "github.com/pratts/go-room-server/models"
)

type RoomService struct {
	roomMap         map[string]models.Room
	numRoomsCreated int
	eventHandler    map[int]interfaces.Extension
}

func (rs *RoomService) InitRoomService() {
	rs.roomMap = make(map[string]models.Room)
	rs.numRoomsCreated = 0
	rs.eventHandler = make(map[int]interfaces.Extension)
}

func (rs *RoomService) GetRoomMap() map[string]models.Room {
	return rs.roomMap
}

func (rs *RoomService) CreateRoom(roomName string, maxUsers int, extension interface{}) {
	room := models.Room{Id: rs.numRoomsCreated + 1, Name: roomName, MaxUserCount: maxUsers, Extension: extension}
	room.InitRoomData()
	rs.numRoomsCreated += 1
	rs.AddRoom(room)
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

func (rs *RoomService) GetUserForRoom(roomName string) map[string]models.User {
	room, ok := rs.roomMap[roomName]
	if ok == true {
		return room.GetUserMap()
	}
	return nil
}
