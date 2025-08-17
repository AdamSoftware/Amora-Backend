package partnerHandler

import ( 
	"Amora/internal/core/CoreHandler"
	"Amora/internal/partner"
	"Amora/internal/partner/partnerService"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"Amora/internal/helper"
	"strconv"
)

type PartnerHandlerInit struct {
	  partnerService *partnerService.PartnerServiceInit
    coreHandler *CoreHandler.CoreHandlerInit[*partner.Partners]
}

func NewpartnerHandlerInit(coreHandler *CoreHandler.CoreHandlerInit[*partner.Partners], partnerService *partnerService.PartnerServiceInit) *PartnerHandlerInit {
    return &PartnerHandlerInit{
				partnerService: partnerService,
        coreHandler: coreHandler,
    }
}


func (u *PartnerHandlerInit) FindById(w http.ResponseWriter, r *http.Request) {
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
		// error if the partner just can't be found
    if entity == nil {
			  helper.WriteError(w, http.StatusNotFound, "partner not found")
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(entity)
}

func (u *PartnerHandlerInit) Get(w http.ResponseWriter, r *http.Request) {
		entities, err := u.coreHandler.Get()
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		// this would should never happen but just in case
		if len(entities) == 0 { 
				helper.WriteError(w, http.StatusNotFound, "No partners found")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}


func (u *PartnerHandlerInit) Insert(w http.ResponseWriter, r *http.Request) {
	var entity *partner.Partners
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

func (u *PartnerHandlerInit) Update(w http.ResponseWriter, r *http.Request) {
	var entity *partner.Partners
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


func (u *PartnerHandlerInit) Delete(w http.ResponseWriter, r *http.Request) {
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


func (u *PartnerHandlerInit) FindByPhoneNumber(w http.ResponseWriter, r *http.Request) {
	phoneNumber := chi.URLParam(r, "phoneNumber") 
	if phoneNumber == "" {
		helper.WriteError(w, http.StatusBadRequest, "Missing phone number parameter")
		return
	}

	user, err := u.partnerService.FindByPhoneNumber(phoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		helper.WriteError(w, http.StatusNotFound, "User not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}


func (u *PartnerHandlerInit) FindByUserId(w http.ResponseWriter, r *http.Request) {
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

	entity, err := u.partnerService.FindByUserId(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if entity == nil {
		helper.WriteError(w, http.StatusNotFound, "Partner not found")
		return
	}
	w.Header().Set("Content-Type", "application/json")	
	json.NewEncoder(w).Encode(entity)
}


func (u *PartnerHandlerInit) FindAllPartnersWithBirthday(w http.ResponseWriter, r *http.Request) {
	partners, err := u.partnerService.FindAllPartnersWithBirthday()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(partners) == 0 {
		helper.WriteError(w, http.StatusNotFound, "No partners with birthday found today nothing to celebrate")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(partners)
}

func (u *PartnerHandlerInit) FindAllPartnersWithAnniversary(w http.ResponseWriter, r *http.Request) {
	partners, err := u.partnerService.FindAllPartnersWithAnniversary()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(partners) == 0 {
		helper.WriteError(w, http.StatusNotFound, "No partners with anniversary found today nothing to celebrate")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(partners)
}
