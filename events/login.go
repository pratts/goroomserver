package events

type LoginEvent struct {
	name string
}

func (e *LoginEvent) getName() string {
	return e.name
}

func (e *LoginEvent) handle(params map[string]interface{}) {

}
