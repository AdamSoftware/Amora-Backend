package eventHandler

import (
	"net/http"
)

type EventHandler interface {
	FindById(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Insert(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindByUserId(w http.ResponseWriter, r *http.Request)
	GetEventsWithGiftsWithUserId(w http.ResponseWriter, r *http.Request)
	GetEventsWithReservationsWithUserId(w http.ResponseWriter, r *http.Request)
}

