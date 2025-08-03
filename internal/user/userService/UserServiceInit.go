package userService

import (
	"Amora/internal/Auth"
	"fmt"
	"Amora/internal/Logging"
	"Amora/internal/user/userRepository"
	"errors"
	"Amora/internal/user"
	"golang.org/x/crypto/bcrypt"
	"Amora/internal/helper"
	"time"
)


var loginAttempts = make(map[string]*helper.LoginAttempt)


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

	var LockDuration = 3 * time.Minute 
	var maxAttempts = 1 


	now := time.Now()

	// checking if the user has exceeded the retry limit for passsword and if the account is locked

	fmt.Println("Checking login attempts", loginAttempts)

	if attempt, exists := loginAttempts[email]; exists {
		if attempt.Count >= maxAttempts && now.Sub(attempt.LastFailed) < LockDuration {
			// log it to the db that the account has been locked and what email was locked also the time it was locked
			fmt.Println("checking attempts for debugging", attempt.Count, maxAttempts, now.Sub(attempt.LastFailed), LockDuration)
			Logging.Logs.Warnf("Account %s is locked due to too many failed login attempts, AtTime: %q", email, time.Now())
			return nil, errors.New("account is locked due to too many failed login attempts")
		}
	}


	
	if email == "" || password == "" {
		return nil, errors.New("email or password cannot be empty")
	}

	User, err := u.userRepo.GetUserByEmail(email)
	if err != nil || User == nil {
		// Record the failed login attempt
		return nil, errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(password)); err != nil {
		recordLastFailedAttempt(email, loginAttempts)
		fmt.Println("checking why its not locking me out", loginAttempts)
		fmt.Println("how many attempts", loginAttempts[email].Count)
		return nil, errors.New("invalid email or password")
	}

	// reset the login attempts and reset the last failed time
	delete(loginAttempts, email)

	// Generate JWT
	token, err := Auth.GenerateJWT(User)
	if err != nil {
		return nil, err
	}

	// the thing is saying that I can use any but I don't know wtf that is and don't want to look it up.
	// fix maybe later
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

func recordLastFailedAttempt(email string, loginAttempts map[string]*helper.LoginAttempt) {
	attempt, exists := loginAttempts[email]
	if !exists {
		attempt = &helper.LoginAttempt{
			Count: 0,
			LastFailed: time.Now(),
		}
		// logs the attempts with the email
		loginAttempts[email] = attempt
	}else {
		fmt.Println("should be hitting this", email, attempt.Count)
		attempt.Count++	
		attempt.LastFailed = time.Now()
	}
}





