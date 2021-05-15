// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

type test_data struct {
	Id   int
	Name string
}

type Payload struct {
	AppName   string      `json:"appName"`
	RoomName  string      `json:"roomName"`
	EventType int         `json:"eventType"`
	Payload   interface{} `json:"payload"`
}

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			log.Println("read data")
			payload := Payload{}
			_, message, err := c.ReadMessage()
			parseError := json.Unmarshal(message, &payload)
			if parseError != nil {
				return
			}
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", payload)
		}
	}()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			fmt.Println("t:", t)
			payload := make(map[string]interface{})
			payload["a"] = 1
			data := Payload{RoomName: "test1", EventType: 1, Payload: payload}
			b, err1 := json.Marshal(&data)
			if err1 != nil {
				fmt.Println("error in parsing")
			}
			err := c.WriteMessage(websocket.TextMessage, b)
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
