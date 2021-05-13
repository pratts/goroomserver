package goroomserver

type EventService struct {
}

func (e *EventService) getEvent(code int) string {
	return EventText(code)
}

func (e *EventService) handleEvent(payload map[string]interface{}) {
	// appName := payload["appName"]
	// roomName := payload["roomName"]
	// eventType := payload["eventType"]
	// remoteAddr := payload["remoteAddr"]
	// data := payload["data"]
}
