package events

type LeaveRoomEvent struct {
	name string
}

func (e *LeaveRoomEvent) getName() string {
	return e.name
}

func (e *LeaveRoomEvent) handle(params map[string]interface{}) {

}
