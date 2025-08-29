package reminderservice

import (
	"Amora/internal/reminders"
)

type ReminderService interface {
	FindByUserID(userID int) ([]reminders.Reminders, error)
	FindByReservationID(reservationID int) ([]reminders.Reminders, error)
}
