package routes


import (
	"Amora/internal/messages/messagehandler"
	"Amora/internal/Auth"
	"github.com/go-chi/chi/v5"
)

type messagesModule struct {
	handler *messagehandler.MessageHandlerInit
}

func NewMessageModule(handler *messagehandler.MessageHandlerInit) RouteModule {
	return &messagesModule{handler: handler}
}

func (m *messagesModule) Prefix() string {
	return "/messages"
}

func (m *messagesModule) Router() *chi.Mux {
	r := chi.NewRouter()
	// this one doesn't need to be protected is just returning messages that is all 

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
