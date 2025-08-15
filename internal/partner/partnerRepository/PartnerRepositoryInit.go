package partnerRepository

import (
	"Amora/internal/Logging"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/partner"
	"errors"
	"gorm.io/gorm"
)


type PartnerRepositoryInit struct {
	Core *CoreRepository.CoreRepositoryInit[*partner.Partners]	
}

func NewuserRepositoryInit(DB *gorm.DB) (*PartnerRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*partner.Partners](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &PartnerRepositoryInit{ Core : coreRepo }, nil
}

func (p *PartnerRepositoryInit) FindByPhoneNumber(phoneNumber string) (*partner.Partners, error) {
	var partner partner.Partners
	err := p.Core.DB.Where("phone_number = ?", phoneNumber).First(&partner).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		Logging.Logs.Errorf("Error retrieving partner by phone number %s: %v", phoneNumber, err)
		return nil, err
	}
	return &partner, nil
}


func (p *PartnerRepositoryInit) FindByUserId(userId int) (*partner.Partners, error) {
	var partner partner.Partners
	err := p.Core.DB.Where("user_id = ?", userId).First(&partner).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		Logging.Logs.Errorf("Error retrieving partner by user ID %d: %v", userId, err)
		return nil, err
	}
	return &partner, nil
}
