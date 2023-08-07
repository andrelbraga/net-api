package router

import (
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	_ "net-api.com/docs" // docs is generated by Swag CLI, you have to import it.

	"net-api.com/internal/controller"
	"net-api.com/internal/middleware"
)

// BooksRouter
type BooksRouter struct {
	*controller.BooksController
}

// NewBooksRouter
func NewBooksRouter(ctrl *controller.BooksController) *BooksRouter {
	return &BooksRouter{
		ctrl,
	}
}

// @title Echo Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9191
// @BasePath /api/v1
// @schemes http
func (ctrl *BooksRouter) RegisterRoutes(route chi.Router) {
	route.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:9191/swagger/doc.json"),
	))
	route.Route("/api/v1", func(r chi.Router) {
		r.Route("/user", func(subRoute chi.Router) {
			subRoute.Post("/hash", ctrl.GetHashCtrl)
		})

		r.Route("/book", func(subRoute chi.Router) {
			subRoute.Use(middleware.Authorization)
			subRoute.Get("/random", ctrl.GetBookRandomCtrl)
			subRoute.Get("/{id}", ctrl.GetBookByIDCtrl)
		})
	})
}