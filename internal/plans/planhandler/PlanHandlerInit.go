package planhandler

import (
	"Amora/internal/core/CoreHandler"
	"Amora/internal/plans"
	"Amora/internal/plans/planservice"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"Amora/internal/helper"
	"strconv"
)

type PlanHandlerInit struct {
	  plansService *planservice.PlanServiceInit
    coreHandler *CoreHandler.CoreHandlerInit[*plans.Plans]
}

func NewPlanHandlerInit(coreHandler *CoreHandler.CoreHandlerInit[*plans.Plans], plansService *planservice.PlanServiceInit) *PlanHandlerInit {
    return &PlanHandlerInit{
				plansService: plansService,
        coreHandler: coreHandler,
    }
}

func (u *PlanHandlerInit) FindById(w http.ResponseWriter, r *http.Request) {
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
		// error if the plans just can't be found
    if entity == nil {
			  helper.WriteError(w, http.StatusNotFound, "Plan not found")
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(entity)
}

func (u *PlanHandlerInit) Get(w http.ResponseWriter, r *http.Request) {
		entities, err := u.coreHandler.Get()
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		// this would should never happen but just in case
		if len(entities) == 0 { 
				helper.WriteError(w, http.StatusNotFound, "No planss found")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}


func (u *PlanHandlerInit) Insert(w http.ResponseWriter, r *http.Request) {
	var entity *plans.Plans
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

func (u *PlanHandlerInit) Update(w http.ResponseWriter, r *http.Request) {
	var entity *plans.Plans
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		helper.WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updatedEntity, err := u.coreHandler.Update(&entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedEntity)
}


func (u *PlanHandlerInit) Delete(w http.ResponseWriter, r *http.Request) {
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
}



