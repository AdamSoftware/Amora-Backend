package routes

import (
	"Amora/internal/user/userHandler"
	"Amora/internal/Auth"
	"github.com/go-chi/chi/v5"
)

type userModule struct {
	handler *userHandler.UserHandlerInit
}

func NewUserModule(handler *userHandler.UserHandlerInit) RouteModule {
	return &userModule{handler: handler}
}

func (m *userModule) Prefix() string {
	return "/users"
}

func (m *userModule) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/", m.handler.Insert)
	r.Post("/login", m.handler.Login)
	r.Post("/forgot/{emailOrPhoneNum}", m.handler.ForgotPassword)
	r.Post("/verify-token", m.handler.VerifyToken)

	r.Group(func(protected chi.Router) {
		protected.Use(Auth.JWTMiddleware)

		protected.Get("/", m.handler.Get)
		protected.Get("/{id}", m.handler.FindById)
		protected.Put("/{id}", m.handler.Update)
		protected.Delete("/{id}", m.handler.Delete)
		protected.Get("/phone/{phoneNumber}", m.handler.GetUserByPhoneNumber)
		protected.Get("/email/{email}", m.handler.GetUserByEmail)
	})

	return r
}
