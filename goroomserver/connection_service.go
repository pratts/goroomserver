package goroomserver

import (
	"github.com/gorilla/websocket"
)

type ConnectionService struct {
	connectionMap map[string]Connection
	numConnection int
	eventService  EventService
}

func (connectionService *ConnectionService) Init(eventService *EventService) {
	connectionService.numConnection = 0
	connectionService.connectionMap = make(map[string]Connection)
	connectionService.eventService = *eventService
}

func (connectionService *ConnectionService) addConnection(connection *websocket.Conn) {
	conn := Connection{Id: connectionService.numConnection + 1, SocketConnection: connection, RemoteAddress: connection.RemoteAddr().String()}
	connectionService.numConnection += 1
	connectionService.connectionMap[connection.RemoteAddr().String()] = conn
}
