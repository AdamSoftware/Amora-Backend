package attachedGiftsService

import ( 
	"Amora/internal/attachedgifts"
	"Amora/internal/helper"
	"Amora/internal/attachedgifts/attachedGiftsRepository"
	"Amora/internal/Logging"
	"fmt"
	"errors"
)

type AttachedGiftsServiceInit struct {
	attachedGiftsRepo *attachedGiftsRepository.AttachedGiftsRepositoryInit
}

func NewAttachedGiftsServiceInit( attachedGiftsRepo *attachedGiftsRepository.AttachedGiftsRepositoryInit) (*AttachedGiftsServiceInit, error) {

	if attachedGiftsRepo == nil {
		Logging.Logs.Errorf("attachedGifts repository is nil")
		return nil, errors.New("attachedGifts repository is nil")
	}

	return &AttachedGiftsServiceInit{
		attachedGiftsRepo: attachedGiftsRepo,
	}, nil
}

func (a *AttachedGiftsServiceInit) UpdateValidation(ag *attachedgifts.AttachedGifts) (*attachedgifts.AttachedGifts, error) {
	if ag == nil {
		Logging.Logs.Errorf("AttachedGifts is nil")
		return nil, errors.New("attachedGifts is nil")
	}

	updatedGift, err := a.attachedGiftsRepo.UpdateValidation(ag)
	if err != nil {
		Logging.Logs.Errorf("Error updating attached gift: %v", err)
		return nil, err
	}

	return updatedGift, nil
}

func (a *AttachedGiftsServiceInit) InsertValidation(ag *attachedgifts.AttachedGifts) (*attachedgifts.AttachedGifts, error) {
	if ag == nil {
		Logging.Logs.Errorf("AttachedGifts is nil")
		return nil, errors.New("attachedGifts is nil")
	}

	insertedGift, err := a.attachedGiftsRepo.InsertValidation(ag)
	if err != nil {
		Logging.Logs.Errorf("Error inserting attached gift: %v", err)
		return nil, err
	}

	return insertedGift, nil
}


func (a *AttachedGiftsServiceInit) GetWithEventIdWithGiftDetails(eventId int) ([]helper.AttachedGiftsWithGiftsDTO, error) {
	if eventId <= 0 {
		Logging.Logs.Errorf("Invalid formate for eventId: %d, Cannot be a nagative number", eventId)
		return nil, fmt.Errorf("invalid eventId: %d", eventId)
	}

	attachedGiftsWithDetails, err := a.attachedGiftsRepo.GetWithEventIdWithGiftDetails(eventId)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving attached gifts with details for eventId %d: %v", eventId, err)
		return nil, err
	}
	return attachedGiftsWithDetails, nil
}
