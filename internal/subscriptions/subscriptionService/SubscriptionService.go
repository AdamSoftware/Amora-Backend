package subscriptionService


import (
	"Amora/internal/subscriptions"
)

type SubscriptionService interface { 
	GetByPlanId(planId int) ([]subscriptions.Subscriptions, error)
	GetByUserId(userId int) (*subscriptions.Subscriptions, error)
	FindByStartDateBefore(date string) ([]subscriptions.Subscriptions, error)
	GetByStatus(status string) ([]subscriptions.Subscriptions, error)
	CheckIfUserHadSubscription(userId int) (bool, error)
	UpdateByUserId(userId int, updatedSubscription subscriptions.Subscriptions) error
	
	// moding this so that if you are chaning the status to cancelled it will also update the cancel date to the current date
	Update(Id int, updatedSubscription subscriptions.Subscriptions) error
}

