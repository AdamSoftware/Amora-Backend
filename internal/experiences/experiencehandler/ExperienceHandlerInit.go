package experiencehandler

import (
	"Amora/internal/core/CoreHandler"
	"Amora/internal/experiences"
	"Amora/internal/experiences/experienceservice"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"Amora/internal/helper"
	"strconv"
)

type ExperienceHandlerInit struct {
	  experiencesService *experienceservice.ExperienceServiceInit
    coreHandler *CoreHandler.CoreHandlerInit[*experiences.Experiences]
}

func NewExperienceHandlerInit(coreHandler *CoreHandler.CoreHandlerInit[*experiences.Experiences], experiencesService *experienceservice.ExperienceServiceInit) *ExperienceHandlerInit {
    return &ExperienceHandlerInit{
				experiencesService: experiencesService,
        coreHandler: coreHandler,
    }
}

func (u *ExperienceHandlerInit) FindById(w http.ResponseWriter, r *http.Request) {
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
		// error if the experiences just can't be found
    if entity == nil {
			  helper.WriteError(w, http.StatusNotFound, "Experience not found")
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(entity)
}

func (u *ExperienceHandlerInit) Get(w http.ResponseWriter, r *http.Request) {
		entities, err := u.coreHandler.Get()
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		// this would should never happen but just in case
		if len(entities) == 0 { 
				helper.WriteError(w, http.StatusNotFound, "No experiencess found")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}


func (u *ExperienceHandlerInit) Insert(w http.ResponseWriter, r *http.Request) {
	var entity *experiences.Experiences
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

func (u *ExperienceHandlerInit) Update(w http.ResponseWriter, r *http.Request) {
	var entity *experiences.Experiences
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


func (u *ExperienceHandlerInit) Delete(w http.ResponseWriter, r *http.Request) {
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

func (u *ExperienceHandlerInit) GetAllActiveExperiences(w http.ResponseWriter, r *http.Request) {
	experiences, err := u.experiencesService.GetAllActiveExperiences()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(experiences)
}

