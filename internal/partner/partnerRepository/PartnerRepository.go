package partnerRepository

import (
	"Amora/internal/partner"
)


type PartnerRepository interface {
	FindByPhoneNumber(phoneNumber string) (*partner.Partners, error)
	FindByUserId(userId int) ([]partner.Partners, error)
}



