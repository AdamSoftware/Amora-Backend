package userHandler

import (
	"strings"
	"os"
	"github.com/golang-jwt/jwt/v5"
	"fmt"
	"Amora/internal/core/CoreHandler"
	"Amora/internal/user"
	"Amora/internal/user/userService"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"Amora/internal/helper"
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
			  helper.WriteError(w, http.StatusBadRequest, "Invalid ID parameter")
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
			  helper.WriteError(w, http.StatusNotFound, "User not found")
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
				helper.WriteError(w, http.StatusNotFound, "No users found")
				return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(entities)
}


func (u *UserHandlerInit) Insert(w http.ResponseWriter, r *http.Request) {
	var entity *user.Users
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

func (u *UserHandlerInit) Update(w http.ResponseWriter, r *http.Request) {
	var entity *user.Users
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


func (u *UserHandlerInit) Delete(w http.ResponseWriter, r *http.Request) {
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


func (u *UserHandlerInit) GetUserByPhoneNumber(w http.ResponseWriter, r *http.Request) {
	phoneNumber := chi.URLParam(r, "phoneNumber") 
	if phoneNumber == "" {
		helper.WriteError(w, http.StatusBadRequest, "Missing phone number parameter")
		return
	}

	user, err := u.userService.GetUserByPhoneNumber(phoneNumber)
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

func (u *UserHandlerInit) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email") 
	if email == "" {
		helper.WriteError(w, http.StatusBadRequest, "Missing email parameter")
		return
	}

	user, err := u.userService.GetUserByEmail(email)
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

// this doesn't work because the software I don't think are working to send it but it does get everything it needs
// and generates the reset token and sends it to the user
// will fully work on this when the frontend is ready to test it
func (u *UserHandlerInit) ForgotPassword(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ForgotPassword called")
	idnetifier := chi.URLParam(r, "emailOrPhoneNum")
	if idnetifier == "" {
		helper.WriteError(w, http.StatusBadRequest, "Missing identifier parameter")
		return
	}

	message, err := u.userService.ForgotPassword(idnetifier)
	fmt.Println("ForgotPassword called with:", idnetifier)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": message,
	})
}


func (u *UserHandlerInit) Login(w http.ResponseWriter, r *http.Request) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Parse request body
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call service layer to login
	responseData, err := u.userService.Login(loginData.Email, loginData.Password)
	if err != nil {
		helper.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	// Respond with token + user
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseData)
}





func (u *UserHandlerInit) VerifyToken(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		helper.WriteError(w, http.StatusUnauthorized, "Missing Authorization header")
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		helper.WriteError(w, http.StatusUnauthorized, "Invalid or expired token")
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"valid": true,
		"user":  claims,
	})
}



