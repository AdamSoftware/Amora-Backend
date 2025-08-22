package giftrepository

import (
	"Amora/internal/gifts"
)

type GiftRepository interface {
	GetAllActiveGifts() ([]gifts.Gifts, error)
}


