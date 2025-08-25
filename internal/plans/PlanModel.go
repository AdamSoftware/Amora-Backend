package plans

type Plans struct {
	Id int `gorm:"primaryKey;column:id"`
	PlanType string `gorm:"column:plantype"`
	Price float64 `gorm:"column:price"`
}

// TableName remains unchanged
func (Plans) TableName() string {
	return "plans"
}

func (u *Plans) GetId() int {
	return u.Id
}
