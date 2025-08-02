package userService

import (
	"Amora/internal/Auth"
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


func (u *UserServiceInit) Login(email, password string) (map[string]interface{}, error) {
	if email == "" || password == "" {
		return nil, errors.New("email or password cannot be empty")
	}

	User, err := u.userRepo.GetUserByEmail(email)
	if err != nil || User == nil {
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT
	token, err := Auth.GenerateJWT(User)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token": token,
		"user": map[string]interface{}{
			"id":         User.Id,
			"email":      User.Email,
			"firstname":  User.Firstname,
			"lastname":   User.Lastname,
			"phoneNumber": User.PhoneNumber,
		},
	}, nil
}






