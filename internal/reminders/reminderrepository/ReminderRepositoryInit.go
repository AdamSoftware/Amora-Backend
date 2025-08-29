package reminderrepository

import (
	"Amora/internal/Logging"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/reminders"
	"errors"
	"gorm.io/gorm"
)

type ReminderRepositoryInit struct {
	Core  *CoreRepository.CoreRepositoryInit[*reminders.Reminders]
}

func NewreminderRepositoryInit(DB *gorm.DB) (*ReminderRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*reminders.Reminders](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &ReminderRepositoryInit{ Core : coreRepo }, nil
}

func (r *ReminderRepositoryInit) FindByUserID(userID int) ([]reminders.Reminders, error) {
	var remindersList []reminders.Reminders
	err := r.Core.DB.Where("user_id = ?", userID).Find(&remindersList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Logging.Logs.Warnf("No reminders found for user ID %s", userID)
			return nil, nil
		}
		Logging.Logs.Errorf("Error retrieving reminders for user ID %s: %v", userID, err)
		return nil, err
	}
	return remindersList, nil
}

func (r *ReminderRepositoryInit) FindByReservationID(ReservationID int) ([]reminders.Reminders, error) {
	var remindersList []reminders.Reminders
	err := r.Core.DB.Where("reservation_id = ?", ReservationID).Find(&remindersList).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Logging.Logs.Warnf("No reminders found for reservation ID %s", ReservationID)
			return nil, nil
		}
		Logging.Logs.Errorf("Error retrieving reminders for reservation ID %s: %v", ReservationID, err)
		return nil, err
	}
	return remindersList, nil
}

