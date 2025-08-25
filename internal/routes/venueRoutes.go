package routes

import (
	"Amora/internal/venues/venuehandler"
	"Amora/internal/Auth"
	"github.com/go-chi/chi/v5"
)


type venuesModule struct {
	handler *venuehandler.VenueHandlerInit
}

func NewVenueModule(handler *venuehandler.VenueHandlerInit) RouteModule {
	return &venuesModule{handler: handler}
}

func (m *venuesModule) Prefix() string {
	return "/venues"
}

func (m *venuesModule) Router() *chi.Mux {
	r := chi.NewRouter()
	// this one doesn't need to be protected is just returning venues that is all 
	r.Get("/active", m.handler.GetAllActiveVenues)

	r.Group(func(protected chi.Router) {
		protected.Use(Auth.JWTMiddleware)

		protected.Get("/{id}", m.handler.FindById)
		protected.Get("/", m.handler.Get)
		protected.Post("/", m.handler.Insert)
		protected.Put("/{id}", m.handler.Update)
		protected.Delete("/{id}", m.handler.Delete)
	})

	return r
}
