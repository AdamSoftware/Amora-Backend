package venueservice

import (
	"Amora/internal/Logging"
	"Amora/internal/venues/venuerepository"
	"errors"
	"Amora/internal/venues"
)

type VenueServiceInit struct {
	venueRepo *venuerepository.VenueRepositoryInit
}

func NewVenueServiceInit( venueRepo *venuerepository.VenueRepositoryInit) (*VenueServiceInit, error) {

	if venueRepo == nil {
		Logging.Logs.Errorf("user repository is nil")
		return nil, errors.New("user repository is nil")
	}

	return &VenueServiceInit{
		venueRepo: venueRepo,
	}, nil
}

func (g *VenueServiceInit) GetAllActiveVenues() ([]venues.Venues, error) {
	ActiveVenues, err := g.venueRepo.GetAllActiveVenues()
	if err != nil {
		Logging.Logs.Errorf("Error retrieving all active venues: %v", err)
		return nil, err
	}
	return ActiveVenues, nil
}
