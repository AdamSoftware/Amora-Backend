package routes

import (
	"Amora/internal/attachedGifts/attachedGiftsHandler"
	"Amora/internal/Auth"
	"github.com/go-chi/chi/v5"
)

type attachedGiftsModule struct {
	handler *attachedGiftsHandler.AttachedGiftsHandlerInit
}

func NewattachedGiftsModule(handler *attachedGiftsHandler.AttachedGiftsHandlerInit) RouteModule {
	return &attachedGiftsModule{handler: handler}
}

func (m *attachedGiftsModule) Prefix() string {
	return "/attachedgifts"
}

func (m *attachedGiftsModule) Router() *chi.Mux {
	r := chi.NewRouter()

		r.Group(func(protected chi.Router) {
		protected.Use(Auth.JWTMiddleware)

		protected.Get("/", m.handler.Get)
		protected.Get("/{id}", m.handler.FindById)
		protected.Delete("/{id}", m.handler.Delete)
		protected.Put("/{id}", m.handler.UpdateValidation)
		protected.Post("/", m.handler.InsertValidation)
		protected.Get("/event/{eventId}", m.handler.GetWithEventIdWithGiftDetails)
	})

	return r
}
