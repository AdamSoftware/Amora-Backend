package attachedGiftsHandler

import (
	"net/http"
)

type AttachedGiftsHandler interface {
	FindById(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	//Insert(w http.ResponseWriter, r *http.Request)
	//Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetWithEventIdWithGiftDetails(w http.ResponseWriter, r *http.Request)
	InsertValidation(w http.ResponseWriter, r *http.Request)
	UpdateValidation(w http.ResponseWriter, r *http.Request)
}
	
