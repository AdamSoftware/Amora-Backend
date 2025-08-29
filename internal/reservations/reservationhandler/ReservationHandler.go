package reservationhandler

import (
	"net/http"
)

type ReservationHandler interface {
	FindById(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Insert(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetByUserId(w http.ResponseWriter, r *http.Request)
	GetByEventId(w http.ResponseWriter, r *http.Request)
}

