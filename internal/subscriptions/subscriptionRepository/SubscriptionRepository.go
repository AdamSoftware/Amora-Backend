package subscriptionRepository

import (
	"Amora/internal/subscriptions"
)

// NOTE: during the update will have to mod the core update in service layer so that when you 
// change the status to cancelled it will also update the cancel date to the current date

// WARNING: most of these are just used for admin purposes and shouuld not be displayed to user for security reasons
// use discretion when calling these methods. 
type SubscriptionRepository interface {
	GetByPlanId(planId int) ([]subscriptions.Subscriptions, error)
	GetByUserId(userId int) (*subscriptions.Subscriptions, error)
	FindByStartDateBefore(date string) ([]subscriptions.Subscriptions, error)
	GetByStatus(status string) ([]subscriptions.Subscriptions, error)
	// will check if the user ever had a subscription 
	CheckIfUserHadSubscription(userId int) (bool, error)
	// will update using the user id same as the core update just allowing you to use user id not just Id
	UpdateByUserId(userId int, updatedSubscription subscriptions.Subscriptions) error
}
