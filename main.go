package main

import (
	"sync"

	"github.com/pratts/go-room-server/goroomserver"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	instance := goroomserver.GetInstance()
	go instance.Init(&wg)
	wg.Wait()
}

// var mainService = GetInstance()
