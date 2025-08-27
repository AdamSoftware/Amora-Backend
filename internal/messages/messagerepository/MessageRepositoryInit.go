package messagerepository

import (
	"Amora/internal/Logging"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/messages"
	"errors"
	"gorm.io/gorm"
)

type MessageRepositoryInit struct {
	Core *CoreRepository.CoreRepositoryInit[*messages.Messages]
}

func NewMessagesRepositoryInit(DB *gorm.DB) (*MessageRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*messages.Messages](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &MessageRepositoryInit{Core: coreRepo}, nil
}
