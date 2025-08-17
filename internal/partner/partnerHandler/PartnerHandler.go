package partnerHandler

import (
    "net/http"
)


type PartnerHandler interface {
	FindById(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Insert(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindByPhoneNumber(w http.ResponseWriter, r *http.Request)
	FindByUserId(w http.ResponseWriter, r *http.Request)
	FindAllPartnersWithBirthday(w http.ResponseWriter, r *http.Request)	
	FindAllPartnersWithAnniversary(w http.ResponseWriter, r *http.Request)
}
	


