package events

type ReconnectionEvent struct {
	name string
}

func (e *ReconnectionEvent) getName() string {
	return e.name
}

func (e *ReconnectionEvent) handle(params map[string]interface{}) {

}
