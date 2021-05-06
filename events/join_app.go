package events

type JoinAppEvent struct {
	name string
}

func (e *JoinAppEvent) getName() string {
	return e.name
}

func (e *JoinAppEvent) handle(params map[string]interface{}) {

}
