package userRepository

import (
	"Amora/internal/Logging"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/user"
	"errors"
	"gorm.io/gorm"
)


type UserRepositoryInit struct {
	Core  *CoreRepository.CoreRepositoryInit[*user.Users]
}

func NewuserRepositoryInit(DB *gorm.DB) (*UserRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*user.Users](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &UserRepositoryInit{ Core : coreRepo }, nil
}
    
// GetUserByPhoneNumber
func (u *UserRepositoryInit) GetUserByPhoneNumber(phoneNumber string) (*user.Users, error) {
    var user user.Users
    err := u.Core.DB.Where("phone_number = ?", phoneNumber).First(&user).Error
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, nil
        }
        Logging.Logs.Errorf("Error retrieving user by phone number %s: %v", phoneNumber, err)
        return nil, err
    }
    return &user, nil
}


// GetUserByEmail
func (u *UserRepositoryInit) GetUserByEmail(email string) (*user.Users, error) {
	var user user.Users
	err := u.Core.DB.Where("email = ?", email).First(&user).Error
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





