package userService


import (
	"Amora/internal/user"
)


type UserService interface {
	GetUserByPhoneNumber(phoneNumber string) (*user.Users, error)
	GetUserByEmail(email string) (*user.Users, error)
	// These will need a auth middleware to check if the user can have a token and then generate a token
	ForgotPassword(EmailOrPhoneNum string) (string, error)
	Login(email, password string) (*user.Users, error)
}



