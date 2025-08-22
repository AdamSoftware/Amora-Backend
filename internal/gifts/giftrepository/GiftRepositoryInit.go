package giftrepository

import (
	"Amora/internal/Logging"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/gifts"
	"errors"
	"gorm.io/gorm"
)

type GiftRepositoryInit struct {
	Core *CoreRepository.CoreRepositoryInit[*gifts.Gifts]
}

func NewGiftsRepositoryInit(DB *gorm.DB) (*GiftRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*gifts.Gifts](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &GiftRepositoryInit{Core: coreRepo}, nil
}

func (g *GiftRepositoryInit) GetAllActiveGifts() ([]gifts.Gifts, error) {
	var AllActiveGifts []gifts.Gifts

	err := g.Core.DB.Where("is_active = ?", true).Find(&AllActiveGifts).Error

	if err != nil{
		Logging.Logs.Errorf("Error retrieving active gifts: %v", err)
		return nil, err
	}

	// if the length is less then or equal to 0 then return nil
	if len(AllActiveGifts) <= 0 {
		Logging.Logs.Warnf("No active gifts found")
		return nil, nil
	}

	return AllActiveGifts, nil
}
