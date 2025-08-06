package subscriptionRepository

import (
	"Amora/internal/subscriptions"
	"gorm.io/gorm"
	"Amora/internal/core/CoreRepository"
	"Amora/internal/Logging"
	"errors"
)

type SubscriptionRepositoryInit struct {
	Core  *CoreRepository.CoreRepositoryInit[*subscriptions.Subscriptions]
}

func NewSubscriptionRepositoryInit(DB *gorm.DB) (*SubscriptionRepositoryInit, error) {
	if DB == nil {
		Logging.Logs.Warnf("Database connection is nil")
	}

	coreRepo, err := CoreRepository.NewCoreRepository[*subscriptions.Subscriptions](DB)
	if err != nil {
		Logging.Logs.Errorf("Failed to initialize core repository: %v", err)
		return nil, errors.New("failed to initialize core repository")
	}

	return &SubscriptionRepositoryInit{ Core : coreRepo }, nil
}





func (s *SubscriptionRepositoryInit) GetByPlanId(planId int) ([]subscriptions.Subscriptions, error) {
	var subscriptionsList []subscriptions.Subscriptions
	err := s.Core.DB.Where("plan_id = ?", planId).Find(&subscriptionsList).Error
	if err != nil {
		Logging.Logs.Errorf("Error retrieving subscriptions by plan ID %d: %v", planId, err)
		return nil, err
	}
	return subscriptionsList, nil
}


func (s *SubscriptionRepositoryInit) GetByUserId(userId int) (subscriptions.Subscriptions, error) {
	var subscription subscriptions.Subscriptions
	err := s.Core.DB.Where("user_id = ?", userId).First(&subscription).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Logging.Logs.Warnf("No subscription found for user ID %d", userId)
			return subscription, nil // Return empty subscription if not found
		}
		Logging.Logs.Errorf("Error retrieving subscription by user ID %d: %v", userId, err)
		return subscription, err
	}
	return subscription, nil
}

// this will allow the admin to see all subscriptions that started at a certain date
func (s *SubscriptionRepositoryInit) FindByStartDateBefore(date string) ([]subscriptions.Subscriptions, error) {
	var subscriptionsList []subscriptions.Subscriptions
	err := s.Core.DB.Where("start_date < ?", date).Find(&subscriptionsList).Error
	if err != nil {
		Logging.Logs.Errorf("Error retrieving subscriptions with start date before %s: %v", date, err)
		return nil, err
	}
	return subscriptionsList, nil
}

// this will allow reporting on how many subscriptions have been cancelled 
func (s *SubscriptionRepositoryInit) GetByStatus(status string) ([]subscriptions.Subscriptions, error) {
	var subscriptionsList []subscriptions.Subscriptions
	err := s.Core.DB.Where("status = ?", status).Find(&subscriptionsList).Error
	if err != nil {
		Logging.Logs.Errorf("Error retrieving subscriptions by status %s: %v", status, err)
		return nil, err
	}
	return subscriptionsList, nil
}

// this is to check if the user had a subscirption before so that we can offer them discount or to just restart the same subscription plan
func (s *SubscriptionRepositoryInit) CheckIfUserHadSubscription(userId int) (bool, error) {
	var count int64
	err := s.Core.DB.Model(&subscriptions.Subscriptions{}).Where("user_id = ?", userId).Count(&count).Error
	if err != nil {
		Logging.Logs.Errorf("Error checking if user %d had a subscription: %v", userId, err)
		return false, err
	}
	return count > 0, nil
}

// this is for easier updating of the subscirption table on the user id instead of the subscirption id 
func (s *SubscriptionRepositoryInit) UpdateByUserId(userId int, updatedSubscription subscriptions.Subscriptions) error {
	err := s.Core.DB.Model(&subscriptions.Subscriptions{}).Where("user_id = ?", userId).Updates(updatedSubscription).Error
	if err != nil {
		Logging.Logs.Errorf("Error updating subscription for user ID %d: %v", userId, err)
		return err
	}
	return nil
}
