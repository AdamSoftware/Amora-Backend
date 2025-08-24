package restaurantservice

import (
	"Amora/internal/restaurants"
)


type RestaurantService interface {
	GetAllActiveRestaurants() ([]restaurants.Restaurants, error)
}
	
