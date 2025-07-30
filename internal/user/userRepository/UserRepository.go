package userRepository

import (
	"Amora/internal/user"
)

type UserRepository interface {
	GetById(id int) (*user.Users, error)
	Create(user *user.Users) (*user.Users, error)
	Update(user *user.Users) (*user.Users, error)
	Delete(id int) error
	Get() ([]*user.Users, error)
	GetUserByPhoneNumber(phoneNumber string) (*user.Users, error)
	GetUserByEmail(email string) (*user.Users, error)
}

