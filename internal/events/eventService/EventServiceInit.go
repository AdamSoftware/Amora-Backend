package eventService

import (
	"Amora/internal/events/eventRepository"
	"Amora/internal/events"
	"Amora/internal/Logging"
	"errors"
	"Amora/internal/helper"
)

type EventServiceInit struct {
	eventRepo *eventRepository.EventRepositoryInit
}

func NewEventServiceInit( eventRepo *eventRepository.EventRepositoryInit) (*EventServiceInit, error) {

	if eventRepo == nil {
		Logging.Logs.Errorf("event repository is nil")
		return nil, errors.New("event repository is nil")
	}

	return &EventServiceInit{
		eventRepo: eventRepo,
	}, nil
}

func (e *EventServiceInit) FindByUserID(userID int) ([]events.Events, error) {
	eventsList, err := e.eventRepo.FindByUserId(userID)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving events by user ID %s: %v", userID, err)
		return nil, err
	}
	return eventsList, nil
}

func (e *EventServiceInit) GetEventsWithGiftsByUserID(userID int) ([]helper.EventWithGiftsDTO, error) {
	eventsWithGifts, err := e.eventRepo.GetEventsWithGiftsByUserID(userID)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving events with gifts by user ID %d: %v", userID, err)
		return nil, err
	}
	return eventsWithGifts, nil
}

func (e *EventServiceInit) GetEventsWithReservationsWithUserId(userId int) ([]helper.EventWithReservationsDTO, error) {
	reservations, err := e.eventRepo.GetEventsWithReservationsWithUserId(userId)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving reservations for user ID %d: %v", userId, err)
		return nil, err
	}
	return reservations, nil
}

