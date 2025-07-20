package models

import (
	"time"
)

type LogEntry struct {
	ID        int       `gorm:"primaryKey"`
	Level     string    `gorm:"size:10;index"`
	Message   string    `gorm:"type:text"`
	CreatedAt time.Time
}

