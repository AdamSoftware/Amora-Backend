package partner

type Partners struct {
	Id int `gorm:"primaryKey;column:id"`
	nickname	string `gorm:"column:nickname"`
	firstname string `gorm:"column:firstname"`
	lastname  string `gorm:"column:lastname"`
	birthdate string `gorm:"column:birthdate"`
	phonenumber     string `gorm:"column:phone_number"`
	user_id int `gorm:"column:user_id"`
}

// this makes sure that the table name is the same as the database one even when the struct name is dif
func (Partners) TableName() string {
	return "partners"
}

func (s *Partners) GetId() int {
	return s.Id
}

