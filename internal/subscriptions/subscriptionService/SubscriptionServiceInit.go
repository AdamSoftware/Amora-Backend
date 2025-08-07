package subscriptionService

import ( 
	"Amora/internal/subscriptions"
	"Amora/internal/subscriptions/subscriptionRepository"
	"Amora/internal/Logging"
	"time"
	"errors"
)

type SubscriptionServiceInit struct {
	subscriptionRepo *subscriptionRepository.SubscriptionRepositoryInit
}

func NewsubscriptionServiceInit( subscriptionRepo *subscriptionRepository.SubscriptionRepositoryInit) (*SubscriptionServiceInit, error) {

	if subscriptionRepo == nil {
		Logging.Logs.Errorf("subscription repository is nil")
		return nil, errors.New("subscription repository is nil")
	}

	return &SubscriptionServiceInit{
		subscriptionRepo: subscriptionRepo,
	}, nil
}


func (s *SubscriptionServiceInit) GetByPlanId(planId int) ([]subscriptions.Subscriptions, error) {
	subscription, err := s.subscriptionRepo.GetByPlanId(planId)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving subscription by plan ID %s: %v", planId, err)
		return nil, err
	}
	return subscription, nil
}

func (s *SubscriptionServiceInit) GetByUserId(userId int) (*subscriptions.Subscriptions, error) {
	subscription, err := s.subscriptionRepo.GetByUserId(userId)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving subscription by user ID %s: %v", userId, err)
		return nil, err
	}
	return subscription, nil
}

func (s *SubscriptionServiceInit) FindByStartDateBefore(date string) ([]subscriptions.Subscriptions, error) {
	subscriptionsList, err := s.subscriptionRepo.FindByStartDateBefore(date)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving subscriptions before date %s: %v", date, err)
		return nil, err
	}
	return subscriptionsList, nil
}

func (s *SubscriptionServiceInit) GetByStatus(status string) ([]subscriptions.Subscriptions, error) {
	subscriptionsList, err := s.subscriptionRepo.GetByStatus(status)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving subscriptions by status %s: %v", status, err)
		return nil, err
	}
	return subscriptionsList, nil
}

func (s *SubscriptionServiceInit) CheckIfUserHadSubscription(userId int) (bool, error) {
	hadSubscription, err := s.subscriptionRepo.CheckIfUserHadSubscription(userId)
	if err != nil {
		Logging.Logs.Errorf("Error checking if user ID %s had subscription: %v", userId, err)
		return false, err
	}
	return hadSubscription, nil
}

func (s *SubscriptionServiceInit) UpdateByUserId(userId int, updatedSubscription subscriptions.Subscriptions) error {
	err := s.subscriptionRepo.UpdateByUserId(userId, updatedSubscription)

	// If the status is being changed to "cancelled", set the CancelDate to the current date
	if updatedSubscription.Status == "cancelled" {
		updatedSubscription.CancelDate = time.Now()
	}
	
	if err != nil {
		Logging.Logs.Errorf("Error updating subscription for user ID %s: %v", userId, err)
		return err
	}
	return nil
}

// moding the update so that if you are chanigng the status to cancelled it will also update the cancel date
func (s *SubscriptionServiceInit) Update(Id int, updatedSubscription subscriptions.Subscriptions) error {
	// If the status is being changed to "cancelled", set the CancelDate to the current date
	if updatedSubscription.Status == "cancelled" {
		updatedSubscription.CancelDate = time.Now()
	}

	err := s.subscriptionRepo.UpdateByUserId(Id, updatedSubscription)
	if err != nil {
		Logging.Logs.Errorf("Error updating subscription for user ID %s: %v", Id, err)
		return err
	}
	return nil
}
