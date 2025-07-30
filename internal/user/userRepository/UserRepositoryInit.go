package userRepository

import (
	"Amora/internal/Logging"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/user"
	"errors"
	"gorm.io/gorm"
)


type userRepositoryInit struct {
	core  *CoreRepository.CoreRepositoryInit[*user.Users]
}

func NewuserRepositoryInit(DB *gorm.DB) (*userRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*user.Users](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &userRepositoryInit{ core : coreRepo }, nil
}

// Generic opperations done 

func (u *userRepositoryInit) GetById(id int) (*user.Users, error) {
	return u.core.FinDById(id) 
}

func (u *userRepositoryInit) Create(user *user.Users) (*user.Users, error) {
	return u.core.Insert(user)
}

func (u *userRepositoryInit) Update(user *user.Users) (*user.Users, error) {
	return u.core.Update(user)
}

func (u *userRepositoryInit) Delete(id int) error {
	return u.core.Delete(id)
}
    
// GetUserByPhoneNumber
func (u *userRepositoryInit) GetUserByPhoneNumber(phoneNumber string) ([]*user.Users, error) {
	var users []*user.Users
	err := u.core.DB.Where("phone_number = ?", phoneNumber).Find(&users).Error
	if err != nil {
		Logging.Logs.Errorf("Error retrieving user by phone number %s: %v", phoneNumber, err)
		return nil, err
	}
	if len(users) == 0 {
		Logging.Logs.Warnf("No user found with phone number %s", phoneNumber)
		return nil, nil
	}
	return users, nil
}

// GetUserByEmail
func (u *userRepositoryInit) GetUserByEmail(email string) (*user.Users, error) {
	var user user.Users
	err := u.core.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Logging.Logs.Warnf("No user found with email %s", email)
			return nil, nil
		}
		Logging.Logs.Errorf("Error retrieving user by email %s: %v", email, err)
		return nil, err
	}
	return &user, nil
}





