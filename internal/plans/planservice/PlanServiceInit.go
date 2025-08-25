package planservice

import (
	"Amora/internal/Logging"
	"Amora/internal/plans/planrepository"
	"errors"
)

type PlanServiceInit struct {
	planRepo *planrepository.PlanRepositoryInit
}

func NewPlanServiceInit( planRepo *planrepository.PlanRepositoryInit) (*PlanServiceInit, error) {

	if planRepo == nil {
		Logging.Logs.Errorf("user repository is nil")
		return nil, errors.New("user repository is nil")
	}

	return &PlanServiceInit{
		planRepo: planRepo,
	}, nil
}
