package userService

import (
	"Amora/internal/Logging"
	"Amora/internal/user/userRepository"
	"errors"
	"Amora/internal/user"
	"golang.org/x/crypto/bcrypt"
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

// the login needs to check if the email matches a user and the passsword matches that usr too 
func (u *UserServiceInit) Login(email, password string) (*user.Users, error) {
	if email == "" || password == "" {
		Logging.Logs.Errorf("Email or password is empty")
		return nil, errors.New("email or password cannot be empty")
	}

	User, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving user by email %s: %v", email, err)
		return nil, err
	}

	if User == nil {
		Logging.Logs.Warnf("No user found with email %s", email)
		return nil, errors.New("user not found")
	}

	// Compare hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(password)); err != nil {
		Logging.Logs.Warnf("Incorrect password for user with email %s", email)
		return nil, errors.New("incorrect password")
	}

	// (Optional) generate and return a token here

	return User, nil
}





