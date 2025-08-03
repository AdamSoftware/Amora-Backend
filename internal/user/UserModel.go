package user

type Users struct {
	Id          int    `gorm:"primaryKey;column:id"`
	Email       string `gorm:"column:email"`
	Password    string `gorm:"column:password"`
	Firstname   string `gorm:"column:firstname"`
	Lastname    string `gorm:"column:lastname"`
	City        string `gorm:"column:city"`
	Address     string `gorm:"column:address"`
	State       string `gorm:"column:state"`
	ZipCode     string `gorm:"column:zip_code"`
	PhoneNumber string `gorm:"column:phone_number"`
	IsActive    bool   `gorm:"column:is_active"`
}

// TableName remains unchanged
func (Users) TableName() string {
	return "users"
}

func (u *Users) GetId() int {
	return u.Id
}

