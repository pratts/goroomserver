package main

import (
	"sync"

	"github.com/pratts/go-room-server/goroomserver"
)

func main() {
	instance := goroomserver.GetInstance()
	instance.Init()
	var wg sync.WaitGroup
	wg.Add(1)
	go instance.StartServer(&wg)
	wg.Wait()
}
