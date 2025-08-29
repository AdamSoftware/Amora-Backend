package reminderrepository

import (
	"Amora/internal/reminders"
)

type ReminderRepository interface {
	FindByUserID(userID int) (*reminders.Reminders, error)
	FindByReservationID(reservationID int) ([]reminders.Reminders, error)
}
