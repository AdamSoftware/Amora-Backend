package attachedGiftsRepository

import ( 
	"Amora/internal/attachedgifts"
	"Amora/internal/helper"
)

type AttachedGiftsRepository interface {
	// these are going to be deviating from the crud so not going to be coming from the core methods
	UpdateValidation(ag *attachedgifts.AttachedGifts) (*attachedgifts.AttachedGifts, error)
	InsertValidation(ag *attachedgifts.AttachedGifts) (*attachedgifts.AttachedGifts, error) 

	// this will get all details of the gifts to desplay them on the event page on the gifts page of the event 
	GetWithEventIdWithGiftDetails(eventId int) ([]helper.AttachedGiftsWithGiftsDTO, error)
}	

