package experiencerepository

import (
	"Amora/internal/Logging"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/experiences"
	"errors"
	"gorm.io/gorm"
)

type ExperienceRepositoryInit struct {
	Core *CoreRepository.CoreRepositoryInit[*experiences.Experiences]
}

func NewExperiencesRepositoryInit(DB *gorm.DB) (*ExperienceRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*experiences.Experiences](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &ExperienceRepositoryInit{Core: coreRepo}, nil
}

func (g *ExperienceRepositoryInit) GetAllActiveExperiences() ([]experiences.Experiences, error) {
	var AllActiveExperiences []experiences.Experiences

	err := g.Core.DB.Where("is_active = ?", true).Find(&AllActiveExperiences).Error

	if err != nil{
		Logging.Logs.Errorf("Error retrieving active resturants: %v", err)
		return nil, err
	}

	// if the length is less then or equal to 0 then return nil
	if len(AllActiveExperiences) <= 0 {
		Logging.Logs.Warnf("No active resturants found")
		return nil, nil
	}

	return AllActiveExperiences, nil
}
