package goroomserver

import (
	"github.com/gorilla/websocket"
)

type Connection struct {
	Id               int
	RemoteAddress    string
	SocketConnection *websocket.Conn
}

func (c *Connection) GetId() int {
	return c.Id
}

func (c *Connection) GetConnection() *websocket.Conn {
	return c.SocketConnection
}

func (c *Connection) GetRemoteAddress() string {
	return c.RemoteAddress
}

func (c *Connection) WriteToSocket(payload Response) {
	GetInstance().webSocketService.WriteToSocket(c.SocketConnection, payload)
}
