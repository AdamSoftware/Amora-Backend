package Logging

import (
	"github.com/sirupsen/logrus"
	"sync"
	"gorm.io/gorm"
)

var Logs *logrus.Logger
var once sync.Once

func InitLogger(db *gorm.DB) {
	once.Do(func() {
		Logs = logrus.New()
		if Logs == nil {
			panic("logger initialization failed")
		}
		Logs.SetLevel(logrus.DebugLevel) // Or InfoLevel in production
		Logs.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})


		Logs.AddHook(NewGormHook(db))
	})
}

