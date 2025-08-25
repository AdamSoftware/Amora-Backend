package routes

import (
	"Amora/internal/plans/planhandler"
	"Amora/internal/Auth"
	"github.com/go-chi/chi/v5"
)

type plansModule struct {
	handler *planhandler.PlanHandlerInit
}

func NewPlanModule(handler *planhandler.PlanHandlerInit) RouteModule {
	return &plansModule{handler: handler}
}

func (m *plansModule) Prefix() string {
	return "/plans"
}

func (m *plansModule) Router() *chi.Mux {
	r := chi.NewRouter()
	// this one doesn't need to be protected is just returning plans that is all 

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
