package reservations

import (
	"time"
)

type Reservations struct {
	Id         int    `gorm:"primaryKey;column:id"`
	UserId     int    `gorm:"column:user_id"`
	PartnerId	int    `gorm:"column:partner_id"`
	ResterantId int    `gorm:"column:restaurant_id"`
	EventId	 int    `gorm:"column:event_id"`
	VenueId	 int    `gorm:"column:venue_id"`
	ExperienceId int    `gorm:"column:experience_id"`
	ReservationTime time.Time `gorm:"column:reservation_time"`
	PartySize  int    `gorm:"column:party_size"`
	SpecialRequests string `gorm:"column:special_requests"`
}

// TableName remains unchanged
func (Reservations) TableName() string {
	return "reservations"
}

func (u *Reservations) GetId() int {
	return u.Id
}

