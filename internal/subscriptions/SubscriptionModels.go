package subscriptions

import (
	"time"
)


type Subscriptions struct {
	Id int `gorm:"primaryKey;column:id"`
	UserId int `gorm:"column:user_id"`
	PlanId int `gorm:"column:plan_id"`
	Status string `gorm:"column:status"`
	StartDate time.Time `gorm:"column:start_date"`
	NextPaymentDate time.Time `gorm:"column:next_payment"`
	CancelDate time.Time `gorm:"column:cancelled_at"`
}

// this makes sure that the table name is the same as the database one even when the struct name is dif
func (Subscriptions) TableName() string {
	return "subscriptions"
}

func (s *Subscriptions) GetId() int {
	return s.Id
}

