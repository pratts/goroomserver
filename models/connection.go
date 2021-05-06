package models

type Connection struct {
	id   int
	port int
	ip   string
}

func (c Connection) GetId() int {
	return c.id
}

func (c Connection) GetPort() int {
	return c.port
}

func (c Connection) GetIP() string {
	return c.ip
}
