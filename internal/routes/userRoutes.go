package routes

import (
	"Amora/internal/user/userHandler"
	"github.com/go-chi/chi/v5"
	"Amora/internal/Auth"
)


func UserRouter(uHandler *userHandler.UserHandlerInit) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/users", uHandler.Insert)                           
	r.Post("/users/login", uHandler.Login)                      
	r.Post("/users/forgot/{emailOrPhoneNum}", uHandler.ForgotPassword)
	r.Post("/verify-token", uHandler.VerifyToken)         

	r.Group(func(protected chi.Router) {
		protected.Use(Auth.JWTMiddleware)

		protected.Get("/users", uHandler.Get)
		protected.Get("/users/{id}", uHandler.FindById)
		protected.Put("/users/{id}", uHandler.Update)
		protected.Delete("/users/{id}", uHandler.Delete)
		protected.Get("/users/phone/{phoneNumber}", uHandler.GetUserByPhoneNumber)
		protected.Get("/users/email/{email}", uHandler.GetUserByEmail)
	})

	return r
}

