package attachedGiftsService

import ( 
	"Amora/internal/attachedGifts"
	"Amora/internal/helper"
)

type AttachedGiftsService interface {
	// these are going to be deviating from the crud so not going to be coming from the core methods
	UpdateValidation(ag *attachedGifts.AttachedGifts) (*attachedGifts.AttachedGifts, error)
	InsertValidation(ag *attachedGifts.AttachedGifts) (*attachedGifts.AttachedGifts, error) 

	// this will get all details of the gifts to desplay them on the event page on the gifts page of the event 
	GetWithEventIdWithGiftDetails(eventId int) ([]helper.AttachedGiftsWithGiftsDTO, error)
}

