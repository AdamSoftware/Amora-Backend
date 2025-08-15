package routes

import (
	"github.com/go-chi/chi/v5"
)


type RouteModule interface {
	Prefix() string
	Router() *chi.Mux
}
