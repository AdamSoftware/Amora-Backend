package attachedGiftsRepository

import (
	"Amora/internal/Logging"
	"Amora/internal/attachedGifts"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/helper"
	"errors"
	"fmt"

	"gorm.io/gorm"
)


type AttachedGiftsRepositoryInit struct {
	Core  *CoreRepository.CoreRepositoryInit[*attachedGifts.AttachedGifts]
}

func NewattachedgiftsRepositoryInit(DB *gorm.DB) (*AttachedGiftsRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*attachedGifts.AttachedGifts](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &AttachedGiftsRepositoryInit{ Core : coreRepo }, nil
}

// UpdateValidation
func (a *AttachedGiftsRepositoryInit) UpdateValidation(ag *attachedGifts.AttachedGifts) (*attachedGifts.AttachedGifts, error) {

	var counts struct {
    GiftCount  int64
    EventCount int64
	}

	err := a.Core.DB.Raw(`
			SELECT 
					(SELECT COUNT(*) FROM gifts WHERE id = ? AND is_active = true) AS gift_count,
					(SELECT COUNT(*) FROM events WHERE id = ? AND is_active = true) AS event_count
			`, ag.GiftId, ag.EventId).Scan(&counts).Error


	if counts.GiftCount == 0 {
		Logging.Logs.Warnf("Gift not found or inactive for GiftId: %d", ag.GiftId)
		return nil,errors.New("gift not found or inactive")
	}

	if counts.EventCount == 0 {
		Logging.Logs.Warnf("Event not found or inactive for EventId: %d", ag.EventId)
		return nil,errors.New("event not found or inactive")
	}

	if err != nil {
		Logging.Logs.Errorf("Error checking gift and event existence: %v", err)
		return nil, errors.New("error checking gift and event existence")
	}

	if ag.Quantity <= 0 {
		Logging.Logs.Warnf("Quantity must be greater than zero: %d", ag.Quantity)
		return nil, errors.New("quantity must be greater than zero")
	}

	if ag.GiftId <= 0 {
		Logging.Logs.Warnf("Invalid GiftId: %d", ag.GiftId)
		return nil, errors.New("invalid GiftId")
	}

	if ag.EventId <= 0 {
		Logging.Logs.Warnf("Invalid EventId: %d", ag.EventId)
		return nil, errors.New("invalid EventId")
	}

	// save everything if there is no error save to DB
	err = a.Core.DB.Save(ag).Error
	if err != nil {
		Logging.Logs.Errorf("Error saving attached gift: %v", err)
		return nil, errors.New("error saving attached gift")
	}

	return ag, nil 

}


// insertValidation
func (a *AttachedGiftsRepositoryInit) InsertValidation(ag *attachedGifts.AttachedGifts) (*attachedGifts.AttachedGifts, error) {
	var counts struct {
    GiftCount  int64
    EventCount int64
	}

	err := a.Core.DB.Raw(`
			SELECT 
					(SELECT COUNT(*) FROM gifts WHERE id = ? AND is_active = true) AS gift_count,
					(SELECT COUNT(*) FROM events WHERE id = ? AND is_active = true) AS event_count
			`, ag.GiftId, ag.EventId).Scan(&counts).Error


	if counts.GiftCount == 0 {
		Logging.Logs.Warnf("Gift not found or inactive for GiftId: %d", ag.GiftId)
		return nil,errors.New("gift not found or inactive")
	}

	if counts.EventCount == 0 {
		Logging.Logs.Warnf("Event not found or inactive for EventId: %d", ag.EventId)
		return nil,errors.New("event not found or inactive")
	}

	if err != nil {
		Logging.Logs.Errorf("Error checking gift and event existence: %v", err)
		return nil, errors.New("error checking gift and event existence")
	}

	if ag.Quantity <= 0 {
		Logging.Logs.Warnf("Quantity must be greater than zero: %d", ag.Quantity)
		return nil, errors.New("quantity must be greater than zero")
	}

	if ag.GiftId <= 0 {
		Logging.Logs.Warnf("Invalid GiftId: %d", ag.GiftId)
		return nil, errors.New("invalid GiftId")
	}

	if ag.EventId <= 0 {
		Logging.Logs.Warnf("Invalid EventId: %d", ag.EventId)
		return nil, errors.New("invalid EventId")
	}


	// Create the thing if everything is AOK
	err = a.Core.DB.Create(ag).Error
	if err != nil {
		Logging.Logs.Errorf("Error saving attached gift: %v", err)
		return nil, errors.New("error saving attached gift")
	}

	return ag, nil 
}

func (a *AttachedGiftsRepositoryInit) GetWithEventIdWithGiftDetails(eventId int) ([]helper.AttachedGiftsWithGiftsDTO, error) {
	var row []struct {
		AttachedGiftID int 
		EventID        int 
		GiftID         int 
		Quantity       int 
		GiftName       string 
		GiftDescription string 
		ImageURL       string 
		AttachedGiftIsActive bool
		Price          float64
	}

	// first pull in the raw data the event and giftid 
	// pull in all the gift details from the gift id searching with this inside of the gift table
	err := a.Core.DB.Raw(`
		SELECT 
			ag.id as attached_gift_id,
			ag.event_id as event_id,
			ag.gift_id as gift_id,
			ag.quantity as quantity,
			ag.is_active as attached_gift_is_active,
			g.name as gift_name,
			g.description as gift_description,
			g.image_url as image_url,
			g.price as Price
		FROM attachedgifts ag
		JOIN gifts g on ag.gift_id = g.id
		WHERE ag.event_id = ? 
		`, eventId).Scan(&row).Error

	if err != nil {
		Logging.Logs.Errorf("Error retrieving attached gifts with gift details for event ID %d: %v", eventId, err)
		return nil, fmt.Errorf("Error retrieving attached gifts with gift details for event ID %d: %v", eventId, err)
	}

	
	attachedGiftMap := make(map[int]*helper.AttachedGiftsWithGiftsDTO)

	for _, r := range row {
		if _, exists := attachedGiftMap[r.AttachedGiftID]; !exists {
			attachedGiftMap[r.AttachedGiftID] = &helper.AttachedGiftsWithGiftsDTO{
				Id:          r.AttachedGiftID,
				EventId:     r.EventID,
				GiftId:      r.GiftID,
				Quantity:    r.Quantity,
				IsActive:    r.AttachedGiftIsActive,
				GiftDetails: []helper.GiftDTO{},
			}
		}

		attachedGiftMap[r.AttachedGiftID].GiftDetails = append(attachedGiftMap[r.AttachedGiftID].GiftDetails, helper.GiftDTO{
			Id:          r.GiftID,
			Name:        r.GiftName,
			Description: r.GiftDescription,
			ImageURL:    r.ImageURL,
			Price:       r.Price,
		})
	}


	var result []helper.AttachedGiftsWithGiftsDTO
	for _, e := range attachedGiftMap {
		result = append(result, *e)
	}

	return result, nil
}





