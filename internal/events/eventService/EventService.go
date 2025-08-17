package eventService

import (
	"Amora/internal/events"
)

type EventService interface {
	FindByUserID(userID int) ([]events.Events, error)
}
