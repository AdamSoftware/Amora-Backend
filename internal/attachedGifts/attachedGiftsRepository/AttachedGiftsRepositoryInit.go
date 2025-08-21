package attachedGiftsRepository

import (
	"Amora/internal/Logging"
	"Amora/internal/attachedgifts"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/helper"
	"errors"
	"fmt"

	"gorm.io/gorm"
)


type AttachedGiftsRepositoryInit struct {
	Core  *CoreRepository.CoreRepositoryInit[*attachedgifts.AttachedGifts]
}

func NewattachedgiftsRepositoryInit(DB *gorm.DB) (*AttachedGiftsRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*attachedgifts.AttachedGifts](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &AttachedGiftsRepositoryInit{ Core : coreRepo }, nil
}

// UpdateValidation
func (a *AttachedGiftsRepositoryInit) UpdateValidation(ag *attachedgifts.AttachedGifts) (*attachedgifts.AttachedGifts, error) {

	// this is the var that the amount would be display for this is just to count if there where any gifts or events that match that id
	var count int64
	err := a.Core.DB.Raw(`
		SELECT COUNT(*) 
		FROM gifts g 
		WHERE gift_id = ? AND event_id = ?`,
		ag.GiftId, ag.EventId,
	).Scan(&count).Error

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

	if count == 0 {
		Logging.Logs.Warnf("No matching gift found for GiftId: %d and EventId: %d", ag.GiftId, ag.EventId)
		return nil, errors.New("no matching gift found for the provided GiftId and EventId")
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
func (a *AttachedGiftsRepositoryInit) InsertValidation(ag *attachedgifts.AttachedGifts) (*attachedgifts.AttachedGifts, error) {
	var count int64
	err := a.Core.DB.Raw(`
		SELECT COUNT(*) 
		FROM gifts g 
		WHERE gift_id = ? AND event_id = ?`,
		ag.GiftId, ag.EventId,
	).Scan(&count).Error

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

	if count == 0 {
		Logging.Logs.Warnf("No matching gift found for GiftId: %d and EventId: %d", ag.GiftId, ag.EventId)
		return nil, errors.New("no matching gift found for the provided GiftId and EventId")
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
		Price          float64
	}

	// first pull in the raw data the event and giftid 
	// pull in all the gift details from the gift id searching with this inside of the gift table
	err := a.Core.DB.Raw(`SELECT ag.id as attached_gift_id, ag.event_id as EventID, ag.gift_id as GiftID, ag.quantity as Quantity
	, g.name as GiftName, g.description as GiftDescription, g.image_url as ImageURL, g.price as Price 
	FROM attachedgifts ag
	JOIN gifts g on ag.gift_id = g.id
	where ag.event_id = ? and ag.is_active = true and g.is_active = true`, eventId).Scan(&row).Error

	if err != nil {
		Logging.Logs.Errorf("Error retrieving attached gifts with gift details for event ID %d: %v", eventId, err)
		return nil, fmt.Errorf("Error retrieving attached gifts with gift details for event ID %d: %v", eventId, err)
	}

	eventMap := make(map[int]*helper.AttachedGiftsWithGiftsDTO)

	for _, r := range row {
		if _, exists := eventMap[r.EventID]; !exists {
			eventMap[r.EventID] = &helper.AttachedGiftsWithGiftsDTO{
				Id:      r.AttachedGiftID,
				EventId: r.EventID,
				GiftId:  r.GiftID,
				GiftDetails:   []helper.GiftDTO{},
			}
		}

		eventMap[r.EventID].GiftDetails = append(eventMap[r.EventID].GiftDetails, helper.GiftDTO{
			Id:          r.GiftID,
			Name:        r.GiftName,
			Description: r.GiftDescription,
			ImageURL:    r.ImageURL,
			Price:       r.Price,
			Quantity:    r.Quantity,
		})
	}

	var result []helper.AttachedGiftsWithGiftsDTO
	for _, e := range eventMap {
		result = append(result, *e)
	}

	return result, nil
}




