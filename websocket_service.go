package goroomserver

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type WebSocketService struct {
	connectionService *ConnectionService
	eventService      *EventService
}

func (webSocketService *WebSocketService) listen(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	webSocketService.connectionService.addConnection(c)
	defer c.Close()
	for {
		payload := make(map[string]interface{})
		err := c.ReadJSON(&payload)
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", payload)
		webSocketService.eventService.handleEvent(payload)
	}
}

func (webSocketService *WebSocketService) WriteToSocket(c *websocket.Conn, payload map[string]interface{}) {
	err := c.WriteJSON(payload)
	if err != nil {
		log.Println("write:", err)
	}
}

func (webSocketService *WebSocketService) StartConnectionService() {
	var addr = flag.String("addr", "localhost:8080", "http service address")
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", webSocketService.listen)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
