package handlers

import (
	"github.com/Meruya-Technology/go-boilerplate/common/configs"
	"github.com/go-chi/chi"
)

type Handlers struct {
	cfg *configs.Config
}

// Provide is the provider of this handler
func Provide(
	cfg *configs.Config) *Handlers {
	return &Handlers{
		cfg: cfg,
	}
}

// Router is the function router
func (h *Handlers) Router(r chi.Router) {
	// if h.cfg.Transport.HTTP.Router.EnabledSyncRouter {
	// 	r.Route("/stats", func(r chi.Router) {
	// 		r.Use(h.auth.InternalCredentials)
	// 		r.Route("/model", func(r chi.Router) {
	// 			r.Put("/accumulate", h.AccumulateModelStats)
	// 		})
	// 	})
	// }

	// r.Route("/catalog", func(r chi.Router) {
	// 	r.Get("/list", h.GetCatalogList)
	// 	r.Get("/product/list/{catalogId}", h.GetCatalogProductList)
	// })

	// r.Route("/", func(r chi.Router) {
	// 	r.Get("/search", h.SearchProduct)
	// 	if h.cfg.Transport.HTTP.Router.EnabledSyncRouter {
	// 		r.Post("/search/index", h.IndexingProductSearch)
	// 	}
	// })
}

// Router is the function router for v2 version
// func (h *Handlers) RouterV2(r chi.Router) {
// 	h.v2.Router(r)
// }
