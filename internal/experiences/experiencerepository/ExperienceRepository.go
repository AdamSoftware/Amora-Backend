package experiencerepository

import (
	"Amora/internal/experiences"
)

type ExperienceRepository interface {
	GetAllActiveVenues() ([]experiences.Experiences, error)
}
