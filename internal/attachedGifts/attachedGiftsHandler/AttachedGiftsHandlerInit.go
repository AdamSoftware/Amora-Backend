package attachedGiftsHandler

import ( 
	"Amora/internal/core/CoreHandler"
	"Amora/internal/attachedGifts"
	"Amora/internal/attachedGifts/attachedGiftsService"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"Amora/internal/helper"
	"strconv"
)

type AttachedGiftsHandlerInit struct {
	attachedGiftsService *attachedGiftsService.AttachedGiftsServiceInit
	coreHandler          *CoreHandler.CoreHandlerInit[*attachedGifts.AttachedGifts]
}

func NewattachedGiftsHandlerInit(
	coreHandler *CoreHandler.CoreHandlerInit[*attachedGifts.AttachedGifts],
	attachedGiftsService *attachedGiftsService.AttachedGiftsServiceInit,
) *AttachedGiftsHandlerInit {
	return &AttachedGiftsHandlerInit{
		attachedGiftsService: attachedGiftsService,
		coreHandler:          coreHandler,
	}
}

func (u *AttachedGiftsHandlerInit) FindById(w http.ResponseWriter, r *http.Request) {
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
		// error if the user just can't be found
    if entity == nil {
			  helper.WriteError(w, http.StatusNotFound, "AttachedGifts not found")
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(entity)
}

func (u *AttachedGiftsHandlerInit) Get(w http.ResponseWriter, r *http.Request) {
		entities, err := u.coreHandler.Get()
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		// this would should never happen but just in case
		if len(entities) == 0 { 
				helper.WriteError(w, http.StatusNotFound, "No users found")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}

func (u *AttachedGiftsHandlerInit) Delete(w http.ResponseWriter, r *http.Request) {
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

func (u *AttachedGiftsHandlerInit) GetWithEventIdWithGiftDetails(w http.ResponseWriter, r *http.Request) {
	eventIdStr := chi.URLParam(r, "eventId")
	if eventIdStr == "" {
		helper.WriteError(w, http.StatusBadRequest, "Missing event ID parameter")
		return
	}

	eventId, err := strconv.Atoi(eventIdStr)
	if err != nil || eventId <= 0 {
		helper.WriteError(w, http.StatusBadRequest, "Invalid event ID parameter")
		return
	}

	details, err := u.attachedGiftsService.GetWithEventIdWithGiftDetails(eventId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(details)
}

func (u *AttachedGiftsHandlerInit) InsertValidation(w http.ResponseWriter, r *http.Request) {
	var entity attachedGifts.AttachedGifts
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	createdEntity, err := u.attachedGiftsService.InsertValidation(&entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdEntity)
}

func (u *AttachedGiftsHandlerInit) UpdateValidation(w http.ResponseWriter, r *http.Request) {
	var entity attachedGifts.AttachedGifts
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updatedEntity, err := u.attachedGiftsService.UpdateValidation(&entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedEntity)
}


