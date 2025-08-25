package experiencehandler

import (
	"net/http"
)

type ExperienceHandler interface {
	FindById(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Insert(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetAllActiveExperiences(w http.ResponseWriter, r *http.Request)
}


