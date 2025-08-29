package routes

import (
	"Amora/internal/Auth"
	"Amora/internal/reservations/reservationhandler"
	"github.com/go-chi/chi/v5"
)

type reservationsModule struct {
	handler *reservationhandler.ReservationHandlerInit
}

func NewreservationsModule(handler *reservationhandler.ReservationHandlerInit) RouteModule {
	return &reservationsModule{handler: handler}
}

func (m *reservationsModule) Prefix() string {
	return "/reservations"
}

func (m *reservationsModule) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Group(func(protected chi.Router) {
		protected.Use(Auth.JWTMiddleware)

		protected.Get("/", m.handler.Get)
		protected.Get("/{id}", m.handler.FindById)
		protected.Post("/", m.handler.Insert)
		protected.Put("/{id}", m.handler.Update)
		protected.Delete("/{id}", m.handler.Delete)
		protected.Get("/user/{userId}", m.handler.GetByUserId)
		protected.Get("/event/{eventId}", m.handler.GetByEventId)
	})

	return r
}

