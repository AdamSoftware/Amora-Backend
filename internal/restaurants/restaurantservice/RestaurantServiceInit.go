package restaurantservice

import (
	"Amora/internal/Logging"
	"Amora/internal/restaurants/restaurantrepository"
	"errors"
	"Amora/internal/restaurants"
)


type RestaurantServiceInit struct {
	restaurantRepo *restaurantrepository.RestaurantRepositoryInit
}

func NewRestaurantServiceInit( restaurantRepo *restaurantrepository.RestaurantRepositoryInit) (*RestaurantServiceInit, error) {

	if restaurantRepo == nil {
		Logging.Logs.Errorf("user repository is nil")
		return nil, errors.New("user repository is nil")
	}

	return &RestaurantServiceInit{
		restaurantRepo: restaurantRepo,
	}, nil
}

func (g *RestaurantServiceInit) GetAllActiveRestaurants() ([]restaurants.Restaurants, error) {
	ActiveRestaurants, err := g.restaurantRepo.GetAllActiveRestaurants()
	if err != nil {
		Logging.Logs.Errorf("Error retrieving all active restaurants: %v", err)
		return nil, err
	}
	return ActiveRestaurants, nil
}
