package events

type LogoutEvent struct {
	name string
}

func (e *LogoutEvent) getName() string {
	return e.name
}

func (e *LogoutEvent) handle(params map[string]interface{}) {

}
