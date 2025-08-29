package reminders

type Reminders struct {
	Id                int       `gorm:"primaryKey;column:id"`
	UserId            int       `gorm:"column:user_id"`
	ReservationId     int       `gorm:"column:reservation_id"`
	ReminderTimeBefore int      `gorm:"column:ReminderTimeBefore"`
	ScheduledTime     *string    `gorm:"column:ScheduleTime"`  // change from time.Time
	ScheduledDate     *string `gorm:"column:ScheduleDate"`
	ReminderTime      *string    `gorm:"column:ReminderTime"`  // change from time.Time
}



// TableName remains unchanged
func (Reminders) TableName() string {
	return "reminders"
}

func (u *Reminders) GetId() int {
	return u.Id
}


