package venuerepository

import (
	"Amora/internal/venues"
)

type VenueRepository interface {
	GetAllActiveVenues() ([]venues.Venues, error)
}

