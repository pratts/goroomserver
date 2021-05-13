package goroomserver

const (
	CONNECTION    int = 1
	DISCONNECTION int = 2
	JOIN_APP      int = 3
	JOIN_ROOM     int = 4
	LEAVE_ROOM    int = 5
	LOGIN         int = 6
	LOGOUT        int = 7
	RECONNECT     int = 8
	SUBSCRIBE     int = 9
	UNSUBSCRIBE   int = 10
)

var eventsText = map[int]string{
	CONNECTION:    "Connection",
	DISCONNECTION: "Disconnection",
	JOIN_APP:      "Join App",
	LEAVE_ROOM:    "Leave Room",
	LOGIN:         "Login",
	LOGOUT:        "Logout",
	RECONNECT:     "Reconnect",
	SUBSCRIBE:     "Subscribe",
	UNSUBSCRIBE:   "Unsubscribe",
}

func EventText(code int) string {
	return eventsText[code]
}
