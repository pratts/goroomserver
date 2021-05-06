package events

type JoinRoomEvent struct {
	name string
}

func (e *JoinRoomEvent) getName() string {
	return e.name
}

func (e *JoinRoomEvent) handle(params map[string]interface{}) {

}
