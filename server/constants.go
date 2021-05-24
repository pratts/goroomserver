package server

const (
	CONNECTION       int = 1
	DISCONNECTION    int = 2
	JOIN_APP         int = 3
	JOIN_ROOM        int = 4
	LEAVE_ROOM       int = 5
	LOGIN            int = 6
	LOGOUT           int = 7
	RECONNECT        int = 8
	SUBSCRIBE_ROOM   int = 9
	UNSUBSCRIBE_ROOM int = 10
	MESSAGE          int = 11
)

var EventsText = map[int]string{
	CONNECTION:       "Connection",
	DISCONNECTION:    "Disconnection",
	JOIN_APP:         "Join App",
	LEAVE_ROOM:       "Leave Room",
	LOGIN:            "Login",
	LOGOUT:           "Logout",
	RECONNECT:        "Reconnect",
	SUBSCRIBE_ROOM:   "Subscribe",
	UNSUBSCRIBE_ROOM: "Unsubscribe",
	MESSAGE:          "Message",
}

func EventText(code int) string {
	return EventsText[code]
}

const (
	APP_NAME_INVALID       int = 1
	USER_NOT_EXISTS        int = 2
	USER_LOGIN_ERROR       int = 3
	USER_ALREADY_LOGGED_IN int = 4
	USER_LOGOUT_ERROR      int = 5
	ROOM_NAME_INVALID      int = 6
	ROOM_NOT_EXISTS        int = 7
	USER_ALREADY_IN_ROOM   int = 8
	USER_NOT_IN_ROOM       int = 9
	INVALID_EVENT          int = 10
	SERVER_ERROR           int = 500
	SUCCESS                int = 200
)

var ErrorMessages = map[int]string{
	APP_NAME_INVALID:       "App name is invalid",
	USER_NOT_EXISTS:        "User unavailable in app",
	USER_LOGIN_ERROR:       "User login error",
	USER_ALREADY_LOGGED_IN: "User already logged in to app",
	USER_LOGOUT_ERROR:      "User logout error",
	ROOM_NAME_INVALID:      "Room name provided is invalid",
	ROOM_NOT_EXISTS:        "Room does not exist in app",
	USER_ALREADY_IN_ROOM:   "User has already joined the room",
	USER_NOT_IN_ROOM:       "User is not available in the room",
	INVALID_EVENT:          "Invalid event provided",
	SERVER_ERROR:           "Server Error",
}
