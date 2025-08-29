package reminderhandler

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"Amora/internal/core/CoreHandler"
	"Amora/internal/reminders"
	"Amora/internal/reminders/reminderservice"
	"encoding/json"
	"strconv"
	"Amora/internal/helper"
)

type ReminderHandlerInit struct {
	  reminderService *reminderservice.ReminderServiceInit
    coreHandler *CoreHandler.CoreHandlerInit[*reminders.Reminders]
}

func NewReminderHandlerInit(coreHandler *CoreHandler.CoreHandlerInit[*reminders.Reminders], reminderService *reminderservice.ReminderServiceInit) *ReminderHandlerInit {
    return &ReminderHandlerInit{
				reminderService: reminderService,
        coreHandler: coreHandler,
    }
}


func (u *ReminderHandlerInit) FindById(w http.ResponseWriter, r *http.Request) {
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
		// error if the reminder just can't be found
    if entity == nil {
			  helper.WriteError(w, http.StatusNotFound, "Reminder not found")
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(entity)
}

func (u *ReminderHandlerInit) Get(w http.ResponseWriter, r *http.Request) {
		entities, err := u.coreHandler.Get()
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		// this would should never happen but just in case
		if len(entities) == 0 { 
				helper.WriteError(w, http.StatusNotFound, "No reminders found")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}


func (u *ReminderHandlerInit) Insert(w http.ResponseWriter, r *http.Request) {
	var entity *reminders.Reminders
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

func (u *ReminderHandlerInit) Update(w http.ResponseWriter, r *http.Request) {
	var entity *reminders.Reminders
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


func (u *ReminderHandlerInit) Delete(w http.ResponseWriter, r *http.Request) {
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


func (u *ReminderHandlerInit) FindByUserId(w http.ResponseWriter, r *http.Request) {
		userIdStr := chi.URLParam(r, "userId") 
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			  helper.WriteError(w, http.StatusBadRequest, "Invalid User ID parameter")
				return
		}

		entities, err := u.reminderService.FindByUserID(userId)
		// this is the internal error
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		// error if the reminder just can't be found
		if len(entities) == 0 {
			  helper.WriteError(w, http.StatusNotFound, "No reminders found for the user")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}

func (u *ReminderHandlerInit) FindByReservationId(w http.ResponseWriter, r *http.Request) {
		ReservationIdStr := chi.URLParam(r, "ReservationId") 
		ReservationId, err := strconv.Atoi(ReservationIdStr)
		if err != nil {
			  helper.WriteError(w, http.StatusBadRequest, "Invalid Reservation ID parameter")
				return
		}

		entities, err := u.reminderService.FindByReservationID(ReservationId)
		// this is the internal error
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		// error if the reminder just can't be found
		if len(entities) == 0 {
			  helper.WriteError(w, http.StatusNotFound, "No reminders found for the reservation")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}

