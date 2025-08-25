package experienceservice

import (
	"Amora/internal/Logging"
	"Amora/internal/experiences/experiencerepository"
	"errors"
	"Amora/internal/experiences"
)

type ExperienceServiceInit struct {
	experienceRepo *experiencerepository.ExperienceRepositoryInit
}

func NewExperienceServiceInit( experienceRepo *experiencerepository.ExperienceRepositoryInit) (*ExperienceServiceInit, error) {

	if experienceRepo == nil {
		Logging.Logs.Errorf("user repository is nil")
		return nil, errors.New("user repository is nil")
	}

	return &ExperienceServiceInit{
		experienceRepo: experienceRepo,
	}, nil
}

func (g *ExperienceServiceInit) GetAllActiveExperiences() ([]experiences.Experiences, error) {
	ActiveExperiences, err := g.experienceRepo.GetAllActiveExperiences()
	if err != nil {
		Logging.Logs.Errorf("Error retrieving all active experiences: %v", err)
		return nil, err
	}
	return ActiveExperiences, nil
}
