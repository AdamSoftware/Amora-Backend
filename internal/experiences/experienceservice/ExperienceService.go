package experienceservice

import (
	"Amora/internal/experiences"
)

type ExperienceService interface {
	GetAllActiveVenues() ([]experiences.Experiences, error)
}
