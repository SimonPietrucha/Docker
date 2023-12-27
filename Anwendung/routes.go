package Anwendung

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes(productHandler *Product) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	router.Route("/product", func(r chi.Router) {
		loadProductRoutes(r, productHandler)
	})

	return router
}

func loadProductRoutes(router chi.Router, productHandler *Product) {
	router.Post("/", productHandler.Create)
	router.Get("/", productHandler.List)
	router.Get("/{id}", productHandler.GetByID)
	router.Put("/{id}", productHandler.UpdateByID)
	router.Delete("/{id}", productHandler.DeleteByID)
	router.Patch("/{id}/aktualisiereBestand", productHandler.AktualisiereBestandByID)
}
