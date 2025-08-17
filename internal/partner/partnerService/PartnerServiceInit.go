package partnerService

import ( 
	"Amora/internal/partner"
	"Amora/internal/partner/partnerRepository"
	"Amora/internal/Logging"
	"errors"
)


type PartnerServiceInit struct {
	partnerRepo *partnerRepository.PartnerRepositoryInit
}

func NewPartnerServiceInit( partnerRepo *partnerRepository.PartnerRepositoryInit) (*PartnerServiceInit, error) {

	if partnerRepo == nil {
		Logging.Logs.Errorf("partner repository is nil")
		return nil, errors.New("partner repository is nil")
	}

	return &PartnerServiceInit{
		partnerRepo: partnerRepo,
	}, nil
}

func (p *PartnerServiceInit) FindByPhoneNumber(phoneNumber string) (*partner.Partners, error) {
	Partner, err := p.partnerRepo.FindByPhoneNumber(phoneNumber)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving partner by phone number %s: %v", phoneNumber, err)
		return nil, err
	}
	return Partner, nil
}

func (p *PartnerServiceInit) FindByUserId(userId int) ([]partner.Partners, error) {
	Partners, err := p.partnerRepo.FindByUserId(userId)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving partners by user ID %d: %v", userId, err)
		return nil, err
	}
	return Partners, nil
}


func (p *PartnerServiceInit) FindAllPartnersWithBirthday() ([]partner.Partners, error) {
	partners, err := p.partnerRepo.FindAllPartnersWithBirthday()
	if err != nil {
		Logging.Logs.Errorf("Error retrieving partners with birthday: %v", err)
		return nil, err
	}
	return partners, nil
}

func (p *PartnerServiceInit) FindAllPartnersWithAnniversary() ([]partner.Partners, error) {
	partners, err := p.partnerRepo.FindAllPartnersWithAnniversary()
	if err != nil {
		Logging.Logs.Errorf("Error retrieving partners with anniversary: %v", err)
		return nil, err
	}
	return partners, nil
}
