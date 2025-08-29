package reminderservice

import (
	"Amora/internal/reminders/reminderrepository"
	"Amora/internal/reminders"
	"Amora/internal/Logging"
	"errors"
)

type ReminderServiceInit struct {
	reminderRepo *reminderrepository.ReminderRepositoryInit
}

func NewReminderServiceInit( reminderRepo *reminderrepository.ReminderRepositoryInit) (*ReminderServiceInit, error) {

	if reminderRepo == nil {
		Logging.Logs.Errorf("reminder repository is nil")
		return nil, errors.New("reminder repository is nil")
	}

	return &ReminderServiceInit{
		reminderRepo: reminderRepo,
	}, nil
}

func (r *ReminderServiceInit) FindByUserID(userID int) ([]reminders.Reminders, error) {
	reminder, err := r.reminderRepo.FindByUserID(userID)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving reminder by user ID %s: %v", userID, err)
		return nil, err
	}
	return reminder, nil
}

func (r *ReminderServiceInit) FindByReservationID(ReservationID int) ([]reminders.Reminders, error) {
	reminder, err := r.reminderRepo.FindByReservationID(ReservationID)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving reminder by reservation ID %s: %v", ReservationID, err)
		return nil, err
	}
	return reminder, nil
}
