package venueservice

import (
	"Amora/internal/venues"
)

type VeneueService interface {
	GetAllActiveVenues() ([]venues.Venues, error)
}


