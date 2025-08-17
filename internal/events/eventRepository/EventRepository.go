package eventRepository

import (
	"Amora/internal/events"
)


type EventRepository interface {
	FindByUserId(userId int) ([]events.Events, error)
}

