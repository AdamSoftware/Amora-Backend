package subscriptionHandler

import ( 
	"net/http"
)


type SubscriptionHandler interface {
	FindById(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Insert(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetByPlanId(w http.ResponseWriter, r *http.Request)
	GetByUserId(w http.ResponseWriter, r *http.Request)
	FindByStartDateBefore(w http.ResponseWriter, r *http.Request)
	GetByStatus(w http.ResponseWriter, r *http.Request)
	CheckIfUserHadSubscription(w http.ResponseWriter, r *http.Request)
	UpdateByUserId(w http.ResponseWriter, r *http.Request)
	GetByPayedStatus(w http.ResponseWriter, r *http.Request)
	CheckSubscriptionWebhook(w http.ResponseWriter, r *http.Request)
}

