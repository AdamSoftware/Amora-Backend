package attachedGifts

type AttachedGifts struct {
	Id int `gorm:"primaryKey;column:id"`
	EventId int `gorm:"column:event_id"`
	GiftId int `gorm:"column:gift_id"`
	Quantity int `gorm:"column:quantity"`
	IsActive bool `gorm:"column:is_active"`
}


// TableName remains unchanged
func (AttachedGifts) TableName() string {
	return "attachedgifts"
}

func (u *AttachedGifts) GetId() int {
	return u.Id
}



