package routes

import (
	"Amora/internal/user/userHandler"
	"github.com/go-chi/chi/v5"
)

func UserRouter() *chi.Mux {
	uHandler := userHandler.NewUserHandler()
	r := chi.NewRouter()

	r.Get("/users", uHandler.Get)                  
	r.Get("/users/{id}", uHandler.FindById)        
	r.Post("/users", uHandler.Insert)              
	r.Put("/users/{id}", uHandler.Update)         
	r.Delete("/users/{id}", uHandler.Delete)     
	r.Get("/users/phone", uHandler.GetUserByPhoneNumber)
	r.Get("/users/email", uHandler.GetUserByEmail)
	r.Post("/users/login", uHandler.Login)
	r.Post("/users/forgot-password", uHandler.ForgotPassword)

	return r
}

