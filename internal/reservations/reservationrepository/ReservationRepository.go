package reservationrepository

import (
	"Amora/internal/reservations"
	"Amora/internal/helper"
)

type ReservationRepository interface {
	GetByUserId(userId int) ([]reservations.Reservations, error)
	GetByEventId(eventId int) ([]helper.ReservationDetailsDTO, error)
}
