package events

type Events struct {
	Id          int    `gorm:"primaryKey;column:id"`
	Name        string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	DateOfEvent        string `gorm:"column:DateOfTheEvent"`
	IsActive    bool   `gorm:"column:is_active"`
	EventId      int    `gorm:"column:user_id"`
}


// TableName remains unchanged
func (Events) TableName() string {
	return "events"
}

func (u *Events) GetId() int {
	return u.Id
}

