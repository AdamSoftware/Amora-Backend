package routes

import (
	"Amora/internal/Auth"
	"Amora/internal/events/eventHandler"
	"github.com/go-chi/chi/v5"
)

type eventsModule struct {
	handler *eventHandler.EventHandlerInit
}

func NeweventsModule(handler *eventHandler.EventHandlerInit) RouteModule {
	return &eventsModule{handler: handler}
}

func (m *eventsModule) Prefix() string {
	return "/events"
}

func (m *eventsModule) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Group(func(protected chi.Router) {
		protected.Use(Auth.JWTMiddleware)

		protected.Get("/", m.handler.Get)
		protected.Get("/{id}", m.handler.FindById)
		protected.Post("/", m.handler.Insert)
		protected.Put("/{id}", m.handler.Update)
		protected.Delete("/{id}", m.handler.Delete)
		protected.Get("/user/{userId}", m.handler.FindByUserId)
		protected.Get("/gift/user/{userId}", m.handler.GetEventsWithGiftsWithUserId)
		protected.Get("/res/user/{userId}", m.handler.GetEventsWithReservationsWithUserId)
	})

	return r
}

