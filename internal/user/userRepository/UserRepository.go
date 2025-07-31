package userRepository

import (
	"Amora/internal/user"
)

type UserRepository interface {
	GetUserByPhoneNumber(phoneNumber string) (*user.Users, error)
	GetUserByEmail(email string) (*user.Users, error)
}

