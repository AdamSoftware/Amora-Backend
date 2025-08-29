package routes

import (
	"Amora/internal/Auth"
	"Amora/internal/reminders/reminderhandler"
	"github.com/go-chi/chi/v5"
)

type remindersModule struct {
	handler *reminderhandler.ReminderHandlerInit
}

func NewremindersModule(handler *reminderhandler.ReminderHandlerInit) RouteModule {
	return &remindersModule{handler: handler}
}

func (m *remindersModule) Prefix() string {
	return "/reminders"
}

func (m *remindersModule) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Group(func(protected chi.Router) {
		protected.Use(Auth.JWTMiddleware)

		protected.Get("/", m.handler.Get)
		protected.Get("/{id}", m.handler.FindById)
		protected.Post("/", m.handler.Insert)
		protected.Put("/{id}", m.handler.Update)
		protected.Delete("/{id}", m.handler.Delete)
		protected.Get("/user/{userId}", m.handler.FindByUserId)
		protected.Get("/reservation/{ReservationId}", m.handler.FindByReservationId)
	})

	return r
}

