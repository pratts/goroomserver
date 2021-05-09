package main

import (
	models "github.com/pratts/go-room-server/models"
)

type EventService struct {
}

func (e *EventService) getEvent(code int) string {
	return models.EventText(code)
}
