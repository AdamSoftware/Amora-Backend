package routes

import (
	"Amora/internal/gifts/gifthandler"
	"Amora/internal/Auth"
	"github.com/go-chi/chi/v5"
)

type giftsModule struct {
	handler *gifthandler.GiftHandlerInit
}

func NewGiftModule(handler *gifthandler.GiftHandlerInit) RouteModule {
	return &giftsModule{handler: handler}
}

func (m *giftsModule) Prefix() string {
	return "/gifts"
}

func (m *giftsModule) Router() *chi.Mux {
	r := chi.NewRouter()
	// this one doesn't need to be protected is just returning gifts that is all 
	r.Get("/active", m.handler.GetAllActiveGifts)

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
