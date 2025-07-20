package Logging 

import (
  "Amora/internal/Logging/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"sync"
)

type GormHook struct {
	db       *gorm.DB
	initOnce sync.Once
}

func NewGormHook(db *gorm.DB) *GormHook {
	return &GormHook{db: db}
}

func (h *GormHook) ensureTable() {
	h.initOnce.Do(func() {
		_ = h.db.AutoMigrate(&models.LogEntry{})
	})
}

func (h *GormHook) Fire(entry *logrus.Entry) error {
	h.ensureTable()
	logEntry := &models.LogEntry{
		Level:   entry.Level.String(),
		Message: entry.Message,
	}
	return h.db.Create(logEntry).Error
}

func (h *GormHook) Levels() []logrus.Level {
	return logrus.AllLevels // Log all levels
}

