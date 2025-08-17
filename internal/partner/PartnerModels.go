package partner

import (
	"time"
)

type Partners struct {
	Id int `gorm:"primaryKey;column:id"`
	Nickname	string `gorm:"column:nickname"`
	Firstname string `gorm:"column:firstname"`
	Lastname  string `gorm:"column:lastname"`
	Birthdate time.Time `gorm:"column:birthdate"`
	Anniversary time.Time `gorm:"column:anniversary"`
	Phonenumber     string `gorm:"column:phone_number"`
	User_id int `gorm:"column:user_id"`
}

// this makes sure that the table name is the same as the database one even when the struct name is dif
func (Partners) TableName() string {
	return "partners"
}

func (s *Partners) GetId() int {
	return s.Id
}

