package main

import (
	"sync"

	"github.com/pratts/goroomserver/server"
)

func main() {
	instance := server.GetInstance()
	instance.Init()
	var wg sync.WaitGroup
	wg.Add(1)
	go instance.StartServer(&wg)
	wg.Wait()
}
