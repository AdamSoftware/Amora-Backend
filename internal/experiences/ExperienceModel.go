package experiences

type Experiences struct {
	Id int `gorm:"primaryKey;column:id"`
	Name string `gorm:"column:name"`
	Description string `gorm:"column:description"`
	Address string `gorm:"column:address"`
	City string `gorm:"column:city"`
	State string `gorm:"column:state"`	
	ZipCode string `gorm:"column:zip_code"`
	PhoneNumber string `gorm:"column:phone_number"`
	Email string `gorm:"column:email"`
	IsActive bool `gorm:"column:is_active"`
	Website string `gorm:"column:website"`
	ImageURL string `gorm:"column:image_url"`
	Price float64 `gorm:"column:price"`
}


// TableName remains unchanged
func (Experiences) TableName() string {
	return "experienceslist"
}

func (u *Experiences) GetId() int {
	return u.Id
}
