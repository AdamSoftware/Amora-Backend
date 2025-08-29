package reservationservice

import (
	"Amora/internal/reservations"
	"Amora/internal/helper"
)

type ReservationService struct {
	GetByUserId   func(userId int) ([]reservations.Reservations, error)
	GetByEventId  func(eventId int) ([]helper.ReservationDetailsDTO, error)
}
