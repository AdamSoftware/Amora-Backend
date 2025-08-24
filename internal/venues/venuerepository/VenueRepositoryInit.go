package venuerepository

import (
	"Amora/internal/Logging"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/venues"
	"errors"
	"gorm.io/gorm"
)


type VenueRepositoryInit struct {
	Core *CoreRepository.CoreRepositoryInit[*venues.Venues]
}

func NewVenuesRepositoryInit(DB *gorm.DB) (*VenueRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*venues.Venues](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &VenueRepositoryInit{Core: coreRepo}, nil
}

func (g *VenueRepositoryInit) GetAllActiveVenues() ([]venues.Venues, error) {
	var AllActiveVenues []venues.Venues

	err := g.Core.DB.Where("is_active = ?", true).Find(&AllActiveVenues).Error

	if err != nil{
		Logging.Logs.Errorf("Error retrieving active resturants: %v", err)
		return nil, err
	}

	// if the length is less then or equal to 0 then return nil
	if len(AllActiveVenues) <= 0 {
		Logging.Logs.Warnf("No active resturants found")
		return nil, nil
	}

	return AllActiveVenues, nil
}
