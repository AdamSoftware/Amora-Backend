package userService

import (
	"Amora/internal/Logging"
	"Amora/internal/user/userRepository"
	"errors"
	"Amora/internal/user"
)


type UserServiceInit struct {
	userRepo *userRepository.UserRepositoryInit
}

func NewUserServiceInit( userRepo *userRepository.UserRepositoryInit) (*UserServiceInit, error) {

	if userRepo == nil {
		Logging.Logs.Errorf("user repository is nil")
		return nil, errors.New("user repository is nil")
	}

	return &UserServiceInit{
		userRepo: userRepo,
	}, nil
}

// I am so good all the bullshit code is gone and generic now 

func (u *UserServiceInit) GetUserByPhoneNumber(phoneNumber string) (*user.Users, error) {
	User, err := u.userRepo.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving user by phone number %s: %v", phoneNumber, err)
		return nil, err
	}
	return User, nil
}

func (u *UserServiceInit) GetUserByEmail(email string) (*user.Users, error) {
	User, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving user by email %s: %v", email, err)
		return nil, err
	}
	return User, nil
}

// ForgotPassword is a placeholder for the logic when the middleware is done
func (u *UserServiceInit) ForgotPassword(EmailOrPhoneNum string) (string, error) {
	return "Reset link sent to your email/phone", nil
}

// Login is a placeholder for the logic when the middleware is done 
func (u *UserServiceInit) Login(email, password string) (*user.Users, error) {
	return nil, nil
}



