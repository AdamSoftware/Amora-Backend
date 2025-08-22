package gifthandler

import (
	"net/http"
)

type GiftHandler interface {
	FindById(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Insert(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetAllActiveGifts(w http.ResponseWriter, r *http.Request)
}
