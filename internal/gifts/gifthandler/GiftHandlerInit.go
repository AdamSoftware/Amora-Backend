package gifthandler

import (
	"Amora/internal/core/CoreHandler"
	"Amora/internal/gifts"
	"Amora/internal/gifts/giftservice"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"Amora/internal/helper"
	"strconv"
)

type GiftHandlerInit struct {
	  giftsService *giftservice.GiftServiceInit
    coreHandler *CoreHandler.CoreHandlerInit[*gifts.Gifts]
}

func NewGiftHandlerInit(coreHandler *CoreHandler.CoreHandlerInit[*gifts.Gifts], giftsService *giftservice.GiftServiceInit) *GiftHandlerInit {
    return &GiftHandlerInit{
				giftsService: giftsService,
        coreHandler: coreHandler,
    }
}

func (u *GiftHandlerInit) FindById(w http.ResponseWriter, r *http.Request) {
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
		// error if the gifts just can't be found
    if entity == nil {
			  helper.WriteError(w, http.StatusNotFound, "Gift not found")
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(entity)
}

func (u *GiftHandlerInit) Get(w http.ResponseWriter, r *http.Request) {
		entities, err := u.coreHandler.Get()
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		// this would should never happen but just in case
		if len(entities) == 0 { 
				helper.WriteError(w, http.StatusNotFound, "No giftss found")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}


func (u *GiftHandlerInit) Insert(w http.ResponseWriter, r *http.Request) {
	var entity *gifts.Gifts
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

func (u *GiftHandlerInit) Update(w http.ResponseWriter, r *http.Request) {
	var entity *gifts.Gifts
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


func (u *GiftHandlerInit) Delete(w http.ResponseWriter, r *http.Request) {
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

func (u *GiftHandlerInit) GetAllActiveGifts(w http.ResponseWriter, r *http.Request) {
	gifts, err := u.giftsService.GetAllActiveGifts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(gifts)
}

