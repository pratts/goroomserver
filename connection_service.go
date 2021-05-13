package main

import (
	"flag"
	"log"
	"net/http"

	models "github.com/pratts/go-room-server/models"

	"github.com/gorilla/websocket"
)

type ConnectionService struct {
	connectionMap map[string]models.Connection
	numConnection int
	eventService  EventService
}

func (connectionService *ConnectionService) Init(eventService *EventService) {
	connectionService.numConnection = 0
	connectionService.connectionMap = make(map[string]models.Connection)
	connectionService.eventService = *eventService
}

func (connectionService *ConnectionService) addConnection(connection *websocket.Conn) {
	conn := models.Connection{Id: connectionService.numConnection + 1, SocketConnection: connection, RemoteAddress: connection.RemoteAddr().String()}
	connectionService.numConnection += 1
	connectionService.connectionMap[connection.RemoteAddr().String()] = conn
}

func (connectionService *ConnectionService) listen(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	log.Printf("Started taking connections")
	connectionService.addConnection(c)
	defer c.Close()
	for {
		payload := make(map[string]interface{})
		err := c.ReadJSON(&payload)
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", payload)
		err = c.WriteJSON(payload)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func (connectionService *ConnectionService) StartConnectionService() {
	var addr = flag.String("addr", "localhost:8080", "http service address")
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", connectionService.listen)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
