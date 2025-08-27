package messages

type Messages struct {
	Id int `gorm:"primaryKey;column:id"`
	Userid int `gorm:"column:user_id"`
	Content string `gorm:"column:content"`
	Status string `gorm:"column:status"`
	SentAt string `gorm:"column:sent_at"`
	CreatedAt string `gorm:"column:created_at"`
	UpdatedAt string `gorm:"column:updated_at"`
}

// TableName remains unchanged
func (Messages) TableName() string {
	return "messages"
}

func (u *Messages) GetId() int {
	return u.Id
}
