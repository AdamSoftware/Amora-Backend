package eventRepository

import (
	"Amora/internal/events"
	"Amora/internal/helper"
)


type EventRepository interface {
	FindByUserId(userId int) ([]events.Events, error)
	GetEventsWithGiftsWithUserId(userId int) ([]events.Events, error)
	GetEventsWithReservationsWithUserId(userId int) ([]helper.EventWithReservationsDTO, error)
}

