package models

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

func (c *Connection) getConnection() *websocket.Conn {
	return c.SocketConnection
}

func (c *Connection) getRemoteAddress() string {
	return c.RemoteAddress
}
