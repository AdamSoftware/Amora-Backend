package routes

import (
	"Amora/internal/Auth"
	"Amora/internal/subscriptions/subscriptionHandler"
	"github.com/go-chi/chi/v5"
)

type subscriptionModule struct {
	handler *subscriptionHandler.SubscriptionHandlerInit
}

func NewSubscriptionModule(handler *subscriptionHandler.SubscriptionHandlerInit) RouteModule {
	return &subscriptionModule{handler: handler}
}

func (m *subscriptionModule) Prefix() string {
	return "/subscriptions"
}

func (m *subscriptionModule) Router() *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(protected chi.Router) {
		protected.Use(Auth.JWTMiddleware)

		protected.Get("/{id}", m.handler.FindById)
		protected.Get("/", m.handler.Get)
		protected.Post("/", m.handler.Insert)
		protected.Put("/{id}", m.handler.Update)
		protected.Delete("/{id}", m.handler.Delete)
		protected.Get("/plan/{planId}", m.handler.GetByPlanId)
		protected.Get("/user/{userId}", m.handler.GetByUserId)
		protected.Get("/start-date-before", m.handler.FindByStartDateBefore)
		protected.Get("/status/{status}", m.handler.GetByStatus)
		protected.Post("/check-user-subscription", m.handler.CheckIfUserHadSubscription)
		protected.Put("/user/{userId}", m.handler.UpdateByUserId)
	})

	return r
}

