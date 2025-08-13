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

func (s *SubscriptionServiceInit) UpdateByUserId(userId int, updatedSubscription *subscriptions.Subscriptions) error {

	now := time.Now()
	// If the status is being changed to "cancelled", set the CancelDate to the current date
	if updatedSubscription.Status == "cancelled" {
		updatedSubscription.CancelDate = &now 
		updatedSubscription.NextPaymentDate = nil // Clear NextPaymentDate when cancelling
		updatedSubscription.StartDate = nil // Clear StartDate when cancelling
	}

	if updatedSubscription.Status == "active" {
		// If the status is being changed to "active", reset the CancelDate
		updatedSubscription.CancelDate = nil
		updatedSubscription.StartDate = &now
		nextPayment := now.Add(30 * 24 * time.Hour)
		updatedSubscription.NextPaymentDate = &nextPayment
	}

	err := s.subscriptionRepo.UpdateByUserId(userId, updatedSubscription)
	if err != nil {
		Logging.Logs.Errorf("Error updating subscription for user ID %s: %v", userId, err)
		return err
	}
	return nil
}

// moding the update so that if you are chanigng the status to cancelled it will also update the cancel date
func (s *SubscriptionServiceInit) Update(Id int, updatedSubscription *subscriptions.Subscriptions) error {
	now := time.Now()
	// If the status is being changed to "cancelled", set the CancelDate to the current date
	if updatedSubscription.Status == "cancelled" {
		updatedSubscription.CancelDate = &now 
		updatedSubscription.NextPaymentDate = nil // Clear NextPaymentDate when cancelling
		updatedSubscription.StartDate = nil // Clear StartDate when cancelling
	}

	if updatedSubscription.Status == "active" {
		// If the status is being changed to "active", reset the CancelDate
		updatedSubscription.CancelDate = nil
		updatedSubscription.StartDate = &now
		nextPayment := now.Add(30 * 24 * time.Hour)
		updatedSubscription.NextPaymentDate = &nextPayment
	}

	err := s.subscriptionRepo.UpdateByUserId(Id, updatedSubscription)
	if err != nil {
		Logging.Logs.Errorf("Error updating subscription for user ID %s: %v", Id, err)
		return err
	}
	return nil
}


func (s *SubscriptionServiceInit) GetByPayedStatus(payed bool) ([]subscriptions.Subscriptions, error) {
	subscriptionsList, err := s.subscriptionRepo.GetByPayedStatus(payed)
	if err != nil {
		Logging.Logs.Errorf("Error retrieving subscriptions by payed status %s: %v", payed, err)
		return nil, err
	}
	return subscriptionsList, nil
}

// NOTE: Hook to check if the subscription has been payed or not and do the logic for either this will run on a different server
// ever day at midnight
func (s *SubscriptionServiceInit) CheckSubscriptionWebhook() ([]subscriptions.Subscriptions, error) {
	var subscriptionsList []subscriptions.Subscriptions

	// grab all the active subscripitons that have a next payment date that is past due
	// this should prob go in the repo layer 
	err := s.subscriptionRepo.Core.DB.
		Where("status = ? AND next_payment < ? AND payed = ?", "active", time.Now(), false).
		Find(&subscriptionsList).Error

	if err != nil {
		Logging.Logs.Errorf("Error retrieving active subscriptions with past next payment date: %v", err)
		return nil, err
	}

	var updatedSubs []subscriptions.Subscriptions

	// update all the fields if it is past due
	for _, sub := range subscriptionsList {
		now := time.Now()
		sub.Status = "past_due"
		sub.StartDate = nil
		sub.NextPaymentDate = nil
		sub.CancelDate = &now

		if err := s.subscriptionRepo.Core.DB.Save(&sub).Error; err != nil {
			Logging.Logs.Errorf("Error updating subscription ID %d: %v", sub.Id, err)
			continue
		}

		updatedSubs = append(updatedSubs, sub)
	}

	return updatedSubs, nil
}

