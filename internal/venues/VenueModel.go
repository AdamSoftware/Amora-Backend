package venues

type Venues struct {
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
}


// TableName remains unchanged
func (Venues) TableName() string {
	return "venueslist"
}

func (u *Venues) GetId() int {
	return u.Id
}
