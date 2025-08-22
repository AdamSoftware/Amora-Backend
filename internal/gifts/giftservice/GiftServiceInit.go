package giftservice

import (
	"Amora/internal/Logging"
	"Amora/internal/gifts/giftrepository"
	"errors"
	"Amora/internal/gifts"
)

type GiftServiceInit struct {
	giftRepo *giftrepository.GiftRepositoryInit
}

func NewGiftServiceInit( giftRepo *giftrepository.GiftRepositoryInit) (*GiftServiceInit, error) {

	if giftRepo == nil {
		Logging.Logs.Errorf("user repository is nil")
		return nil, errors.New("user repository is nil")
	}

	return &GiftServiceInit{
		giftRepo: giftRepo,
	}, nil
}

func (g *GiftServiceInit) GetAllActiveGifts() ([]gifts.Gifts, error) {
	ActiveGifts, err := g.giftRepo.GetAllActiveGifts()
	if err != nil {
		Logging.Logs.Errorf("Error retrieving all active gifts: %v", err)
		return nil, err
	}
	return ActiveGifts, nil
}
