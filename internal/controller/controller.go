package controller

import (
	"github.com/go-chi/chi"
	"net-api.com/internal/middleware"
	"net-api.com/internal/service"
)

// TODO: Inclusao de interface Controler e NewController -> DI

func RegisterRoutes(route chi.Router) {
	route.Route("/user", func(r chi.Router) {
		r.Post("/hash", service.GetHash)
	})

	route.Route("/book", func(r chi.Router) {
		r.Use(middleware.Authorization)
		r.Get("/book/random", service.GetBookRandom)
		r.Get("/book/{id}", service.GetBookByID)
	})
}
