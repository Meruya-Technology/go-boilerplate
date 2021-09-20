package router

import (
	"github.com/evermos/evm-product/internal/handlers"
	"github.com/go-chi/chi"
)

// Router is the router struct containing handlers.
type Router struct {
	handlers *handlers.Handlers
}

// ProvideRouter is the provider function for this router.
func ProvideRouter(handlers *handlers.Handlers) Router {
	return Router{
		handlers: handlers,
	}
}

// SetupRoutes sets up all routing for this server.
func (r *Router) SetupRoutes(mux *chi.Mux) {
	mux.Route("/v3", func(rc chi.Router) {
		r.handlers.Router(rc)
	})
	mux.Route("/v2", func(rc chi.Router) {
		r.handlers.RouterV2(rc)
	})
}
