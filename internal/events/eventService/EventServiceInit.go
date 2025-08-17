package eventService

import (
	"Amora/internal/events/eventRepository"
	"Amora/internal/events"
	"Amora/internal/Logging"
	"errors"
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

