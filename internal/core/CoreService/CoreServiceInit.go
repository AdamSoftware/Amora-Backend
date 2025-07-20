package CoreService

import (
	"Amora/internal/core"
	"Amora/internal/Core/CoreRepository"
	"Amora/internal/Logging"
	"fmt"
)

type CoreServiceInit[T core.CoreModel] struct {
	repo *CoreRepository.CoreRepositoryInit[T]
}

func NewCoreServiceInit[T core.CoreModel](repo *CoreRepository.CoreRepositoryInit[T]) *CoreServiceInit[T] {
	return &CoreServiceInit[T]{repo: repo}
}

// FindById retrieves an entity by its ID
func (s *CoreServiceInit[T]) FindById(id int) (T, error) {
	entity, err := s.repo.FindById(id)
	if err != nil {
		var zero T
		return zero, fmt.Errorf("error finding entity with ID %d: %v", id, err)
	}
	return entity, nil
}

// Get retrieves all entities
func (s *CoreServiceInit[T]) Get() ([]T, error) {
	entities, err := s.repo.Get()
	if err != nil {
		Logging.Logs.Error(fmt.Sprintf("Error retrieving entities: %v", err))
		return nil, err
	}

	return entities, nil
}

// Insert adds a new entity
func (s *CoreServiceInit[T]) Insert(entity T) (T, error) {
	createdEntity, err := s.repo.Insert(entity)
	if err != nil {
		Logging.Logs.Error(fmt.Sprintf("Error inserting entity: %v", err))
		var zero T
		return zero, err
	}
	return createdEntity, nil
}

// Update modifies an existing entity
func (s *CoreServiceInit[T]) Update(entity T) (T, error) {
	updatedEntity, err := s.repo.Update(entity)
	if err != nil {
		Logging.Logs.Error(fmt.Sprintf("Error updating entity: %v", err))
		return updatedEntity, err
	}
	return updatedEntity, nil
}

// Delete removes an entity by its ID
func (s *CoreServiceInit[T]) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		Logging.Logs.Error(fmt.Sprintf("Error deleting entity with ID %d: %v", id, err))
		return err
	}
	return nil
}
