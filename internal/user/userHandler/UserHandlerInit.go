package userHandler

import (
	"Amora/internal/core/CoreHandler"
	"Amora/internal/user"
	"Amora/internal/user/userService"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"strconv"
)


type UserHandlerInit struct {
	  userService *userService.UserServiceInit
    coreHandler *CoreHandler.CoreHandlerInit[*user.Users]
}

func NewUserHandlerInit(coreHandler *CoreHandler.CoreHandlerInit[*user.Users], userService *userService.UserServiceInit) *UserHandlerInit {
    return &UserHandlerInit{
				userService: userService,
        coreHandler: coreHandler,
    }
}

func (u *UserHandlerInit) FindById(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id") 
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    entity, err := u.coreHandler.FindById(id)
		// this is the internal error
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
		// error if the user just can't be found
    if entity == nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(entity)
}

func (u *UserHandlerInit) Get(w http.ResponseWriter, r *http.Request) {
		entities, err := u.coreHandler.Get()
		if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
		}
		// this would should never happen but just in case
		if len(entities) == 0 { 
				http.Error(w, "No users found", http.StatusNotFound)
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}


func (u *UserHandlerInit) Insert(w http.ResponseWriter, r *http.Request) {
	var entity *user.Users
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
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

func (u *UserHandlerInit) Update(w http.ResponseWriter, r *http.Request) {
	var entity *user.Users
	if err := json.NewDecoder(r.Body).Decode(&entity); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
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


func (u *UserHandlerInit) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	if idStr == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		return
	}

	err = u.coreHandler.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (u *UserHandlerInit) GetUserByPhoneNumber(w http.ResponseWriter, r *http.Request) {
	phoneNumber := chi.URLParam(r, "phoneNumber") 
	if phoneNumber == "" {
		http.Error(w, "Missing phone number parameter", http.StatusBadRequest)
		return
	}

	user, err := u.userService.GetUserByPhoneNumber(phoneNumber)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (u *UserHandlerInit) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email") 
	if email == "" {
		http.Error(w, "Missing email parameter", http.StatusBadRequest)
		return
	}

	user, err := u.userService.GetUserByEmail(email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (u *UserHandlerInit) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	EmailOrPhoneNum := chi.URLParam(r, "emailOrPhoneNum") 
	if EmailOrPhoneNum == "" {
		http.Error(w, "Missing email or phone number parameter", http.StatusBadRequest)
		return
	}

	message, err := u.userService.ForgotPassword(EmailOrPhoneNum)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

func (u *UserHandlerInit) Login(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	user, err := u.userService.Login(loginData.Email, loginData.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user == nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}




