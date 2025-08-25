package routes

import (
	"Amora/internal/experiences/experiencehandler"
	"Amora/internal/Auth"
	"github.com/go-chi/chi/v5"
)


type experiencesModule struct {
	handler *experiencehandler.ExperienceHandlerInit
}

func NewExperienceModule(handler *experiencehandler.ExperienceHandlerInit) RouteModule {
	return &experiencesModule{handler: handler}
}

func (m *experiencesModule) Prefix() string {
	return "/experiences"
}

func (m *experiencesModule) Router() *chi.Mux {
	r := chi.NewRouter()
	// this one doesn't need to be protected is just returning experiences that is all 
	r.Get("/active", m.handler.GetAllActiveExperiences)

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
