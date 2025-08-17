package partnerRepository

import (
	"Amora/internal/partner"
)

type PartnerRepository interface {
	FindByPhoneNumber(phoneNumber string) (*partner.Partners, error)
	FindByUserId(userId int) ([]partner.Partners, error)
	// will find all the partners that have a birthday on the givin date
	FindAllPartnersWithBirthday() ([]partner.Partners, error)
	FindAllPartnersWithAnniversary() ([]partner.Partners, error)
}



