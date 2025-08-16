package partnerService

import (
	"Amora/internal/partner"
)


type PartnerService interface {
	FindByPhoneNumber(phoneNumber string) (*partner.Partners, error)
	FindByUserId(userId int) ([]partner.Partners, error)
}


