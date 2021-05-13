package main

import (
	"github.com/pratts/go-room-server/goroomserver"
)

func main() {
	instance := goroomserver.GetInstance()
	instance.Init()
}

// var mainService = GetInstance()
