package eventHandler

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"Amora/internal/core/CoreHandler"
	"Amora/internal/events"
	"Amora/internal/events/eventService"
	"encoding/json"
	"strconv"
	"Amora/internal/helper"
)

type EventHandlerInit struct {
	  eventService *eventService.EventServiceInit
    coreHandler *CoreHandler.CoreHandlerInit[*events.Events]
}

func NewEventHandlerInit(coreHandler *CoreHandler.CoreHandlerInit[*events.Events], eventService *eventService.EventServiceInit) *EventHandlerInit {
    return &EventHandlerInit{
				eventService: eventService,
        coreHandler: coreHandler,
    }
}

func (u *EventHandlerInit) FindById(w http.ResponseWriter, r *http.Request) {
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
		// error if the event just can't be found
    if entity == nil {
			  helper.WriteError(w, http.StatusNotFound, "Event not found")
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(entity)
}

func (u *EventHandlerInit) Get(w http.ResponseWriter, r *http.Request) {
		entities, err := u.coreHandler.Get()
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		// this would should never happen but just in case
		if len(entities) == 0 { 
				helper.WriteError(w, http.StatusNotFound, "No events found")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}


func (u *EventHandlerInit) Insert(w http.ResponseWriter, r *http.Request) {
	var entity *events.Events
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

func (u *EventHandlerInit) Update(w http.ResponseWriter, r *http.Request) {
	var entity *events.Events
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


func (u *EventHandlerInit) Delete(w http.ResponseWriter, r *http.Request) {
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

func (u *EventHandlerInit) FindByUserId(w http.ResponseWriter, r *http.Request) {
		userIdStr := chi.URLParam(r, "userId") 
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			  helper.WriteError(w, http.StatusBadRequest, "Invalid User ID parameter")
				return
		}

		entities, err := u.eventService.FindByUserID(userId)
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		if len(entities) == 0 {
			  helper.WriteError(w, http.StatusAccepted, "No events found for this user")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}

func (u *EventHandlerInit) GetEventsWithGiftsWithUserId(w http.ResponseWriter, r *http.Request) {
		userIdStr := chi.URLParam(r, "userId") 
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			  helper.WriteError(w, http.StatusBadRequest, "Invalid User ID parameter")
				return
		}

		eventsWithGifts, err := u.eventService.GetEventsWithGiftsByUserID(userId)
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		if len(eventsWithGifts) == 0 {
			  helper.WriteError(w, http.StatusNotFound, "No events with gifts found for this user")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(eventsWithGifts)
}

func (u *EventHandlerInit) GetEventsWithReservationsWithUserId(w http.ResponseWriter, r *http.Request) {
		userIdStr := chi.URLParam(r, "userId") 
		userId, err := strconv.Atoi(userIdStr)
		if err != nil {
			  helper.WriteError(w, http.StatusBadRequest, "Invalid User ID parameter")
				return
		}

		reservations, err := u.eventService.GetEventsWithReservationsWithUserId(userId)
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		if len(reservations) == 0 {
			  helper.WriteError(w, http.StatusNotFound, "No reservations found for this user")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reservations)
}



