package CoreRepository

import (
	"gorm.io/gorm"
	"Amora/internal/Logging"
	"Amora/internal/core"
	"fmt"
	"errors"
)

// this is to init the new repository with the DB connection
// I forgot the lowercase letters are not exported outside of the package in GO fucking stupid 
// why can't we just have public and private like WTF
type CoreRepositoryInit[T core.CoreModel] struct {
	DB  *gorm.DB
}

func NewCoreRepository[T core.CoreModel](DB *gorm.DB) (*CoreRepositoryInit[T], error) {

	// if the database was never found then return an error
	if DB == nil {
		err := errors.New("The database was never initialized")
		return nil, fmt.Errorf("Failed to create GenericRepo: %w", err)
	}


	// if everything looks ok for the database then just return the new repositrory and continue
	return &CoreRepositoryInit[T]{DB: DB}, nil
}

// FindById returns a single entity by id well it should 
func (r *CoreRepositoryInit[T]) FindById(id int) (T, error) {
	var entity T
	err := r.DB.First(&entity, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		Logging.Logs.Warnf("Entity with id %d not found", id)
		var zero T
		return zero, fmt.Errorf("Entity with id %d not found: %w", id, err)
	}

	return entity, nil
}

// Return all entities of type T
func (r *CoreRepositoryInit[T]) Get() ([]T, error) {
	var entities []T
	err := r.DB.Find(&entities).Error
	if err != nil {
		Logging.Logs.Errorf("Error retrieving entities: %v", err)
		return nil, fmt.Errorf("Error retrieving entities: %v", err)
	}

	return entities, nil
}


// Insert adds all new entitys to the database table
func (r *CoreRepositoryInit[T]) Insert(entity T) (T, error) {
    err := r.DB.Create(entity).Error // Don't pass a pointer because entity is already a pointer 
    if err != nil {
				Logging.Logs.Errorf("Error inserting entity: %v", err)
        var zero T
        return zero, fmt.Errorf("Error inserting entity: %v", err)
    }

    return entity, nil
}



// Update modifies an existing entity in the database
func (r *CoreRepositoryInit[T]) Update(entity T) (T, error) {
    Logging.Logs.Infof("Got to the Update: %+v", entity)

    // Log the state before saving
    Logging.Logs.Infof("Before save: %+v", entity)

    // Perform the update
    if err := r.DB.Save(&entity).Error; err != nil {
        var zero T
				Logging.Logs.Errorf("Error updating entity: %v", err)
        return zero, err
    }

    // Log the state after saving
    Logging.Logs.Infof("After save: %+v", entity)

    // Return the updated entity
    return entity, nil
}

// Delete removes an entity from the database by id
func (r *CoreRepositoryInit[T]) Delete(id int) error {
	var entity T
	err := r.DB.First(&entity, id).Error
	// check if the entity was found if not throw error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		Logging.Logs.Warnf("Entity with id %d not found for deletion", id)
		return fmt.Errorf("Entity with id %d not found: %w", id, err)
	} else if err != nil {
		// check to see if there was any other error found
		Logging.Logs.Errorf("Something happened: %v", err)
		return fmt.Errorf("Something happened: %v", err)
	}

	// if the database never delete the entity return error
	err = r.DB.Delete(&entity).Error
	if err != nil {
		Logging.Logs.Errorf("Error deleting entity: %v", err)
		return fmt.Errorf("Error deleting entity: %v", err)
	}

	return nil
}
