package restaurantrepository

import (
	"Amora/internal/restaurants"
)


type RestaurantRepository interface {
	GetAllActiveRestaurants() ([]restaurants.Restaurants, error)
}
