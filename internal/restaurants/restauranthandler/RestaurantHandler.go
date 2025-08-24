package restauranthandler

import (
	"net/http"
)

type RestaurantHandler interface {
	FindById(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Insert(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetAllActiveRestaurants(w http.ResponseWriter, r *http.Request)
}


