package events

type ConnectionEvent struct {
	name string
}

func (c *ConnectionEvent) getName() string {
	return c.name
}

func (c *ConnectionEvent) handEvent(params map[string]interface{}) {

}
