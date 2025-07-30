package user

type Users struct {
	id int `gorm:"primaryKey;column:id"`
	email string `gorm:"column:email;"`
	password string `gorm:"column:password;"`
	firstname string `gorm:"column:firstname;"`
	lastname string `gorm:"column:lastname;"`
	city string `gorm:"column:city;"`
	address string `gorm:"column:address;"`
	state string `gorm:"column:state;"`
	zip_code string `gorm:"column:zip_code;"`
	phone_number string `gorm:"column:phone_number;"`
	// I might not keep this field, but it is useful incase of a user ending there subs
	is_active bool `gorm:"column:is_active;"`
}

// this is to change the name of the table to undercase Users so that it 
// matches in the database
func (Users) TableName() string {
	return "users"
}

func (u *Users) GetId() int {
	return u.id
}
