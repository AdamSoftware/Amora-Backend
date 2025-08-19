package routes

import (
	"Amora/internal/Auth"
	"Amora/internal/partner/partnerHandler"
	"github.com/go-chi/chi/v5"
)

type partnerModule struct {
	handler *partnerHandler.PartnerHandlerInit
}

func NewpartnerModule(handler *partnerHandler.PartnerHandlerInit) RouteModule {
	return &partnerModule{handler: handler}
}

func (m *partnerModule) Prefix() string {
	return "/partners"
}

func (m *partnerModule) Router() *chi.Mux {
	r := chi.NewRouter()
	// public routes
	r.Get("/birthday/{SpicyContent}", m.handler.FindAllPartnersWithBirthday)
	r.Get("/anniversary/{SpicyContent}", m.handler.FindAllPartnersWithAnniversary)

	r.Group(func(protected chi.Router) {
		protected.Use(Auth.JWTMiddleware)

		protected.Get("/", m.handler.Get)
		protected.Get("/{id}", m.handler.FindById)
		protected.Post("/", m.handler.Insert)
		protected.Put("/{id}", m.handler.Update)
		protected.Delete("/{id}", m.handler.Delete)
		protected.Get("/phone/{phoneNumber}", m.handler.FindByPhoneNumber)
		protected.Get("/user/{userId}", m.handler.FindByUserId)
	})

	return r
}

