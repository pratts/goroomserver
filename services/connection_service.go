package services

import (
	"github.com/gorilla/websocket"
	models "github.com/pratts/go-room-server/models"
)

type ConnectionService struct {
	connectionMap map[string]models.Connection
	numConnection int
	eventService  EventService
}

func (connectionService *ConnectionService) Init(eventService *EventService) {
	connectionService.numConnection = 0
	connectionService.connectionMap = make(map[string]models.Connection)
	connectionService.eventService = *eventService
}

func (connectionService *ConnectionService) addConnection(connection *websocket.Conn) {
	conn := models.Connection{Id: connectionService.numConnection + 1, SocketConnection: connection, RemoteAddress: connection.RemoteAddr().String()}
	connectionService.numConnection += 1
	connectionService.connectionMap[connection.RemoteAddr().String()] = conn
}
