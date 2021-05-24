package server

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	config "github.com/pratts/goroomserver/config"
)

type WebSocketService struct {
	eventService *EventService
}

func (webSocketService *WebSocketService) listen(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	conn := Connection{Id: 0, SocketConnection: c, RemoteAddress: c.RemoteAddr().String()}
	payload := Payload{EventType: CONNECTION, Connection: conn}
	webSocketService.notifyEventHandler(payload)
	defer c.Close()
	for {
		payload := Payload{}
		_, message, err := c.ReadMessage()
		if err != nil {
			payload := Payload{EventType: DISCONNECTION, RemoteAddress: c.RemoteAddr().String()}
			webSocketService.notifyEventHandler(payload)
			log.Println("Error while reading message:", err)
			break
		}
		parseError := json.Unmarshal(message, &payload)
		if parseError != nil {
			log.Println("Error while parsing data:", parseError)
		}
		payload.RemoteAddress = c.RemoteAddr().String()
		webSocketService.notifyEventHandler(payload)
	}
}

func (webSocketService *WebSocketService) notifyEventHandler(payload Payload) {
	go webSocketService.eventService.handleEvent(payload)
}

func (webSocketService *WebSocketService) WriteToSocket(c *websocket.Conn, res Response) {
	data, parseError := json.Marshal(&res)
	if parseError != nil {
		log.Println("error in parsing")
	}
	err := c.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		log.Println("write:", err)
	}
}

func (webSocketService *WebSocketService) startWebSocketServer(c *config.ServerConfiguration) {
	var url = fmt.Sprintf("%s:%d", c.GetHost(), c.GetPort())
	var addr = flag.String("addr", url, "http service address")
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc(c.GetRequestPattern(), webSocketService.listen)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
