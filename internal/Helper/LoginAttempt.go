package helper

import (
	"time"
)

type LoginAttempt struct {
	Count int
	LastFailed time.Time 
}
