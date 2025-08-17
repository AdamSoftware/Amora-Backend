package partnerService

import ( 
	"Amora/internal/helper"
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


func (p *PartnerServiceInit) FindAllPartnersWithBirthday() ([]helper.PartnerMessage, error) {
	partners, err := p.partnerRepo.FindAllPartnersWithBirthday()
	if err != nil {
		Logging.Logs.Errorf("Error retrieving partners with birthday: %v", err)
		return nil, err
	}


	client := helper.GetOpenAiClient()
	if client == nil {
		Logging.Logs.Errorf("OpenAI client is not initialized")
		return nil, errors.New("OpenAI client is not initialized")
	}

	var messages []helper.PartnerMessage

	for _, partner := range partners {
		response, err := helper.SendOpenAiMessage(client, partner.Firstname, partner.Nickname, true)
		if err != nil {
			Logging.Logs.Errorf("Error generating message from OpenAI for partner ID %d: %v", partner.Id, err)
			continue // just skip the one that errored
		}

		messages = append(messages, helper.PartnerMessage{
			PartnerId:    partner.Id,
			PartnerName:  partner.Firstname,
			Nickname:     partner.Nickname,
			PhoneNumber:  partner.Phonenumber,
			UserId:      	partner.User_id,
			Message:      response.Message,
		})
	}

	return messages, nil
}

func (p *PartnerServiceInit) FindAllPartnersWithAnniversary() ([]helper.PartnerMessage, error){ 
	partners, err := p.partnerRepo.FindAllPartnersWithAnniversary()
	if err != nil {
		Logging.Logs.Errorf("Error retrieving partners with anniversary: %v", err)
		return nil, err
	}

	client := helper.GetOpenAiClient()
	if client == nil {
		Logging.Logs.Errorf("OpenAI client is not initialized")
		return nil, errors.New("OpenAI client is not initialized")
	}

	var messages []helper.PartnerMessage
	for _, partner := range partners {
		response, err := helper.SendOpenAiMessage(client, partner.Firstname, partner.Nickname, false)
		if err != nil {
			Logging.Logs.Errorf("Error generating message from OpenAI for partner ID %d: %v", partner.Id, err)
			continue // just skip the one that errored
		}

		messages = append(messages, helper.PartnerMessage{
			PartnerId:    partner.Id,
			PartnerName:  partner.Firstname,
			PhoneNumber:  partner.Phonenumber,
			Nickname:     partner.Nickname,
			UserId:      	partner.User_id,
			Message:      response.Message,
		})
	}
	return messages, nil
}





