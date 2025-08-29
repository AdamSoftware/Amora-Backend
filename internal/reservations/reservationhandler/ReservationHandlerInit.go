package reservationhandler

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"Amora/internal/core/CoreHandler"
	"Amora/internal/reservations"
	"Amora/internal/reservations/reservationservice"
	"encoding/json"
	"strconv"
	"Amora/internal/helper"
)

type ReservationHandlerInit struct {
	  reservationService *reservationservice.ReservationServiceInit
    coreHandler *CoreHandler.CoreHandlerInit[*reservations.Reservations]
}

func NewReservationHandlerInit(coreHandler *CoreHandler.CoreHandlerInit[*reservations.Reservations], reservationService *reservationservice.ReservationServiceInit) *ReservationHandlerInit {
    return &ReservationHandlerInit{
				reservationService: reservationService,
        coreHandler: coreHandler,
    }
}


func (u *ReservationHandlerInit) FindById(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id") 
    id, err := strconv.Atoi(idStr)
    if err != nil {
			  helper.WriteError(w, http.StatusBadRequest, "Invalid ID parameter")
        return
    }

    entity, err := u.coreHandler.FindById(id)
		// this is the internal error
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
		// error if the reservation just can't be found
    if entity == nil {
			  helper.WriteError(w, http.StatusNotFound, "Reservation not found")
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(entity)
}

func (u *ReservationHandlerInit) Get(w http.ResponseWriter, r *http.Request) {
		entities, err := u.coreHandler.Get()
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		// this would should never happen but just in case
		if len(entities) == 0 { 
				helper.WriteError(w, http.StatusNotFound, "No reservations found")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}


func (u *ReservationHandlerInit) Insert(w http.ResponseWriter, r *http.Request) {
	var entity *reservations.Reservations
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	createdEntity, err := u.coreHandler.Insert(&entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdEntity)
}

func (u *ReservationHandlerInit) Update(w http.ResponseWriter, r *http.Request) {
	var entity *reservations.Reservations
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updatedEntity, err := u.coreHandler.Update(&entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedEntity)
}


func (u *ReservationHandlerInit) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		helper.WriteError(w, http.StatusBadRequest, "Missing ID parameter")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		helper.WriteError(w, http.StatusBadRequest, "Invalid ID parameter")
		return
	}

	err = u.coreHandler.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (u *ReservationHandlerInit) GetByUserId(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "userId") 
		id, err := strconv.Atoi(idStr)
		if err != nil {
			  helper.WriteError(w, http.StatusBadRequest, "Invalid User ID parameter")
				return
		}

		entities, err := u.reservationService.GetByUserId(id)
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		// this would should never happen but just in case
		if len(entities) == 0 { 
				helper.WriteError(w, http.StatusNotFound, "No reservations found for the user")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}

func (u *ReservationHandlerInit) GetByEventId(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "eventId") 
		id, err := strconv.Atoi(idStr)
		if err != nil {
			  helper.WriteError(w, http.StatusBadRequest, "Invalid Event ID parameter")
				return
		}

		entities, err := u.reservationService.GetByEventId(id)
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		// this would should never happen but just in case
		if len(entities) == 0 { 
				helper.WriteError(w, http.StatusNotFound, "No reservations found for the event")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}

