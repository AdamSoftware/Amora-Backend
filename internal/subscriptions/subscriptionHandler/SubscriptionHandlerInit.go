package subscriptionHandler

import ( 
	"Amora/internal/core/CoreHandler"
	"Amora/internal/subscriptions"
	"Amora/internal/subscriptions/subscriptionService"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"Amora/internal/helper"
	"strconv"
)


type SubscriptionHandlerInit struct {
	  subscriptionService *subscriptionService.SubscriptionServiceInit
    coreHandler *CoreHandler.CoreHandlerInit[*subscriptions.Subscriptions]
}

func NewsubscriptionHandlerInit(coreHandler *CoreHandler.CoreHandlerInit[*subscriptions.Subscriptions], subscriptionService *subscriptionService.SubscriptionServiceInit) *SubscriptionHandlerInit {
    return &SubscriptionHandlerInit{
				subscriptionService: subscriptionService,
        coreHandler: coreHandler,
    }
}

func (u *SubscriptionHandlerInit) FindById(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id") 
    id, err := strconv.Atoi(idStr)
    if err != nil {
			  helper.WriteError(w, http.StatusBadRequest, "Invalid ID parameter")
        return
    }

    entity, err := u.coreHandler.FindById(id)
		// this is the internal error
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
		// error if the subsctions just can't be found
    if entity == nil {
			  helper.WriteError(w, http.StatusNotFound, "Subscription not found")
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(entity)
}

func (u *SubscriptionHandlerInit) Get(w http.ResponseWriter, r *http.Request) {
		entities, err := u.coreHandler.Get()
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		// this would should never happen but just in case
		if len(entities) == 0 { 
				helper.WriteError(w, http.StatusNotFound, "No subsctionss found")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}


func (u *SubscriptionHandlerInit) Insert(w http.ResponseWriter, r *http.Request) {
	var entity *subscriptions.Subscriptions
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	createdEntity, err := u.coreHandler.Insert(&entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdEntity)
}

func (u *SubscriptionHandlerInit) Update(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	if idStr == "" {
		helper.WriteError(w, http.StatusBadRequest, "Missing ID parameter")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		helper.WriteError(w, http.StatusBadRequest, "Invalid ID parameter")
		return
	}

	var entity subscriptions.Subscriptions
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	defer r.Body.Close()

	// Call service layer
	err = u.subscriptionService.Update(id, &entity)
	if err != nil {
		helper.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Subscription updated successfully"}`))
}



func (u *SubscriptionHandlerInit) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		helper.WriteError(w, http.StatusBadRequest, "Missing ID parameter")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		helper.WriteError(w, http.StatusBadRequest, "Invalid ID parameter")
		return
	}

	err = u.coreHandler.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	// return the id of the item deleted
	w.Write([]byte(`{"message": "Subscription deleted successfully", "id": ` + strconv.Itoa(id) + `}`))
}

func (u *SubscriptionHandlerInit) GetByPlanId(w http.ResponseWriter, r *http.Request) {
	planIdStr := chi.URLParam(r, "planId")
	if planIdStr == "" {
		helper.WriteError(w, http.StatusBadRequest, "Missing plan ID parameter")
		return
	}

	planId, err := strconv.Atoi(planIdStr)
	if err != nil || planId <= 0 {
		helper.WriteError(w, http.StatusBadRequest, "Invalid plan ID parameter")
		return
	}

	subscriptions, err := u.subscriptionService.GetByPlanId(planId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(subscriptions) == 0 {
		helper.WriteError(w, http.StatusNotFound, "No subscriptions found for the given plan ID")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subscriptions)
}

func (u *SubscriptionHandlerInit) GetByUserId(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userId")
	if userIdStr == "" {
		helper.WriteError(w, http.StatusBadRequest, "Missing user ID parameter")
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil || userId <= 0 {
		helper.WriteError(w, http.StatusBadRequest, "Invalid user ID parameter")
		return
	}

	subscriptions, err := u.subscriptionService.GetByUserId(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subscriptions)
}

func (u *SubscriptionHandlerInit) FindByStartDateBefore(w http.ResponseWriter, r *http.Request) {
	startDateStr := chi.URLParam(r, "startDate")
	if startDateStr == "" {
		helper.WriteError(w, http.StatusBadRequest, "Missing start date parameter")
		return
	}

	subscriptions, err := u.subscriptionService.FindByStartDateBefore(startDateStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(subscriptions) == 0 {
		helper.WriteError(w, http.StatusNotFound, "No subscriptions found before the given start date")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subscriptions)
}

func (u *SubscriptionHandlerInit) GetByStatus(w http.ResponseWriter, r *http.Request) {
	status := chi.URLParam(r, "status")
	if status == "" {
		helper.WriteError(w, http.StatusBadRequest, "Missing status parameter")
		return
	}

	subscriptions, err := u.subscriptionService.GetByStatus(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(subscriptions) == 0 {
		helper.WriteError(w, http.StatusNotFound, "No subscriptions found with the given status")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(subscriptions)
}

func (u *SubscriptionHandlerInit) CheckIfUserHadSubscription(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userId")
	if userIdStr == "" {
		helper.WriteError(w, http.StatusBadRequest, "Missing user ID parameter")
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil || userId <= 0 {
		helper.WriteError(w, http.StatusBadRequest, "Invalid user ID parameter")
		return
	}

	hasSubscription, err := u.subscriptionService.CheckIfUserHadSubscription(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]bool{"hasSubscription": hasSubscription}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (u *SubscriptionHandlerInit) UpdateByUserId(w http.ResponseWriter, r *http.Request) {
	userIdStr := chi.URLParam(r, "userId")
	if userIdStr == "" {
		helper.WriteError(w, http.StatusBadRequest, "Missing user ID parameter")
		return
	}

	userId, err := strconv.Atoi(userIdStr)
	if err != nil || userId <= 0 {
		helper.WriteError(w, http.StatusBadRequest, "Invalid user ID parameter")
		return
	}

	var entity subscriptions.Subscriptions
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}
	defer r.Body.Close()

	err = u.subscriptionService.UpdateByUserId(userId, &entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Subscription updated successfully for user ID ` + userIdStr + `"}`))
}
