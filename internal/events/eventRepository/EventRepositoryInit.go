package eventRepository


import (
	"Amora/internal/events"
	"Amora/internal/Logging"
	"Amora/internal/core/CoreRepository"
	"errors"
	"gorm.io/gorm"
)


type EventRepositoryInit struct {
	Core  *CoreRepository.CoreRepositoryInit[*events.Events]
}

func NewuserRepositoryInit(DB *gorm.DB) (*EventRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*events.Events](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &EventRepositoryInit{ Core : coreRepo }, nil
}


func (e *EventRepositoryInit) FindByUserId(userId int) ([]events.Events, error) {
	var eventsList []events.Events
	err := e.Core.DB.Where("user_id = ?", userId).Find(&eventsList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Logging.Logs.Warnf("No events found for user ID %d", userId)
			return nil, nil
		}
		Logging.Logs.Errorf("Error retrieving events for user ID %d: %v", userId, err)
		return nil, err
	}
	return eventsList, nil
}
