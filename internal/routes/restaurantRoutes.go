package routes

import (
	"Amora/internal/restaurants/restauranthandler"
	"Amora/internal/Auth"
	"github.com/go-chi/chi/v5"
)


type restaurantsModule struct {
	handler *restauranthandler.RestaurantHandlerInit
}

func NewRestaurantModule(handler *restauranthandler.RestaurantHandlerInit) RouteModule {
	return &restaurantsModule{handler: handler}
}

func (m *restaurantsModule) Prefix() string {
	return "/restaurants"
}

func (m *restaurantsModule) Router() *chi.Mux {
	r := chi.NewRouter()
	// this one doesn't need to be protected is just returning restaurants that is all 
	r.Get("/active", m.handler.GetAllActiveRestaurants)

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
