package CoreHandler

import (
	"Amora/internal/Logging"
	"Amora/internal/core"
	"Amora/internal/core/CoreService"
	"fmt"
	"errors"
)

type CoreHandlerInit[T core.CoreModel] struct {
	Service *CoreService.CoreServiceInit[T]
}

func NewCoreHandlerInit[T core.CoreModel](service *CoreService.CoreServiceInit[T]) (*CoreHandlerInit[T], error) {
	// if the service was never given then return an error and write the error to the logs table
	if service == nil {
		err := errors.New("The service was never initialized")
		Logging.Logs.Errorf("Failed to create CoreHandler: %v", err)
		return nil, fmt.Errorf("Failed to create CoreHandler: %w", err)
	}

	return &CoreHandlerInit[T]{Service: service}, nil
}

// FindById handler this will act a core method for all other handlers
func (h *CoreHandlerInit[T]) FindById(id int) (*T, error) {
	entity, err := h.Service.FindById(id)
	if err != nil {
		Logging.Logs.Errorf("Error finding entity with ID %d: %v", id, err)
		return nil, fmt.Errorf("error finding entity with ID %d: %w", id, err)
	}
	return &entity, nil
}

// Get handler will act as a core method for all the otehr get handler methods on all other tables
func (h *CoreHandlerInit[T]) Get() ([]*T, error) {
	entities, err := h.Service.Get()
	if err != nil {
		Logging.Logs.Errorf("Error retrieving entities: %v", err)
		return nil, fmt.Errorf("error retrieving entities: %w", err)
	}

	var result []*T
	// this will append all the entities to the results slice so that I can return the whole slice to the user/caller
	for i := range entities {
		result = append(result, &entities[i])
	}
	return result, nil
}

// Insert handler
func (h *CoreHandlerInit[T]) Insert(entity *T) (*T, error) {
	createdEntity, err := h.Service.Insert(*entity)
	if err != nil {
		Logging.Logs.Errorf("Error inserting entity: %v", err)
		return nil, fmt.Errorf("error inserting entity: %w", err)
	}
	return &createdEntity, nil
}

// Update handler
func (h *CoreHandlerInit[T]) Update(entity *T) (*T, error) {
	updatedEntity, err := h.Service.Update(*entity)
	if err != nil {
		Logging.Logs.Errorf("Error updating entity: %v", err)
		return nil, fmt.Errorf("error updating entity: %w", err)
	}
	return &updatedEntity, nil
}

// Delete handler
func (h *CoreHandlerInit[T]) Delete(id int) error {
	err := h.Service.Delete(id)
	if err != nil {
		Logging.Logs.Errorf("Error deleting entity with ID %d: %v", id, err)
		return fmt.Errorf("error deleting entity with ID %d: %w", id, err)
	}
	return nil
}

