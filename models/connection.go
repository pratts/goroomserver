package models

type Connection struct {
	id   int
	port int
	ip   string
}

func (c Connection) getId() int {
	return c.id
}

func (c Connection) getPort() int {
	return c.port
}

func (c Connection) getIP() string {
	return c.ip
}
