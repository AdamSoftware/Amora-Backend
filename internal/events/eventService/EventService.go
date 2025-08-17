package eventService

import (
	"Amora/internal/events"
	"Amora/internal/helper"
)

type EventService interface {
	FindByUserID(userID int) ([]events.Events, error)
	GetEventsWithGiftsByUserID(userID int) ([]helper.EventWithGiftsDTO, error)
	GetEventsWithReservationsWithUserId(userId int) ([]helper.EventWithReservationsDTO, error)
}
