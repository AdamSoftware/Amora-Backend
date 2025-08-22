package giftservice

import ( 
	"Amora/internal/gifts"
)

type GiftService interface {
	GetAllActiveGifts() ([]gifts.Gifts, error)
}
