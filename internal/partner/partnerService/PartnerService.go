package partnerService

import (
	"Amora/internal/partner"
)


type PartnerService interface {
	FindByPhoneNumber(phoneNumber string) (*partner.Partners, error)
	FindByUserId(userId int) ([]partner.Partners, error)
	// will find all the partners that have a birthday on the givin date
	FindAllPartnersWithBirthday(SpicyContent int) ([]partner.Partners, error)
	FindAllPartnersWithAnniversary(SpicyContent int) ([]partner.Partners, error)

}


