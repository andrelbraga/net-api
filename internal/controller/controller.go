package controller

import (
	"github.com/go-chi/chi"
	"net-api.com/internal/middleware"
	"net-api.com/internal/service"
)

type BooksController struct {
	srv *service.BookService
}

func NewBooksController(srv *service.BookService) *BooksController {
	return &BooksController{
		srv: srv,
	}
}

func (ctrl *BooksController) RegisterRoutes(route chi.Router) {
	route.Route("/api/v1", func(r chi.Router) {
		r.Route("/user", func(subRoute chi.Router) {
			subRoute.Post("/hash", ctrl.srv.GetHash)
		})

		r.Route("/book", func(subRoute chi.Router) {
			subRoute.Use(middleware.Authorization)
			subRoute.Get("/random", ctrl.srv.GetBookRandom)
			subRoute.Get("/{bookId}", ctrl.srv.GetBookByID)
		})
	})
}
