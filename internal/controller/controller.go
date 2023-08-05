package controller

import (
	"github.com/go-chi/chi"
	"net-api.com/internal/middleware"
	"net-api.com/internal/service"
)

func RegisterRoutes(route chi.Router) {
	route.Route("/api/v1", func(r chi.Router) {
		r.Route("/user", func(subRoute chi.Router) {
			subRoute.Post("/hash", service.GetHash)
		})

		r.Route("/book", func(subRoute chi.Router) {
			subRoute.Use(middleware.Authorization)
			subRoute.Get("/random", service.GetBookRandom)
			subRoute.Get("/{bookId}", service.GetBookByID)
		})
	})
}
