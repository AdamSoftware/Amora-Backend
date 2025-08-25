package planrepository

import (
	"Amora/internal/Logging"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/plans"
	"errors"
	"gorm.io/gorm"
)

type PlanRepositoryInit struct {
	Core *CoreRepository.CoreRepositoryInit[*plans.Plans]
}

func NewPlansRepositoryInit(DB *gorm.DB) (*PlanRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*plans.Plans](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &PlanRepositoryInit{Core: coreRepo}, nil
}
