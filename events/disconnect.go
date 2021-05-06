package events

type DisconnectionEvent struct {
	name string
}

func (c *DisconnectionEvent) getName() string {
	return c.name
}

func (c *DisconnectionEvent) handEvent(params map[string]interface{}) {

}
