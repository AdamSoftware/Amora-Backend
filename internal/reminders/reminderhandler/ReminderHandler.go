package reminderhandler

import (
	"net/http"
)

type ReminderHandler interface {
	FindById(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Insert(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindByUserId(w http.ResponseWriter, r *http.Request)
	FindByReservationId(w http.ResponseWriter, r *http.Request)
}
