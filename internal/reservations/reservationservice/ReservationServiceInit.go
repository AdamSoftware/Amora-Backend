package reservationservice

import (
	"Amora/internal/reservations/reservationrepository"
	"Amora/internal/reservations"
	"Amora/internal/Logging"
	"errors"
	"Amora/internal/helper"
)

type ReservationServiceInit struct {
	reservationRepo *reservationrepository.ReservationRepositoryInit
}

func NewReservationServiceInit( reservationRepo *reservationrepository.ReservationRepositoryInit) (*ReservationServiceInit, error) {

	if reservationRepo == nil {
		Logging.Logs.Errorf("reservation repository is nil")
		return nil, errors.New("reservation repository is nil")
	}

	return &ReservationServiceInit{
		reservationRepo: reservationRepo,
	}, nil
}

func (r *ReservationServiceInit) GetByUserId(userId int) ([]reservations.Reservations, error) {
	reservationsList, err := r.reservationRepo.GetByUserId(userId)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving reservations by user ID %d: %v", userId, err)
		return nil, err
	}
	return reservationsList, nil
}

func (r *ReservationServiceInit) GetByEventId(eventId int) ([]helper.ReservationDetailsDTO, error) {
	reservationsDetails, err := r.reservationRepo.GetByEventId(eventId)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving reservations details by event ID %d: %v", eventId, err)
		return nil, err
	}
	return reservationsDetails, nil
}
