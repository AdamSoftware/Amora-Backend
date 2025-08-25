package restaurantrepository

import (
	"Amora/internal/Logging"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/restaurants"
	"errors"
	"gorm.io/gorm"
)

type RestaurantRepositoryInit struct {
	Core *CoreRepository.CoreRepositoryInit[*restaurants.Restaurants]
}

func NewRestaurantsRepositoryInit(DB *gorm.DB) (*RestaurantRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*restaurants.Restaurants](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &RestaurantRepositoryInit{Core: coreRepo}, nil
}

func (g *RestaurantRepositoryInit) GetAllActiveRestaurants() ([]restaurants.Restaurants, error) {
	var AllActiveRestaurants []restaurants.Restaurants

	err := g.Core.DB.Where("is_active = ?", true).Find(&AllActiveRestaurants).Error

	if err != nil{
		Logging.Logs.Errorf("Error retrieving active resturants: %v", err)
		return nil, err
	}

	// if the length is less then or equal to 0 then return nil
	if len(AllActiveRestaurants) <= 0 {
		Logging.Logs.Warnf("No active resturants found")
		return nil, nil
	}

	return AllActiveRestaurants, nil
}
