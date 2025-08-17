package partnerRepository

import (
	"time"
	"Amora/internal/Logging"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/partner"
	"errors"
	"gorm.io/gorm"
)


type PartnerRepositoryInit struct {
	Core *CoreRepository.CoreRepositoryInit[*partner.Partners]	
}

func NewPartnerRepositoryInit(DB *gorm.DB) (*PartnerRepositoryInit, error) {
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


func (p *PartnerRepositoryInit) FindByUserId(userId int) ([]partner.Partners, error) {
	// this is a slice to handle the possibility for multiple partners for less bugs looking forward.
	var partner []partner.Partners
	err := p.Core.DB.Where("user_id = ?", userId).First(&partner).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		Logging.Logs.Errorf("Error retrieving partner by user ID %d: %v", userId, err)
		return nil, err
	}
	return partner, nil
}


// FindAllPartnersWith Birthday will give a list back of all partners that have a birthday on the given date
// this will only look at the month and day to check
func (p *PartnerRepositoryInit) FindAllPartnersWithBirthday() ([]partner.Partners, error) {
	var partners []partner.Partners

	today := time.Now()
	month := int(today.Month())
	day := today.Day()

	err := p.Core.DB.
		Where("EXTRACT(MONTH FROM birthdate) = ? AND EXTRACT(DAY FROM birthdate) = ?", month, day).
		Find(&partners).Error

	if err != nil {
		Logging.Logs.Errorf("Error retrieving partners with birthday on %02d-%02d: %v", month, day, err)
		return nil, err
	}

	return partners, nil
}

// FindAllPartnersWithAnniversary will give a list back of all partners that have an anniversary on the given date
func (p *PartnerRepositoryInit) FindAllPartnersWithAnniversary() ([]partner.Partners, error) {
	var partners []partner.Partners

	today := time.Now()
	month := int(today.Month())
	day := today.Day()

	err := p.Core.DB.
		Where("EXTRACT(MONTH FROM anniversary) = ? AND EXTRACT(DAY FROM anniversary) = ?", month, day).
		Find(&partners).Error

	if err != nil {
		Logging.Logs.Errorf("Error retrieving partners with anniversary on %02d-%02d: %v", month, day, err)
		return nil, err
	}

	return partners, nil
}

