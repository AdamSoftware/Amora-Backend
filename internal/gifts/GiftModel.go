package gifts

type Gifts struct {
	Id int `gorm:"primaryKey;column:id"`
	Name string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	ImageURL string `gorm:"column:image_url"`
	Price float64 `gorm:"column:price"`
	IsActive bool `gorm:"column:is_active"`
}

// TableName remains unchanged
func (Gifts) TableName() string {
	return "gifts"
}

func (u *Gifts) GetId() int {
	return u.Id
}

